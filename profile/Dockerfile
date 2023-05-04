FROM golang:1.20 as builder

# modules
WORKDIR $GOPATH/src/github.com/shanehowearth/heidi/profile
ADD . $GOPATH/src/github.com/shanehowearth/heidi/profile

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# build time
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /go/bin/profile cmd/main.go

# Reduce the size of the image
FROM scratch
WORKDIR /root/
COPY --from=0 /go/bin/profile .
# run options
ENTRYPOINT ["./profile"]
