package tags

import (
	"reflect"

	"github.com/elgs/gojq"
)

// TagName is the name of the tag used to annotate gojq queries
const TagName = "tags"

// Parse parses a reflect.Type
func Parse(t reflect.Type, data string) (value reflect.Value, err error) {
	value = reflect.New(t)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagValue, ok := field.Tag.Lookup(TagName)
		if !ok {
			continue
		}

		parser, err := gojq.NewStringQuery(data)
		if err != nil {
			return value, err
		}

		result, err := parser.Query(tagValue)
		if err != nil {
			return value, err
		}

		switch result.(type) {
		case string:
			value.Elem().Field(i).SetString(result.(string))
		case float64:
			value.Elem().Field(i).SetInt(int64(result.(float64)))
		}
	}
	return value, err
}
