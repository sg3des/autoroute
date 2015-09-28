# AutoRoute is http router for Go

Http routing based on your controllers and their methods, no need to write routes manually

	for url: "http://site.com/books/read", where the path is `/books/read`

	will be open controller and method: Books.Read()

	package controllers

	type Books autoroute.Controller

	func (c *Books) Read(){
		c.W.Write([]byte{"this is Books read page"})
	}

autoroute.Controller is struct contain http.ResponseWriter, http.Request and list of arguments

	package autoroute

	{***}

	type Controller struct {
		W http.ResponseWriter
		R *http.Request
		A []string
	}


Default controller `Main` and method `index` mean that url path `/` calls to `Main.index()`, which incidentally is similar to `/main/index` and `/index`

AutoRoute impose two-level url path where first node is controller and second node is method, in most cases this is enough.

##USAGE

1) Write controller, example:

	package controllers

	import (
		"net/http"
		"***autoroute"
	)

	type Main autoroute.Controller

	func (c *Main) Index() {
		c.W.Write([]byte("Main Index"))
	}

	func (c *Main) Echo() {
		c.W.Write([]byte("Main Echo"))
	}


2) Declare your controllers to AutoRoute:

	autoroute.Controllers = map[string]interface{}{
		"Main": &controllers.Main{},
	}

3) Start server with one route:

	http.Handle("/", http.HandlerFunc(autoroute.Route))
	http.ListenAndServe("127.0.0.1:8000", nil)


##FEATURES

Default controller and method is Main.Index, you may replace them with own values `autoroute.Default = autoroute.Def{"MyController", "MyDefaultMethod"}`


url path `/` call default controller and default method.