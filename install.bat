@echo off
echo ========================================
echo YT-69 YouTube Downloader - Setup Script
echo ========================================
echo.

echo Checking if Python is installed...
python --version >nul 2>&1
if %errorlevel% neq 0 (
    echo ERROR: Python is not installed or not in PATH
    echo Please install Python from https://python.org
    echo Then run this script again.
    pause
    exit /b 1
)

echo Python found! Installing yt-dlp...
pip install yt-dlp

if %errorlevel% neq 0 (
    echo ERROR: Failed to install yt-dlp
    echo Please try running: pip install yt-dlp
    pause
    exit /b 1
)

echo.
echo Checking if ffmpeg is available...
ffmpeg -version >nul 2>&1
if %errorlevel% neq 0 (
    echo WARNING: ffmpeg not found
    echo For MP3 conversion, install ffmpeg:
    echo   - Using Chocolatey: choco install ffmpeg
    echo   - Using Scoop: scoop install ffmpeg
    echo   - Manual: https://ffmpeg.org/download.html
    echo.
    echo You can still use YT-69 without ffmpeg for M4A/WebM audio downloads.
) else (
    echo ffmpeg found! Full functionality available.
)

echo.
echo ========================================
echo Setup Complete!
echo ========================================
echo.
echo You can now use YT-69:
echo   yt-69.exe -url "https://youtube.com/watch?v=VIDEO_ID"
echo.
echo For help: yt-69.exe --help
echo.
pause
