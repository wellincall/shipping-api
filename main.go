package main

import (
	"fmt"
	"os"

	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func getEuMembers() map[string]bool {
	membersFile, err := os.Open("euMembers.json")
	if err != nil {
		fmt.Println(err)
	}
	fileInBytes, _ := ioutil.ReadAll(membersFile)
	membersFile.Close()
	var membersMap map[string]bool
	json.Unmarshal([]byte(fileInBytes), &membersMap)
	return membersMap
}

func main() {
	r := gin.Default()

	r.GET("/shipments", GetHandler)
	r.POST("/shipments", PostHandler)

	r.Run()
}
