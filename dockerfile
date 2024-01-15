FROM golang:alpine

WORKDIR /app
RUN apt-get autoclean && apt-get clean && apt-get autoremove
RUN apt update && apt upgrade -y
RUN apt install nodejs -y
RUN apt install npm -y
COPY . .
RUN go install

# RUN go build -o main .
# RUN ls

RUN npm install
RUN npm install -g nodemon

EXPOSE 8080

# CMD ["nodemon --exec go run main.go --signal SIGTERM"]
