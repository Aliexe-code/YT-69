# YT-69 Installation Guide

## Simple Installation Steps

1. Create a folder (let's call it "X")
2. Download these 3 files from the pre-release:
   - The .exe file
   - The install file
   - FFmpeg

3. Download Python
4. Inside folder X, create another folder named "ffmpeg"
5. Paste the ffmpeg.exe file you downloaded into this ffmpeg folder

Your folder structure should look like this:
```
Folder-X
  ├── yt-69.exe
  ├── install
  └── ffmpeg/
      └── ffmpeg.exe
```
Run the install file (python should be installed)
## How to Use

1. Open PowerShell in the location of folder X

2. For video downloads:
```bash
./yt-69.exe -url "video-url" -quality 720p
```
You can choose quality: 720p, 1080p, 2k, 4k

3. For audio downloads:
```bash
./yt-69.exe -url "video-url" -audio
```

## Examples

```bash
# Download 720p video
./yt-69.exe -url "link" -quality 720p

# Download 1080p video
./yt-69.exe -url "link" -quality 1080p

# Download 2K video
./yt-69.exe -url "link" -quality 2k

# Download 4K video
./yt-69.exe -url "link" -quality 4k

# Download audio only
./yt-69.exe -url "link" -audio
```
