#go image
FROM golang:1.23.8-alpine
WORKDIR /exporterapp
COPY go.mod go.sum ./
RUN go mod tidy
COPY . . 
RUN go build -o main .
EXPOSE 2112
ENTRYPOINT ["./main"] 