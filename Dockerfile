#FROM golang:alpine
#WORKDIR /go/cmd
#COPY . .
#RUN go build -o /go/bin/myapp ./cmd/main.go

#FROM golang:1.19
#WORKDIR /usr/src/app
#COPY go.mod go.sum ./
#RUN go mod download && go mod verify
#COPY docker .
#RUN go build -v -o /usr/local/go/src/cmd/main.go
#CMD ["app"]



FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY *.go ./

RUN go build -o /docker-gs-ping cmd/main.go

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
