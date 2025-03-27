.PHONY: install-tailwind-linux
install-tailwind-linux:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
	chmod +x tailwindcss-linux-x64
	mv tailwindcss-linux-x64 tailwindcss

.PHONY: tailwind-build
tailwind-generate:
	./tailwindcss -i ./src/css/custom.css -o ./static/css/style.css

.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./src/css/custom.css -o ./static/css/style.css --watch

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch
