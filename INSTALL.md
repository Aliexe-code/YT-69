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
# Download a video
yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ"

# Download audio only
yt-69.exe -url "https://youtube.com/watch?v=dQw4w9WgXcQ" -audio

# Get help
yt-69.exe --help
```

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
