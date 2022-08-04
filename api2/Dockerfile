# https://www.bacancytechnology.com/blog/dockerize-golang-application

#====================
# STEP1 - BUILD
#====================
FROM golang:1.18.2-alpine3.15 as build-env

# Set environment variable
ENV APP_NAME theapp
ENV CMD_PATH main.go

# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

# Install go packages
RUN go mod tidy

# Build application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

#====================
# STEP2 - PACKAGE
#====================
FROM alpine:3.15

# Set environment variable
ENV APP_NAME theapp

# Copy only required data into this image
COPY --from=build-env /$APP_NAME .

# Expose application port
EXPOSE 80

# Start app
CMD ./$APP_NAME
