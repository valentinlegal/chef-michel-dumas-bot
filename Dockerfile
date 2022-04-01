FROM golang:1.17-buster

WORKDIR /app

COPY . .

RUN apt-get update -y

CMD [ "/app/chef-michel-dumas-bot" ]
