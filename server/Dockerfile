FROM golang:1.21-alpine AS build

WORKDIR /
COPY . /

RUN go build -o ./server/cmd/main ./server/cmd/main.go

FROM scratch

WORKDIR /app
COPY --from=build /server/cmd/main .

ENTRYPOINT ["/app/main"]
