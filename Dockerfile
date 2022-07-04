# build stage
FROM golang:1.18-alpine AS build-env
RUN mkdir ../home/app
WORKDIR /../home/app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY api ./api
COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg
COPY test ./test
RUN go build -o /bin/main github.com/TekCatZ/imgour-authen-service/cmd/imgour-authen

# run stage
FROM alpine
EXPOSE 8018 9812
WORKDIR /app
COPY --from=build-env /bin/main /app/
COPY /configs /app/
ENTRYPOINT ./main