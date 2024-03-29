# ldetool means line data extraction tool

`ldetool` is a command line utility to generate Go code for fast log files parsing.

```bash
go install github.com/sirkon/ldetool@latest
```

1. [Contributors](CONTRIBUTORS.md)
1. [Rationale](RATIONALE.md)
2. [Typical operations and formal set of rules](TOOL_RULES.md)
3. [Performance comparison against regexp and Ragel](PERFORMANCE.md)
5. [Usage examples](EXAMPLES.md)


### How it works.
1. First write extraction script, we usually name it `<something>.lde`
2. Generate go code with `ldetool <something.lde> --package main`. Of course
   you can use your own package name, not only `main`
3. Use it via the generated extraction method `Parse(line []byte)`.

> It turned out we like using it even for non-performant tasks, where we are dealing with strings, not slices of bytes 
> and it would be handy to use it for strings as well without manual type casting. There's an option to generate code
> that use string, just put an option `--go-string`

##### CLI utility options
1. `--go-string` generates code that uses `string` everywhere instead of `[]byte`. You better not to use it for log processing as it may lead to excessive memory allocations.
2. `--yaml-dict` or `--json-dict` sets translation rules for names. For instance, if we have YAML file with
    ```yaml
    http: HTTP
    ```
    and feed this file to the `ldetool` then every name (of field or rule itself) like `GetHttpHandle` or `get_http_handle` will be translated into `GetHTTPHandle`
3. `--package <pkg name>` name of the package to use in generated code. If a directory of `*.lde` file has other Go files package name will automatically setup with these files' package name.
4. `--big-endian` or `--little-endian` sets the target architecture to be either big or little endian. This
    enables prefix check optimization 

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
The recommended way is to put something like `//go:generate ldetool --package main Line.lde` in `generate.go` of a package and then generate a code with
```bash
go generate <project path>
```
It will be written into `line_lde.go` file in the same directory. It will look like [this](SAMPLE.md)

Now, we have
1. Data extractor type
    ```go
    // Line autogenerated parser
    type Line struct {
        Rest   []byte
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
    Take a look at return data. First bool signals if the data was successfully matched and error that is not nil signals if there were
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
before and then feed `Extract` method with lines:
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

#### custom types

> Special thanks to Matt Hook (github.com/hookenz) who proposed this feature

It is possible to use custom types in generated structure. You should declare them first via

```perl
type pkg.Type from "pkgpath";
```

for external types and 

```perl
type typeName;
```

for local types before all rules definitions and you can use them as field types. The parsing to be done via

```go
p.unmarshal<FieldName>([]byte) (Type, error)
``` 

function.

Example:

```perl
type time.Time from "time";
type net.IP from "net";

Custom = Time(time.Time) ' ' ?Addr(^"addr: " IP(ip.IP) ' ');
```

Now, two parsing functions will be needed to parse this (they are to be written manually):

```go
func (p *Custom) unmarshalTime(s string) (time.Time, error) { … }

func (p *Custom) unmarshalAddrIP(s string) (net.IP, error) { … }
```
