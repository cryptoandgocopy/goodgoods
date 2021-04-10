package api

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"goodgoods/data"
	"goodgoods/utils"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

/*
Data structures for API response
*/
type externalAdapterResponse struct {
	JobRunID string              `json:"jobRunID"`
	Data     externalAdapterData `json:"data"`
	Status   string              `json:"status,omitempty"`
	Error    string              `json:"error,omitempty"`
}

type externalAdapterData struct {
	IsGood bool `json:"isGood"`
}

type apiRequest struct {
	ID   string         `json:"id"`
	Data apiRequestData `json:"data"`
}

type apiRequestData struct {
	Origin string `json:"origin"`
	Goods  string `json:"goods"`
}

/*
Create a server for rest API
*/
func Create() {
	// server setup
	app := fiber.New()

	// logging
	app.Use(logger.New())

	// handlers
	app.Post("/isGood", isGood)

	// start
	port := ":3000"
	if os.Getenv("PORT") != "" {
		var sb strings.Builder
		sb.WriteString(":")
		sb.WriteString(os.Getenv("PORT"))
		port = sb.String()
	}
	log.Fatal(app.Listen(port))
}

/*
Handle api request for isGood
*/
func isGood(c *fiber.Ctx) error {
	// parse request
	request := parseRequest(c.Body())

	// check data against database
	responseDOL := data.IsGood(request.Data.Origin, request.Data.Goods)

	// build JSON response
	ead := externalAdapterData{responseDOL}
	ear := externalAdapterResponse{request.ID, ead, "", ""}
	jsonData, err := json.Marshal(ear)
	utils.CheckErr(err)

	// display json response
	return c.SendString(string(jsonData))
}

/*
Parse JSON request data
*/
func parseRequest(request []byte) apiRequest {
	// read into struct
	var result apiRequest

	err := json.Unmarshal(request, &result)
	utils.CheckErr(err)

	return result
}
