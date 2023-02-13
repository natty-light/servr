FROM golang:1.20

WORKDIR /usr/src/server

# Install R
RUN apt-get update && apt-get install r-base -y

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN Rscript ./scripts/install.R
RUN go build 

EXPOSE 3000
RUN source .env
CMD ["./generate"]
# CMD ["bash"]


# docker build -t generate .