FROM golang:1.11.4
WORKDIR /go/src/github.com/ilgooz/service-location
COPY . .
RUN go install ./...
CMD location \
    --maxmindDBPath /go/src/github.com/ilgooz/service-location/data/GeoLite2-City/GeoLite2-City.mmdb