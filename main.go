package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os/exec"
)

const pactl = "pactl"
const setSinkVolume = "set-sink-volume"
const upStep = "+10%"
const downStep = "-10%"

//go:embed static/*
var static embed.FS

func main() {

	// strip the "static" prefix from the file server
	fsDir, err := fs.Sub(static, "static")
	if err != nil {
		log.Fatal("error getting static content", err)
	}
	// Create file server
	staticFileServer := http.FileServer(http.FS(fsDir))

	// Routing
	mux := http.NewServeMux()
	mux.Handle("/", staticFileServer)
	mux.HandleFunc("/vol", volHandler)
	mux.HandleFunc("/play", playHandler)

	// Start server
	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", mux))

}

type VolumeAdjustment struct {
	Direction string `json:"direction"`
	Amount    int    `json:"amount"`
}

func volHandler(w http.ResponseWriter, r *http.Request) {
	var volume VolumeAdjustment
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
	}(r.Body)
	err = json.Unmarshal(b, &volume)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		log.Println("malformed json")
		return
	}
	if volume.Direction == "down" {
		cmd := exec.Command(pactl, "--", setSinkVolume, "0", downStep)
		stdout, _ := cmd.Output()
		fmt.Println("volume down: ", string(stdout))
	} else if volume.Direction == "up" {
		cmd := exec.Command(pactl, "--", setSinkVolume, "0", upStep)
		stdout, _ := cmd.Output()
		fmt.Println("volume up: ", string(stdout))
	}
	fmt.Println(&volume)
}

type playCommand struct {
	Command string `json:"command"`
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	var play playCommand
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
	}(r.Body)
	err = json.Unmarshal(b, &play)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	if play.Command == "play-pause" {
		cmd := exec.Command("playerctl", "play-pause")
		stdout, _ := cmd.Output()
		fmt.Println(string(stdout))
	} else if play.Command == "next" {
		cmd := exec.Command("playerctl", "next")
		stdout, _ := cmd.Output()
		fmt.Println(string(stdout))
	} else if play.Command == "previous" {
		cmd := exec.Command("playerctl", "previous")
		stdout, _ := cmd.Output()
		fmt.Println(string(stdout))
	}
}
