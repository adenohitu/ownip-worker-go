package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/openrdap/rdap"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		client := &rdap.Client{}
		domain, err := client.QueryIP(ctx.ClientIP())

		if err == nil {
			fmt.Printf("Handle=%s\n", domain.Handle)

			var Organization string = ""

			for _, v := range domain.Remarks {
				if v.Title == "description" {
					Organization = v.Description[0]
				}
			}

			ctx.JSON(http.StatusOK, gin.H{
				"status":       "ok",
				"ClientIP":     ctx.ClientIP(),
				"Name":         domain.Name,
				"Organization": Organization,
			})

		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status":       "rdapError",
				"ClientIP":     ctx.ClientIP(),
				"Name":         "",
				"Organization": "",
			})
		}

	})

	r.GET("/rdap/all", func(ctx *gin.Context) {
		client := &rdap.Client{}
		domain, err := client.QueryIP(ctx.ClientIP())

		if err == nil {
			fmt.Printf("Handle=%s\n", domain.Handle)
		}

		ctx.JSON(http.StatusOK, domain)
	})
	return r
}
func main() {
	r := setupRouter()
	r.Run()
}
