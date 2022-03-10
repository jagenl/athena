package main

import (
	"athena/bootstarp"
	"athena/routers"
	"github.com/gin-gonic/gin"
)

func main()  {

	r := gin.Default();
	bootstarp.SetupDB()

	routers.InitApiRouter(r);
	_ = r.Run(":8888")

}
