# YT-69 Installation Guide

🚀 **Simple installation guide for YT-69 YouTube Downloader**

## 📋 System Requirements

- **Operating System**: Windows 10/11 (64-bit)
- **Internet Connection**: Broadband recommended (Monster Mode requires high-speed)
- **Disk Space**: 200MB for application + space for downloads
- **Permissions**: Standard user (no admin rights required)

## 📥 Download Files

### Option 1: GitHub Release (Recommended)
1. Go to the [YT-69 GitHub Releases](https://github.com/YOUR_USERNAME/yt-69/releases)
2. Download the latest release ZIP file
3. Extract to your desired folder (e.g., `C:\Tools\YT-69\`)

### Option 2: Direct Download
Download these files to the same folder:
```
yt-69.exe              # Main executable (required)
install-ffmpeg.bat     # FFmpeg installer (required)
README.md              # Documentation (optional)
INSTALL.md             # This guide (optional)
```

## ⚙️ Installation Steps

### Step 1: Create Installation Folder
```bash
# Create a folder for YT-69 (example locations):
C:\Tools\YT-69\
C:\Users\YourName\YT-69\
D:\Software\YT-69\
```

### Step 2: Extract Files
1. Extract all downloaded files to your chosen folder
2. Ensure `yt-69.exe` and `install-ffmpeg.bat` are in the same directory

### Step 3: Install FFmpeg (Critical!)
```bash
# Navigate to your YT-69 folder and run:
install-ffmpeg.bat
```

**What this does:**
- Downloads portable FFmpeg (no admin rights needed)
- Installs to local `ffmpeg/` folder
- Enables automatic video+audio merging
- Required for single-file downloads

### Step 4: Test Installation
```bash
# Test with a short video:
yt-69.exe -url "https://www.youtube.com/watch?v=dQw4w9WgXcQ" -quality 720p --turbo
```

## 🎯 Folder Structure After Installation

```
YT-69/
├── yt-69.exe              # Main executable
├── install-ffmpeg.bat     # FFmpeg installer
├── ffmpeg/                # FFmpeg installation (created after step 3)
│   ├── ffmpeg.exe
│   ├── ffprobe.exe
│   └── ...
├── downloads/             # Default download folder (auto-created)
├── README.md              # Documentation
└── INSTALL.md             # This installation guide
```

## 🔧 Configuration Options

### Default Download Location
By default, videos download to `downloads/` folder in the same directory as `yt-69.exe`.

To change the download location:
```bash
# Download to custom folder:
yt-69.exe -url "VIDEO_URL" -output "C:\MyVideos"

# Download to current directory:
yt-69.exe -url "VIDEO_URL" -output "."
```

### Adding to PATH (Optional)
To use `yt-69` from anywhere in Command Prompt:

1. **Copy installation path** (e.g., `C:\Tools\YT-69\`)
2. **Open System Properties** → Advanced → Environment Variables
3. **Edit PATH variable** → Add your YT-69 folder path
4. **Restart Command Prompt**

Now you can use:
```bash
# From anywhere:
yt-69 -url "VIDEO_URL" -quality 2k --turbo
```

## ✅ Verification

### Test All Quality Levels
```bash
# Test 1080p download:
yt-69.exe -url "https://www.youtube.com/watch?v=dQw4w9WgXcQ" -quality 1080p --turbo

# Test audio download:
yt-69.exe -url "https://www.youtube.com/watch?v=dQw4w9WgXcQ" -quality audio --turbo

# Test video info:
yt-69.exe -url "https://www.youtube.com/watch?v=dQw4w9WgXcQ" --info
```

### Expected Results
- **Video downloads**: Single MP4 files with video+audio
- **Audio downloads**: MP3 files
- **File sizes**: Vary by quality (see README.md for typical sizes)
- **Playback**: Files should play in any media player

## 🚨 Troubleshooting

### FFmpeg Installation Failed
```bash
# Symptoms: Downloads create separate video/audio files
# Solution: Re-run FFmpeg installer
install-ffmpeg.bat

# Alternative: Manual FFmpeg installation
# 1. Download FFmpeg from https://ffmpeg.org/download.html
# 2. Extract to ffmpeg/ folder in YT-69 directory
```

### Downloads Fail
```bash
# Check internet connection
ping youtube.com

# Test with different video
yt-69.exe -url "https://www.youtube.com/watch?v=9bZkp7q19f0" -quality 720p

# Check video URL is valid
yt-69.exe -url "YOUR_VIDEO_URL" --info
```

### Permission Errors
```bash
# Ensure you have write permissions to installation folder
# Try installing to user directory instead:
C:\Users\YourName\YT-69\
```

### Antivirus Warnings
Some antivirus software may flag the executable:
1. **Add exception** for YT-69 folder in your antivirus
2. **Whitelist** yt-69.exe specifically
3. **Download from official GitHub** releases only

## 🔄 Updates

### Updating YT-69
1. **Download** latest release from GitHub
2. **Replace** old `yt-69.exe` with new version
3. **Keep** existing `ffmpeg/` folder and settings
4. **Test** with a quick download

### Updating FFmpeg
```bash
# Re-run installer to get latest FFmpeg:
install-ffmpeg.bat
```

## 🗑️ Uninstallation

### Complete Removal
1. **Delete** YT-69 installation folder
2. **Remove** PATH entry (if added)
3. **Delete** downloaded videos (if desired)

### Keep Downloads Only
1. **Move** videos from `downloads/` folder to safe location
2. **Delete** YT-69 installation folder

## 💡 Tips for Best Experience

### Performance Tips
1. **Use SSD** for installation and downloads (faster file operations)
2. **Close other downloads** when using Monster Mode
3. **Stable internet** required for high-quality downloads

### Storage Tips
1. **Monitor disk space** - 4K videos can be 100-200MB each
2. **Regular cleanup** of downloads folder
3. **External drive** for large video collections

### Usage Tips
1. **Start with 720p** to test before downloading 4K
2. **Use --info** to check available qualities
3. **Audio-only** for music and podcasts saves space

---

## 🆘 Support

If you encounter issues:
1. **Check this guide** for common solutions
2. **Read README.md** for usage examples
3. **Visit GitHub repository** for latest updates
4. **Check GitHub Issues** for known problems

**Happy downloading! 🎬**
