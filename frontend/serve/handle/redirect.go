package handle

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/short/frontend/serve/shortapi"
	"github.com/short-d/short/frontend/serve/ssr"
)

// Redirect redirects user to the corresponding long link with the given alias.
func Redirect(redirectPage ssr.RedirectPage, gRPC shortapi.GRPC, rootDir string) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		alias := params["alias"]
		openGraphTags, err := gRPC.GetOpenGraphTags(alias)
		if err != nil {
			serveIndex(rootDir, w, r)
			return
		}

		twitterTags, err := gRPC.GetTwitterTags(alias)
		if err != nil {
			serveIndex(rootDir, w, r)
			return
		}

		page, err := redirectPage.Render(openGraphTags, twitterTags)
		if err != nil {
			serveIndex(rootDir, w, r)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(page))
	}
}

func serveIndex(rootDir string, w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadFile(filepath.Join(rootDir, "index.html"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(buf)
}
