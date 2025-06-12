FROM golang:1.24-alpine AS build

WORKDIR /src

COPY go.mod .

RUN go mod download

COPY . .

# CGO_ENABLED=0 means a static build, so there is no dynamic linking

RUN CGO_ENABLED=0 GOOS=linux go build -o validator-service ./cmd

FROM alpine:edge

WORKDIR /src

COPY --from=build /src/validator-service .

RUN chmod +x /src/validator-service

RUN apk --no-cache add ca-certificates

EXPOSE 1234

ENTRYPOINT [ "/src/validator-service" ]