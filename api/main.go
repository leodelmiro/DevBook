package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/* func init() {
	key := make([]byte, 64)

	if _, initError := rand.Read(key); initError != nil {
		log.Fatal(initError)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println(stringBase64)
} */

func main() {
	config.LoadEnvironments()
	r := router.Generate()
	
	fmt.Printf("Running API on port %d\n", config.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), r))
}
