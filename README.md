# AutoRoute is http router for Go

Http routing based on your controllers and their methods, no need to write routes manually

	
	for url: "http://site.com/books/read", where the path is `/books/read`

	will be open controller and method: Books.Read()

	package controllers

	type Books struct {
		W http.ResponseWriter
		R *http.Request
	}

	func (c *Books) Read(){
		c.W.Write([]byte{"this is Books read page"})
	}



Default controller `Main` and method `index` mean that url path `/` calls to `Main.index()`, which incidentally is similar to `/main/index`

