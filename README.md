
# kiedit

KiEdit: the ultimate solution for hassle-free WhatsApp video sharing!

Tired of dealing with the frustration of syncing audio and video when sharing longer videos on WhatsApp status? Look no further! KiEdit is the innovative tool you've been waiting for. It empowers you to effortlessly split your videos into seamless 30-second segments, ensuring your content is flawlessly synchronized every time you share it with your family and friends.


## Tech Stack

- Golang
- rabbitmq
- docker
- ffmpeg (install on your machine)


## Run Locally

Clone the project

```bash
  git clone https://github.com/kiedit/server.git
```

Go to the project directory

```bash
  cd server
```

Install dependencies

```bash
  go mod download
```

Start the rabbitmq instance

```bash
  docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

Start the queue consumer

```bash
  go run cmd/events/main.go
```

Start the server

```bash
  go run cmd/server/main.go
```

Upload mp4 file via curl or rest api application such as postman
