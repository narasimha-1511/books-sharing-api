package main

import (
	"github.com/gin-gonic/gin"
	"github.com/narasimha-1511/zolo-backend/config"
	"github.com/narasimha-1511/zolo-backend/routes"
)


func main(){
	router := gin.Default()

	config.Connect();//For database Connection to postgres

	routes.Routes(router);//For Routes

	router.Run(":3000")
}