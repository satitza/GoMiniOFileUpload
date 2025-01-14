package main

import (
	"fmt"
	"go-minio-file-upload/configuration"
	"log"
	"os"
)

func main() {
	fmt.Println("Golang minio file upload")

	appConfig, err := configuration.GetConfig()
	if err != nil {
		log.Fatalf("error read configuration. %s\n", err.Error())
	}

	fmt.Printf("port %d\n", appConfig.Port)
	os.Exit(0)
}
