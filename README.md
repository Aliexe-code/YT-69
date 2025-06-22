# YT-69 - YouTube Downloader CLI

A fast and clean YouTube video downloader CLI tool built with Go. Downloads only video/audio files - no extra clutter!

## 🚀 Features

- **🎬 4K/2K Ultra High Quality** - Download videos in 4K (2160p), 2K (1440p), 1080p and more
- **🚀 Turbo Mode** - Maximum download speed with 8 concurrent fragments
- **Single Video Downloads** - Download any YouTube video
- **Playlist Support** - Download entire playlists at once
- **Audio Extraction** - Download audio-only files in various formats
- **Smart Quality Selection** - Automatic best quality detection with fallbacks
- **Format Conversion** - Convert to different video/audio formats
- **Video Information** - Get detailed metadata without downloading
- **Cross-Platform** - Works on Windows, macOS, and Linux
- **Beautiful CLI** - Intuitive command-line interface with helpful messages

## 📦 Installation

### 🚀 Quick Start
**For easy installation instructions, see [INSTALL.md](INSTALL.md)**

**TL;DR:**
1. Download `yt-69.exe`
2. Run `install.bat` (Windows) or `pip install yt-dlp` (Mac/Linux)
3. Start downloading: `yt-69.exe -url "https://youtube.com/watch?v=VIDEO_ID"`

### Building from Source
```bash
git clone https://github.com/Aliexe-code/YT-69.git
cd YT-69
make build
```

## 🎯 Usage

### Basic Examples

```bash
# Download a video (best quality available)
./yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ"

# 🎬 Download 4K Ultra HD quality
./yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" -quality 4k

# 🎬 Download 2K Quad HD quality
./yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" -quality 2k

# 🎬 Download 1080p Full HD quality
./yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" -quality 1080p

# Download audio only
./yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" -audio -format mp3

# 🚀 TURBO MODE - Maximum download speed
./yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" --turbo

# 🔥 ULTIMATE COMBO - 4K + Turbo Mode
./yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" -quality 4k --turbo

# Download to specific directory
./yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" -output "./my-videos"

# Download entire playlist
./yt-69.exe -url "https://youtube.com/playlist?list=PLrAXtmRdnEQy6nuLMHjMZOz59Oq8VGLrG" -playlist
```

### Advanced Usage

```bash
# Get video information without downloading
./yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" --info

# List available formats
./yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" --list-formats

# Download video only (no audio)
./yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" -video -format mp4

# 2K quality with turbo speed
./yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" -quality 2k --turbo

# Custom resolution
./yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" -quality 900p

# Verbose output for debugging
./yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" -verbose
```

### Using with Makefile

```bash
# Quick download
make run URL="https://youtube.com/watch?v=dQw4w9WgXcQ"

# With options
make run URL="https://youtube.com/watch?v=dQw4w9WgXcQ" OPTIONS="-quality 720p -audio"

# Show help
make run
```

## 📋 Command Line Options

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--url` | `-u` | YouTube video or playlist URL (required) | - |
| `--output` | `-o` | Output directory for downloads | `./downloads` |
| `--quality` | `-q` | Video quality (4k, 2k, 1080p, best, etc.) | `best` |
| `--format` | `-f` | Output format (mp4, mp3, etc.) | `mp4` |
| `--audio` | `-a` | Download audio only | `false` |
| `--video` | `-v` | Download video only (no audio) | `false` |
| `--playlist` | `-p` | Download entire playlist | `false` |
| `--turbo` | `-t` | Enable turbo mode for maximum speed | `false` |
| `--verbose` | - | Enable verbose output | `false` |
| `--info` | - | Show video info without downloading | `false` |
| `--list-formats` | - | List available formats | `false` |

## 🎨 Supported Formats

### Video Formats
- **MP4** (recommended) - Universal compatibility
- **WebM** - Good compression, web-friendly
- **MKV** - High quality, supports multiple streams
- **AVI** - Legacy format, wide compatibility
- **MOV** - Apple's format

### Audio Formats
- **MP3** (recommended) - Universal compatibility
- **M4A** - High quality, Apple devices
- **WAV** - Uncompressed, highest quality
- **AAC** - Good compression, quality balance
- **OGG** - Open source, good compression

### Quality Options

#### 🎬 Ultra High Quality
- **4k/2160p** - 4K Ultra HD (best for large screens and TVs)
- **2k/1440p** - 2K Quad HD (great balance of quality/size)

#### 🎬 High Quality
- **1080p/fhd** - Full HD (standard high quality)
- **720p/hd** - HD (good for most uses)

#### 📺 Standard Quality
- **480p/sd** - Standard Definition
- **360p** - Low quality
- **240p** - Very low quality
- **144p** - Minimum quality

#### 🎯 Smart Options
- **best** - Automatically picks highest available quality
- **worst** - Lowest available quality
- **Custom** - Any resolution like 900p, 1200p, etc.

## � Turbo Mode & 4K Quality

### 🔥 Ultimate Performance Features

**Turbo Mode** - Maximum download speed optimization:
```bash
# Enable turbo mode for any download
yt-69.exe -url "VIDEO_URL" --turbo

# Turbo mode + specific quality
yt-69.exe -url "VIDEO_URL" --turbo -quality 1080p
```

**What Turbo Mode Does:**
- ✅ Downloads 8 fragments simultaneously
- ✅ Uses 10MB chunks for faster transfer
- ✅ Aggressive retry settings (10 retries)
- ✅ Optimized buffer sizes (16KB)
- ✅ Maximum bandwidth utilization
- ✅ Skip unnecessary file operations

**4K/2K Ultra Quality:**
```bash
# 🎬 4K Ultra HD (if available)
yt-69.exe -url "VIDEO_URL" -quality 4k

# 🎬 2K Quad HD (great balance)
yt-69.exe -url "VIDEO_URL" -quality 2k

# 🔥 ULTIMATE COMBO - 4K + Turbo
yt-69.exe -url "VIDEO_URL" -quality 4k --turbo
```

**Perfect For:**
- 🎬 Large video files (1GB+)
- 🚀 High-speed internet connections
- 📺 Premium viewing on large screens
- ⚡ When you want downloads ASAP

## �🛠️ Makefile Commands

| Command | Description |
|---------|-------------|
| `make build` | Build the executable |
| `make run URL=<url>` | Run with a YouTube URL |
| `make test` | Run tests |
| `make clean` | Remove built executable |
| `make help` | Show available commands |
| `make install-deps` | Install yt-dlp dependency |

## 🚨 Troubleshooting

### Common Issues

**"yt-dlp not found" error:**
```bash
# Install yt-dlp
pip install yt-dlp
# or use the Makefile
make install-deps
```

**Permission denied errors:**
```bash
# Make sure output directory is writable
chmod 755 ./downloads
```

**Download fails:**
```bash
# Try with verbose output to see detailed error
./ytdl.exe -url "..." -verbose

# Update yt-dlp to latest version
pip install --upgrade yt-dlp
```

### Getting Help

```bash
# Show help message
./ytdl.exe --help

# Show Makefile commands
make help
```

## 📁 Project Structure

```
ytdl-cli/
├── main.go          # Complete CLI application in one file
├── go.mod           # Go module definition
├── Makefile         # Build automation
├── README.md        # This file
└── ytdl.exe         # Built executable (after make build)
```

## 🔮 Future Features

- [ ] Batch download from file
- [ ] Resume interrupted downloads
- [ ] Download subtitles
- [ ] Custom naming templates
- [ ] Configuration file support
- [ ] Progress bar improvements

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test your changes
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## ⚠️ Legal Notice

This tool is for educational and personal use only. Please respect YouTube's Terms of Service and copyright laws. Only download content you have permission to download.

## 🙏 Acknowledgments

- [yt-dlp](https://github.com/yt-dlp/yt-dlp) - The powerful backend that makes this possible
- [Go](https://golang.org/) - The amazing programming language
- YouTube - For providing the platform (please support content creators!)

---

**Happy downloading! 🎬✨**
