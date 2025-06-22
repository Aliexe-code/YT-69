# 🚀 Quick Installation Guide - YT-69

## Option 1: Download & Run (Easiest)

### Windows Users:
1. **Download** `yt-69.exe`
   - Click on the file in this repository → Click "Download" button
2. **Download** `install.bat`
   - Click on the file in this repository → Click "Download" button
3. **Put both files in the same folder**
4. **Double-click** `install.bat` (this installs yt-dlp automatically)
5. **Done!** You can now use the tool

```bash
# Example usage:
yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ"
```

---

## Option 2: Manual Installation

### Step 1: Download the Tool
- Download `yt-69.exe` from this repository

### Step 2: Install yt-dlp
**Windows:**
```bash
pip install yt-dlp
```

**Mac:**
```bash
brew install yt-dlp
```

**Linux:**
```bash
pip install yt-dlp
```

### Step 3: (Optional) Install ffmpeg for MP3 conversion
**Windows:**
```bash
choco install ffmpeg
# or
scoop install ffmpeg
```

**Mac:**
```bash
brew install ffmpeg
```

**Linux:**
```bash
sudo apt install ffmpeg  # Ubuntu/Debian
```

---

## Quick Test

```bash
# Download a video (best quality available)
yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ"

# 🎬 Download 4K Ultra HD quality
yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" -quality 4k

# 🎬 Download 2K Quad HD quality
yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" -quality 2k

# 🎬 Download 1080p Full HD quality
yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" -quality 1080p

# Download audio only
yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" -audio

# 🚀 TURBO MODE - Maximum download speed for large videos
yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" --turbo

# 🔥 ULTIMATE COMBO - 4K + Turbo Mode
yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" -quality 4k --turbo

# Get help
yt-69.exe --help
```

---

## 🚀 TURBO MODE - For Large Videos

**Your friend wants maximum download speed?** Use turbo mode!

```bash
# Enable turbo mode for any download
yt-69.exe -url "VIDEO_URL" --turbo

# Turbo mode + audio download
yt-69.exe -url "VIDEO_URL" --turbo -audio

# Turbo mode + specific quality
yt-69.exe -url "VIDEO_URL" --turbo -quality 1080p
```

**What Turbo Mode Does:**
- ✅ Downloads 8 fragments simultaneously
- ✅ Uses 10MB chunks for faster transfer
- ✅ Aggressive retry settings
- ✅ Optimized buffer sizes
- ✅ Maximum bandwidth utilization

**Quality Options:**
- 🎬 **4K/2160p** - Ultra HD (best for large screens)
- 🎬 **2K/1440p** - Quad HD (great balance of quality/size)
- 🎬 **1080p** - Full HD (standard high quality)
- 📺 **720p** - HD (good for most uses)
- 🎯 **best** - Automatically picks highest available

**Perfect for:**
- Large video files (1GB+)
- High-speed internet connections
- When you want downloads ASAP
- 4K/2K content for premium viewing

---

## Troubleshooting

**"yt-dlp not found" error:**
```bash
pip install yt-dlp
```

**Need help?**
```bash
yt-69.exe --help
```

---

## That's it! 🎉

For detailed usage examples and advanced features, see the main [README.md](README.md) file.
