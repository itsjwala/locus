FROM golang:1.14-alpine as builder
LABEL stage=locus_web-builder

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go get github.com/rakyll/statik \
    && $GOPATH/bin/statik -src="frontend" 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /src/main .

EXPOSE 8000

CMD ["./main"] 
