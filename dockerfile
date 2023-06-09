FROM golang:1.20.3
WORKDIR /app
COPY go.mod go.sum /
RUN go mod download
COPY . .
WORKDIR /app/src/main
RUN CG0_ENABLED=0  go build -o main .
EXPOSE 8080
CMD ["./main", "serve"]