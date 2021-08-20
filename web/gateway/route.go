package gateway

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wI2L/fizz"
	"github.com/zxzixuanwang/go-forum/web/gateway/api/v1/auth"
	"github.com/zxzixuanwang/go-forum/web/gateway/config"
	"github.com/zxzixuanwang/go-forum/web/pkg/logzap"
)

var log = logzap.GetLogger()

func LoadRoute() {
	log.Infof("读取路由,环境:%s", config.ConfigCollection.Services.Env)
	g := gin.New()
	gin.SetMode(config.ConfigCollection.Services.Env)
	f := fizz.NewFromEngine(g)
	v1Group := f.Group("/api", "api", "This is for api.")
	auth.RegisterRouter(v1Group)
	log.Fatal(http.ListenAndServe(config.ConfigCollection.Services.Port, f))
}
