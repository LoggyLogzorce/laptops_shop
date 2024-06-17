package main

import (
	"laptopsShop/api"
	"laptopsShop/db"
	"laptopsShop/engine"
	_ "laptopsShop/entity"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
)

var hdl *api.Handler
var apiMap map[string]map[string]reflect.Value

func init() {
	log.Println("Initializing handler")
	hdl = &api.Handler{}
	log.Println("Initializing apiMap")
	apiMap = map[string]map[string]reflect.Value{
		"GET": {
			"user":  reflect.ValueOf(hdl).MethodByName("GetUser"),
			"users": reflect.ValueOf(hdl).MethodByName("GetAllUsers"),
		},
		"POST": {
			"auth": reflect.ValueOf(hdl).MethodByName("PostIndex"),
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

	log.Println("Connecting to database `laptops`")
	db.Connect()
	db.Migrate()

	log.Println("Listening port 8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// createHandler is a helper function that uses reflection to call the appropriate method
func createHandler(apiMap map[string]map[string]reflect.Value) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &engine.Context{
			Response: w,
			Request:  r,
		}

		methodMap, ok := apiMap[r.Method]
		if !ok {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		log.Println("go to the page: ", r.URL.Path)
		path := r.URL.Path[1:]
		pathArr := strings.Split(path, "/")
		if pathArr[0] != "api" {
			front(ctx, path)
			return
		}
		method, ok := methodMap[path[4:]]
		if !ok {
			http.Error(w, "Path not found", http.StatusNotFound)
			return
		}

		// Call the method with the appropriate arguments
		log.Println("method: ", method)
		method.Call([]reflect.Value{reflect.ValueOf(ctx)})
	}
}

func front(ctx *engine.Context, path string) {
	if path == "" {
		http.ServeFile(ctx.Response, ctx.Request, "./static/html/index.html")
		return
	}
	http.ServeFile(ctx.Response, ctx.Request, "./static/"+path)
}
