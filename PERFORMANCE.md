# General thoughts
Provided method of data extraction is basically a split of known output length on steroids:
* The most common operation is looking for char/string in a string - this is what string splitting utilities like awk/gawk do when looking for the next field separator.
* There's nothing like in depth lookup like in regex matching. We just find first character/string in the rest of line. It is enough in a huge majority of cases (and it is rather something wrong with the input if this is not enough).

So, the conclusion is it would be basically enough to check performance on regular column-formatted files as it reflects common usage. The file can be generated with [utility](https://github.com/glossina/ldetool/blob/master/columngen.7z). Don't forget to get a package needed:
```bash
go get github.com/glossina/message
```
And here is the [benchmark code](https://github.com/glossina/ldetool/blob/master/benchmarker.7z). Following packages are needed:
```bash
go get github.com/glossina/message
go get github.com/youtube/vitess/go/cgzip
```
Compile both and make preparations:
```bash
go install columngen
go install main
time columngen | gzip -c > data.gz
time zcat data.gz | wc -l 
```
#### Now a raw performance comparison (the file is cached) on 1.3Gb of gzipped data.
```
$ time zcat data.gz | wc -l
100000000

real	0m19,936s
user	0m19,212s
sys	0m0,668s
```
1.3G is just a matter or of seconds on modern hardware, definitely less than 19s, thus it is decompressing performance that limits us.

#### Now output 1st and 4th column separated by `|` with gawk:
```
$ time zcat data.gz | gawk -F '|' '{ print $1 "|" $4 }' | wc -l
100000000

real	1m10,401s
user	1m28,048s
sys	0m1,476s
```
Not very fast, 70 seconds, about 3.5 times slower than plain file decompression.
CPU usage in this pipe

|utility|CPU usage %|
|-------|-----------|
|zcat|29%|
|gawk|99%|

#### Finally, `main` program doing the same (output 1st and 4th columns separated by `|`)
```
$ time zcat data.gz | main | wc -l
100000000

real	0m20,102s
user	0m31,376s
sys	0m1,464s
```
The bottleneck for Go version is data decompression.

|utility|CPU usage %|
|-------|-----------|
|zcat|95%|
|gawk|60%|

#### Will try sed
```
$ time zcat data.gz | sed -E 's/^([^|]*)\|[^|]*\|[^|]*\|[^|]*\|([^|]*)\|.*$/\1|\2/g' | wc -l
^C
real	5m40,856s
user	5m54,020s
sys	0m2,884s
```
I need to go, can't wait when it will complete.

|utility|CPU usage %|
|-------|-----------|
|zcat|5.3%|
|sed|99%|

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

	"github.com/glossina/message"
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
	r := regexp.MustCompile(`^([^|]*)\|[^|]*\|[^|]*\|[^|]*\|([^|]*)\|.*$`)
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
$ time zcat data.gz | regmain | wc -l
100000000

real	2m11,926s
user	2m27,044s
sys	0m3,444s
```

|utility|CPU usage %|
|-------|-----------|
|zcat|15.5%|
|regmain|97%|
