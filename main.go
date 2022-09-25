// package main

// import (
// 	"fmt"

// 	"rsc.io/quote"
// )

// func main() {
// 	fmt.Println(quote.Go())
// }

package main

import (
	"fmt"
	"net/url"
	"reflect"
)

type Item struct {
	Limit  int
	Skip   int
	Fields string
}

func main() {
	v := url.Values{}

	i := Item{1, 2, "foo"}

	fmt.Printf("%v\n", i)  // {1 2 foo}
	fmt.Printf("%+v\n", i) // {Limit:1 Skip:2 Fields:foo}
	fmt.Printf("%#v\n", i) // main.Item{Limit:1, Skip:2, Fields:"foo"}
	fmt.Println(i)         // {1 2 foo}

	ri := reflect.ValueOf(i)
	ri.FieldByNameFunc(func(name string) bool {
		v.Set(name, fmt.Sprintf("%v", ri.FieldByName(name).Interface()))
		return false
	})

	fmt.Println(v.Encode()) // Fields=foo&Limit=1&Skip=2
}
