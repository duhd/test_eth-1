package main

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func main() {
	router := routing.New()

	router.Get("/", func(c *routing.Context) error {
		fmt.Fprintf(c, "Hello, world!")
		return nil
	})

  fmt.Println("Start listening")
	panic(fasthttp.ListenAndServe(":8080", router.HandleRequest))
}
