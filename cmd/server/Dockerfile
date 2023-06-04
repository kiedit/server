# Build
FROM golang:alpine AS build
WORKDIR /build

RUN mkdir -p /dist

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server

# Runtime
FROM alpine

RUN apk add ffmpeg

COPY --from=build /build/server /usr/bin

EXPOSE 8080

CMD ["server"]