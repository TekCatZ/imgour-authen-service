# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
COPY . .
RUN go build github.com/TekCatZ/imgour-authen-service/cmd/imgour-authen -o /bin/main

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /bin/goapp /app/
ENTRYPOINT ./main