package utils

import "log"

func CheckError(err error) {
	if err != nil {
		log.Println("Fatal Error : ", err.Error())
	}
}
