# Build
FROM golang:alpine AS build
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o events ./cmd/events

# Runtime
FROM alpine

RUN apk add ffmpeg

COPY --from=build /build/events /usr/bin

CMD ["events"]