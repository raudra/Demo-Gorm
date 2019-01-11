package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"log"
)


func NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
        handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf(
            "\n%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            name,
        )

        inner.ServeHTTP(w, r)
        log.Printf(
            "%s Response Time: %s",
            "Completed Request with ",
            time.Since(start),
        )
    })
}