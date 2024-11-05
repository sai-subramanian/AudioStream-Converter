package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
    pathData,err := os.ReadFile("filename.txt")
    if(err != nil){
        log.Fatal("Error file name not present in filename.txt:", err)
    }
    path := string(pathData)
	fmt.Println("Starting websocket")

	ws, _, err := websocket.DefaultDialer.Dial("ws://localhost:3000/ws/audio", nil)
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer ws.Close()

    

	wavData, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error reading WAV file:", err)
	}

	err = ws.WriteMessage(websocket.BinaryMessage, wavData)
	if err != nil {
		log.Fatal("Error sending WAV data:", err)
	}

	_, flacData, err := ws.ReadMessage()
	if err != nil {
		log.Fatal("Error receiving FLAC data:", err)
	}

	err = os.WriteFile("converted.flac", flacData, 0644)
	if err != nil {
		log.Fatal("Error saving FLAC file:", err)
	}

	fmt.Println("Converted FLAC file saved as 'converted.flac'")
}
