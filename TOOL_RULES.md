# Typical operations and full set of rules

## Typical operations
1. Pass all data to the certain string or character including it – bounds are rarely needed per se and something is wrong otherwise. Applying
    ```perl
    _ "anchor"
    ```
    on `"1234anchor1234"` will be left `1234` in the rest 
2. Take all data up to the certain bounding string or character and then pass both data and bound.
    ```perl
    Field(int) ' '
    ```
    on `"1234 4321"` will put 1234 into field `Field` and `"4321"` will be left in the rest
3. Just take the rest. Easy as
    ```perl
    Rest(string)
    ```
4. Type conversion (text to number) might be needed on data retrieval. See #2 as an example.
5. Check if the rest starts with the certain string or character and pass it
    ```perl
    ^"prefix"
    ```
    Applying it on `"prefix1234"` will pass `prefix` and `1234` will be left in the rest. 
6. Just pass first N characters
    ```perl
    _[123:]
    ```
    cuts first 123 characters from the rest
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
8. Error text generation on mismatch. Obviously, this cannot be done too successful as error messages are rather kind of art than something than can be generated. Anyway, they turned to be surprisingly helpful on diagnostic, as they always return the rest of line that coulnd't be extracted.
9. Take right to the certain bounding character or a string or all the rest if it is not found
    ```perl
    Rule = Rest(string) ?'@'
    ```
10. etc

## Rule and capture names
Rule and capture names must be public and starts from capital letter (i.e. `Name`, not `name`)

## Passing rules
|Syntax|Description|Example of rule application<br>``rule(rest) → new rest``|
|:-----|-----------|--------------------------------------------------------|
|``_[N:]``|Pass first N characters of the rest.<br>Signal error if the rest is shorter than N|``_[2:]("12ab") → "ab"``|
|``^'c'``|Check if the rest starts with the given character *c* and pass it.<br>Signal error otherwise|``^'@'("@usr") → "usr"``|
|``?^`c'``|If the rest starts with the given character *c* then pass it|``?^'@'("@usr") → "usr"``<br>``?^'@'("usr") → "usr"``|
|``^"t"``|Check if the rest starts with the given text *t* and pass it.<br> Signal error otherwise|``^"ab"("ab12") → "12"``|
|``?^"t"``|If the rest starts with the given text *t* then pass it|``^"ab"("ab12") → "12"``<br>``^"ab"("a12") → "a12"``|
|``_'c'``|Look for the character *c* in the rest and pass it.<br>Signal error when it was not found|``_'1'("a12") → "2"``|
|``_?'c'``|Works exactly like ``_'c'`` if the character *c* was found.<br>Do nothing otherwise|``_'1'("a12") → "2"``<br>``_'1'("a2") → "a2"``|
|``_'c'[:N]``|Look for the character *c* in first N characters the rest and pass it.<br>Signal error when it was not found|``_'1'[:2]("a12") → "2"``<br>``_'1'[:2]("aa12") → error``<br>``_'1'[:3]("aa123c") → "23c"``|
|``_?'c'[:N]``|Look for the character *c* in first N characters the rest and pass it.<br>Ignore when text *t* was not found|``_'1'[:2]("a12") → "2"``<br>``_'1'[:2]("aa12") → "aa12"``<br>``_'1'[:3]("aa123c") → "23c"``|
|``_'c'[M:N]``|Look for the character *c* in the M..N-1 characters of the rest<br>and pass it.<br>Signal error when it was not found|``_'1'[1:2]("a12") → "2"``<br>``_'1'[1:2]("12") → error``<br>``_'1'[0:2]("123c") → "23c"``|
|``_'c'[M:]``|Look for the character *c* in the M, M+1, etc characters of the rest<br>and pass it.<br>Signal error when it was not found|``_'1'[1:]("a12") → "2"``<br>``_'1'[1:]("12") → error``<br>``_'1'[0:]("123c") → "23c"``|
|``_'c'[M]``|Only check for character *c* at the M+1-th place of the rest||
|``_?'c'[M:N]``|Just like with other ?-powered searches:<br>Pass or ignore rather than pass or error|``_'1'[1:2]("a12") → "2"``<br>``_'1'[1:2]("12") → "12"``<br>``_'1'[0:2]("123c") → "23c"``|
|``_?'c'[M:]``|Just like with other ?-powered searches:<br>Pass or ignore rather than pass or error|``_'1'[1:]("a12") → "2"``<br>``_'1'[1:]("12") → "12"``<br>``_'1'[0:]("123c") → "23c"``|
|``_"t"``|Look for the text *t* in the rest and pass it.<br>Signal error when it was not found|``_"ab"("1ab2") → "2"``|
|``_?"t"``|Works exactly like ``_"t"`` if the text *t* was fou-------|nd.<br>Do nothing otherwise|``_?"ab"("1ab2") → "2"``<br>``_?"ab"("1b2") → "1b2"``|
|``_"t"[:N]``|Here and for the rest of lookups: same as for character lookup|``_"ab"[:3]("1ab2") → "2"``<br>``_"ab"[:2]("1ab2") → error``|
|``_?"t"[:N]``| |``_?"ab"[:3]("1ab2") → "2"``<br>``_?"ab"[:2]("1ab2") → "1ab2"``|
|``_"t"[M:N]``| |``_"ab"[1:3]("1ab2") → "2"``<br>``_?"ab"[2:4]("1ab2") → error``|
|``_"t"[M:]``| |``_"ab"[1:]("1ab2") → "2"``<br>``_?"ab"[2:]("1ab2") → error``|
|``_"t"[M]``|Symbols of the rest from (M+1)-th position must starts with *t* |``_"ab"[1:3]("1ab2") → "2"``<br>``_?"ab"[2:4]("1ab2") → error``|
|``_?"t"[M:N]``| |``_?"ab"[1:3]("1ab2") → "2"``<br>``_?"ab"[2:4]("1ab2") → "1ab2"``|

#### Note
You can put `~` sign before a char or string you are looking for. This means "short" lookup: for loop will be used for
char lookup instead of `bytes.IndexByte`:
``_~'c'`` will be translated into
```go
pos = -1
for i, char := range p.rest {
	if char == 'c' {
		pos = i
		break
	}
}
``` 
This gives a speedup in case of short lookup. Something like ×4 for one character. `bytes.IndexByte` becomes nearly as
fast as loop version at 5 characters and faster at 6. There is no optimization for strings though, code for strings will
be the same.

## Capturing rules
1. There's currently only Go code generator, so I will base the further description on Go-specific syntax. Capturing rules are all named and these names are mapped into Go struct field names.
2. Capturing can be limited and unlimited. Limited capture takes all symbols right to the start of some boundary (text or character) or all symbols to the rest. Captured value can be stored in one of the following type. Capturing as numeric type can cause number parsing errors and these are always treated as "serious" ones.

    |int|int8|int16|int32|int64|uint|uint8|uint16|uint32|uint64|float32|float64|string|str|
    |---|----|-----|-----|-----|----|-----|------|------|------|-------|-------|------|---|
    
    Remember, type `string` is treated as `[]byte` by default for the sake of performance and switches to 
    `string` with flag `--go-string`. You may use `str` to have exactly `string`.
    
    There is support for hexadecimal and octal values extraction:
    
    |LDE type|hex|hex8|hex16|hex32|hex64|oct|oct8|oct16|oct32|oct64|
    |--------|---|----|-----|-----|-----|---|----|-----|-----|-----|
    |Go type|uint|uint8|uint16|uint32|uint64|uint|uint8|uint16|uint32|uint64|

    Also, there's support for "decimal" types in the following form:
    
    ```perl
    Rule = Data(dec4.3);
    ```
    
    `dec4.3` means decimal number with 4 digits where 3 of them are in fraction, e.g.
    number `1.123` fits into `dec4.3`, while `12.1` or `1.1111` don't. 
    So, the generic form is `decX.Y` where 1 ≤ X ≤ 38 and Y ≤ X.
    Go type behind these `decX.Y` depends on the X:
    
    |X range|Go type|
    |-------|-------|
    |1…9|int32|
    |10…18|int64|
    |19…38|`(uint64, uint64)`|
    
    `decX.Y` mirrors [`Decimal`](https://clickhouse.yandex/docs/en/data_types/decimal/) type in Clickhouse
    

3. As I mentioned, capture can be limited with char or text. The rules are absolutely the same as with character or string unconditional lookup, just replace ``_`` symbol with named capture description ``FieldName(type)``. There's a difference though in ``?....`` treatment. `?` for capturing will mean try to limit a capture area and if no boundary was found take everything to the rest.
    
So, the rules

|Syntax|Description|Example of rule application|Field value|
|:-----|-----------|---------------------------|---------------|
|``ID(τ)``|Put the rest into ID as τ|``ID(string)("abc")``|``ID="abc"``<br>``rest=""``|
|``ID(τ) 'c'``|Put the rest until *c* character into ID |``ID(τ)"t"``||as τ|``ID(int16)'@'("1234@kkk")``|``ID=1234``<br>``rest="kkk"``|
|``ID(τ)?'c'``|Put the rest until *c* character into ID as τ<br>If character was not found<br>put all symbols of the rest as τ|``ID(int16)?'@'("1234")``<br>``ID(int16)?'@'("1234@m")``|``ID=1234``<br>``rest=""``<br>and<br>``ID=1234``<br>``rest='m'``|
|``ID(τ)'c'[:N]``|Passing and capturing rules are<br> exactly the same as in cases above|
|``ID(τ)?'c'[:N]``||
|``ID(τ)'c'[M:N]``||
|``ID(τ)'c'[M:]``||
|``ID(τ)?'c'[M:N]``||
|``ID(τ)?'c'[M:]``||
|``ID(τ)"t"``||
|``ID(τ)?"t"``||
|``ID(τ)"t"[:N]``||
|``ID(τ)?"t"[:N]``||
|``ID(τ)"t"[M:N]``||
|``ID(τ)"t"[M:]``||
|``ID(τ)?"t"[M:N]``||
|``ID(τ)?"t"[M:]``||
|``?ID(...)``|Optional named group. When its own set of rules failed it marks its ``ID.Valid`` field to false|``?ID(V(int8)' ')("1 b")``|``ID.Valid=true``<br>``ID.V=1``<br>``rest="b"``|
|``ID[τ] 'c'``|This works exactly the same as regular ``ID(τ) 'c'`` with one excpetion: the `c` character will be consumed as well|``ID(string)'c'("1cb")``|``ID=1c``<br>``rest=b``|
|``ID[τ]?'c'``||
|``ID[τ]'c'[:N]``||
|``ID[τ]?'c'[:N]``||
|``ID[τ]'c'[M:N]``||
|``ID[τ]'c'[M:]``||
|``ID[τ]?'c'[M:N]``||
|``ID[τ]?'c'[M:]``||
|``ID[τ]"tau"``|This works exactly the same as regular ``ID(τ) "tau"`` with one exception: the "tau" string will be consumed as well|``ID(string)"t"("1ta")``|``ID=1t``<br>``rest=a``|
|``ID[τ]?"t"``||
|``ID[τ]"t"[:N]``||
|``ID[τ]?"t"[:N]``||
|``ID[τ]"t"[M:N]``||
|``ID[τ]"t"[M:]``||
|``ID[τ]?"t"[M:N]``||
|``ID[τ]?"t"[M:]``||

It is possible to use short lookup sign as well:
``ID(τ)~'c'``



## Miscellaenous rules

|Syntax|Description|
|------|-----------|
|``$``|Checks if the rest is empty. Signal error otherwise|
|``!``|All mismatches after this symbol are treated as "serious" errors, i.e. ones producing error messages rather than just exit parsing|
|``%<N``|Checks if less than N symbols left in the rest|
|``%N``|Checks if exactly N symbols left in the rest|
|``%>N``|Checks if more than N symbols left in the rest|
