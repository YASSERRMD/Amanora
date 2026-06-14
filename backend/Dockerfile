FROM golang:1.25-alpine AS build

ARG AMANORA_COMMAND=api

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/amanora ./cmd/${AMANORA_COMMAND}

FROM alpine:3.22

RUN addgroup -S amanora && adduser -S amanora -G amanora

WORKDIR /app
COPY --from=build /out/amanora /app/amanora

USER amanora
EXPOSE 8080

ENTRYPOINT ["/app/amanora"]
