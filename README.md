# ldegen means line data extraction generator

### Preamble

There's a traditional solution for this kind of tasks: regular expression with capture groups. But it has numerous generic and Go-specific disadvantages:

1. Syntax. Hard to debug and read.
2. Speed. While simple non-capturing regular expressions can be speedy, they quickly becomes slow as the complexity of the regular expression grows
3. They are overpowered. In our experience with log processing we are not looking for patterns within the line. We just need to find some substring, then take everthing between this substring and the next comma as a number, or a string without allocation, just pointing at the fragment extracted from the line. It may cost us additional CPU time to parse and allocate on capture. Not a good thing when we have billions of lines to process.
4. Go regular expressions are slow. Go regular expressions with group capture are even slower.

### Proposal

Look at lines

```
[19234 2017-07-15T15:41:39] FETCH_EVENTS first=1 format=json responseTime=1233 hidden=1 user_agent="Android 7.0.1 App 10.1.0alpha" country=RU /libs/network/fetch_handler.cpp:408
[19234 2017-07-15T15:41:39] FETCH_EVENTS first=0 format=proto responseTime=1233 user_agent="Android 6.0.1 App 9.8.24" country=RU
```

How would we extract all needed data without regexes?
This is a possible way:
* We have the rest of line. What can we do with that, basic operations:
  1. We can look for the some string or character in the rest or in the first N characters in the rest.
  2. We can check if the rest starts with the string or character.
  3. We can pass first N characters
  4. We can take all characters of the rest right to the string or character
  5. We can take the rest of the string
  6. We can take all characters of the rest right to the string of character if it found or take the rest otherwise.
* In practice, we extremely rarely need matched boundary string or character, thus if the string or character matched on operations I, II, IV and VI we pass it and the rest is taken after that match. Obviously, if we took all characters to the end the rest becomes empty.
* There should be a possibility to put subsequences of these operations into optional groups.

#### Syntax for extraction of needed data for these particular lines

```perl
line = 
  _ ' '                                  # Pass to the space (x20) character
  Time(string) ']'                       # Take everything as a record for Time right to ']' character
  ^" FETCH_EVENTS "                      # Current rest must starts with " FETCH_EVENTS " string
  ^"first=" First(uint8) ' '             # The rest must starts with "first=" characters, then take the rest until ' ' as uint8
                                         # under the name of First
  ^"format=" Format(string) ' '          # Take format id
  ^"responseTime" Duration(string) ' '   # Take mandatory response time
  ?Hidden (^"hidden=" Value(uint8) ' ')  # Optionally look for "hidden=\d+"
  ^"user_agent=\"" UserAgent(string) '"' # User agent data
  ^"country=" Country(string) ?? ' ';    # Take data as country to the rest or right to the first space character
```

And what would like to have from it:
* Code must be easy for comprehension and manual extension
* There should be as least dependencies as possible..
* Error messages should be helpful, i.e. mismatch cases must be easy to spot via error messages.
* Extracted data must be accessible via names (using struct fields)
* Unneccessary allocations should be avoided. For instance, when scanning logs we use []byte buffer as a temporary storage. Usually these fields only needed within the lifetime of current line, so extracted substrings better be []byte themselves
  
And this is example of generated code produced by closed source version of the utility. It is not ready to be released yet because of crappy code. We don't need crappy code in open source :)
```go
package main

import (
	"bytes"
	"fmt"
	"strconv"
	"unsafe"
)

var const0 = []byte(" FETCH_EVENTS ")
var const1 = []byte("first=")
var const2 = []byte("format=")
var const3 = []byte("responseTime")
var const4 = []byte("hidden=")
var const5 = []byte("user_agent=\"")
var const6 = []byte("country=")

type line struct {
	rest     []byte
	Time     []byte
	First    uint8
	Format   []byte
	Duration []byte
	Hidden   struct {
		Valid bool
		Value uint8
	}
	UserAgent []byte
	Country   []byte
}

func (p *line) Parse(line []byte) (bool, error) {
	p.rest = line
	var pos int
	var tmp []byte
	var restHidden []byte
	if pos = bytes.IndexByte(p.rest, ' '); pos < 0 {
		return false, nil
	}
	p.rest = p.rest[pos+1:]
	if pos = bytes.IndexByte(p.rest, ']'); pos < 0 {
		return false, nil
	}
	p.Time = p.rest[:pos]
	p.rest = p.rest[pos+1:]
	if !bytes.HasPrefix(p.rest, const0) {
		return false, nil
	}
	p.rest = p.rest[len(const0):]
	if !bytes.HasPrefix(p.rest, const1) {
		return false, nil
	}
	p.rest = p.rest[len(const1):]
	if pos = bytes.IndexByte(p.rest, ' '); pos < 0 {
		return false, nil
	}
	tmp = p.rest[:pos]
	if value, err := strconv.ParseUint(*(*string)(unsafe.Pointer(&tmp)), 10, 8); err == nil {
		p.First = uint8(value)
	} else {
		return false, fmt.Errorf("Cannot convert `%s` into uint8 (field First)", string(p.rest[:pos]))
	}
	p.rest = p.rest[pos+1:]
	if !bytes.HasPrefix(p.rest, const2) {
		return false, nil
	}
	p.rest = p.rest[len(const2):]
	if pos = bytes.IndexByte(p.rest, ' '); pos < 0 {
		return false, nil
	}
	p.Format = p.rest[:pos]
	p.rest = p.rest[pos+1:]
	if !bytes.HasPrefix(p.rest, const3) {
		return false, nil
	}
	p.rest = p.rest[len(const3):]
	if pos = bytes.IndexByte(p.rest, ' '); pos < 0 {
		return false, nil
	}
	p.Duration = p.rest[:pos]
	p.rest = p.rest[pos+1:]
	restHidden = p.rest
	if !bytes.HasPrefix(p.rest, const4) {
		p.Hidden.Valid = false
		p.rest = restHidden
		goto outHidden
	}
	p.rest = p.rest[len(const4):]
	if pos = bytes.IndexByte(p.rest, ' '); pos < 0 {
		p.Hidden.Valid = false
		p.rest = restHidden
		goto outHidden
	}
	tmp = p.rest[:pos]
	if value, err := strconv.ParseUint(*(*string)(unsafe.Pointer(&tmp)), 10, 8); err == nil {
		p.Hidden.Value = uint8(value)
	} else {
		return false, fmt.Errorf("Cannot convert `%s` into uint8 (field Hidden.Value)", string(p.rest[:pos]))
	}
	p.rest = p.rest[pos+1:]
	p.Hidden.Valid = true
outHidden:
	if !bytes.HasPrefix(p.rest, const5) {
		return false, nil
	}
	p.rest = p.rest[len(const5):]
	if pos = bytes.IndexByte(p.rest, '"'); pos < 0 {
		return false, nil
	}
	p.UserAgent = p.rest[:pos]
	p.rest = p.rest[pos+1:]
	if !bytes.HasPrefix(p.rest, const6) {
		return false, nil
	}
	p.rest = p.rest[len(const6):]
	if pos = bytes.IndexByte(p.rest, ' '); pos < 0 {
		p.Country = p.rest
		p.rest = p.rest[len(p.rest):]
	} else {
		p.Country = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	}
	return true, nil
}

func (p *line) GetHiddenValue() (res uint8) {
	if p.Hidden.Valid {
		res = p.Hidden.Value
	}
	return
}
```
