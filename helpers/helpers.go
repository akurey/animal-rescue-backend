package helpers

import "log"

// manage error.
func HandleErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
}
