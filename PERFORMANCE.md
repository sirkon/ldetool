# PERFORMANCE

##### Manual comparison against CLI tools
There a comparison against `gawk`, `sed` and Go's `regex` implementation in processing 1.3Gb of data: [here](PERFORMANCE_MANUAL.md)
##### Automated Go comparsion against Ragel and stdlib regex
```
go test -v github.com/sirkon/ldetool/benchmarking
```
There's parameter on a [line](https://github.com/sirkon/ldetool/blob/6be94610ca6da1fbf0cfe8e2c18e27792622a320/benchmarking/performance_test.go#L29)
where you can tweak first field's maximal length. Generally, the longer field (and thus the further character lookup bounding it),
the more advantage has LDE generated code because it uses highly optimized `bytes.IndexByte` function for lookups.
Using Ragel for log parsing hardly makes any sense though, because it doesn't seem any easier to write Ragel actions and rules instead of looking up for bounding characters and strings manually since the amount of boilerplate is equal or even more with Ragel, while advanced features of finite state machines are rarely needed at all and it is rather a reason to tweak log output if they are instead of utilizing Ragel.

1. 16 symbols
    ```
    $ go test -v -bench . github.com/sirkon/ldetool/benchmarking
    BenchmarkLDE-4     	   30000	     54751 ns/op
    BenchmarkRagel-4   	   10000	    113695 ns/op
    BenchmarkRegex-4   	     500	   3141558 ns/op
    PASS
    ok  	github.com/sirkon/ldetool/benchmarking	5.244s
    ```
2. 64 symbols
    ```
    $ go test -v -bench . github.com/sirkon/ldetool/benchmarking
    BenchmarkLDE-4     	   20000	     62158 ns/op
    BenchmarkRagel-4   	   10000	    141991 ns/op
    BenchmarkRegex-4   	     500	   3944421 ns/op
    PASS
    ok  	github.com/sirkon/ldetool/benchmarking	5.686s

    ```
3. 256 symbols
    ```
    $ go test -v -bench . github.com/sirkon/ldetool/benchmarking
    BenchmarkLDE-4     	   20000	     69599 ns/op
    BenchmarkRagel-4   	    5000	    241497 ns/op
    BenchmarkRegex-4   	     200	   7212705 ns/op
    PASS
    ok  	github.com/sirkon/ldetool/benchmarking	5.513s
    ```
4. 1024 symbols
    ```
    $ go test -v -bench . github.com/sirkon/ldetool/benchmarking
    BenchmarkLDE-4     	   20000	     90019 ns/op
    BenchmarkRagel-4   	    2000	    626500 ns/op
    BenchmarkRegex-4   	     100	  20325788 ns/op
    PASS
    ok  	github.com/sirkon/ldetool/benchmarking	6.106s
    ```

##### Automated comparison against Go regex on real world sample
1. LDE rule first
    ```perl
    CRMod = !
        ^'[' _' ' Time(string) ']'
        _'[' ChatID(uint64) '.'
        _"reqid '" ReqID(string) '\''
        _"from" _'(' UIN(string) ')'
        _"FLAGS[set:" FlagsSet(string) ','
        ^" unset:" FlagsUnset(string) ']'
        ^" FIELDS[changed:" FieldsChanged(string) ']'
        ?AnkVer (^" ank_ver[" Value(string) ']')
        ?ListVer (^" list_ver[" Value(string) ']')
        ^" name[" Name(string) ']'
        ?About (^" about[" Value(string) ~']')
        ?Rules (^" rules[" Value(string) ~']')
        ?Nick (^" nick[" Value(string) ']')
        ?Location(^" location[" Value(string) ']')
        ?Stamp(^" stamp[" Value(string) ']')
        ?Regions(^" regions[" Value(string) ~']')
        ?Flags(^" flags[" Value(string) ']')
        ^" created[" Created(int64) '='
        ?Creator(_"creator[" Value(string) ']')
        ?AvatarLastCheck(_"avatars_lastcheck[" Value(int64) ']')
        ?AvatarsLastMod(_"cavatar_lastmod[" Value(int64) ']')
        ^" origin[" Origin(string) ~']'
        ^" abuse" ^"drugs["[1] Drugs(int16) ~']'
        ^" abuse" ^"spam["[1] Spam(int16) ~']'
        ^" abuse" ^"porno["[1] Pron(int16) ~']'
        ?Violation  (^" abuse" ^"violation["[1] Value(int16) ~']')
        ?AbuseOther (^" abuse" ^"other["[1] Value(int16) ~']');
    ```
    You see, we have all data translated into types needed and usable error messages.
2. Regex, about 490 characters hard to debug language without any boilerplate. Good luck optimizing this mess.
    ```
    \[\S* (.*?)\][^[]*\[(\d+)\..*?reqid '(.*?)' from.*?\((.*?)\).*?FLAGS\[set:(.*?), unset:(.*?)\] FIELDS\[changed:(.*?)\](:? ank_ver\[(.*?)\])?(:? list_ver\[(.*?)\])? name\[(.*?)\](:? about\[(.*?)\])?(:? stamp\[(.*?)\])?(:? regions\[(.*?)\])?(:? flags\[(.*?)\])? created\[(\d+)=.*?\](:?.*?creator\[(.*?)\])?(:?.*?avatars_lastcheck\[(\d+)\])?(:?.*?cavatar_lastmod\[(\d+)\])? origin\[(.*?)\] abuse.*drugs\[(\d+)\] abuse.*spam\[(\d+)\] abuse.*porno\[(\d+)\](:? abuse.violation\[(\d+)\])?(:? abuse.other\[(\d+)\])?
    ```
```
$ go test -v -bench '.*Complex.*' github.com/sirkon/ldetool/benchmarking

BenchmarkLDEComplex-4            1000000              2116 ns/op
BenchmarkRegexComplex-4             1000           2169577 ns/op
PASS
ok      github.com/sirkon/ldetool/benchmarking  4.537s
```
You see, specialized solution about 1000 times faster, much easier to write and debug and does a lot of boilerplating
beneath — we have numeric fields converted on successful extraction, we have errors where we failed on action processing, etc.
Regexes can't do anything of that. 

##### Automated comparsion against Ragel on a close to real world sample
1. Lines will look like
    ```
    [12345 2017-10-10T21:11:12] PRESENCE uid=123423546455654 ua='App.Com Samsung i7300 Android/8.0.0/RU' Geo={Lat: 12.0, Lon: 13.0} Acitivity=0
    [12345 2017-10-10T21:11:12] PRESENCE uid=123423546455654 ua='App.Com iPhone 8+ iOS/11.0.2/DE' Geo={Lat: 14.0, Lon: 15.0} Acitivity=0
    [12345 2017-10-10T21:11:12] PRESENCE uid=123423546455654 ua='App.Com Windows x86-64/7/BY' Acitivity=1
    ```
    where `uid` is user identifier, `ua` is user agent, `Geo` is optional geo information and `Activity` is a kind of activity the user was making (background, typing, etc)
2. LDE rule without type conversion
    ```perl
    Presence =
        _' ' Time(string) ']'
        ^" PRESENCE uid=" ! UID(string) ' '
        ^"ua='" UA(string) '\''
        ?Geo (
            ^" Geo={Lat: " Lat(string) ','
            ^" Lon: " Lon(string) '}'
        )
        ^" Activity=" Activity(string);
    ```
2. LDE rule with type conversion
    ```perl
    PresenceFloats =
    _' ' Time(string) ']'
    ^" PRESENCE uid=" ! UID(string) ' '
    ^"ua='" UA(string) '\''
    ?Geo (
        ^" Geo={Lat: " Lat(float64) ','
        ^" Lon: " Lon(float64) '}'
    )
    ^" Activity=" Activity(uint8);
    ```
3. Ragel template without type conversion
    ```perl
    package main

    // Easy based parsing
    type Easy struct {
    	Time []byte
    	UID  []byte
    	UA   []byte
    	Geo  struct {
    		Valid bool
    		Lat   []byte
    		Lon   []byte
    	}
    	Activity []byte
    }

    %% machine easy;
    %% write data;

    // Extract extracts field from
    func (r *Easy) Extract(data []byte) (ok bool, error error) {
        cs, p, pe := 0, 0, len(data)
        var pos = 0
        r.Geo.Valid = false

        %%{
            action shot       { pos = p + 1                }
            action take_time  { r.Time = data[pos:p+1]     }
            action take_uid   { r.UID = data[pos:p+1]      }
            action take_ua    { r.UA = data[pos:p+1]       }
            action take_lat   { r.Geo.Lat = data[pos:p+1]  }
            action take_lon   { r.Geo.Lon = data[pos:p+1]  }
            action take_act   { r.Activity = data[pos:p+1] }
            action set_geo    { r.Geo.Valid = true         }

            ns = (any -- " ")*;
            main :=
                 ns " "@shot ((any -- "]")*)@take_time "] PRESENCE uid="@shot
                 ns@take_uid " ua='"@shot ((any -- "'")*)@take_ua "' "@shot
                 (
                    "Geo={Lat: "@set_geo@shot ((any -- ",")*)@take_lat ", Lon: "@shot ((any -- "}")*)@take_lon "} "@shot
                 )?
                 "Activity="@shot (any*)@take_act
                 ;
            write init;
            write exec;
        }%%
        return true, nil
    }
    ```
4. Ragel template with type conversion
    ```perl
    package main
    
    
    import (
        "unsafe"
        "strconv"
    )
    
    // EasyFloat based parsing
    type EasyFloat struct {
    	Time []byte
    	UID  []byte
    	UA   []byte
    	Geo  struct {
    		Valid bool
    		Lat   float64
    		Lon   float64
    	}
    	Activity uint8
    }
    
    %% machine easyfloats;
    %% write data;
    
    // Extract extracts field from
    func (r *EasyFloat) Extract(data []byte) (ok bool, err error) {
        cs, p, pe := 0, 0, len(data)
        var pos = 0
        r.Geo.Valid = false
        var tmpFloat float64
        var tmpUint uint64
        var tmp []byte
    
        %%{
            action shot       { pos = p + 1                }
            action take_time  { r.Time = data[pos:p+1]     }
            action take_uid   { r.UID = data[pos:p+1]      }
            action take_ua    { r.UA = data[pos:p+1]       }
    	action tmp_float  { 
                tmp = data[pos:p+1]
                if tmpFloat, err = strconv.ParseFloat(*(*string)(unsafe.Pointer(&tmp)), 64); err != nil {
                    return false, err
                }
    	}
            action take_lat   { r.Geo.Lat = tmpFloat       }
            action take_lon   { r.Geo.Lon = tmpFloat       }
            action take_act   {
                tmp = data[pos:p+1]
                if tmpUint, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&tmp)), 10, 8); err != nil {
                    return false, err
                }
                r.Activity = uint8(tmpUint)
            }
            action set_geo    { r.Geo.Valid = true         }
    
            ns = (any -- " ")*;
            main :=
                 ns " "@shot ((any -- "]")*)@take_time "] PRESENCE uid="@shot
                 ns@take_uid " ua='"@shot ((any -- "'")*)@take_ua "' "@shot
                 (
                    "Geo={Lat: "@set_geo@shot ((any -- ",")*)@tmp_float@take_lat ", Lon: "@shot ((any -- "}")*)@tmp_float@take_lon "} "@shot
                 )?
                 "Activity="@shot (any*)@take_act
                 ;
            write init;
            write exec;
        }%%
        return true, nil
    }
    ```
5. And regex with data extraction without conversion
```regexp
^\S*\s([^\]]+)] PRESENCE uid=(\S*) ua='([^']*)' (:?Geo=\{Lat: ([^,]+), Lon: ([^,]+)\} )?Activity=(.*)$
```

> both these Ragel templates only does processing without error handling, so generated code is not production ready.
> The problem here we will need to handle type conversion and error processing manually each time writing Ragel rules.
> The LDE tool makes this automatically. This alone is a #1 in a list of *pros* for using LDE, even if the code generated
> with Ragel would be a bit faster.

Now, let's benchmark:

```
$ go test -v -bench '.*RealWorld.*' github.com/sirkon/ldetool/benchmarking

BenchmarkLDEEasyRealWorld-4                     10     172518252 ns/op
BenchmarkLDEEasyFloatsRealWorld-4                5     217304418 ns/op
BenchmarkRagelEasyRealWorld-4                    5     295341158 ns/op
BenchmarkRagelEasyFloatsRealWorld-4              2     626229546 ns/op
BenchmarkRegexEasyRealWorld-4                    1    3308693182 ns/op
PASS
ok  	github.com/sirkon/ldetool/benchmarking	9.218s
```

You see, not only LDE generated code does a lot more than straight Ragel, it is actually faster, something like several
times faster. Notice a two times performance drop with type conversions on Ragel sample, when the LDE generated code
suffers only %30 speed decrease in the same circumstances: it looks like Ragel works best when all actions are done within
generated finite state machine, probably something with cache locality. It slows down immediately after there was an
"external" function call. Notice, the regexp is not THAT bad as it was in the previous example: only 19 times slower than
the code generated with LDE
