package token

import "fmt"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
