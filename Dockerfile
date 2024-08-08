FROM golang:1.22-alpine
LABEL maintainer="Me"


# Setting up Dev environment

WORKDIR /myapp/
# note this file, go.mod exists locally. and contain reference 
# to direct/indirect dependencies. this step allows to download 
# dependencies and speedup build for docker images (if it used 
# to build artifacts, and not as dev env).  
COPY go.mod  /myapp/go.mod
RUN go mod download


EXPOSE 1323