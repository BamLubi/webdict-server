FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

WORKDIR /webdict-server

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o /webdict-server/output/app ./src/main/main.go

EXPOSE 8000
ENTRYPOINT ["/webdict-server/output/app", "-DEV"]