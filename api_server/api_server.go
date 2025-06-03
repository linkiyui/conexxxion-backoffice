package api_server

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/flowchartsman/swaggerui"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.com/conexxxion/conexxxion-backoffice/config"

	auth_controller "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/auth/infra/controller"
	user_controller "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/user/infra/controller"

	auth_middleware "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/auth/infra/middleware"
	user_middleware "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/user/infra/middleware"
)

var requestsDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "api_request_duration",
		Help: "API requests duration in milliseconds",
	},
	[]string{"status"},
)

// go:embed swagger_doc.yaml
var swagger_doc []byte

var apiServer *APIServer

type APIServer struct {
	ginInstance *gin.Engine
}

func NewApiServer() *APIServer {
	conf := config.GetConfig()
	ginInstance := gin.New()
	errorsTemplate := filepath.Join(config.GetConfig().TemplatesPath, "errors.html")
	getcodeTemaplate := filepath.Join(config.GetConfig().TemplatesPath, "get_code.html")
	ginInstance.LoadHTMLFiles(errorsTemplate, getcodeTemaplate)
	ginInstance.Use(gin.Recovery())
	prometheus.MustRegister(requestsDuration)
	registerRoutes(ginInstance)
	if conf.IsDev() {
		// fmt.Println("dev mode")
		registerDevRoutes(ginInstance)
	}
	// registerDevRoutes(ginInstance)
	apiServer = &APIServer{
		ginInstance: ginInstance,
	}
	return apiServer
}

func (s *APIServer) Start() {
	url := fmt.Sprintf(":%d", config.GetConfig().APIListenPort)
	fmt.Println(url)
	if err := s.ginInstance.Run(url); err != nil {
		log.Fatal("can't start api server on", url)
	}
}

func registerDevRoutes(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.Data(200, "text/html; charset=utf-8", []byte("Welcome to Conexxxion API"))
	})

	// auth for documentation endpoints
	docAuth := r.Group("", gin.BasicAuth(gin.Accounts{
		"conexxxion": "nomedalaganaquelasepas",
	}))

	docAuth.GET("/swagger/*any",
		gin.WrapH(http.StripPrefix("/swagger", swaggerui.Handler(swagger_doc))))

}

func registerRoutes(r *gin.Engine) {
	r.Use(LoggerMiddleware)

	v1 := r.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/create-user/", auth_middleware.VerifyLoginToken, auth_controller.CreateUser)

	auth.POST("/login/", auth_controller.Login)

	auth.POST("/refresh-token/", auth_controller.RefreshToken)

	user := v1.Group("/user")
	userWithAuth := user.Use(auth_middleware.VerifyLoginToken)
	userWithAuth.GET("/me/", user_controller.GetMeInfo)
	userWithAuth.PATCH("/pasword/", user_controller.UpdatePassword)
	userWithAuth.PATCH("/", user_controller.UpdateUser)
	userWithAuth.GET("/:id", user_controller.GetUserInfo)
	userWithAuth.DELETE("/", user_controller.DeleteMeUser)

	// solo admin y super_admin
	userWithAuthAdmin := user.Use(user_middleware.VerifyAdmin)
	userWithAuthAdmin.DELETE("/:id", user_controller.DeleteUser)

}
