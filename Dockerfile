FROM golang:1.13.6
WORKDIR /go/src/gowebapp/
RUN go get -d -v golang.org/x/net/html  
COPY /go/src/gowebapp/app.go	.
RUN go get github.com/gorilla/mux
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add tzdata
WORKDIR /root/
COPY --from=0 /go/src/gowebapp/app .
CMD ["./app"]   
