FROM golang:1.19.7-bullseye

WORKDIR /app
COPY . .
RUN go install
RUN go build -o main .
RUN ls
RUN apt update && apt upgrade -y
RUN apt install nodejs -y
RUN apt install npm -y
RUN npm install -g nodemon

EXPOSE 8080

CMD ["nodemon --exec go run main.go --signal SIGTERM"]