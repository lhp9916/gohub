package main

import (
	"fmt"
	"gohub/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	bootstrap.SetRouter(router)
	err := router.Run(":8000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
