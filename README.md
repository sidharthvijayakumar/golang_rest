# golang_rest

## Prerequisite to build Docker image of this application
1. Install docker desktop

## Build GOLANG_REST Docker image

To Build the app using docker run the below command to create a docker image.

**docker build -t golang-rest .**

## To deploy GOLANG_REST Docker image

To deploy use the below command

**docker run -itd -p 8080:8080 golang-rest**