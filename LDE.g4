grammar LDE;

rules
    : rules atomicRule
    | atomicRule
    ;

atomicRule
    : Identifier '=' baseAction ';';

baseAction
    : '(' baseAction ')' baseAction
    | '(' baseAction ')'
    | '!' baseAction
    | atomicAction baseAction
    | atomicAction
    ;

atomicAction
    : passStringPrefix
    | passCharPrefix
    | mayPassStringPrefix
    | mayPassCharPrefix
    | passChars
    | passUntil
    | mayPassUntil
    | takeUntil
    | takeUntilOrRest
    | takeUntilRest
    | optionalNamedArea
    | optionalArea
    | atEnd;

passStringPrefix
    : '^' StringLit;

passCharPrefix
    : '^' CharLit;

mayPassStringPrefix
    : '?' '^' StringLit;

mayPassCharPrefix
    : '?' '^' CharLit;

passChars
    : '_' '[' IntLit ':' ']';

passUntil
    : '_' target;

mayPassUntil
    : '?' '_' target;

takeUntil
    : Identifier '(' fieldType ')' target;

takeUntilOrRest
    : Identifier '(' fieldType ')' '?' target;

takeUntilRest
    : Identifier '(' fieldType ')';

optionalNamedArea
    : '?' Identifier '(' baseAction ')';

optionalArea
    : '?' '(' baseAction ')';

atEnd
    : '$';
    
    
target
    : targetLit bound
    | targetLit limit
    | targetLit
    | '~' target;

targetLit
    : CharLit
    | StringLit;

bound
    : '[' IntLit ':' IntLit ']'
    ;

limit
    : '[' ':' IntLit ']';

fieldType
    : 'int8'
    | 'int16'
    | 'int32'
    | 'int64'
    | 'uint8'
    | 'uint16'
    | 'uint32'
    | 'uint64'
    | 'float32'
    | 'float64'
    | 'string'
    ;


Identifier
    : [a-zA-Z_] ([a-zA-Z0-9_]*)
    ;

IntLit
    : [0-9]+
    ;

fragment EscapedQuote : '\\"';
StringLit :   '"' ( EscapedQuote | ~('\n'|'\r'|'\t') )*? '"'
    ;

fragment EscapedApo : '\\\'';

CharLit
    : '\'' ( EscapedApo | ~('\n'|'\r'|'\t') )*? '\''
    ;

WS
    : [ \n\t\r] -> skip
    ;

LineComment
    : '#' ~[\r\n]* -> skip
    ;