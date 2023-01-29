package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

const pactl = "pactl"
const setSinkVolume = "set-sink-volume"
const upStep = "+10%"
const downStep = "-10%"

func main() {

	fs := http.FileServer(http.Dir("static"))
	mux := http.NewServeMux()

	// Routing
	mux.Handle("/", fs)
	mux.HandleFunc("/vol", volHandler)
	mux.HandleFunc("/play", playHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))

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
		return
	}
	if volume.Direction == "down" {
		cmd := exec.Command(pactl, "--", setSinkVolume, "0", downStep)
		stdout, _ := cmd.Output()
		fmt.Println(string(stdout))
	} else if volume.Direction == "up" {
		cmd := exec.Command(pactl, "--", setSinkVolume, "0", upStep)
		stdout, _ := cmd.Output()
		fmt.Println(string(stdout))
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
