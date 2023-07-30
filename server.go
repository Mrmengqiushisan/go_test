package main

import (
	"os"

	"github.com/Mrmengqiushisan/go_test/cotroller"
	"github.com/Mrmengqiushisan/go_test/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := Init("./data/"); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(ctx *gin.Context) {
		topicId := ctx.Param("id")
		data := cotroller.QueryPageInfo(topicId)
		ctx.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}

func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		return err
	}
	return nil
}
