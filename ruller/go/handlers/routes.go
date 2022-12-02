package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(mux.CORSMethodMiddleware(router))
	// Handle all preflight request
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})
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
	fmt.Fprintf(w, "Suricata ruleset management service")
}

var routes = Routes{
	//root does nothing
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"RulesetGetAll",
		strings.ToUpper("Get"),
		"/ruleset",
		RulesetGetAll,
	},
	Route{
		"RulesetGetById",
		strings.ToUpper("Get"),
		"/ruleset/{id}",
		RulesetGetById,
	},
	Route{
		"RuleValidate",
		strings.ToUpper("Post"),
		"/rule/validate",
		RuleValidate,
	},
	Route{
		"RuleGetAll",
		strings.ToUpper("Get"),
		"/rule",
		RuleGetAll,
	},
	Route{
		"RuleGetById",
		strings.ToUpper("Get"),
		"/rule/{id}",
		RuleGetById,
	},
	Route{
		"RuleAdd",
		strings.ToUpper("Post"),
		"/rule",
		RuleAdd,
	},
	Route{
		"RuleUpdate",
		strings.ToUpper("Put"),
		"/rule/{id}",
		RuleUpdate,
	},
	Route{
		"RuleEnable",
		strings.ToUpper("Put"),
		"/rule/{id}/enable",
		RuleEnable,
	},
	Route{
		"RuleDisable",
		strings.ToUpper("Put"),
		"/rule/{id}/disable",
		RuleDisable,
	},

}

func writeDefaultHeaders(w *http.ResponseWriter, httpType string) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", httpType)
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}
