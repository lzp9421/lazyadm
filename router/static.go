package router

import (
	"github.com/lazygo/lazygo"
	"github.com/lazygo/lazygo/utils"
	"lazyadm/asset"
	"net/http"
	"strings"
)

func StaticRouter(router *lazygo.Router) {
	router.GetHandle().HandleFunc("/frontend/", HandleStatic)
	router.GetHandle().HandleFunc("/static/", HandleStatic)

}

func HandleStatic (res http.ResponseWriter, req *http.Request) {
	path := strings.TrimPrefix(req.URL.Path, "/")
	data, err := asset.Asset(path)
	if err != nil {
		res.WriteHeader(404)
		data = []byte("404 page not found")
	}
	mimetype := utils.GetMimeType(path)
	res.Header().Set("Content-Type", mimetype)
	res.Write(data)
}

func init()  {
	registry["static"] = StaticRouter
}
