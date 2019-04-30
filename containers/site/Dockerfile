FROM golang:1.11 as builder
RUN mkdir /app 
ADD . /go/src/github.com/jwtea/site
WORKDIR /go/src/github.com/jwtea/site
RUN go get ./...
RUN go build -o site . 

FROM alpine:3.7  
CMD ["./site"]
COPY --from=builder /go/src/github.com/jwtea/site .
