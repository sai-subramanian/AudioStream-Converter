FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

# Install FFmpeg
RUN apt-get update && apt-get install -y ffmpeg

RUN go mod download

COPY . .

RUN go build -o websocket-server .



EXPOSE 3000

CMD ["./websocket-server"]
