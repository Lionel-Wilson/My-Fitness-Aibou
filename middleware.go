package main

import (
	"fmt"
	"net/http"
)

/*
Below middleware function will will act on all https request before they reach
servemux(the router of our application)

Order will look like this:
secureHeaders → servemux → application handler

It’s important to know that when the last handler in the chain returns,
control is passed back up the chain in the reverse direction. So when our
code is being executed the flow of control actually looks like this:
secureHeaders → servemux → application handler → servemux → secureHeaders
*/
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")
		// Any code here will execute on the way down the chain.
		next.ServeHTTP(w, r)
		// Any code here will execute on the way back up the chain.
	})
}

/*
Flow:
logRequest ↔ secureHeaders ↔ servemux ↔ application handler

this middleware function was made into a function onto the application.
this is to allow it to have access to the handler dependencies such as the info logger.
*/
func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a deferred function (which will always be run in the event
		// of a panic as Go unwinds the stack).
		defer func() {
			// Use the builtin recover function to check if there has been a
			// panic or not. If there has...
			if err := recover(); err != nil {
				// Set a "Connection: close" header on the response.
				w.Header().Set("Connection", "close")
				// Call the app.serverError helper method to return a 500
				// Internal Server response.
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
