package types

import "reflect"

func Contains(arr interface{}, target interface{}) bool {
	arrValue := reflect.ValueOf(arr)
	if arrValue.Kind() != reflect.Slice {
		panic("not a slice")
	}

	targetValue := reflect.ValueOf(target)
	for i := 0; i < arrValue.Len(); i++ {
		if reflect.DeepEqual(arrValue.Index(i).Interface(), targetValue.Interface()) {
			return true
		}
	}

	return false
}
