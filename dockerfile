FROM golang:1.17-alpine as build
ADD . /src
WORKDIR /src
RUN go build

FROM alpine:latest as execute
ENV BASE="base.yaml"
ENV OPTIONS=""
WORKDIR /project
COPY --from=build /src/oas-project-bootstraper /usr/local/bin
ADD ./example /project
CMD oas-project-bootstraper -b /project/$BASE -w /project $OPTIONS
