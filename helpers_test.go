package benchmark

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
)

func GetExpectedResultStom() map[string]interface{} {
	return map[string]interface{}{
		"id":       1,
		"name":     "item_1",
		"number":   11,
		"created":  time.Unix(10000, 0),
		"updated":  mysql.NullTime{time.Unix(11000, 0), true},
		"discount": 111.0,
		"price":    1111.0,
		"reserved": sql.NullBool{true, true},
		"points":   nil,
		"rating":   sql.NullFloat64{1.0, true},
		"visible":  true,
	}
}

func AssertMapsEqual(t *testing.T, expected, actual map[string]interface{}) {
	if len(expected) != len(actual) {
		t.Fatalf("size of expected map %d\n%+v\ndoes not match size of generated map %d\n%+v",
			len(expected),
			expected,
			len(actual),
			actual)
	}

	for key, e := range expected {
		a, ok := actual[key]
		if !ok {
			t.Fatalf("could not find key '%s' in map:\n%v\nexpected map:\n%v", key, actual, expected)
		}
		if !reflect.DeepEqual(e, a) {
			t.Fatalf("expected value by key '%s' is %v, got %v", key, e, a)
		}

	}
}
