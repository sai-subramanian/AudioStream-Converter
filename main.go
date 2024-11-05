package main

import (
	"log"
	"net/http"
	"strconv"

	"bytes"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	ID      string
	Conn    *websocket.Conn
	OutConn *websocket.Conn
}

var (
	clients        = make(map[string]*Client)
	ClientIDcounter = 1 // Initialize the global client ID counter
)

func handleAudioStream(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to set WebSocket upgrade:", err)
		return
	}
	
	ClientIDcounter++ // Increment the counter for each new client
	clientID := "testID" + strconv.Itoa(ClientIDcounter)

	client := &Client{
		ID:   clientID,
		Conn: conn,
	}

	clients[clientID] = client

	go handleClientAudioStream(client)
}

func handleClientAudioStream(client *Client) {
	defer func() {
		client.Conn.Close()
		delete(clients, client.ID)
	}()

	for {

		_, data, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		flacData, err := convertWAVToFLAC(data)
		if err != nil {
			log.Println("Error Converting from WAV to FLAC:", err)
			break
		}

		if err := client.Conn.WriteMessage(websocket.BinaryMessage, flacData); err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}

func convertWAVToFLAC(wavData []byte) ([]byte, error) {

	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-f", "flac", "pipe:1")

	var flacBuffer bytes.Buffer
	cmd.Stdin = bytes.NewReader(wavData)
	cmd.Stdout = &flacBuffer

	err := cmd.Run()
	if err != nil {
		log.Println("Error running FFmpeg:", err)
		return nil, err
	}

	return flacBuffer.Bytes(), nil
}

func main() {
	r := gin.Default()

	r.GET("/ws/audio", handleAudioStream)

	r.Run(":3000")
}
