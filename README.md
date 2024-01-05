# Rest API with Go, Fiber, Docker and Postgre

## Dev environment
Create `Dockerfile` and `docker-compose.yaml`

run `docker-compose up` while docker is running

When working with docker we need to run all the commands in the container bash, run the following command to get into the container bash
```sh
docker compose run --service-ports web bash
```

in the Go container shell, initialize a new Go module in your project using go mod init. This will create a go.mod file that will be used to
install and manage dependencies
```sh
go mod init github.com/JiCodes/go-rest-api
```

then install the fiber package
```sh
go get github.com/gofiber/fiber/v2
```

create a dir called cmd to hold main.go file
```sh
mkdir cmd
touch cmd/main.go
```

copy the following code into main.go
```go
package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Go Rest API!")
    })

    app.Listen(":3000")
}
```

run the app with binding on localhost
```sh
go run cmd/main.go -b 0.0.0.0
```

check the app is running by visiting http://localhost:3000

## modify the Dockerfile
```Dockerfile

COPY . . // copy all the files from host to the container

RUN go install github.com/cosmtrek/air@latest // install air to watch the changes in the code and restart the app automatically

RUN go mod tidy // to install all the dependencies

```

** create a new file called .air.toml **
copy the confif from air github repo air example config file
`https://github.com/cosmtrek/air/blob/master/air_example.toml`

with only one change
```toml
cmd = "go build -o ./tmp/main ./cmd"
```

## modify the docker-compose.yaml 
```yaml
...
  command: air cmd/main.go -b 0.0.0.0 
```
allow us to launch the app with docker compose up command from the host machine (now we don't need to get into the container bash to run the app)

rebuild the image
```sh
docker compose build
```
then run the app
```sh
docker compose up
```

update docker-compose yaml file with postgres and add .env file for it
```yaml
...
  db:
    image: postgres:alpine
    environment:
      - POSTGRES_USER=${DB_USER}
      - POSTGRES_PASSWORD=${DB_PASSWORD}
      - POSTGRES_DB=${DB_NAME}
    ports:
      - 5432:5432
    volumes:
      - postgres-db:/var/lib/postgresql/data
  
  volumes:
    postgres-db:

```
enter the docker shell
```sh
docker compose run --service-ports web bash
```
then run the following command to install the gorm ORM package and postgres driver
```dockershell
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

create the models dir and models.go file
```sh
mkdir models
touch models/models.go
```

create the database directory and database.go file
```sh
mkdir database
touch database/database.go
```

## Routes and Endpoints

```sh
touch cmd/routes.go # to hold all the routes
```

```sh
mkdir handlers
touch handlers/handlers.go # to hold all the handlers for the routes
```

then add some routes and handlers to the files


## enable templates 

```dockershell
go get github.com/gofiber/template/html/v2

mdkir views

touch views/index.html
```

```sh
docker compose up
```
create a layout file in layouts dir
```sh
mkdir layouts
touch layouts/main.html
```

Add several views and partials to the layout's main.html file
