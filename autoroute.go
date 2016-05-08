//Package autoroute is http interlayer implements routing based on existing methods of your controllers
//Example:
//
//    c := autoroute.NewControllers(map[string]interface{}{"Name":&yourcontroller.structname{},et.c.})
//
// this set map of your controllers
//
// after, all as usual:
//
//    http.Handle("/", http.HandlerFunc(autoroute.Route))
//    http.ListenAndServe("127.0.0.1:8000", nil)
//
// no need write route manually, all simple - if controller have appropriate method for requested url path, it will be open, otherwise - error 404.
// Example:
//    http://site.com/books/open - where path is `/books/open` for it will be need controller `Books` with method `Open`
//
// in detailes see exmaples on github
//
// Advantages of this approach is that there is no need to spend time on thinking and write routes.
//
package autoroute

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

var (
	//Default settings - you may change predefined values
	Default = Settings{"Main", "Index"}
)

//Settings struct
type Settings struct {
	Controller string
	Method     string
}

//Controller struct contain http.ResponseWriter, http.Request and list of arguments. Arguments is url path split by slash and crop two first element, ex: http://www.site.com/books/read/45 - where 45 will be added to Args
type Controller struct {
	W    http.ResponseWriter
	R    *http.Request
	Args []string
	Json interface{}
}

//C is Controllers type
type C struct {
	Controllers map[string]interface{}
}

//NewControllers bind controllers
func NewControllers(cons map[string]interface{}) *C {
	return &C{Controllers: cons}
}

//ListenAndServe start http server
func (c *C) ListenAndServe(path, addr string) error {
	http.Handle(path, http.HandlerFunc(c.Route))
	return http.ListenAndServe(addr, nil)
}

//ListenAndServeTLS start https server
func (c *C) ListenAndServeTLS(path, addr, certFile, keyFile string) error {
	http.Handle(path, http.HandlerFunc(c.Route))
	return http.ListenAndServeTLS(addr, certFile, keyFile, nil)
}

//Route is main function, parse url path and call obtained method
func (c *C) Route(w http.ResponseWriter, r *http.Request) {
	urlpath := strings.Split(strings.Trim(r.URL.String(), "/"), "/")

	if len(urlpath) <= 1 {
		urlpath = []string{Default.Controller, urlpath[0]}
	}

	if urlpath[1] == "" {
		urlpath[1] = Default.Method
	}

	//uppercase first letter
	for i := 0; i < 2; i++ {
		urlpath[i] = strings.Title(urlpath[i])
	}

	response, err := c.call(w, r, urlpath)
	if err != nil {
		error404(w)
		return
	}

	w.Write(response)
}

//Call to method by path returned []byte output or error if page not found
func (c *C) call(w http.ResponseWriter, r *http.Request, urlpath []string) ([]byte, error) {

	if icontoller, ok := c.Controllers[urlpath[0]]; ok {
		controller := reflect.ValueOf(icontoller)

		method := controller.MethodByName(urlpath[1])

		if !method.IsValid() {
			return []byte{}, errors.New("page not found")
		}

		data := controller.Elem()

		//fill standard controller
		if data.FieldByName("W").IsValid() {
			data.FieldByName("W").Set(reflect.ValueOf(w))
		}
		if data.FieldByName("R").IsValid() {
			data.FieldByName("R").Set(reflect.ValueOf(r))
		}
		if data.FieldByName("Args").IsValid() {
			data.FieldByName("Args").Set(reflect.ValueOf(Args(urlpath)))
		}

		//parse json request
		if strings.Contains(strings.ToLower(r.Header.Get("Content-Type")), "application/json") {
			var jsonStruct interface{}

			if data.FieldByName("Json").IsValid() {
				jsonStruct = data.FieldByName("Json").Interface()
			} else {
				jsonStruct = controller.Interface() //data.Interface()
			}

			if err := RequestJSON(r, &jsonStruct); err != nil {
				return []byte{}, err
			}
		}

		//call controller method
		values := method.Call([]reflect.Value{})
		if len(values) == 0 {
			return []byte{}, nil
		}
		return values[0].Bytes(), nil
	}
	return []byte{}, errors.New("page not found")
}

//RequestJSON function parse incoming request in json format
func RequestJSON(r *http.Request, i interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return json.Unmarshal(body, &i)
}

//Args split url path to arguments
func Args(urlpath []string) []string {
	if len(urlpath) > 2 {
		return urlpath[2:]
	}
	return []string{}
}

//error404 page not found
func error404(w http.ResponseWriter) {
	http.Error(w, "404 page not found", 404)
}
