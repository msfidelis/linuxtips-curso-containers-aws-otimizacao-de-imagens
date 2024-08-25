# Primeiro Stage: Build do Artefato
FROM golang:1.22 as build

WORKDIR /src/app

COPY . ./

RUN go get -u
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


# Segundo Stage: Uso de uma imagem minimalista
FROM alpine:3.18

WORKDIR /app

COPY --from=build /src/app/main .

EXPOSE 8080

CMD ["./main"]


