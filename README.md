# YT-69 - YouTube Downloader

🎬 **High-performance YouTube video and audio downloader with automatic quality selection and ffmpeg integration.**

## ✨ Features

- 🎯 **Multiple Quality Options**: 4K, 2K, 1080p, 720p, 480p
- 🎵 **Audio-Only Downloads**: Extract MP3 audio from videos
- ⚡ **Monster Mode**: Maximum internet bandwidth utilization
- 🔧 **Auto-Merging**: Automatic video+audio combination
- 📱 **Single File Output**: No separate video/audio files
- 🎬 **Universal Compatibility**: Works with all media players
- 📋 **Playlist Support**: Download entire playlists
- 🚀 **Portable**: No installation required

## 🚀 Quick Start

### 1. Download Files
```bash
# Download these files from GitHub:
yt-69.exe              # Main executable
install-ffmpeg.bat     # FFmpeg installer (run once)
```

### 2. Install FFmpeg (One-time setup)
```bash
# Run this once to enable automatic video+audio merging:
install-ffmpeg.bat
```

### 3. Start Downloading
```bash
# Download 2K video with audio:
yt-69.exe -url "https://www.youtube.com/watch?v=VIDEO_ID" -quality 2k --turbo

# Download audio only:
yt-69.exe -url "https://www.youtube.com/watch?v=VIDEO_ID" -quality audio --turbo
```

## 📖 Usage Guide

### Basic Video Downloads

```bash
# Download best quality available
yt-69.exe -url "VIDEO_URL"

# Download specific quality with Monster Mode
yt-69.exe -url "VIDEO_URL" -quality 2k --turbo

# Download to specific folder
yt-69.exe -url "VIDEO_URL" -output "C:\Downloads\Videos"
```

### Quality Options

| Quality | Resolution | Typical Size | Command |
|---------|------------|--------------|---------|
| **4K** | 3840x2160 | 100-200MB | `-quality 4k` |
| **2K** | 2304x1440 | 30-60MB | `-quality 2k` |
| **1080p** | 1920x1080 | 15-30MB | `-quality 1080p` |
| **720p** | 1280x720 | 5-15MB | `-quality 720p` |


### Audio Downloads

```bash
# Method 1: Using quality flag
yt-69.exe -url "VIDEO_URL" -quality audio --turbo

# Method 2: Using audio flag
yt-69.exe -url "VIDEO_URL" --audio --turbo

# Both create MP3 files (1-10MB depending on length)
```

### Playlist Downloads

```bash
# Download entire playlist
yt-69.exe -url "PLAYLIST_URL" --playlist -quality 1080p --turbo

# Download playlist audio only
yt-69.exe -url "PLAYLIST_URL" --playlist --audio --turbo
```

## 🛠️ Command Reference

### Required Parameters
- `-url "VIDEO_URL"` - YouTube video or playlist URL

### Quality Parameters
- `-quality 4k` - Download 4K quality (when available)
- `-quality 2k` - Download 2K quality
- `-quality 1080p` - Download 1080p quality
- `-quality 720p` - Download 720p quality
- `-quality 480p` - Download 480p quality
- `-quality audio` - Download audio only (MP3)

### Performance Parameters
- `--turbo` - Enable Monster Mode (maximum bandwidth)
- `--verbose` - Show detailed download information

### Output Parameters
- `-output "PATH"` - Specify download directory
- `--playlist` - Download entire playlist

### Audio/Video Parameters
- `--audio` or `-a` - Download audio only
- `--video` or `-v` - Download video only (no audio)

### Information Parameters
- `--info` - Show video information without downloading
- `--help` - Show help message

## 📁 File Structure

```
YT-69/
├── yt-69.exe              # Main executable
├── install-ffmpeg.bat     # FFmpeg installer
├── README.md              # This documentation
├── INSTALL.md             # Installation guide
└── downloads/             # Default download folder
```

## 🎯 Examples

### Download 2K Video
```bash
yt-69.exe -url "https://www.youtube.com/watch?v=dQw4w9WgXcQ" -quality 2k --turbo
# Output: Rick Astley - Never Gonna Give You Up.mp4 (44MB)
```

### Download Audio Only
```bash
yt-69.exe -url "https://www.youtube.com/watch?v=dQw4w9WgXcQ" -quality audio --turbo
# Output: Rick Astley - Never Gonna Give You Up.mp3 (6.7MB)
```

### Download Playlist
```bash
yt-69.exe -url "https://www.youtube.com/playlist?list=PLAYLIST_ID" --playlist -quality 1080p --turbo
# Output: Multiple videos in playlist folder
```

### Check Video Info
```bash
yt-69.exe -url "https://www.youtube.com/watch?v=VIDEO_ID" --info
# Shows available qualities and video information
```

## ⚡ Monster Mode

Monster Mode (`--turbo`) maximizes your internet bandwidth:
- **16 parallel downloads** for faster speeds
- **50MB chunk sizes** for optimal throughput
- **Automatic retry** on connection issues
- **Speed optimization** for your connection

```bash
# Normal download
yt-69.exe -url "VIDEO_URL" -quality 2k

# Monster Mode (faster)
yt-69.exe -url "VIDEO_URL" -quality 2k --turbo
```

## 🔧 Technical Details

### Supported Formats
- **Video**: MP4 (H.264, VP9 codecs)
- **Audio**: MP3, M4A, WebM
- **Containers**: MP4, WebM

### Quality Fallback
If requested quality isn't available:
1. **4K** → Falls back to best available (2K/1080p)
2. **2K** → Falls back to 1080p
3. **1080p** → Falls back to 720p
4. **720p** → Falls back to 480p

### Codec Compatibility
- **H.264**: Universal compatibility (all players)
- **VP9**: Modern browsers and players
- **AV01**: Latest browsers (auto-avoided for compatibility)

## 🚨 Troubleshooting

### Video Won't Play
```bash
# Solution: Re-run install-ffmpeg.bat
install-ffmpeg.bat

# Then re-download the video
yt-69.exe -url "VIDEO_URL" -quality 1080p --turbo
```

### Download Fails
```bash
# Check video URL is correct
yt-69.exe -url "VIDEO_URL" --info

# Try different quality
yt-69.exe -url "VIDEO_URL" -quality 720p --turbo
```

### Slow Downloads
```bash
# Enable Monster Mode
yt-69.exe -url "VIDEO_URL" -quality 2k --turbo

# Check internet connection
# Monster Mode requires stable high-speed internet
```

## 📋 Requirements

- **Windows 10/11** (64-bit)
- **Internet Connection** (high-speed recommended for Monster Mode)
- **FFmpeg** (auto-installed via install-ffmpeg.bat)
- **Disk Space** (varies by video quality and length)

## 🎉 Tips & Tricks

### Best Practices
1. **Run install-ffmpeg.bat first** for automatic video+audio merging
2. **Use Monster Mode** for faster downloads on high-speed internet
3. **Check video info** before downloading large files
4. **Use 1080p** for best balance of quality and file size

### Quality Recommendations
- **4K**: For premium displays and archival
- **2K**: For high-quality viewing on modern screens
- **1080p**: Best balance for most users
- **720p**: For mobile devices and limited storage
- **Audio**: For music and podcasts

---

**Made with ❤️ for the YouTube downloading community**

*For support and updates, visit our GitHub repository*
