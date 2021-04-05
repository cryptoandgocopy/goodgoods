package utils

import (
	"log"
)

/*
CheckErr checks the status of the error provided
*/
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
