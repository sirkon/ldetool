# General
Provided method of data extraction is basically an `awk` on steroids:
* The most common operation is looking for char/string in a string - this is what string splitting utilities like awk/gawk do when looking for the next field separator.
* There's nothing like in depth lookup like in regex matching. We just find first character/string in the rest of line. It is enough in a huge majority of cases (and it is rather something wrong with the input if this is not enough).

So, the conclusion is it would be basically enough to check performance on regular column-formatted files as it reflects common usage. The file can be generated with [utility](https://github.com/sirkon/ldetool/blob/master/columngen.7z). Don't forget to get a package needed:
```bash
go get github.com/sirkon/message
```
And here is the [benchmark code](https://github.com/sirkon/ldetool/blob/master/benchmarker.7z). Following packages are needed:
```bash
go get github.com/sirkon/message
go get github.com/youtube/vitess/go/cgzip
```
Compile both and make preparations:
```bash
go install columngen
go install main
time columngen 100000000 > data
wc -l data
```
Now, `data` file is cached and we can test performance. We will extract 1st and 4th columns and output them separated again with | in the stdout.

#### Generated utility
```
$ time ./main < data | wc -l

real	0m10.219s
user	0m10.704s
sys	0m0.752s
```
#### Gawk
```
$ time  gawk -F '|' '{ print $1 "|" $4 }' data | wc -l
100000000

real	1m2.773s
user	1m3.208s
sys	0m1.520s
```
#### Mawk, should be quite fast
```
$ time  mawk -F '|' '{ print $1 "|" $4 }' data | wc -l
100000000

real	0m20.785s
user	0m20.992s
sys	0m1.008s
```
It is indeed, only two times slower. Still, it doesn't signal errors, it only can work upon columned files.
#### sed
```
$ time sed -E 's/^(.*?)\|.*?\|.*?\|.*?\|(.*?)\|.*$/\1|\2/g' data | wc -l
100000000

real	7m44.861s
user	7m47.444s
sys	0m5.600s
```
OMG, that was SLOW
#### Go with regular expression with group capture 
Program
```go
package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"regexp"

	"github.com/sirkon/message"
	"github.com/youtube/vitess/go/cgzip"
)

func main() {
	var zreader io.Reader
	if len(os.Args) != 2 {
		zreader = os.Stdin
	} else {
		rawreader, err := os.Open(os.Args[1])
		if err != nil {
			message.Critical(err)
		}
		defer rawreader.Close()
		zreader, err = cgzip.NewReaderBuffer(rawreader, 512*1024)
		if err != nil {
			message.Critical(err)
		}
	}
	reader := bufio.NewReaderSize(zreader, 128*1024)
	scanner := bufio.NewScanner(reader)

	dest := bufio.NewWriter(os.Stdout)
	defer dest.Flush()
	buf := &bytes.Buffer{}
	r := regexp.MustCompile(`^(.*?)\|.*?\|.*?\|.*?\|(.*?)\|.*$`)
	for scanner.Scan() {
		data := r.FindSubmatch(scanner.Bytes())
		if len(data) == 3 {
			buf.Reset()
			buf.Write(data[1])
			buf.WriteByte('|')
			buf.Write(data[2])
			buf.WriteByte('\n')
			dest.Write(buf.Bytes())
		}
	}
	if scanner.Err() != nil {
		message.Critical(scanner.Err())
	}
}
```
Launching
```bash
$ time ./goregex < data | wc -l

100000000

real	2m4.064s
user	2m6.928s
sys	0m5.616s
```
Harder to use than LDE generated code, slower, harder to reason when something goes wrong. It was a bit easier to write LDE 
rule than a regexp and significantly easier to access extracted data
```perl
parser =
    Name(string) '|'     # Take a text until | into Name as a string ([]byte, actually), then pass |
    _ '|'                # We are at the start of column 2, find | and pass it
    _ '|'                # find | and pass it to go at column 4
    _ '|'                # find | and pass again
    Count(string) '|';   # take the content of 5th column right to the | and exit
``` 
vs
```
^(.*?)\|.*?\|.*?\|.*?\|(^.*?)\|.*$
```
