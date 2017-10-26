package sw

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	// App			*Application
}

type Routes []Route

func NewRouter() *mux.Router {
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

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index2",
		"GET",
		"/gantony/CorkDevQuestions/1.0.0/",
		Index2,
	},

	Route{
		"AddQuestion",
		"POST",
		"/gantony/CorkDevQuestions/1.0.0/questions",
		AddQuestion,
	},

	Route{
		"ListQuestions",
		"GET",
		"/gantony/CorkDevQuestions/1.0.0/questions",
		ListQuestions,
	},

	Route{
		"Upvote",
		"POST",
		"/gantony/CorkDevQuestions/1.0.0/questions/{id}/votes",
		Upvote,
	},

}