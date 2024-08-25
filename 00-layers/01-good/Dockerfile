FROM debian:latest

RUN apt-get update && \
    apt-get install curl wget vim nginx -y && \
    apt-get clean autoclean && \
    rm -rf /var/lib/{apt,dpkg,cache,log}/

EXPOSE 8080