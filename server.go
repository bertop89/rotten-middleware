package main

import (
	"net/http"
	"github.com/go-martini/martini"
	"io/ioutil"
)

func main() {
  m := martini.Classic()

  bo,up := read_files()

  m.Get("/box_office", func(writer http.ResponseWriter) string {
  	writer.Header().Set("Content-Type", "application/json")
    return bo
  })

  m.Get("/upcoming_movies", func(writer http.ResponseWriter) string {
    writer.Header().Set("Content-Type", "application/json")
    return up
  })

  m.Get("/top_dvds", func() string {
    return "Top DVDs"
  })

  m.Get("/upcoming_dvds", func() string {
    return "Upcoming DVDs"
  })

  

  m.Run()
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func read_files() (string,string) {
	bo, err := ioutil.ReadFile("responses/box_office.json")
  	check(err)
  	up, err := ioutil.ReadFile("responses/upcoming.json")
  	check(err)

  	return string(bo),string(up)
}