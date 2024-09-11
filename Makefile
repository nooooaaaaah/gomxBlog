.PHONY: all prepare build-go build-css run run-go watch watch-go watch-refresh watch-css clean stop

# Ensure directories are created
prepare:
	@mkdir -p tmp bin

# Define the entirety of the build process
all: prepare build-go build-css

# Compile Go binaries
build-go:
	@echo "Building Go binaries..."
	@go build -o ./bin/blog ./cmd/main.go

# Compile Tailwind CSS
build-css:
	@echo "Compiling Tailwind CSS..."
	@npx tailwindcss -i ./ui/static/css/input.css -o ./ui/static/css/site.css --config ./configs/tailwind.config.js --minify || { echo "CSS compilation failed"; exit 1; }

build-docker:
	@echo "Building Docker container..."
	@docker build -t gomex . || { echo "Docker build failed"; exit 1; }

push-docker:
	@echo "Pushing Docker image to Docker Hub..."
	@docker push gomex:latest || { echo "Docker push failed"; exit 1; }
	# The general command to start the application


run: watch

# Specific command to start the Go application
run-go:
	@echo "Running Go application..."
	@./bin/main & echo $$! > .pidfile

# Watches all relevant files for changes
watch: prepare watch-go watch-refresh watch-css

# Watches for UI changes to reload the browser
watch-refresh:
	@echo "Starting RefreshMeDaddy for live browser reloading..."
	@RefreshMeDaddy -p 6900 -w ./ui -v & echo $$! >> .pidfile
	@if [ $$? -eq 0 ]; then \
					echo "RefreshMeDaddy started on port 6900..."; \
	else \
					echo "Failed to start RefreshMeDaddy"; \
	fi

# Watches Go files for changes
watch-go:
	@echo "Watching Go files for changes..."
	@air -c ./configs/.air.toml & echo $$! >> .pidfile
	@echo "Air started..."

# Watches CSS files for changes
watch-css:
	@echo "Watching CSS files for changes..."
	@npm run watch-css & echo $$! >> .pidfile
	@echo "Tailwind CSS watch started..."

# Clean up builds and logs
clean:
	@echo "Cleaning up..."
	@rm -rf ./bin/* ./ui/static/css/site.css tmp/* .pidfile
	@echo "Clean completed."

# Stops all processes started for development
stop:
	@echo "Stopping all processes..."
	@kill $$(cat .pidfile) 2>/dev/null || true
	@rm -f .pidfile
	@kill $$(lsof -ti:4200,6900) 2>/dev/null || true
	@echo "All development processes stopped."
