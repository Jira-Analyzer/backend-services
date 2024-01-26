package service

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	server "github.com/Jira-Analyzer/backend-services/internal/server/backend"
	"github.com/sirupsen/logrus"
)

type ProxyService struct {
	*httputil.ReverseProxy
}

func NewProxyService(conf *server.ServerConfig) *ProxyService {
	url, err := url.Parse("http://" + *conf.ConnectorHost)
	if err != nil {
		logrus.Fatal(err)
		return nil
	}
	return &ProxyService{
		httputil.NewSingleHostReverseProxy(url),
	}
}

func (proxy *ProxyService) ProxyRequestHandler(http.ResponseWriter, *http.Request) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}
