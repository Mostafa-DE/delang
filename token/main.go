package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // For any unknown token/character
	EOFILE  = "EOFILE"  // Tells the parser that it should stop

	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"   // 1343456
	FLOAT  = "FLOAT" // 1.234
	STRING = "STRING"

	// Operators
	ASSIGN        = "="
	PLUS          = "+"
	MINUS         = "-"
	EXCLAMATION   = "!"
	ASTERISK      = "*"
	SLASH         = "/"
	LESSTHAN      = "<"
	GREATERTHAN   = ">"
	LESSTHANEQ    = "<="
	GREATERTHANEQ = ">="
	EQUAL         = "=="
	NOTEQUAL      = "!="
	MOD           = "%"

	// Logical Operators
	AND = "and"
	OR  = "or"

	// Delimiters
	COMMA       = ","
	SEMICOLON   = ";"
	LEFTPAR     = "("
	RIGHTPAR    = ")"
	LEFTBRAC    = "{"
	RIGHTBRAC   = "}"
	COLON       = ":"
	LEFTSQPRAC  = "["
	RIGHTSQPRAC = "]"
	UNDERSCORE  = "_"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	CONST    = "CONST"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	DURING   = "DURING"
	BREAK    = "BREAK"
	SKIP     = "SKIP"
	FOR      = "FOR"
	IN       = "IN"
)

var keywords = map[string]TokenType{
	"fun":    FUNCTION,
	"let":    LET,
	"const":  CONST,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"during": DURING,
	"break":  BREAK,
	"skip":   SKIP,
	"for":    FOR,
	"in":     IN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
