.PHONY: all prepare build-go build-css run run-go watch watch-go watch-refresh watch-css clean stop

LOGFILE := tmp/build.log

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
	@echo "Compiling Tailwind CSS..." | tee -a $(LOGFILE)
	@npx tailwindcss -i ./ui/static/css/input.css -o ./ui/static/css/site.css --config ./configs/tailwind.config.js --minify >> $(LOGFILE) 2>&1 || { echo "CSS compilation failed" | tee -a $(LOGFILE); exit 1; }

build-docker:
	@echo "Building Docker container..." | tee -a $(LOGFILE)
	@docker build -t gomex . >> $(LOGFILE) 2>&1 || { echo "Docker build failed" | tee -a $(LOGFILE); exit 1; }

push-docker:
	@echo "Pushing Docker image to Docker Hub..." | tee -a $(LOGFILE)
	@docker push gomex:latest >> $(LOGFILE) 2>&1 || { echo "Docker push failed" | tee -a $(LOGFILE); exit 1; }
	# The general command to start the application


run: watch

# Specific command to start the Go application
run-go:
	@echo "Running Go application..." | tee -a $(LOGFILE)
	@./bin/main >> $(LOGFILE) 2>&1 & echo $$! > .pidfile

# Watches all relevant files for changes
watch: prepare watch-go watch-refresh watch-css

# Watches for UI changes to reload the browser
watch-refresh:
	@echo "Starting RefreshMeDaddy for live browser reloading..." | tee -a $(LOGFILE)
	@RefreshMeDaddy -p 6900 -w ./ui -v >> tmp/refresh-watch.log 2>&1 & echo $$! >> .pidfile
	@if [ $$? -eq 0 ]; then \
	    echo "RefreshMeDaddy started on port 6900..." | tee -a $(LOGFILE); \
	else \
	    echo "Failed to start RefreshMeDaddy" | tee -a $(LOGFILE); \
	fi

# Watches Go files for changes
watch-go:
	@echo "Watching Go files for changes..." | tee -a $(LOGFILE)
	@air -c ./configs/.air.toml >> tmp/go-watch.log 2>&1 & echo $$! >> .pidfile
	@echo "Air started..." | tee -a $(LOGFILE)

# Watches CSS files for changes
watch-css:
	@echo "Watching CSS files for changes..." | tee -a $(LOGFILE)
	@npm run watch-css >> tmp/css-watch.log 2>&1 & echo $$! >> .pidfile
	@echo "Tailwind CSS watch started..." | tee -a $(LOGFILE)

# Clean up builds and logs
clean:
	@echo "Cleaning up..." | tee -a $(LOGFILE)
	@rm -rf ./bin/* ./ui/static/css/site.css tmp/* .pidfile
	@echo "Clean completed." | tee -a $(LOGFILE)

# Stops all processes started for development
stop:
	@echo "Stopping all processes..." | tee -a $(LOGFILE)
	@kill $$(cat .pidfile) 2>/dev/null || true
	@rm -f .pidfile
	@kill $$(lsof -ti:4200,6900) 2>/dev/null || true
	@echo "All development processes stopped." | tee -a $(LOGFILE)
