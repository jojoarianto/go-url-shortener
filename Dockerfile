FROM golang:alpine as build-env

# All these steps will be cached
RUN mkdir /url-shorter
WORKDIR /url-shorter
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/url-shorter
FROM scratch
COPY --from=build-env /go/bin/url-shorter /go/bin/url-shorter
ENTRYPOINT ["/go/bin/url-shorter"]
EXPOSE 3000