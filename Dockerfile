FROM golang:1.19 as dependencies

WORKDIR /gomod
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM golang:1.19 as build

COPY --from=dependencies /go /go
WORKDIR /build
COPY main.go .
COPY api_test.go .
COPY go.mod .
COPY go.sum .

RUN go build -o app

FROM golang:1.19 as test

COPY --from=dependencies /go /go
WORKDIR /build
RUN curl -sSL "https://github.com/gotestyourself/gotestsum/releases/download/v1.8.2/gotestsum_1.8.2_linux_amd64.tar.gz" | tar -xz -C /usr/local/bin gotestsum
COPY . .