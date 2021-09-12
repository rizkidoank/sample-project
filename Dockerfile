FROM golang:1.16
WORKDIR /src
COPY go.mod main.go ./
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:3.13  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /src/app ./
EXPOSE 8080
CMD ["./app"]
