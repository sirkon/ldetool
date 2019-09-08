Let we have two files __a.info__ and __b.info__:

##### `a.info`
```a.info a.info
1 2 3 4
12 13 14 15
1 1 1 1
```

##### `b.info`

```
1 2 3 4 5
5 4 3 2 1
```

You see, format in `b.info` clearly extends the one in `a.info` and we would
love to process both in a similar way (with the last column defaulted to some
value in case of `a.info`, most likely to `0`)

How can we achieve this? Tricky, but simple:

### Rules


```perl
A = First(int) ' ' Second(int) ' ' Third(int) ' ' Fourth(int);
B = First(int) ' ' Second(int) ' ' Third(int) ' ' Fourth(int) ` ` Fifth(int);
```

and compile it

```
ldetool --package main rule.lde
```

### Usage

`b.info` parses in a simple manner

```go
for bInfo.Next() {
    var b B
    if ok, err := b.Extract(bInfo.Bytes()); !ok {
        if err != nil {
            return …
        }
        return …
    }
    processInfo(b)
}
```

parsing `a.info`:

```go
for aInfo.Next() {
    var b B
    if ok, err := (*A)(unsafe.Pointer(&b)).Extract(aInfo.Bytes()); !ok {
        if err != nil {
            return …
        }
        return …
    }
    processInfo(b)
}
```

This will work as memory layout in A looks exactly like the memory layout of the first four fields
of B 