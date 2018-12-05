package tags

import (
	"reflect"

	"github.com/elgs/gojq"
)

// TagName is the name of the tag used to annotate gojq queries
const TagName = "tags"

func queryJSON(data, query string) (interface{}, error) {
	parser, err := gojq.NewStringQuery(data)
	if err != nil {
		return nil, err
	}

	result, err := parser.Query(query)
	if err != nil {
		return nil, err
	}

	return result, err
}

func setField(fieldNr int, value interface{}, instance *reflect.Value) error {
	switch value.(type) {
	case string:
		instance.Elem().Field(fieldNr).SetString(value.(string))
	case float64:
		instance.Elem().Field(fieldNr).SetInt(int64(value.(float64)))
	}

	return nil
}

// Parse parses a reflect.Type
func Parse(t reflect.Type, data string) (value reflect.Value, err error) {
	value = reflect.New(t)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagValue, ok := field.Tag.Lookup(TagName)
		if !ok {
			continue
		}

		result, err := queryJSON(data, tagValue)
		if err != nil {
			return value, err
		}

		err = setField(i, result, &value)
		if err != nil {
			return value, err
		}
	}
	return value, err
}
