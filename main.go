package main

import (
	"log"
	"net/http"
	"reflect"
	"time"
	"vapingShop/api"
	"vapingShop/engine"
)

var hdl *api.Handler
var apiMap map[string]map[string]reflect.Value

func init() {
	hdl = &api.Handler{}
	apiMap = map[string]map[string]reflect.Value{
		"GET": {
			"/":      reflect.ValueOf(hdl).MethodByName("GetUser"),
			"/users": reflect.ValueOf(hdl).MethodByName("GetAllUsers"),
		},
		"POST": {
			"/": reflect.ValueOf(hdl).MethodByName("PostIndex"),
		},
	}

}

// Main function to set up the server and routes
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", createHandler(apiMap))

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Listening 8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// createHandler is a helper function that uses reflection to call the appropriate method
func createHandler(apiMap map[string]map[string]reflect.Value) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		methodMap, ok := apiMap[r.Method]
		if !ok {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		ctx := &engine.Context{
			Response: w,
			Request:  r,
		}

		path := r.URL.Path
		method, ok := methodMap[path]
		if !ok {
			http.Error(w, "Path not found", http.StatusNotFound)
			return
		}

		// Call the method with the appropriate arguments
		method.Call([]reflect.Value{reflect.ValueOf(ctx)})
	}
}
