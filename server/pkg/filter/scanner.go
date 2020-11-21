package filter

import (
	"fmt"
	"strings"
	"unicode"
)

type tag string

const (
	tagEOF tag = "EOF"

	tagLParen tag = "("
	tagRParen tag = ")"

	tagAnd tag = "AND"
	tagOr  tag = "OR"

	tagEq   tag = "="
	tagNeq  tag = "!="
	tagLike tag = "~="
	tagGt   tag = ">"
	tagGe   tag = ">="
	tagLe   tag = "<="
	tagLt   tag = "<"

	tagField  tag = "FIELD"
	tagInt    tag = "INT"
	tagFloat  tag = "FLOAT"
	tagBool   tag = "BOOL"
	tagString tag = "STRING"
	tagNull   tag = "NULL"
)

var keywords = map[string]tag{
	"and":   tagAnd,
	"or":    tagOr,
	"true":  tagBool,
	"false": tagBool,
	"null":  tagNull,
}

type token struct {
	tag    tag
	lexeme string
}

func Scan(str string) ([]token, error) {
	scanner := scanner{
		input:  []rune(str),
		pos:    0,
		offset: 0,
	}

	tokens := make([]token, 0)
	for {
		tok, err := scanner.Next()
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, tok)

		if tok.tag == tagEOF {
			break
		}
	}
	return tokens, nil
}

type scanner struct {
	input  []rune
	pos    int
	offset int
}

func (s *scanner) Next() (token, error) {
	s.skipWhitespace()

	if s.atEOF() {
		return s.emit(tagEOF), nil
	}

	switch r := s.nextRune(); r {
	case '(':
		return s.emit(tagLParen), nil
	case ')':
		return s.emit(tagRParen), nil
	case '=':
		return s.emit(tagEq), nil
	case '!':
		if s.matchNextRune('=') {
			return s.emit(tagNeq), nil
		}
		return s.error("expected '='")
	case '~':
		if s.matchNextRune('=') {
			return s.emit(tagLike), nil
		}
		return s.error("expected '='")
	case '>':
		if s.matchNextRune('=') {
			return s.emit(tagGe), nil
		}
		return s.emit(tagGt), nil
	case '<':
		if s.matchNextRune('=') {
			return s.emit(tagLe), nil
		}
		return s.emit(tagLt), nil
	case '"':
		return s.scanString()
	default:
		if isLetter(r) {
			field, err := s.scanField()
			if err != nil {
				return token{}, err
			}
			return field, nil
		}
		if isStartOfNumber(r) {
			return s.scanNumber(), nil
		}
		return s.error("unknown entity")
	}
}

func (s *scanner) nextRune() rune {
	r := s.input[s.pos]
	s.pos++
	return r
}

// matchNextRune will match the next rune of a multi-run tag e.g. >= !=
func (s *scanner) matchNextRune(r rune) bool {
	if s.atEOF() || s.peek() != r {
		return false
	}
	s.nextRune()
	return true
}

func (s *scanner) skipWhitespace() {
	for !s.atEOF() && unicode.IsSpace(s.peek()) {
		s.nextRune()
	}
	s.offset = s.pos
}

func (s *scanner) scanField() (token, error) {
	for !s.atEOF() && (isLetter(s.peek()) || isNumber(s.peek()) || s.peek() == '.') {
		r := s.nextRune()
		if r == '.' && (s.atEOF() || isLetter(s.peek())) {
			return s.error("invalid field")
		}
	}
	tok := s.emit(tagField)
	if tag, ok := keywords[tok.lexeme]; ok {
		tok.tag = tag
	}
	return tok, nil
}

func (s *scanner) scanNumber() token {
	hasDecimal := false
	for !s.atEOF() && (isNumber(s.peek()) || s.peek() == '.' && !hasDecimal) {
		r := s.nextRune()
		hasDecimal = hasDecimal || r == '.'
	}
	if hasDecimal {
		return s.emit(tagFloat)
	}
	return s.emit(tagInt)
}

func (s *scanner) scanString() (token, error) {
	for !s.matchNextRune('"') {
		if s.atEOF() {
			return s.error("unclosed double quote")
		}
		s.nextRune()
	}
	return trimTokenLexeme(s.emit(tagString), `""`), nil
}

func (s *scanner) atEOF() bool {
	return s.pos >= len(s.input)
}

func (s *scanner) peek() rune {
	return s.input[s.pos]
}

func (s *scanner) emit(tag tag) token {
	lexeme := string(s.input[s.offset:s.pos])
	s.offset = s.pos
	return token{tag: tag, lexeme: lexeme}
}

func (s *scanner) error(reason string) (token, error) {
	return token{}, fmt.Errorf("failed to scan string at position %d ('%s'): %s", s.pos, string(s.input[s.offset:s.pos]), reason)
}

func isStartOfNumber(r rune) bool {
	return isNumber(r) || r == '-'
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
}

func trimTokenLexeme(t token, trimSet string) token {
	t.lexeme = strings.Trim(t.lexeme, trimSet)
	return t
}
