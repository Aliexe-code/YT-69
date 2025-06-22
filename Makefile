# YT-69 - YouTube Downloader CLI Makefile

# Build the application
build:
	@echo "Building YT-69 YouTube downloader..."
	@go build -o yt-69.exe main.go

# Run the application
run:
	@echo "Usage: make run URL=<youtube_url> [OPTIONS]"
	@echo "Example: make run URL=https://youtube.com/watch?v=dQw4w9WgXcQ"
	@powershell -Command "if ('$(URL)' -eq '') { go run main.go --help } else { go run main.go -url '$(URL)' $(OPTIONS) }"

# Test the application
test:
	@echo "Running tests..."
	@go test -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@powershell -Command "if (Test-Path yt-69.exe) { Remove-Item yt-69.exe }"

# Show help
help:
	@echo "YT-69 - YouTube Downloader CLI - Available commands:"
	@echo ""
	@echo "  make build                    - Build the executable"
	@echo "  make run URL=<url>           - Run with a YouTube URL"
	@echo "  make test                    - Run tests"
	@echo "  make clean                   - Remove built executable"
	@echo "  make help                    - Show this help"
	@echo ""
	@echo "Examples:"
	@echo "  make build"
	@echo "  make run URL=https://youtube.com/watch?v=dQw4w9WgXcQ"
	@echo "  make run URL=https://youtube.com/watch?v=dQw4w9WgXcQ OPTIONS=\"-quality 720p -audio\""

# Install dependencies (yt-dlp)
install-deps:
	@echo "Installing yt-dlp..."
	@powershell -Command "try { pip install yt-dlp; Write-Output 'yt-dlp installed successfully!' } catch { Write-Output 'Failed to install yt-dlp. Please install manually.' }"

.PHONY: build run test clean help install-deps
