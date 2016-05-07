package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sg3des/autoroute"
	"github.com/sg3des/autoroute/example/controllers"
)

var (
	addr = ":7000"
)

func main() {
	fmt.Println("start")
	c := autoroute.NewControllers(map[string]interface{}{
		"Main":  &controllers.Main{},
		"Users": &controllers.Users{},
		"City":  &controllers.City{Json: &controllers.CityJson{}},
	})

	//now need start listening
	//is can be done with c.ListenAndServe or c.ListenAndServeTLS
	//or your method

	// http.Handle("/", http.HandlerFunc(c.Route))
	fmt.Printf("starting web server on addr '%s'...\n", addr)

	err := http.ListenAndServe(addr, http.HandlerFunc(c.Route))
	if err != nil {
		log.Fatalln(err)
	}

}
