package core

import (
	"github.com/gorilla/mux"
	"lazyadm/core/library"
	"net/http"
	"reflect"
	"unsafe"
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

func (r *Router) HandleFunc(path string, c interface{}) {
	r.router.HandleFunc(path, r.Handler(c))
}

func (r *Router) Handler(c interface{}) func(res http.ResponseWriter, req *http.Request) {
	v := reflect.ValueOf(c)
	t := reflect.Indirect(v).Type()
	h := func(res http.ResponseWriter, req *http.Request) {
		action := mux.Vars(req)["action"]
		c := reflect.New(t)
		pC := unsafe.Pointer(c.Pointer())
		base := (*Controller)(pC)
		base.Init(t.Name(), c.Interface().(IControllerRun), res, req)
		base.Run(action)
	}
	return h
}

func (r *Router) HandleStatic(asset func(string) ([]byte, error)) {
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

func (r *Router) GetHandle() *mux.Router {
	return r.router
}
