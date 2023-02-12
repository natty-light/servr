FROM golang:1.20

WORKDIR /usr/src/web

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

RUN go build 

EXPOSE 5000
CMD ["./web"]