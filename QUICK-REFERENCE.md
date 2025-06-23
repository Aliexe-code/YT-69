# YT-69 Quick Reference

⚡ **Fast reference for experienced users**

## 🚀 Essential Commands

```bash
# Basic download
yt-69.exe -url "VIDEO_URL"

# Quality + Monster Mode
yt-69.exe -url "VIDEO_URL" -quality 2k --turbo

# Audio only
yt-69.exe -url "VIDEO_URL" -quality audio --turbo

# Playlist
yt-69.exe -url "PLAYLIST_URL" --playlist -quality 1080p --turbo

# Video info
yt-69.exe -url "VIDEO_URL" --info
```

## 📊 Quality Matrix

| Flag | Resolution | Size | Codec | Compatibility |
|------|------------|------|-------|---------------|
| `4k` | 3840x2160 | 100-200MB | H.264 | Universal |
| `2k` | 2304x1440 | 30-60MB | VP9 | Modern |
| `1080p` | 1920x1080 | 15-30MB | H.264 | Universal |
| `720p` | 1280x720 | 5-15MB | H.264 | Universal |
| `480p` | 854x480 | 3-8MB | H.264 | Universal |
| `audio` | N/A | 1-10MB | MP3 | Universal |

## 🛠️ All Flags

### Core
- `-url "URL"` - Video/playlist URL (required)
- `-quality QUALITY` - Video quality (4k, 2k, 1080p, 720p, 480p, audio)
- `-output "PATH"` - Download directory

### Performance  
- `--turbo` - Monster Mode (16 parallel, 50MB chunks)
- `--verbose` - Detailed logging

### Content Type
- `--audio` / `-a` - Audio-only download
- `--video` / `-v` - Video-only (no audio)
- `--playlist` - Download entire playlist

### Information
- `--info` - Show video info without downloading
- `--help` - Show help message

## 🎯 Common Use Cases

### High Quality Archive
```bash
# 4K for archival
yt-69.exe -url "VIDEO_URL" -quality 4k --turbo -output "D:\Archive"
```

### Mobile/Streaming
```bash
# 720p for mobile devices
yt-69.exe -url "VIDEO_URL" -quality 720p --turbo
```

### Music Collection
```bash
# Audio-only for music
yt-69.exe -url "MUSIC_VIDEO_URL" -quality audio --turbo -output "C:\Music"
```

### Batch Playlist
```bash
# Entire playlist in 1080p
yt-69.exe -url "PLAYLIST_URL" --playlist -quality 1080p --turbo
```

### Quick Preview
```bash
# Check what's available
yt-69.exe -url "VIDEO_URL" --info
```

## ⚡ Monster Mode Details

**Normal Mode:**
- 1 connection
- 1MB chunks
- Basic retry logic

**Monster Mode (`--turbo`):**
- 16 parallel connections
- 50MB chunks  
- Advanced retry with exponential backoff
- Bandwidth optimization
- Progress tracking per fragment

## 🔧 Technical Notes

### Format Selection Priority
1. **4K**: H.264 > VP9 > AV01 (fallback to best available)
2. **2K**: VP9 2K > H.264 1080p > fallback
3. **1080p**: H.264 > VP9 > AV01
4. **720p/480p**: H.264 > VP9 > AV01
5. **Audio**: Opus > AAC > MP3

### Codec Compatibility
- **H.264**: All players, all devices
- **VP9**: Modern browsers, VLC, recent players
- **AV01**: Latest browsers only (auto-avoided)

### File Naming
- **Videos**: `{Title}.mp4`
- **Audio**: `{Title}.mp3`
- **Playlists**: `{Playlist}/{Title}.{ext}`

## 🚨 Quick Troubleshooting

| Problem | Solution |
|---------|----------|
| Video won't play | Re-run `install-ffmpeg.bat` |
| Separate video/audio files | Install FFmpeg first |
| Download fails | Try different quality |
| Slow downloads | Use `--turbo` flag |
| Permission error | Check folder permissions |

## 📁 File Structure

```
YT-69/
├── yt-69.exe              # Main executable
├── install-ffmpeg.bat     # Run once for setup
├── ffmpeg/                # Auto-created by installer
└── downloads/             # Default output folder
```

## 🎮 Power User Tips

### Batch Operations
```bash
# Multiple videos (create batch file):
@echo off
yt-69.exe -url "VIDEO1_URL" -quality 1080p --turbo
yt-69.exe -url "VIDEO2_URL" -quality 1080p --turbo
yt-69.exe -url "VIDEO3_URL" -quality 1080p --turbo
```

### Custom Output Structure
```bash
# Organized by quality
yt-69.exe -url "VIDEO_URL" -quality 4k --turbo -output "D:\Videos\4K"
yt-69.exe -url "VIDEO_URL" -quality 1080p --turbo -output "D:\Videos\1080p"
```

### Quality Testing
```bash
# Test all qualities for comparison
yt-69.exe -url "VIDEO_URL" -quality 4k --turbo -output "test\4k"
yt-69.exe -url "VIDEO_URL" -quality 2k --turbo -output "test\2k"  
yt-69.exe -url "VIDEO_URL" -quality 1080p --turbo -output "test\1080p"
```

### Performance Monitoring
```bash
# Verbose mode for debugging
yt-69.exe -url "VIDEO_URL" -quality 2k --turbo --verbose
```

## 🔄 Update Workflow

1. **Download** new `yt-69.exe`
2. **Replace** old executable
3. **Keep** `ffmpeg/` folder
4. **Test** with quick download

## 📊 Expected Performance

### Download Speeds (Monster Mode)
- **100 Mbps**: ~12 MB/s sustained
- **50 Mbps**: ~6 MB/s sustained  
- **25 Mbps**: ~3 MB/s sustained

### File Size Estimates
- **4K (10 min)**: ~500MB-1GB
- **2K (10 min)**: ~200-400MB
- **1080p (10 min)**: ~100-200MB
- **720p (10 min)**: ~50-100MB
- **Audio (10 min)**: ~10-20MB

---

**For detailed documentation, see README.md and INSTALL.md**
