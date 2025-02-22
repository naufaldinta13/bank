FROM alpine:latest

RUN apk add --no-cache tzdata
# ENV TZ="Asia/Jakarta"

ARG COMMIT=""

WORKDIR /root/

# Copy the Pre-built binary
# the binary should be run by
COPY bank .

# Not Secret env variable
ENV DEBUG_MODE=true
ENV REST_SERVER="0.0.0.0:80"
ENV APP_VERSION=${COMMIT}

# Expose port 80 to the outside world
EXPOSE 80

# Command to run the executable
CMD ["./bank"]