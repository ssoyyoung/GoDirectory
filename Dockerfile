FROM golang:1.14

WORKDIR /go/src/github.com/ssoyyoung.p/GoDirectory

ENV key=value

RUN go get -v -u go.mongodb.org/mongo-driver/mongo
RUN go get github.com/labstack/echo
RUN go get github.com/dgrijalva/jwt-go
RUN git clone --branch 7.8 https://github.com/elastic/go-elasticsearch.git $GOPATH/src/github.com/elastic/go-elasticsearch

#RUN go get github.com/swaggo/swag/cmd/swag
#RUN go get -u github.com/swaggo/echo-swagger
#RUN go get -u github.com/alecthomas/template
#RUN swag init

CMD go run main.go
