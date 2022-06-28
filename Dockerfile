FROM golang:1.15

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE=on

WORKDIR /

COPY go.mod go.sum /
RUN go mod download

COPY . .
RUN go build -o /gorvel main.go

CMD /gorvel