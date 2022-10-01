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
RUN go build -v -o /usr/local/bin/ ./...

#Start Execution
CMD [ \
    "sh", "-c", \
    "closeddoors \
    -MYSQL-Host $MYSQL_HOST \
    -MYSQL-Port $MYSQL_PORT \
    -MYSQL-User $MYSQL_USER \
    -MYSQL-Passwd $MYSQL_PASSWD \
    -MYSQL-DB-name $MYSQL_DBNAME " \
    ]