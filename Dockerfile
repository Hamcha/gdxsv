FROM golang:alpine

# Install dependencies
RUN apk add protoc git build-base

# Add code
ADD . /app

# Build
RUN cd /app && make install-tools build

VOLUME /data

CMD /app/bin/gdxsv