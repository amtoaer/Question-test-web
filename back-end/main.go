package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	mayuan, _ := ioutil.ReadFile("./mayuan.json")
	maogai, _ := ioutil.ReadFile("./maogai.json")
	var mayuanQuestionList []map[string]interface{}
	var maogaiQuestionList []map[string]interface{}
	json.Unmarshal(mayuan, &mayuanQuestionList)
	json.Unmarshal(maogai, &maogaiQuestionList)
	router := gin.Default()
	router.GET("/api/:subject/position/:num", func(c *gin.Context) { //按位置返回
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		var questionList []map[string]interface{}
		subject := c.Param("subject")
		if subject == "mayuan" {
			questionList = mayuanQuestionList
		} else if subject == "maogai" {
			questionList = maogaiQuestionList
		} else {
			c.JSON(404, "subject error!")
		}
		numStr := c.Param("num")
		num, err3 := strconv.Atoi(numStr)
		if err3 != nil {
			c.JSON(404, "num error!")
		}
		if num >= 0 && num < len(questionList) {
			c.JSON(200, questionList[num])
		} else {
			c.JSON(404, "out of range!")
		}
	})
	router.GET("/api/:subject/random/*type", func(c *gin.Context) { //随机返回
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		var questionList []map[string]interface{}
		subject := c.Param("subject")
		if subject == "mayuan" {
			questionList = mayuanQuestionList
		} else if subject == "maogai" {
			questionList = maogaiQuestionList
		} else {
			c.JSON(404, "subject error!")
		}
		rand.Seed(time.Now().UnixNano())
		randomNum := rand.Intn(len(questionList))
		questionType := c.Param("type")
		if questionType == "/" {
			c.JSON(200, questionList[randomNum])
		} else if questionType == "/radio" {
			for {
				if len(questionList[randomNum]["Answer"].(string)) == 1 {
					c.JSON(200, questionList[randomNum])
					break
				} else {
					randomNum = rand.Intn(len(questionList))
				}
			}
		} else if questionType == "/checkbox" {
			for {
				if len(questionList[randomNum]["Answer"].(string)) != 1 {
					c.JSON(200, questionList[randomNum])
					break
				} else {
					randomNum = rand.Intn(len(questionList))
				}
			}
		}
	})
	router.Run(":8080")
}
