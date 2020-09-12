FROM golang:latest

WORKDIR /usr/src/app
COPY . .
RUN go build -o server cmd/server/main.go
RUN cp configs/settings.json settings.json

EXPOSE 8888
CMD ["./server"]




