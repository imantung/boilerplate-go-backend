package repokit

import "reflect"

func IsZero(v interface{}) bool {
	val := reflect.ValueOf(v)
	return val.IsValid() && val.IsZero()
}
