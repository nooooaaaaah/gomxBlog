.PHONY: all prepare build-go build-css run run-go watch watch-go watch-css clean stop

LOGFILE := tmp/build.log

# Ensure directories are created
prepare:
	@mkdir -p tmp bin

# Define the entirety of the build process
all: prepare build-go build-css

# Compile Go binaries
build-go:
	@echo "Building Go binaries..." | tee -a $(LOGFILE)
	@go build -o ./bin/main ./src/main.go >> $(LOGFILE) 2>&1 || echo "Go build failed" | tee -a $(LOGFILE)

# Compile Tailwind CSS
build-css:
	@echo "Compiling Tailwind CSS..." | tee -a $(LOGFILE)
	@npx tailwindcss -i ./ui/static/css/input.css -o ./ui/static/css/site.css --config ./configs/tailwind.config.js >> $(LOGFILE) 2>&1 || echo "CSS compilation failed" | tee -a $(LOGFILE)

# The general command to start the application
run: watch

# Specific command to start the Go application
run-go:
	@echo "Running Go application..." | tee -a $(LOGFILE)
	@./bin/main >> $(LOGFILE) 2>&1 & echo $$! > .pidfile

# Watches all relevant files for changes
watch: watch-go watch-css watch-refresh

# Watches for ui changes to reload browser
watch-refresh:
	@echo "Starting RefreshMeDaddy for live browser reloading..." | tee -a $(LOGFILE)
	@RefreshMeDaddy -p 6900 -w ./ui -v >> tmp/refresh-watch.log 2>&1 & echo $$! >> .pidfile
	@echo "RefreshMeDaddy started on port 6900..." | tee -a $(LOGFILE)

# Watches Go files for changes
watch-go:
	@echo "Watching Go files for changes..." | tee -a $(LOGFILE)
	@air -c ./configs/.air.toml >> tmp/go-watch.log 2>&1 & echo $$! >> .pidfile
	@echo "Air started..." | tee -a $(LOGFILE)

# Watches CSS files for changes
watch-css:
	@echo "Watching CSS files for changes..." | tee -a $(LOGFILE)
	@npx tailwindcss --config ./configs/tailwind.config.js -i ./ui/static/css/input.css -o ./ui/static/css/site.css --watch >> tmp/css-watch.log 2>&1 || echo "Failed to watch CSS" >> tmp/css-watch.log echo $$! >> .pidfile
	@echo "Tailwind CSS watch started..." | tee -a $(LOGFILE)

# Clean up builds and logs
clean:
	@echo "Cleaning up..." | tee -a $(LOGFILE)
	@rm -rf ./bin/* ./ui/static/css/site.css tmp/* .pidfile

# Stops all processes started for development
stop:
	@echo "Stopping all processes..." | tee -a $(LOGFILE)
	@kill $$(cat .pidfile) 2>/dev/null || true
	@rm -f .pidfile
	@kill $$(lsof -ti:4200,6900) 2>/dev/null || true
	@echo "All development processes stopped." | tee -a $(LOGFILE)
