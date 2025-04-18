.PHONY: install-tailwind-linux
install-tailwind-linux:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
	chmod +x tailwindcss-linux-x64
	mv tailwindcss-linux-x64 tailwindcss

.PHONY: tailwind-generate
tailwind-generate:
	./tailwindcss -i ./src/css/custom.css -o ./static/css/style.css --minify

.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./src/css/custom.css -o ./static/css/style.css --watch

.PHONY: tailwind-watch-background
tailwind-watch-background:
	@echo "tailwindcss is running in the background. You can stop it with 'pkill tailwindcss'"
	./tailwindcss -i ./src/css/custom.css -o ./static/css/style.css --watch=always &

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: build
build: 
	go build -o bin/goHome ./cmd/api/main.go

.PHONY: run
run: templ-generate tailwind-generate
	go run ./cmd/api/main.go

.PHONY: run-watch
run-watch: tailwind-watch-background
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run ./cmd/api" --proxyport=8081
	