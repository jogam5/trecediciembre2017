package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//func main() {
//	fs := http.FileServer(http.Dir("static"))
//	http.Handle("/", fs)
//
//	log.Println("Listening...")
//	http.ListenAndServe(":3000", nil)
//}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World FIM")
	//fs := http.FileServer(http.Dir("static"))
	//http.Handle("/", fs)

}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	//http.HandleFunc("/", hello)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
