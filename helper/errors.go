package helper

import (
	"fmt"
	"log"
)

func PanicIfError(err error) {
	if err != nil {
		log.Panic(err)
		fmt.Println(err)
	}
}
