- http.Handle on the "/" pattern will act as a catch-all route, so we define that route last. 

- http.FileServer returns an http.Handler so we use http.Handle to map a pattern string to a handler

- r.FormValue("body") is very common to get input from the http.Request object that the http.HandlerFunc receives as an argument.
	- where r is r *http.Request.
	- and rw is a http.ResponseWriter

- Why is http.ResponseWriter not a pointer, and http.Request is?
	- ResponseWriter is an interface, and Request is a struct.
	- r is a concrete struct.
	- Reference: http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
		- An interface is two things: it is a set of methods, but it is also a type
		- Instead of designing our abstractions in terms of what kinds of data our type holds, we design our abstractions in terms of what actions our types can execute. This is core to Go's type system.

- http.ServeMux can go pretty far for simple applications, but if you need more power in how you parse URL endpoints and route them to the proper handler, you may need to pull in a third party routing framework.
	- github.com/gorilla/mux
		- It has an interface that is familiar for http.ServeMux users, yet has a ton of extra features built around the idea of finding, the right http.Handler for the given URL path.

- interface{} (via http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)
	- The interface{} type, the empty interface, is the source of much confusion.
	- The interface{} type is the interface that has no methods. Since there is no implements keyword, all types implement at least zero methods, and satisfying an interface is done automatically, all types satisfy the empty interface.
	- That means that if you write a function that takes an interface{} value as a parameter, you can supply that function with any value.

So, this function:

```
func DoSomething(v interface{}) {
   // ...
}
```

will accept any parameter whatsoever.

- An interface value is constructed of two words of data;
	- one word is used to point to a method table for the value’s underlying type,
	- the other word is used to point to the actual data being held by that value.
	- interface value is two words wide and it contains a pointer to the underlying data, that’s typically enough to avoid common pitfalls.

This is how you are able to use a []string to pass an interface.

```
func PrintAll(vals []interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}

func main() {
    names := []string{"stanley", "david", "oscar"}
    vals := make([]interface{}, len(names))
    for i, v := range names {
        vals[i] = v
    }
    PrintAll(vals)
}
```

- That may sound cryptic, but it makes sense when you remember the following: everything in Go is passed by value. Every time you call a function, the data you’re passing into it is copied.
	- method receivers are passed into the function by value just like any other parameter.

- Twitter’s API was originally written in Ruby, and the default format for Ruby is not the same as the default format for Go

- Using `github.com/gorilla/mux` there is `mux.Vars(r)["id"]` to view id.

- reflect.TypeOf is how to find types. You've to import reflect