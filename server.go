package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Mrmengqiushisan/go_test/cotroller"
	"github.com/Mrmengqiushisan/go_test/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := Init("./data/"); err != nil {
		os.Exit(-1)
	}
	fmt.Println("话题数为：", *repository.GetTopicCount())
	fmt.Println("贴子数为：", *repository.GetPostCount())

	r := gin.Default()
	r.GET("/community/page/get/:id", func(ctx *gin.Context) {
		topicId := ctx.Param("id")
		for {
			fmt.Println("请问是否有帖子需要输入呢")
			reader := bufio.NewReader(os.Stdin)
			put, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("input error", err)
				break
			}
			putstr := strings.Trim(put, "\r\n")
			if strings.ToUpper(putstr) == "Y" {
				if err := InputPost(); err != nil {
					break
				}
			} else {
				break
			}
		}
		if err := Init("./data/"); err != nil {
			fmt.Println("再次初始化失败", err)
		}
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

func InputPost() error {
	for {
		fmt.Println("结束请输入E")
		fmt.Println("我需要以下信息：")
		fmt.Println("请输入发布帖子的内容：")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input Please try again", err)
			continue
		}
		topicContext := strings.Trim(input, "\r\n")
		if strings.ToUpper(topicContext) == "E" {
			break
		}
		fmt.Println("请输入您选择的话题：")
		reader = bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input Please try again", err)
			continue
		}
		parentIdStr := strings.Trim(input, "\r\n")
		parent, err := strconv.ParseInt(parentIdStr, 10, 64)
		if err != nil {
			fmt.Println("Invalid input,Please enter an integer value")
			continue
		} //构建帖子结构
		now := time.Now()
		id := repository.GetPostCount()
		(*id)++
		fmt.Println("当前id值为")
		posTmp := repository.Post{
			Id:         *id,
			ParentId:   parent,
			Content:    topicContext,
			CreateTime: now.Unix(),
		}
		open, err := os.OpenFile("./data/post", os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("file can not open ")
			return errors.New("file can not open")
		}
		buf, err := json.Marshal(posTmp)
		if err != nil {
			fmt.Println("序列化失败")
			return err
		}
		_, err = open.Write([]byte{'\n'})
		_, err = open.Write(buf)
		if err != nil {
			fmt.Println("写入失败", err)
			return err
		}
		fmt.Println("数据写入成功")
	}
	return nil
}
