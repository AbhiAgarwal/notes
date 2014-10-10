- http.Handle on the "/" pattern will act as a catch-all route, so we define that route last. 
- http.FileServer returns an http.Handler so we use http.Handle to map a pattern string to a handler
- r.FormValue("body") is very common to get input from the http.Request object that the http.HandlerFunc receives as an argument.
	- where r is r *http.Request.

