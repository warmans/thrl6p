package psql

import (
	"bytes"
	"fmt"
	"github.com/warmans/thrl6p/server/pkg/filter"
	"github.com/warmans/thrl6p/server/pkg/util"
	"io"
	"strings"
)

func JSONCondition(f filter.Filter, field string, validPathSegments []string) (string, []interface{}, error) {
	if f == nil {
		return "", []interface{}{}, nil
	}
	buff := &bytes.Buffer{}
	q := &JsonQuery{
		sql:               buff,
		params:            make([]interface{}, 0),
		validPathSegments: validPathSegments,
		jsonField:         field,
	}
	if err := f.Accept(q); err != nil {
		return "", nil, err
	}
	return buff.String(), q.params, nil
}

type JsonQuery struct {
	sql               io.Writer
	params            []interface{}
	validPathSegments []string // this should be a schema instead of a list
	jsonField         string
}

func (j *JsonQuery) VisitCompFilter(f *filter.CompFilter) (filter.Visitor, error) {
	cond, err := j.jsonCondition(f.Field, f.Value)
	if err != nil {
		return nil, err
	}
	if f.Value.Type().Equal(filter.NullType) {
		if _, err := fmt.Fprint(j.sql, cond); err != nil {
			return nil, err
		}
		return j, nil
	}
	if _, err := fmt.Fprintf(j.sql, "%s %s $%d", cond, f.Op, len(j.params)+1); err != nil {
		return nil, err
	}
	j.params = append(j.params, f.Value.Value())
	return j, nil
}

func (j *JsonQuery) VisitBoolFilter(filter *filter.BoolFilter) (filter.Visitor, error) {
	needsLeftParen := filter.LHS.Precedence() < filter.Precedence()

	if needsLeftParen {
		if _, err := fmt.Fprint(j.sql, `(`); err != nil {
			return nil, err
		}
	}
	if err := filter.LHS.Accept(j); err != nil {
		return nil, err
	}
	if needsLeftParen {
		if _, err := fmt.Fprint(j.sql, `)`); err != nil {
			return nil, err
		}
	}

	if _, err := fmt.Fprintf(j.sql, " %s ", j.sqlOp(filter.Op)); err != nil {
		return nil, err
	}

	needsRightParen := filter.RHS.Precedence() < filter.Precedence()

	if needsRightParen {
		if _, err := fmt.Fprint(j.sql, `(`); err != nil {
			return nil, err
		}
	}
	if err := filter.RHS.Accept(j); err != nil {
		return nil, err
	}
	if needsRightParen {
		if _, err := fmt.Fprint(j.sql, `)`); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (j *JsonQuery) jsonCondition(path string, value filter.Value) (string, error) {

	jsonPath, err := j.jsonPath(path)
	if err != nil {
		return "", err
	}
	switch value.Type() {
	case filter.NullType:
		return fmt.Sprintf(`("%s"::JSON #>> %s) IS NULL`, j.jsonField, jsonPath), nil
	case filter.IntType, filter.FloatType:
		return fmt.Sprintf(`("%s"::JSON #>> %s)::FLOAT`, j.jsonField, jsonPath), nil
	case filter.StringType:
		return fmt.Sprintf(`("%s"::JSON #>> %s)::TEXT`, j.jsonField, jsonPath), nil
	case filter.BoolType:
		return fmt.Sprintf(`("%s"::JSON #>> %s)::BOOL`, j.jsonField, jsonPath), nil
	default:
		return "", fmt.Errorf("unknown value type: %s", value.Type().String())
	}
}

func (j *JsonQuery) jsonPath(path string) (string, error) {
	pathParts := strings.Split(path, ".")
	for _, v := range pathParts {
		if !util.InStrings(v, j.validPathSegments...) {
			return "", fmt.Errorf("%s is not a valid field path segment", v)
		}
	}
	return fmt.Sprintf("'{%s}'", strings.Join(pathParts, ",")), nil
}

func (j *JsonQuery) sqlOp(op filter.BoolOp) string {
	switch op {
	case filter.BoolOpAnd:
		return "AND"
	case filter.BoolOpOr:
		return "OR"
	}
	panic(fmt.Sprintf("unknown bool op: %s", op))
}
