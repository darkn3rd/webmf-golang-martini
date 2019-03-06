package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	
  m.Get("/", func() string {
    return "Hello world!\n"
	})
	
	m.Get("/hello", func() string {
    return "Hello world!\n"
	})

	m.Get("/hello/:name", func(params martini.Params) string {
		return "Why Hello " + params["name"] + "!\n"
	})

  m.Run()
}