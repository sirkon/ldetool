# Tool rules that can be used for parsing

## Rule and capture names
Rule and capture names must be public and goish

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
|``_'c'[M:N]``|Look for the character *c* in the M..N-1 characters the rest<br>and pass it.<br>Signal error when it was not found|``_'1'[1:2]("a12") → "2"``<br>``_'1'[1:2]("12") → error``<br>``_'1'[0:2]("123c") → "23c"``|
|``_'c'[M]``|Only check for character *c* at the M+1-th place of the rest||
|``_?'c'[M:N]``|Just like with other ?-powered searches:<br>Pass or ignore rather than pass or error|``_'1'[1:2]("a12") → "2"``<br>``_'1'[1:2]("12") → "12"``<br>``_'1'[0:2]("123c") → "23c"``|
|``_"t"``|Look for the text *t* in the rest and pass it.<br>Signal error when it was not found|``_"ab"("1ab2") → "2"``|
|``_?"t"``|Works exactly like ``_"t"`` if the text *t* was fou-------|nd.<br>Do nothing otherwise|``_?"ab"("1ab2") → "2"``<br>``_?"ab"("1b2") → "1b2"``|
|``_"t"[:N]``|Here and for the rest of lookups: same as for character lookup|``_"ab"[:3]("1ab2") → "2"``<br>``_"ab"[:2]("1ab2") → error``|
|``_?"t"[:N]``| |``_?"ab"[:3]("1ab2") → "2"``<br>``_?"ab"[:2]("1ab2") → "1ab2"``|
|``_"t"[M:N]``| |``_"ab"[1:3]("1ab2") → "2"``<br>``_?"ab"[2:4]("1ab2") → error``|
|``_"t"[M]``|Symbols of the rest from (M+1)-th position must starts with *t* |``_"ab"[1:3]("1ab2") → "2"``<br>``_?"ab"[2:4]("1ab2") → error``|
|``_?"t"[M:N]``| |``_?"ab"[1:3]("1ab2") → "2"``<br>``_?"ab"[2:4]("1ab2") → "1ab2"``|

## Capturing rules
1. There's currently only Go code generator, so I will base the further description on Go-specific syntax. Capturing rules are all named and these names are mapped into Go struct field names.
2. Capturing can be limited and unlimited. Limited capture takes all symbols right to the start of some boundary (text or character) or all symbols to the rest. Captured value can be stored in one of the following type. Capturing as numeric type can cause number parsing errors and these are always treated as "serious" ones.

    |int8|int16|int32|int64|uint8|uint16|uint32|uint64|float32|float64|string|
    |----|-----|-----|-----|-----|------|------|------|-------|-------|------|

3. As I mentioned, capture can be limited with char or text. The rules are absolutely the same as with character or string unconditional lookup, just replace ``_`` symbol with named capture description ``FieldName(type)``. There's a difference though in ``?....`` treatment. `?` for capturing will mean try to limit a capture area and if no boundary was found take everything to the rest.
    
So, the rules

|Syntax|Description|Example of rule application|Field value|
|:-----|-----------|---------------------------|---------------|
|``ID(τ)``|Put the rest into ID as τ|``ID(string)("abc")``|``ID="abc"``<br>``rest=""``|
|``ID(τ) 'c'``|Put the rest until *c* character into ID |``ID(τ)"t"``||as τ|``ID(int16)'@'("1234@kkk")``|``ID=1234``<br>``rest="kkk"``|
|``ID(τ)?'c'``|Put the rest until *c* character into ID as τ<br>If character was not found<br>put all symbols of the rest as τ|``ID(int16)?'@'("1234")``|``ID=1234``<br>``rest=""``|
|``ID(τ)'c'[:N]``|Passing and capturing rules are<br> exactly the same as in cases above|
|``ID(τ)?'c'[:N]``||
|``ID(τ)'c'[M:N]``||
|``ID(τ)?'c'[M:N]``||
|``ID(τ)"t"``||
|``ID(τ)?"t"``||
|``ID(τ)"t"[:N]``||
|``ID(τ)?"t"[:N]``||
|``ID(τ)"t"[M:N]``||
|``ID(τ)?"t"[M:N]``||
|``?ID(...)``|Optional named group. When its own set of rules failed it marks its ``ID.Valid`` field to false|``?ID(V(int8)' ')("1 b")``|``ID.Valid=true``<br>``ID.V=1``<br>``rest="b"``|

## Miscellaenous rules

|Syntax|Description|
|------|-----------|
|``$``|Checks if the rest is empty. Signal error otherwise|
|``!``|All mismatches after this symbol are treated as "serious" errors, i.e. ones producing error messages rather than just exit parsing|
