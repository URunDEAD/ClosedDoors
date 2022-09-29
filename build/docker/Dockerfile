FROM golang:1.19.1-alpine3.16

#Enter build dir
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

#Copy files
COPY . .
RUN go build -v -o /usr/local/bin/closeddoors ./...

CMD [ "closeddoors" ]

# Copy files
