package peek

import (
	"errors"
	"reflect"
	"strings"
)

func Peek(path string, value interface{}) (interface{}, error) {
	v := reflect.ValueOf(value)
	p := strings.Split(path, ".")
	_, r, err := peek(p, v)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func peek(path []string, v reflect.Value) (reflect.Value, interface{}, error) {

	if len(path) == 1 {
		switch v.Kind() {
		case reflect.Ptr:
			return peek(path, v.Elem())
		case reflect.Interface:
			return peek(path, v.Elem())
		case reflect.Struct:
			f := v.FieldByName(path[0])
			return f, f.Interface(), nil
		default:
			return reflect.ValueOf(nil), nil, errors.New("not yet implemeted")
		}
	}

	switch v.Kind() {
	case reflect.Ptr:
		return peek(path, v.Elem())
	case reflect.Interface:
		return peek(path, v.Elem())
	case reflect.Struct:
		f := v.FieldByName(path[0])
		return peek(path[1:], f)
	default:
		return reflect.ValueOf(nil), nil, errors.New("not yet implemeted")
	}

}
