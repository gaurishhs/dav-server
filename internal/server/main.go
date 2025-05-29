package server

import (
	"net/http"

	"github.com/gaurishhs/dav-server/internal/config"
	"github.com/gaurishhs/dav-server/internal/web"
	"github.com/gaurishhs/dav-server/internal/web/pages"
	"github.com/gaurishhs/gor"
	"github.com/rs/zerolog/log"
	"maragu.dev/gomponents"
	ghttp "maragu.dev/gomponents/http"
)

type DAVServer struct {
	// calDavBackend        caldav.Backend
	// cardDavBackend       carddav.Backend
	// userPrincipalBackend webdav.UserPrincipalBackend
	config     *config.Config
	HttpServer *http.Server
}

func setupRouter() *gor.Router {
	log.Debug().Msg("setting up router")
	router := gor.NewRouter()

	router.Get("/admin", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (gomponents.Node, error) {
		return pages.HomePage(true), nil
	}))
	router.Handle("/assets/", http.FileServer(http.FS(web.AssetFiles)))

	return router
}

func NewDAVServer() (*DAVServer, error) {
	config, err := config.LoadConfig("/Users/gaurish/projects/dav-server/config.toml")
	if err != nil {
		return nil, err
	}

	return &DAVServer{
		config: config,
		HttpServer: &http.Server{
			Addr:    config.Server.Addr,
			Handler: setupRouter(),
		},
	}, nil
}
