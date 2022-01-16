package main

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	GinEngine *gin.Engine
	Service   Service
}

func (s Server) setUp() {
	s.GinEngine = gin.Default()

	api := s.GinEngine.Group("/api")
	{
		api.GET("/", s.regHandler)
	}
}

func (s Server) regHandler(c *gin.Context) {
	s.Service.registerUser()
}
