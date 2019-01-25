package peek

import (
	"errors"
	"log"
	"reflect"
)

func Peek(path string, value interface{}) (interface{}, error) {
	v := reflect.ValueOf(value)
	return peek(path, v)
}

func peek(path string, v reflect.Value) (interface{}, error) {
	switch v.Kind() {
	case reflect.Ptr:
		log.Println("is pointer!", v.Elem())
		return peek(path, v.Elem())
	case reflect.Struct:
		f := v.FieldByName(path)

		return f.Interface(), nil
	}

	return nil, errors.New("not yet implemeted")

}
