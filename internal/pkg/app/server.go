package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"lab3/internal/app/middlewares"
	"log"
)

func (a *Application) StartServer() {
	log.Println("Server start up")
	log.Println("Server start up")
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:3000"}
	config.AllowMethods = []string{"PUT", "PATCH", "GET", "POST", "DELETE"}
	r.Use(cors.New(config))

	public := r.Group("/api")
	public.POST("/register", a.Register)
	public.POST("/login", a.Login)

	protected := r.Group("api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", a.CurrentUser)
	protected.GET("/bucket/:id_user")
	r.GET("/goods", a.GetAll)
	r.GET("/goods/:id", a.GetProduct)
	r.POST("/goods", a.PostProduct)
	r.PUT("/goods/:id", a.ChangePrice)
	r.DELETE("goods/:id", a.DeleteProduct)
	r.GET("/ping/:name", a.Ping)
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Println("Run failed")
	}
	log.Println("Server down")
}

type AnswerJSON struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

type pingResp struct {
	Status string `json:"status"`
}
