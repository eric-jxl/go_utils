package test

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetValue(p Person) {
	v := reflect.ValueOf(p)
	count := v.NumField()
	for i := 0; i < count; i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.String:
			fmt.Println(f.String())
		case reflect.Int:
			fmt.Println(f.Int())
		}
	}
}
