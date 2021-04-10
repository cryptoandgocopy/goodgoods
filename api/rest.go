package api

import (
	"encoding/json"
	"log"

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

/*
Create a server for rest API
*/
func Create() {
	// server setup
	app := fiber.New()

	// logging
	app.Use(logger.New())

	// handlers
	app.Get("/isGood/:origin/:goods", isGood)

	// start
	log.Fatal(app.Listen(":3000"))
}

/*
Handle api request for isGood
*/
func isGood(c *fiber.Ctx) error {
	// check data
	responseDOL := true //data.IsGood(c.Params("origin"), c.Params("goods"))

	// build JSON response
	ead := externalAdapterData{responseDOL}
	ear := externalAdapterResponse{"abc123", ead, "", ""}
	jsonData, err := json.Marshal(ear)
	utils.CheckErr(err)

	return c.SendString(string(jsonData))
}
