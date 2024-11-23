package main

import (
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

type TensorFlowSert struct {

	`json:"instances"`

}

type TensorFlowResult struct {

	`json:"predictions"`
	
}

func GetResult(C *gin.Context){

	var DataTest TensorFlowSert

	var Result TensorFlowResult

	C.JSON(200, Result)
}

func main(){

	r := gin.Default()

	r.POST("/GetResult", GetResult)

	r.Run(":8080")
}