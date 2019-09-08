# Table of contents


# What if some parts may be missing

Example:

```
1905-04-15 name[Sergey Mihoparov] died[2010-07-12] education[school4] jailed[1951-11-02_1959-11-01] moderate-negative single experience[35]
1989-03-29 name[Nikita Zverev] education[school10+university6] marred1- experience[6]
2017-09-14 name[Aleksey Bezborodko] single experience[]
```

See at `moderate-negative` charactestics – imagine we don't care about it. In this case a rule that will consume both these lines may look like

```perl
Rule =
    Time(string) ' ' 
    ^"name[" Name(string) ']' ^' '
    ??Death(
        ^"died[" Date(string) ']' ^' '
    )
    ^"education[" Education(string) ']' ^' '
    ?Jailed(
        ^"jailed[" Intervals(string) ']' ^' '
    )
    ?( ^"moderate-negative ")
    Status(string) ' '
    ^"experience[" 
    ??Experience(
        Years(int) ']'
    );
```

Please take a look at optional areas. The area `?Jailed` will create a substructure with `Valid bool` field, which is set to true if it was matched successfully and set it to
false otherwise. Remember though, if you have some numeric taker and conversion failed the whole `Rule` failed. Use `??` to ignore conversion errors as well.

And so called ___anonymous area___ `?( ^"moderate-negative ")`. It is checking if the rest starts with `moderate-negative ` and passes it in case it is true. Or go to the next action otherwise.  


# Case when we need all data right to the some character or all the data if the character was not found

Lines:

```
Vladimir 37 186 91 15
Mikhail 35 184 76
```

And we don't need the last number `Vladimir' has. So, use this rule

```perl
Rule =
    Name(string) ' '
    Age(int) ' '
    Height(int) ' '
    Weight(int) ?' '
    ;
```

See the `?` before the last `' '` – it commands to take everything as `Weight` right to the end or the first `' '` 

# Case when bounds needed too

Let we have something like

```
2019-09-08T23:47:11.671194+03:00 id=some_id version=1.1.1 coords={"lat":11.11, "lon":22.22} locale=en-EN
```

and we `' '` won't work as a bound for coords as we have one in JSON itself. The solution is to use
`[]` taker, it will store bound itself:

```perl
Rule = 
    Time(string) " id="
    ID(string) " version="
    Version(string) " coords="
    Coords[string] '}' ^" locale="
    Locale(string)
    ; 
``` 

# Case of extended data format, where lines in B looks like lines from A with some more data appended

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