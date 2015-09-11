# AutoRoute is http router for Go

Http routing based on your controllers and their methods, no need to manually write routes

	
	for url: "http://site.com/books/read" 

	will be open controoler and method: Books.Read()

	package controllers

	type Books struct {
		W http.ResponseWriter
		R *http.Request
	}

	func (c *Books) Read(){
		c.W.Write([]byte{"this is Books read page"})
	}



Default controller `Main` and method `index`, it means thar url path `/` call to `Main.index()`, which, incidentally, is similar `/main/index`

