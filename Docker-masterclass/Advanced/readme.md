### Docker Masterclass - Advanced


#### Dockerfile creation

### Reactjs App Dockerfile
```
FROM node:16-alpine

WORKDIR /app/web-app

COPY ./package*.json ./

RUN npm install

COPY ./ .

RUN npm run build

CMD ["npm","start"]
```

```
docker build -t web-app:v1.0 .

docker run -p 3000:3000 web-app:v1.0 

```


### Golang App Dockerfile
```
FROM golang:1.19-alpine

WORKDIR /app/backend

COPY go.mod go.sum ./

RUN go mod download

COPY ./ .

RUN go build -o main .

EXPOSE 8000

CMD ["./main"]
```

```
docker build -t go-api:v1.0 .

docker run -p 8000:8000 go-api:v1.0 

```


#### Docker-compose file creation

```
version: '3'

services:
  frontend:
    build:
      context: ./web-app
      dockerfile: Dockerfile
    ports:
      - "3000:3000"
  backend:
    build:
      context: ./go-api
      dockerfile: Dockerfile
    ports:
      - "8000:8000"
    depends_on:
      mongo:
        condition: service_healthy
  mongo:
    image: mongo:4.4.4-bionic
    volumes:
      - mongo-data:/data/db
    ports:
      - "27017:27017"
    healthcheck:
      test: mongo --norc --quiet --host=localhost:27017 --eval "db.getMongo()"
      interval: 30s
      timeout: 2s
      retries: 3
      start_period: 15s

volumes:
  mongo-data:


```





