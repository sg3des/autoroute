# AutoRoute is http router for Go

Http routing based on your controllers and their methods, no need to write routes manually.

AutoRoute impose two-level url path where first node is controller and second node is method, in most cases this is enough.

For url: "http://site.com/books/read", where the path is `/books/read` will be open controller and method: Books.Read()

example of controller:
	
	package controllers

	type Books autoroute.Controller

	func (c *Books) Read(){
		c.W.Write([]byte{"this is Books read page"})
	}

autoroute.Controller is struct contains: W - http.ResponseWriter, R - http.Request, Args - list of arguments, Json - for json requests

	type Controller struct {
		W http.ResponseWriter
		R *http.Request
		Args []string
		Json interface{}
	}


Default controller `Main` and method `Index` mean that url path `/` calls to `Main.Index()`, which incidentally is similar to `/main/index` and `/index`


AutoRoute can automate parse json requests to your structs: 

if request contains in field *"Content-Type"* of header - *"application/json"*, it's trying parse body of request to your controller struct by condition: if exists field with name *"Json"* - json parsed to it, else to the controller itself.

##USAGE

1) Write controller, example:

	package controllers

	import (***)

	type Main autoroute.Controller

	func (c *Main) Index() {
		c.W.Write([]byte("Main Index"))
	}

	func (c *Main) Echo() {
		c.W.Write([]byte("Main Echo"))
	}


2) Declare your controllers to AutoRoute:

	c := autoroute.NewControllers(map[string]interface{}{
		"Main":  &controllers.Main{},
	})

3) Start server with one route:

	c.ListenAndServe("/","127.0.0.1:8000")

**For more examples look into the appropriate directory**

##FEATURES

Default controller and method is Main.Index, you may replace them with own values `autoroute.Default = autoroute.Settings{"YourController", "YourDefaultMethod"}`


url path `/` call default controller and default method.


##LIMITATIONS

AutoRoute routing method does not provide a build true CRUD API, so as not distinguish http-methods of requests(GET,POST,etc...).
