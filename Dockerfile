# build executable binary
FROM golang:alpine3.16  as builder

WORKDIR $GOPATH/src/github.com/sanix-darker/
# We only copy our app
COPY main.go app/main.go
COPY go.mod app/go.mod

RUN apk add git-lfs

WORKDIR $GOPATH/src/github.com/sanix-darker/app

# We compile
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/sglfs

####################################################################

# Let's build our small image
FROM scratch as prod

# Copy our static executable.
COPY --from=builder /go/bin/sglfs /sglfs
COPY --from=builder /usr/bin/git-lfs /git-lfs

EXPOSE 3000

# Run the hello binary.
ENTRYPOINT ["/sglfs"]
