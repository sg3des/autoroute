# AutoRoute is http router for Go

AutoRoute is just an **experiment** - another view on web-routing, but **it works**! 

Idea of AutoRoute based on web-routing using your own controllers and their methods, and does not require writing routes manually.

AutoRoute impose only two-level url path where the first node is controller and the second node is method, in most cases this is enough.

For url: "http://site.com/books/read", where the path is `/books/read` will be opened controller and method: Books.Read()

For url: "http://site.com/books/read/1", will be opened same Books.Read() and sent *"1"* how arguments in special field *Args*

AutoRoute uses structs for presentation controllers, you can use any structs, but should understand, that AutoRoute fills fields by their name: W - http.ResponseWriter, R - http.Request, Args - list of arguments, Json - for json requests. This structure already exists in AutoRoute package, and it can be used in most cases.

	type Controller struct {
		W http.ResponseWriter
		R *http.Request
		Args []string
		Json interface{}
	}

AutoRoute can automate parse json requests where header field *"Content-Type"* contains *"application/json"* - trying parse body to your controller struct by condition: if in struct exists field with name *"Json"* - json parsed to it, else to the controller itself.

example of controller for standard struct:

	package controllers

	type Books autoroute.Controller

	func (c *Books) Read() {
		c.W.Write([]byte{"this is Books Read page"})
	}

as aforesaid struct may be another, furthermore is possible to return an answer as []byte:

	package controllers

	type Books struct {
		ID string
		Name string
		Page int
	}

	func (c *Books) Read() []byte {
		return []byte{"this is Books Read page by own struct"}
	}


For access to `/` url path need use default controller `Main` and method `Index` which calls to `Main.Index()`, which incidentally is similar to `/main/index` and `/index`


AutoRoute is fully compatible with native net/http package.

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

Default controller and method is Main.Index, you may replace them with is own values `autoroute.Default = autoroute.Settings{"YourController", "YourDefaultMethod"}`


url path `/` calls default controller and default method.



##LIMITATIONS

AutoRoute routing method does not provide to build the true CRUD API, as so it does not distinguish http-methods of requests(GET,POST,etc...).


##BENCHMARK

AutoRoute does not pretend to be the fastest, and is not suitable for high-load projects. It may be used only for small services or websites.

benchmark simple requests:

	BenchmarkMainIndex-8	    5000	    318426 ns/op

benchmark json requests:

	BenchmarkCity-8     	    5000	    342184 ns/op
