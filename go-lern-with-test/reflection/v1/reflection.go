package main

import (
	"reflect"
	"fmt"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

    for i:=0; i<val.NumField(); i++ {
        field := val.Field(i)
        fn(field.String())
    }
}




// 	package main

// import (
// 	"reflect"
// 	"fmt"
// )

// func walk(x interface{}, fn func(input string)) {
	
// 	val := reflect.ValueOf(x)
// 	field := val.Field(0)
// 	fn(field.String())
// 	// fn("12345")
// }

func main () {

	cases := []struct{
        Name string
        Input interface{}
        ExpectedCalls []string
    } {
		{
			"Struct with two string fields",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
    }

    for _, test := range cases {
        // t.Run(test.Name, func(t *testing.T) {
            var got []string
            walk(test.Input, func(input string) {
				
                got = append(got, input)
				
            })
			fmt.Println(got)
            // if !reflect.DeepEqual(got, test.ExpectedCalls) {
            //     t.Errorf("got %v, want %v", got, test.ExpectedCalls)
            // }
        // })
    }

	// cases := []struct{
    //     Name string
    //     Input interface{}
    //     ExpectedCalls []string
    // } {
    //     {
    //         "Struct with one string field",
    //         struct {
    //             Name string
    //         }{ "Chris"},
    //         []string{"Chris"},
    //     },
    // }

    // for _, test := range cases {
    //     // t.Run(test.Name, func(t *testing.T) {
    //         var got []string
    //         walk(test.Input, func(input string) {
	// 			fmt.Println(input)
    //             got = append(got, input)
    //         })

    //         if !reflect.DeepEqual(got, test.ExpectedCalls) {
    //             t.Errorf("got %v, want %v", got, test.ExpectedCalls)
    //         }
    //     // })
    // }

// 	x := []struct {
// 		Name          string
// 		Input         interface{}
// 		ExpectedCalls []string
// 	}{
// 		{
// 			"Struct with one string field",
// 			struct {
// 				Name string
// 			}{"Chris"},
// 			[]string{"Chris"},
// 		},
// 	}
// 	// fmt.Println(input)
// 	for _, test := range x {
// 		walk(x, func(input string) {
// 			// fmt.Println(x)
// 			fmt.Println(test)
// 			// got = append(got, input)
// 		})
// 	}

	// cases := []struct {
	// 	Name          string
	// 	Input         interface{}
	// 	ExpectedCalls []string
	// }{
	// 	{
	// 		"Struct with one string field",
	// 		struct {
	// 			Name string
	// 		}{"Chris"},
	// 		[]string{"Chris"},
	// 	},
	// }

	// for _, test := range cases {
		// t.Run(test.Name, func(t *testing.T) {
			// var got []string

			
			// walk(test.Input, func(input string) {
			// 	fmt.Println(test.Input)
			// 	// got = append(got, input)
			// })

			// if !reflect.DeepEqual(got, test.ExpectedCalls) {
			// 	t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			// }
		// })
	// }
}