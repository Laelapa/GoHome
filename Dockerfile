FROM golang:1.24.2 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN apt-get update && apt-get install -y curl
RUN go install github.com/a-h/templ/cmd/templ@latest
COPY . .
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
RUN chmod +x tailwindcss-linux-x64
RUN mv tailwindcss-linux-x64 tailwindcss
RUN templ generate
RUN ./tailwindcss -i ./src/css/custom.css -o ./static/css/style.css --minify
RUN CGO_ENABLED=0 go build -o gohome ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/gohome .
COPY --from=build /app/static ./static
EXPOSE 8080
CMD ["./gohome"]

# Licensing information for the included software is available in the LICENSE and NOTICE files in the project repository.