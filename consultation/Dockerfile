FROM golang:1.20 as builder

# modules
WORKDIR $GOPATH/src/github.com/shanehowearth/heidi/consultation
ADD . $GOPATH/src/github.com/shanehowearth/heidi/consultation

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# build time
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /go/bin/consultation cmd/main.go

# Reduce the size of the image
FROM scratch
WORKDIR /root/
COPY --from=0 /go/bin/consultation .
# run options
ENTRYPOINT ["./consultation"]
