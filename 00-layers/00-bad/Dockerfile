FROM debian:latest

RUN apt-get update

RUN apt-get install curl wget vim -y 

RUN apt-get install nginx -y

RUN apt-get clean autoclean

RUN rm -rf /var/lib/{apt,dpkg,cache,log}/

EXPOSE 8080