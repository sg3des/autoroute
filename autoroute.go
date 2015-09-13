//Package Autoroute is http interlayer implements routing based on existing methods of your controllers
//Example:
//
//    autoroute.Controllers = map[string]interface{}{"Name":&yourcontroller.structname{},et.c.}
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
	"net/http"
	"reflect"
	"strings"
)

var (
	//list of your controllers
	Controllers map[string]interface{}

	//you may change predefined values
	Default = Def{"Main", "Index"}
)

type Def struct {
	Controller string
	Method     string
}

//main function, parse url path and Call method
func Route(w http.ResponseWriter, r *http.Request) {
	urlpath := strings.Split(strings.Trim(r.URL.String(), "/"), "/")

	if len(urlpath) == 1 {
		urlpath = []string{Default.Controller, urlpath[0]}
	}

	if len(urlpath[1]) == 0 {
		urlpath[1] = Default.Method
	}

	for i, u := range urlpath {
		urlpath[i] = strings.ToUpper(string(u[0])) + strings.ToLower(u[1:])
	}

	Call(w, r, urlpath)
}

//Call to method by path or not found
func Call(w http.ResponseWriter, r *http.Request, urlpath []string) {
	for p, m := range Controllers {
		if p == urlpath[0] {
			//check if controller.method can be call
			if reflect.ValueOf(m).MethodByName(urlpath[1]).CanInterface() {
				//add request and responce writer to controller structur
				reflect.ValueOf(m).Elem().FieldByName("W").Set(reflect.ValueOf(w))
				reflect.ValueOf(m).Elem().FieldByName("R").Set(reflect.ValueOf(r))
				//call
				reflect.ValueOf(m).MethodByName(urlpath[1]).Call([]reflect.Value{})
			} else {
				error404(w)
				return
			}
			return
		}
	}
	error404(w)
}

//error not found
func error404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Error 404 page not found"))
}
