FROM golang:1.19.1-alpine3.16
ENV MYSQL_HOST=127.0.0.1
ENV MYSQL_PORT=3306
ENV MYSQL_USER=root
ENV MYSQL_DBNAME="doors"

#Enter build dir
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

#Copy files
COPY . .
RUN go build -v -o /usr/local/bin/closeddoors ./...

#Start Execution
CMD [ \
    "sh", "-c", \
    "closeddoors \
    -host $MYSQL_HOST \
    -port $MYSQL_PORT \
    -user $MYSQL_USER \
    -passwd $MYSQL_PASSWD \
    -db-name $MYSQL_DBNAME " \
    ]