# ldetool means line data extraction tool
`ldetool` is a command line utility to generate Go code for log files parsing. 

```bash
go get -u github.com/sirkon/ldetool
```

### Preamble

There's a traditional solution for this kind of tasks: regular expression with capture groups. But it has numerous generic and Go-specific disadvantages:

1. Regexes are hard to read and debug.
2. Speed. While simple non-capturing regular expressions can be speedy, they quickly becomes slow as the complexity of the regular expression grows
3. They are overpowered for simple log parsing. In our experience with log processing we are not looking for patterns within the line. Usually our data is well structured and it is easier to think (and compute!) in terms of bounds and separators. And if the data is not well structured then it is a good idea to make it so, just for the sake of readability.
4. Go regular expressions are slow. Go regular expressions with group capture are even slower.
5. There are no cheap way in Go regexes what would give us a convenient way to access a group's value, we must use arrays instead of access to captured value by group name, thus it is hard for reading and comprehension.

There is another traditional approach: manual data extraction. We manually command to find a symbol or substring and pass
it or take everything before it and put into variable, it also has his share of generic disadvantages:

1. It is annoying as hell to write it
2. It can be hard to read

Still, the major advantage is:
1. It can be fast

We had severe shortage of resources at my last job, we couldn't just buy some more power, so we had no choice. We had to write it manually.
It turned out most of things to retrieve data are repetitive and we are writing nearly the same things again and again.

##### Typical operations:
1. Check if the rest starts with the certain string or character and pass it
2. Just pass first N characters
3. Look for the char or substring (in what follows we shall call this character or string the __target__) in the rest and pass it
4. Take all data from the rest up to the target and save it somewhere having some type with conversion if possible (string to numeric)
5. Take the rest and save it somewhere
6. Take until the target or the the whole rest if not found and save it somewhere
7. Optional areas: there should be a possibility to ignore some subset of the extraction rule if it wasn't successful for the rest and roll the rest back to
    the start of failed attempt, like this:
    ```perl
    Rule = 
        ?Start( 
            ^" start=" Time(string) ' '
        )
        ^" rest=\"" Rest(string) '"';
    ```
    Both lines
    ```
    start=2017-09-28T16:58:23 rest="The rest is here"
    rest="Just the rest"
    ```
    must should pass the extraction with the result
    ```
    (Start="2017-09-28T16:58:23", Rest="The rest is here"),
    (Start="", Rest="Just the rest")
    ```
8. Target can bo not just a string or character: there might be some kind of limitations, such as apriori knowledge of
    where the target actually is, some range or even exact position. Also, `bytes.IndexByte` works great for character
    lookup, it can use vector CPU instruction for speed bump, but it is not [inlined](https://github.com/golang/go/issues/21759#issuecomment-327124437)
    and call overhead can be quite significant for closer targets, where the naive loop search can be faster.

So, we wrote a code generator for this purpose. The code turned to be even faster than one we used to write, since we actually
were trying to reduce amount of code introducing helper abstractions what have some cost while the generator just put raw code.

### How it works.
1. Write extraction script.
2. Generate go code using extraction script.
3. Use it via the generated extraction method

#### Example

Take a look at these two lines

```
[2017-09-02T22:48:13] FETCH first[1] format[JSON] hidden[0] userAgent[Android App v1.0] rnd[21341975] country[MA]
[2017-09-02T22:48:14] FETCH first[0] format[JSON] userAgent[Android App v1.0] rnd[10000000] country[LC]
```

We likely need a time, value of parameter `first`, `format`, `hidden`, `userAgent` and `country`. We obviously don't need `rnd` 

##### Extraction script syntax
See [more details](https://github.com/sirkon/ldetool/blob/master/TOOL_RULES.md) on extraction rules

```perl
# filename: line.lde
Line =                                   # Name of the extraction object' type
  ^'[' Time(string) ']'                  # The line must start with [, then take everything as a struct field Time string right to ']' character
  ^" FETCH "                             # Current rest must starts with " FETCH " string
  ^"first[" First(uint8) ']'[1]          # The rest must starts with "first[" characters, then take the rest until ']' as uint8. It is
                                         # known First is the single character, thus the [1] index.
                                         # under the name of First
  ^" format[" Format(string) ~']'        # Take format id. Format is a short word: XML, JSON, BIN. ~ before lookup oobject suggests
                                         # generator to use for loop scan rather than IndexByte, which is although fast
                                         # has call overhead as it cannot be inlined by Go compiler.
  ?Hidden (^" hidden[" Value(uint8) ']') # Optionally look for " hidden[\d+]"
  ^" user_agent[" UserAgent(string) ']'  # User agent data
  _ "country[" Country(string)  ']'      # Look for the piece starting with country[
;
```

##### Code generation
The easiest way is to put `//go:generate ldetool generate --package main Line.lde` somewhere in the Go file and then generate a code with 
```bash
go generate <project path>
```
It will be written into `line_lde.go` file in the same directory. It will look like [this](SAMPLE.md)

Now, we have
1. Data extractor type
    ```go
    // Line autogenerated parser
    type Line struct {
        rest   []byte
        Time   []byte
        First  uint8
        Format []byte
        Hidden struct {
            Valid bool
            Value uint8
        }
        UserAgent []byte
        Country   []byte
    }
    ```
2. Parse method
    ```go
    // Extract autogenerated method of Line
    func (p *Line) Extract(line []byte) (bool, error) {
       …
    }
    ```
    Take a look at return data. First bool signals if the data was successfully matched and error signals if there were
    any error. String to numeric failures are always treated as errors, you can put `!` into extraction script and all
    mismatches after the sign will be treated as errors
3. Helper to access optional `Hidden` area returning default Go value if the the area was not matched
    ```go
    // GetHiddenValue retrieves optional value for HiddenValue.Name
    func (p *Line) GetHiddenValue() (res uint8) {
        if !p.Hidden.Valid {
            return
        }
        return p.Hidden.Value
    }    
    ```
    
##### Generated code usage
It is easy: put
```go
l := &Line{}
```
before and then feed `Parse` method with lines:
```go
scanner := bufio.NewScanner(reader)
for scanner.Scan() {
    ok, err := l.Extract(scanner.Bytes())
    if !ok {
        if err != nil {
            return err
        }
        continue
    }
    …
    l.Format
    l.Time
    l.GetHiddenValue()
    …
}
```

### Performance
##### Manual comparison against CLI tools
There a comparison against `gawk`, `sed` and Go's `regex` implementation in processing 1.3Gb of data: [here](https://github.com/sirkon/ldetool/blob/master/PERFORMANCE.md)
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
beneath — we have numeric fields converted on successful extraction, we have errors where we failed on action processing, etc. We have nothing of it with regexes.

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
4. Ragel template with type conversion is [here](benchmarking/easy_floats.ragel).

Note, both these Ragel templates only does processing without error handling, so generated code is not production ready.
The problem here we will need to handle type conversion and error processing manually each time writing Ragel rules.
The LDE tool makes this automatically. This alone is a #1 in a list of *pros* for using LDE, even if the code generated
with Ragel will be a bit faster. Now, let's benchmark: 

```
$ go test -v -bench '.*RealWorld.*' github.com/sirkon/ldetool/benchmarking

BenchmarkLDEEasyRealWorld-4           	      10	 168081059 ns/op
BenchmarkLDEEasyFloatsRealWorld-4     	       5	 216827185 ns/op
BenchmarkRagelEasyRealWorld-4         	       5	 301412608 ns/op
BenchmarkRagelEasyFloatsRealWorld-4   	       2	 619422100 ns/op
PASS
ok  	github.com/sirkon/ldetool/benchmarking	9.218s
```

You see, not only LDE generated code does a lot more than straight Ragel, it is actually faster, something like several
times faster. Notice a two times performance drop with type conversions on Ragel sample, when the LDE generated code
suffers only %30 speed decrease in the same circumstances: it looks like Ragel works best when all actions are done within 
generated finite state machine, probably something with cache locality. It slows down immediately after there was an 
"external" function call.
