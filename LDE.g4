grammar LDE;

rules
    : typeDeclaration* atomicRule* EOF
    ;

typeDeclaration
    : 'type' TypeName 'from' StringLit ';'
    | 'type' IdentifierMayStar ';'
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
    | checkTargetPrefix
    | passHeadingCharacters
    | mayBePassTargetPrefix
    | passChars
    | passUntil
    | mayPassUntil
    | goUntil
    | mayGoUntil
    | takeUntil
    | takeUntilIncluding
    | takeUntilOrRest
    | takeUntilIncludingOrRest
    | takeUntilRest
    | optionalNamedArea
    | optionalNamedSilentArea
    | optionalArea
    | restCheck
    | atEnd;

passHeadingCharacters
    : '*' CharLit;

passTargetPrefix
    : '^' targetLit '[' IntLit ']'
    | '^' targetLit
    ;

checkTargetPrefix
    : '@' targetLit '[' IntLit ']'
    | '@' targetLit
    ;

mayBePassTargetPrefix
    : '?' '^' targetLit '[' IntLit ']'
    | '?' '^' targetLit
    ;

passChars
    : '_' '[' IntLit ':' ']';

goUntil
    : '..' target;

mayGoUntil
    : '?' '..' target;

passUntil
    : '_' target;

mayPassUntil
    : '?' '_' target;

takeUntil
    : Identifier '(' fieldType ')' target;

takeUntilIncluding
    : Identifier '[' fieldType ']' target;

takeUntilOrRest
    : Identifier '(' fieldType ')' '?' target;

takeUntilIncludingOrRest
    : Identifier '[' fieldType ']' '?' target;

takeUntilRest
    : Identifier '(' fieldType ')';

optionalNamedArea
    : '?' Identifier '(' baseAction ')';

optionalNamedSilentArea
    : '??' Identifier '(' baseAction ')';

optionalArea
    : '?' '(' baseAction ')';

restCheck
    : '%' IntLit
    | '%' ComparisonOperator IntLit;

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
    : IdentifierWithFraction
    | Identifier
    | DollarIdentifier
    | TypeName
    ;

ComparisonOperator
    : [<>]
    ;

DollarIdentifier
    : '$' [a-zA-Z_] ([a-zA-Z0-9_]*)
    ;

Identifier
    : [a-zA-Z_] ([a-zA-Z0-9_]*)
    ;

TypeName
    : '*'* [a-zA-Z_] ([a-zA-Z0-9_]*) '.' [a-zA-Z_] ([a-zA-Z0-9_]*)
    ;

IdentifierMayStar
    :  '*'* [a-zA-Z_] ([a-zA-Z0-9_]*)
    ;

IdentifierWithFraction
    : [a-zA-Z_] ([a-zA-Z0-9_]*) '.' [0-9]+
    ;

IntLit
    : [0-9]+
    ;

fragment EscapedQuote : '\\"';
StringLit :   '"' ( EscapedQuote | ~('\n'|'\r'|'\t') ) ( EscapedQuote | ~('\n'|'\r'|'\t') )*? '"'
    ; // remember this is actually not empty string, as empty strings has no sense in this task

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
