package core

import (
	"github.com/gorilla/mux"
	"net/http"
	"unsafe"
	"reflect"
	"lazyadm/core/library"
)

type Router struct {
	router *mux.Router
}

func NewRouter() (*Router, error) {
	router := mux.NewRouter()
	return &Router{
		router: router,
	}, nil
}

func (r *Router) HandleFunc(path string, f interface{}) {
	r.router.HandleFunc(path, r.Handler(f))
}

func (r *Router) Handler(n interface{}) func(res http.ResponseWriter, req *http.Request) {
	fn := reflect.ValueOf(n)
	h := func(res http.ResponseWriter, req *http.Request) {
		action := mux.Vars(req)["action"]
		params := []reflect.Value{reflect.ValueOf(res), reflect.ValueOf(req)}
		c := fn.Call(params)[0]
		pC := unsafe.Pointer(c.Pointer())
		base := (*Controller)(pC)
		base.Run(action)
	}
	return h
}

func (r *Router)HandleStatic(asset func(string) ([]byte, error)) {
	r.router.HandleFunc("/{path:static/.+}", func(res http.ResponseWriter, req *http.Request) {
		path := mux.Vars(req)["path"]
		data, err := asset(path)
		if err != nil {
			res.WriteHeader(404)
			data = []byte("404 page not found")
		}
		mimetype := library.GetMimeType(path)
		res.Header().Set("Content-Type", mimetype)
		res.Write(data)
	})
}

func (r *Router) GetHandle() (*mux.Router) {
	return r.router
}
