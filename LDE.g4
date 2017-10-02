grammar LDE;

rules
    : rules atomicRule
    | atomicRule
    | EOF
    ;

atomicRule
    : Identifier '=' baseAction ';';

baseAction
    : Stress baseAction
    | '(' baseAction ')' baseAction
    | '(' baseAction ')'
    | atomicAction baseAction
    | atomicAction
    ;

atomicAction
    : passTargetPrefix
    | mayBePassTargetPrefix
    | passChars
    | passUntil
    | mayPassUntil
    | takeUntil
    | takeUntilOrRest
    | takeUntilRest
    | optionalNamedArea
    | optionalArea
    | atEnd;

passTargetPrefix
    : '^' targetLit '[' IntLit ']'
    | '^' targetLit
    ;

mayBePassTargetPrefix
    : '?' '^' targetLit '[' IntLit ']'
    | '?' '^' targetLit
    ;

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
    | targetLit exact
    | targetLit jump
    | targetLit
    | '~' target;

targetLit
    : CharLit
    | StringLit;

bound
    : '[' IntLit ':' IntLit ']'
    ;

limit
    : '[' ':' IntLit ']'
    ;

jump
    : '[' IntLit ':' ']'
    ;

exact
    : '[' IntLit ']'
    ;

fieldType
    : Identifier;


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

Stress
    : '!'
    ;