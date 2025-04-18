FROM golang:1.24.2-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest
COPY . .
RUN make install-tailwind-linux
RUN templ generate
RUN make tailwind-generate
RUN go build -o gohome ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/gohome .
COPY --from=build /app/static ./static
EXPOSE 8080
CMD ["./gohome"]
