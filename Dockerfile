FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o ./build/app_srv_tic ./cmd/main.go

EXPOSE 8000

CMD ["./build/app_srv_tic"]