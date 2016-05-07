package controllers

import (
	"fmt"
	"net/http"

	"github.com/sg3des/autoroute"
)

//Main controller available without path
type Main autoroute.Controller

//Index page is available without path
func (c *Main) Index() []byte {
	//if arguments received
	if len(c.Args) > 0 {
		args := fmt.Sprintf("/Main/Index received arguments: %v\n", c.Args)
		return []byte(args)
	}

	return []byte("/Main/Index\n")
}

//Users controller is example of struct where standard http`s fields mix with json fields
type Users struct {
	W    http.ResponseWriter
	Name string `json:"name"`
	IP   string `json:"ip"`
}

//GetMixJson example method for json requests
func (c *Users) GetMixJson() {
	//so as in struct Users contains field W(http.ResponseWriter),this makes it possible write directly to it, insead returning an response
	fmt.Fprintln(c.W, "/Users/GetJson/ request data:", c.Name, c.IP)
}

//City controller displayed how embed your struct to standard autoroute controller structure to special field "Json"
//
//for this case need to use next entry for adding controller:
//	"City":  &controllers.City{Json: &controllers.CityJson{}},
type City autoroute.Controller

//CityJson is example your structure for controller City
type CityJson struct {
	Name    string
	Country string
}

func (c *City) GetFieldJson() []byte {
	//reveal json request
	req := c.Json.(*CityJson)

	return []byte(fmt.Sprintf("/City/GetJson/ request data: %s %s\n", req.Name, req.Country))
}
