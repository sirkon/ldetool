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
#### Now a bit pointless raw performance comparison (the file is cached) on 1.3Gb of gzipped data - something like 5-10 seconds to read if not cached, will keep it in mind.
```
$ time zcat data.gz | wc -l
100000000

real	0m19,936s
user	0m19,212s
sys	0m0,668s
```
So, 25-30 second on non-cached data, which is basically the case.

#### Now output 1st and 4th column separated by `|` with gawk:
```
$ time zcat data.gz | gawk -F '|' '{ print $1 "|" $4 }' | wc -l
100000000

real	1m10,401s
user	1m28,048s
sys	0m1,476s
```
Not very fast, you see. 75-80s if not cached, about 3 times slower than plain file decompression.

#### Finally, `main` program doing the same (output 1st and 4th columns separated by `|`)
```
$ time zcat data.gz | main | wc -l
100000000

real	0m20,102s
user	0m31,376s
sys	0m1,464s
```
The bottleneck for Go version is data decompression.

#### Will try sed
```
$ time zcat data.gz | sed -E 's/^([^|]*)\|[^|]*\|[^|]*\|[^|]*\|([^|]*)\|.*$/\1|\2/g' | wc -l
^C
real	5m40,856s
user	5m54,020s
sys	0m2,884s
```
I need to go, can't wait when it will complete.
