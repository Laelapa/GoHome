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
	

# For running the server locally in a docker container

.PHONY: docker-build docker-run docker-test docker-stop

docker-build:
	docker build -t gohome:local .

# Run the container with proper environment variables
docker-run:
	docker stop gohome-local 2>/dev/null || true
	docker rm gohome-local 2>/dev/null || true
	docker run -d -p 8080:8080 \
		-e ENVIRONMENT=production \
		-e SERVER_PORT=8080 \
		-e STATIC_DIR=/app/static \
		-e SERVER_SHUTDOWN_TIMEOUT=5 \
		--name gohome-local \
		gohome:local

# View logs from the container
docker-logs:
	docker logs -f gohome-local

# Build and run in one command
docker-test: docker-build docker-run
	@echo "Container running at http://localhost:8080"
	@echo "View logs with: make docker-logs"

# Stop and remove the container
docker-stop:
	docker stop gohome-local 2>/dev/null || true
	docker rm gohome-local 2>/dev/null || true