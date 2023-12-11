FROM golang:1.21-alpine AS build

WORKDIR /
COPY . /

RUN go build -o ./api/cmd/main ./api/cmd/main.go

FROM scratch

WORKDIR /app
COPY --from=build /api/cmd/main .

ENTRYPOINT ["/app/main"]
