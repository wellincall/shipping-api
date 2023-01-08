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
		return
	}
	var shipments []ShipmentResponse

	for _, line := range csvLines {
		weight, _ := strconv.ParseFloat(line[2], 64)
		price, _ := strconv.ParseFloat(line[3], 64)
		shipment := ShipmentResponse{
			Sender:   line[0],
			Receiver: line[1],
			Weight:   float64(weight),
			Price:    float64(price),
		}
		shipments = append(shipments, shipment)
	}
	csvFile.Close()
	c.JSON(200, gin.H{"data": shipments})
}

func calculatePrice(s *ShipmentInput) (float64, string) {
	shipmentPrice, err := WeightConstraint(s)
	if err != "" {
		return 0, err
	}
	regionPrice, err := RegionConstraint(s)
	if err != "" {
		return 0, err
	}
	return shipmentPrice * regionPrice, ""
}

func storeShipment(s *ShipmentInput, price float64) string {
	csvFile, err := os.OpenFile("shipments.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return "error when opening csv file"
	}
	writer := csv.NewWriter(csvFile)
	weight := strconv.FormatFloat(s.Weight, 'E', -1, 64)
	shipmentPrice := strconv.FormatFloat(price, 'E', -1, 64)
	var newLine = []string{
		s.Sender, s.Receiver, weight, shipmentPrice,
	}
	writer.Write(newLine)
	writer.Flush()
	csvFile.Close()
	return ""
}

func PostHandler(c *gin.Context) {
	var shipment ShipmentInput
	if err := c.ShouldBind(&shipment); err != nil {
		c.JSON(400, gin.H{"error": "bad request"})
		return
	}
	price, err := calculatePrice(&shipment)
	if err != "" {
		c.JSON(422, gin.H{"error": err, "input": shipment})
		return
	}
	err = storeShipment(&shipment, price)
	if err != "" {
		c.JSON(500, gin.H{"error": err, "input": shipment})
		return
	}

	c.JSON(200, gin.H{"input": shipment, "price": price})
}
