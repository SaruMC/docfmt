FROM golang:1.22.0-alpine3.19 as build
WORKDIR /opt/build

COPY . .
RUN go build -o main main.go

FROM alpine
WORKDIR /opt/app

COPY --from=build /opt/build/main .

ENTRYPOINT ["/opt/app/main"]