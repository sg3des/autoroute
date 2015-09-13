package main

import (
	"fmt"
	"net/http"

	"github.com/sg3des/autoroute"
	"github.com/sg3des/autoroute/example/controllers"
)

func main() {
	fmt.Println("start")
	autoroute.Controllers = map[string]interface{}{
		"Main": &controllers.Main{},
		"Urls": &controllers.Urls{},
	}

	http.Handle("/", http.HandlerFunc(autoroute.Route))
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	fmt.Println(err)
}
