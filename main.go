package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("static"))
	mux := http.NewServeMux()

	// Routing
	mux.Handle("/", fs)
	mux.HandleFunc("/vol", volHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))

}

type VolumeAdjustment struct {
	Direction string `json:"direction"`
	Amount    int    `json:"amount"`
}

func volHandler(w http.ResponseWriter, r *http.Request) {
	var volume VolumeAdjustment
	b, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	_ = json.Unmarshal(b, &volume)
	fmt.Println("testHandler used")
	fmt.Println(&volume)
}
