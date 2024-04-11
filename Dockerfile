# syntax=docker/dockerfile:1

#Build Stage

FROM golang:1.19-alpine AS BuildStage

RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN go build -o ascii-art-web .

# Deploy Stage

FROM alpine

ENV PORT="3000" \
    ADDRESS="localhost"

ARG BUILD_DATE

LABEL name="ascii art web" \
    maintainer="Semyon Serbulov <sonnenblumenglas@gmail.com>" \
    org.label-schema.schema-version="1.0" \
    org.label-schema.build-date=$BUILD_DATE \
    org.label-schema.name="ascii-art-web/dockerize" \
    org.label-schema.description="Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. This educational project of the allem school is aimed at learning http and docker" \
    org.label-schema.vcs-url="https://01.alem.school/git/sserbulo/ascii-art-web-dockerize" \
    org.label-schema.docker.cmd="docker run -p 8000:3000 -d --name ascii-art-web-container ascii-art-web"

COPY --from=BuildStage /build/ascii-art-web .
RUN mkdir /templates  && mkdir /ascii-art && mkdir /ascii-art/banners
COPY /templates /templates
COPY /ascii-art/banners /ascii-art/banners
EXPOSE 3000

ENTRYPOINT ["./ascii-art-web"]
