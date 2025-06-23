@echo off
echo 🔧 Installing ffmpeg for automatic video+audio merging...
echo.

REM Check if ffmpeg is already available
ffmpeg -version >nul 2>&1
if %errorlevel% == 0 (
    echo ✅ ffmpeg is already installed and working!
    echo Your downloads will automatically merge video+audio into single files.
    pause
    exit /b 0
)

echo 📦 Downloading portable ffmpeg (no admin rights needed)...
echo This will install ffmpeg in the current directory.
echo.

REM Create ffmpeg directory
if not exist "ffmpeg" mkdir ffmpeg
cd ffmpeg

REM Download ffmpeg portable version
echo Downloading ffmpeg-master-latest-win64-gpl.zip...
powershell -Command "& {[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; Invoke-WebRequest -Uri 'https://github.com/BtbN/FFmpeg-Builds/releases/download/latest/ffmpeg-master-latest-win64-gpl.zip' -OutFile 'ffmpeg.zip'}"

if %errorlevel% neq 0 (
    echo ❌ Download failed. Trying alternative method...
    echo Please download manually from: https://ffmpeg.org/download.html
    pause
    exit /b 1
)

echo 📂 Extracting ffmpeg...
powershell -Command "Expand-Archive -Path 'ffmpeg.zip' -DestinationPath '.' -Force"

REM Find the extracted directory and copy ffmpeg.exe
for /d %%i in (ffmpeg-master-*) do (
    copy "%%i\bin\ffmpeg.exe" ".\" >nul 2>&1
    copy "%%i\bin\ffprobe.exe" ".\" >nul 2>&1
)

REM Clean up
del ffmpeg.zip >nul 2>&1
for /d %%i in (ffmpeg-master-*) do rmdir /s /q "%%i" >nul 2>&1

cd ..

REM Add ffmpeg to PATH for current session
set PATH=%CD%\ffmpeg;%PATH%

REM Test if ffmpeg works
ffmpeg\ffmpeg.exe -version >nul 2>&1
if %errorlevel% == 0 (
    echo.
    echo ✅ ffmpeg installed successfully!
    echo 📍 Location: %CD%\ffmpeg\ffmpeg.exe
    echo.
    echo 🎬 Now your downloads will automatically merge video+audio into single files!
    echo.
    echo 🚀 Try downloading a video:
    echo yt-69.exe -url "https://youtube.com/watch?v=VIDEO_ID" -quality 2k --turbo
    echo.
    echo 💡 Note: ffmpeg is installed locally in this directory.
    echo    Keep the ffmpeg folder with your yt-69.exe file.
) else (
    echo ❌ Installation failed. Please try manual installation:
    echo 1. Download from: https://ffmpeg.org/download.html
    echo 2. Extract ffmpeg.exe to the same folder as yt-69.exe
)

echo.
pause
