# AudioStream-Converter
### Real-Time Audio Streaming Server

This project provides a backend service built with Go and Gin, designed to handle real-time audio streaming. It converts audio data from WAV to FLAC format using WebSockets and FFmpeg. This service can handle multiple clients streaming audio simultaneously, providing converted FLAC data back to each client in real-time.

### Installation
- Clone the Repository:
    ```bash
    git clone https://github.com/sai-subramanian AudioStream-Converter.git
    cd AudioStream-Converter
    ````
- Install Dependencies: This project uses Go modules. Run the following command to install dependencies:
    ```bash
    go mod tidy
    ````
- Run the Server:
    ```bash
    go run main.go
    ````
This starts the server at http://localhost:3000

**Endpoint**: `GET /ws/audio` Connect to this WebSocket endpoint to start streaming WAV audio data for real-time conversion to FLAC.

### Testing the Websocket

There is an go Script written to test the websocket, in order to run the test script 



- run the test script
    ```bash
    cd test
    go run main.go
    ```
- This will convert WAV audio to FLAC format and save it in the test directory

- (OPTIONAL) change the file to desired file (.wav) , and provide path in the file test/filename.txt, currently this path is set to the example .wav file present in the directory i.e test/file_example_WAV_2MG.wav

