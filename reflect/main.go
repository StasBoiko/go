package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type User struct {
	Name string
	Age  int64
}

type City struct {
	Name       string `json:"name"`
	Population int64  `json:"population"`
	GDP        int64  `json:"gdp"`
	Mayor      string `json:"mayor"`
}

func main() {
	var u User = User{"bob", 10}

	res, err := JSONEncode(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	c := City{"sf", 5000000, 567896, "mr jones"}
	res, err = JSONEncode(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

func JSONEncode(v interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString("{")
	rv := reflect.ValueOf(v)
	// TODO: check if v is a struct else return error
	if rv.Kind() != reflect.Struct {
		return buf.Bytes(), fmt.Errorf("must be struct")
	}

	rt := reflect.TypeOf(v)
	result := make([]string, 0)
	// TODO: iterate over v`s reflect value using NumField()
	// use type switch to create string result of "{field}" + ": " + "{value}"
	// start with just 2 types - reflect.String and reflect.Int64
	for i := 0; i < rv.NumField(); i++ {
		switch rv.Field(i).Kind() {
		case reflect.Int64:
			result = append(result, rt.Field(i).Name+`: `+strconv.FormatInt(rv.Field(i).Interface().(int64), 10))
		case reflect.String:
			result = append(result, rt.Field(i).Name+`: `+rv.Field(i).Interface().(string))
		default:
			return buf.Bytes(), fmt.Errorf("must be int64 or string")
		}
	}

	buf.WriteString(strings.Join(result, ", "))
	buf.WriteString("}")
	return buf.Bytes(), nil
}
