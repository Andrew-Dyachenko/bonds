package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

func yearPost(c *gin.Context) {

	yearMap := make(map[string]string)

	err := c.BindJSON(&yearMap)
	if err != nil {
		fmt.Println("err: ", err)
	}

	var year string
	for _, value := range yearMap {
		year = string(value)
	}

	test2 := TakeData(year)
	c.JSON(http.StatusOK, struct {
		AllInfos AllInfo `json:"allInfos"`
	}{
		test2.Coupons,
	})
}

func bondsPost(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("err: ", err)
	}
	//fmt.Println(string(jsonData))

	bondsMap := make(map[string]string)
	err = json.Unmarshal(jsonData, &bondsMap)
	if err != nil {
		fmt.Println("err: ", err)
	}

	//fmt.Println("bondsMap: ", bondsMap)
	var bond string
	var count float64
	for i := 1; i < len(bondsMap); i++ {
		bond = bondsMap["bond"]
		count, _ = strconv.ParseFloat(bondsMap["count"], 64)
		bonds = append(bonds, Bond{Name: bond, Count: count})
	}

	for i := 0; i < len(bonds); i++ {
		var bondInfo = Bond{
			bonds[i].Name,
			bonds[i].Count,
		}
		addToDB(bondInfo)
		//deleteBond(bondInfo.Name)
	}
}

func delete(c *gin.Context) {
	deleteMap := make(map[string]string)

	err := c.BindJSON(&deleteMap)
	if err != nil {
		fmt.Println("err: ", err)
	}

	var name string
	for _, value := range deleteMap {
		name = string(value)
	}
	deleteBond(name)
}

func HandleRequest() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))
	router.POST("/year", yearPost)
	router.POST("/bonds", bondsPost)
	router.POST("/delete", delete)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("err: ", err)
	}
}
