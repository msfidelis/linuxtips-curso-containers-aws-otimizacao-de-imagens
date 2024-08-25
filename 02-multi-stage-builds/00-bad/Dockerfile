FROM golang:1.22 

WORKDIR $GOPATH/src/app

COPY . ./

RUN go get -u
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


EXPOSE 8080

ENTRYPOINT ["./main"]