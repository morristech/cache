package redis

import "fmt"

func Example() {
	// init  driver
	c := Connect("127.0.0.1:6379", "", 0)

	// set
	c.Set("name", "cache value", 60)
	// get
	val := c.Get("name")
	// del
	c.Del("name")

	// get: "cache value"
	fmt.Print(val)
}
