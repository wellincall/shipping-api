package main

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	csvFile, err := os.Open("shipments.csv")

	if err != nil {
		c.JSON(500, gin.H{"error": "unable to open csv file"})
		return
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		c.JSON(500, gin.H{"error": "file does not have CSV structure"})
	}
	var shipments []ShipmentResponse

	for _, line := range csvLines {
		weight, _ := strconv.ParseFloat(line[2], 32)
		price, _ := strconv.ParseFloat(line[3], 32)
		shipment := ShipmentResponse{
			Sender:   line[0],
			Receiver: line[1],
			Weight:   float32(weight),
			Price:    float32(price),
		}
		shipments = append(shipments, shipment)
	}
	defer csvFile.Close()
	c.JSON(200, gin.H{"data": shipments})
}
