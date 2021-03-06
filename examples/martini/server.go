package main

import (
	"github.com/RainingClouds/yaag/martiniyaag"
	"github.com/RainingClouds/yaag/yaag"
	"github.com/go-martini/martini"
)

func main() {
	yaag.Init(&yaag.Config{On: true, DocTitle: "Gorilla Mux", DocPath: "apidoc.html"})
	m := martini.Classic()
	m.Use(martiniyaag.Document)
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
