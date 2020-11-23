package psql

import (
	"github.com/stretchr/testify/require"
	"github.com/warmans/thrl6p/server/pkg/filter"
	"testing"
)

func TestJSONCondition(t *testing.T) {

	validPaths := []string{"foo", "bar", "baz", "@asset"}

	tests := map[string]struct {
		expectSql    string
		expectParams []interface{}
		expectError  bool
	}{
		`foo.bar = 1`: {
			expectSql:    `("fname"::JSON #>> '{foo,bar}')::FLOAT = $1`,
			expectParams: []interface{}{int64(1)},
			expectError:  false,
		},
		`foo = "bar"`: {
			expectSql:    `("fname"::JSON #>> '{foo}')::TEXT = $1`,
			expectParams: []interface{}{"bar"},
			expectError:  false,
		},
		`foo.bar = 2.5`: {
			expectSql:    `("fname"::JSON #>> '{foo,bar}')::FLOAT = $1`,
			expectParams: []interface{}{float64(2.5)},
			expectError:  false,
		},
		`bar.baz = null`: {
			expectSql:    `("fname"::JSON #>> '{bar,baz}') IS NULL`,
			expectParams: []interface{}{},
			expectError:  false,
		},
		`foo = "bar" and (bar > 1 or baz != 2)`: {
			expectSql:    `("fname"::JSON #>> '{foo}')::TEXT = $1 AND (("fname"::JSON #>> '{bar}')::FLOAT > $2 OR ("fname"::JSON #>> '{baz}')::FLOAT != $3)`,
			expectParams: []interface{}{"bar", int64(1), int64(2)},
			expectError:  false,
		},
		`foo.@asset = "speakerSimulator"`: {
			expectSql:    `("fname"::JSON #>> '{foo,@asset}')::TEXT = $1`,
			expectParams: []interface{}{"speakerSimulator"},
			expectError:  false,
		},
	}
	for cond, test := range tests {
		t.Run(cond, func(t *testing.T) {
			sql, params, err := JSONCondition(filter.MustParse(cond), "fname", validPaths)
			if test.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.EqualValues(t, test.expectSql, sql)
			require.EqualValues(t, test.expectParams, params)
		})
	}
}
