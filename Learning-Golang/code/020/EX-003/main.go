package main

import "fmt"

// There is a case where we don't know in advance what kind of value we will be receiving,
// EX: JSON Payload (HTTP endpoint)
// Unmarshalling that into a type is ok
// but it's coming from outside of boundary of your API

func main() {
	person := make(map[string]interface{}, 0)

	person["name"] = "Mario"
	person["age"] = 99

	for _, val := range person {
		switch x := val.(type) {
		case string:
			DoString(x)
		}
	}
}

// We are only applying string for now..

func DoString(v string) {
	fmt.Println("this is a string", v)
}

// We do not know what type this function will be receiving
// We gotta modify this switch-case manually all the time
// We will be implementing it with Generics for Go 1.19 in the future
func DoMany(v interface{}) {
	switch x := v.(type) {
	case string:
		fmt.Println("this is a string", x)
	}
}
