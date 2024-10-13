package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/openrdap/rdap"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		client := &rdap.Client{}
		domain, err := client.QueryIP(ctx.ClientIP())

		if err == nil {
			fmt.Printf("Handle=%s\n", domain.Handle)
		}

		var Organization string = "----"

		for _, v := range domain.Remarks {
			if v.Title == "description" {
				Organization = v.Description[0]
			}
		}
		ctx.JSON(http.StatusOK, gin.H{
			"ClientIP":     ctx.ClientIP(),
			"Name":         domain.Name,
			"Organization": Organization,
		})
	})

	r.GET("/rdap/all", func(ctx *gin.Context) {
		client := &rdap.Client{}
		domain, err := client.QueryIP(ctx.ClientIP())

		if err == nil {
			fmt.Printf("Handle=%s\n", domain.Handle)
		}

		ctx.JSON(http.StatusOK, domain)
	})
	r.Run()
}
