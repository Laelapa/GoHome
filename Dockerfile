FROM golang:1.24.2-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN apk add --no-cache curl
RUN go install github.com/a-h/templ/cmd/templ@latest
COPY . .
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
RUN chmod +x tailwindcss-linux-x64
RUN mv tailwindcss-linux-x64 tailwindcss
RUN templ generate
RUN ./tailwindcss -i ./src/css/custom.css -o ./static/css/style.css --minify
RUN go build -o gohome ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/gohome .
COPY --from=build /app/static ./static
EXPOSE 8080
CMD ["./gohome"]
