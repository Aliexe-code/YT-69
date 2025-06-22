package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"
)

const (
	version = "1.0.0"
	banner  = `
██╗   ██╗████████╗      ██████╗  █████╗
╚██╗ ██╔╝╚══██╔══╝      ██╔══██╗██╔══██╗
 ╚████╔╝    ██║   █████╗██████╔╝╚██████║
  ╚██╔╝     ██║   ╚════╝██╔══██╗ ╚═══██║
   ██║      ██║         ██████╔╝ █████╔╝
   ╚═╝      ╚═╝         ╚═════╝  ╚════╝
YT-69 - YouTube Downloader v%s
`
)

// Config holds the CLI configuration
type Config struct {
	URL         string
	OutputDir   string
	Quality     string
	AudioOnly   bool
	VideoOnly   bool
	Format      string
	Playlist    bool
	Verbose     bool
	ShowInfo    bool
	ListFormats bool
}

// Downloader handles YouTube video downloads
type Downloader struct {
	options Options
	client  *http.Client
}

// Options configures the downloader behavior
type Options struct {
	OutputDir string
	Quality   string
	AudioOnly bool
	VideoOnly bool
	Format    string
	Verbose   bool
}

// VideoInfo contains metadata about a YouTube video
type VideoInfo struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Duration    interface{} `json:"duration"`
	Author      string      `json:"uploader"`
	ViewCount   interface{} `json:"view_count"`
	UploadDate  string      `json:"upload_date"`
	Description string      `json:"description"`
	Thumbnail   string      `json:"thumbnail"`
}

// FormatInfo contains information about available video formats
type FormatInfo struct {
	Format  string `json:"format"`
	Quality string `json:"quality"`
	Size    string `json:"filesize"`
	Ext     string `json:"ext"`
}

func main() {
	config := parseFlags()

	if config.ShowInfo {
		showInfo(config.URL)
		return
	}

	if config.ListFormats {
		listFormats(config.URL)
		return
	}

	downloader := NewDownloader(Options{
		OutputDir: config.OutputDir,
		Quality:   config.Quality,
		AudioOnly: config.AudioOnly,
		VideoOnly: config.VideoOnly,
		Format:    config.Format,
		Verbose:   config.Verbose,
	})

	if config.Playlist {
		err := downloader.DownloadPlaylist(config.URL)
		if err != nil {
			log.Fatalf("Failed to download playlist: %v", err)
		}
	} else {
		err := downloader.DownloadVideo(config.URL)
		if err != nil {
			log.Fatalf("Failed to download video: %v", err)
		}
	}

	fmt.Println("✅ Download completed successfully!")
}

// NewDownloader creates a new YouTube downloader instance
func NewDownloader(opts Options) *Downloader {
	if opts.OutputDir == "" {
		opts.OutputDir = "./downloads"
	}
	if opts.Quality == "" {
		opts.Quality = "best"
	}
	if opts.Format == "" {
		opts.Format = "mp4"
	}

	return &Downloader{
		options: opts,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// DownloadVideo downloads a single YouTube video
func (d *Downloader) DownloadVideo(videoURL string) error {
	// Check if yt-dlp is available
	if !d.isYtDlpAvailable() {
		return d.downloadWithFallback(videoURL)
	}

	return d.downloadWithYtDlp(videoURL)
}

// DownloadPlaylist downloads all videos in a YouTube playlist
func (d *Downloader) DownloadPlaylist(playlistURL string) error {
	if !d.isYtDlpAvailable() {
		return fmt.Errorf("playlist download requires yt-dlp to be installed")
	}

	return d.downloadPlaylistWithYtDlp(playlistURL)
}

// GetVideoInfo retrieves metadata about a YouTube video
func (d *Downloader) GetVideoInfo(videoURL string) (*VideoInfo, error) {
	if d.isYtDlpAvailable() {
		return d.getVideoInfoWithYtDlp(videoURL)
	}
	return d.getVideoInfoWithAPI(videoURL)
}

// ListFormats lists available formats for a video
func (d *Downloader) ListFormats(videoURL string) ([]FormatInfo, error) {
	if !d.isYtDlpAvailable() {
		return nil, fmt.Errorf("listing formats requires yt-dlp to be installed")
	}
	return d.listFormatsWithYtDlp(videoURL)
}

// isYtDlpAvailable checks if yt-dlp is installed and available
func (d *Downloader) isYtDlpAvailable() bool {
	_, err := exec.LookPath("yt-dlp")
	return err == nil
}

// isFFmpegAvailable checks if ffmpeg is installed and available
func (d *Downloader) isFFmpegAvailable() bool {
	_, err := exec.LookPath("ffmpeg")
	return err == nil
}

// downloadWithYtDlp downloads using yt-dlp (preferred method)
func (d *Downloader) downloadWithYtDlp(videoURL string) error {
	args := []string{
		"--no-warnings",
		"--quiet",
		"--progress",
		"-o", filepath.Join(d.options.OutputDir, "%(title)s.%(ext)s"),
	}

	// Add format selection (always specify format for proper audio/video selection)
	args = append(args, "-f", d.buildFormatSelector())

	// Add format conversion if needed
	if d.options.AudioOnly {
		// For audio-only downloads, handle format conversion
		if isAudioFormat(d.options.Format) && d.options.Format != "m4a" && d.options.Format != "webm" {
			// User wants specific audio format (mp3, wav, etc.)
			if d.isFFmpegAvailable() {
				args = append(args, "--extract-audio", "--audio-format", d.options.Format)
			} else {
				// Only show this message if user specifically requested a format that needs conversion
				if d.options.Verbose {
					fmt.Printf("⚠️  Note: Audio will be downloaded in native format (M4A/WebM)\n")
					fmt.Printf("   Install ffmpeg for %s conversion: choco install ffmpeg\n", strings.ToUpper(d.options.Format))
				}
			}
		}
	} else if d.options.Format != "mp4" && isVideoFormat(d.options.Format) {
		// Video format conversion
		args = append(args, "--recode-video", d.options.Format)
	}

	args = append(args, videoURL)

	// Only show command in verbose mode
	if d.options.Verbose {
		fmt.Printf("🔧 Running: yt-dlp %s\n", strings.Join(args, " "))
	}

	cmd := exec.Command("yt-dlp", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// downloadPlaylistWithYtDlp downloads a playlist using yt-dlp
func (d *Downloader) downloadPlaylistWithYtDlp(playlistURL string) error {
	args := []string{
		"--no-warnings",
		"--quiet",
		"--progress",
		"-o", filepath.Join(d.options.OutputDir, "%(playlist_title)s/%(title)s.%(ext)s"),
		"--yes-playlist",
	}

	// Add quality and format options
	if d.options.Quality != "best" {
		args = append(args, "-f", d.buildFormatSelector())
	}

	if d.options.AudioOnly {
		args = append(args, "--extract-audio", "--audio-format", d.options.Format)
	}

	args = append(args, playlistURL)

	cmd := exec.Command("yt-dlp", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// getVideoInfoWithYtDlp gets video info using yt-dlp
func (d *Downloader) getVideoInfoWithYtDlp(videoURL string) (*VideoInfo, error) {
	args := []string{
		"--dump-json",
		"--no-warnings",
		videoURL,
	}

	cmd := exec.Command("yt-dlp", args...)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get video info: %v", err)
	}

	var info VideoInfo
	if err := json.Unmarshal(output, &info); err != nil {
		return nil, fmt.Errorf("failed to parse video info: %v", err)
	}

	return &info, nil
}

// listFormatsWithYtDlp lists available formats using yt-dlp
func (d *Downloader) listFormatsWithYtDlp(videoURL string) ([]FormatInfo, error) {
	args := []string{
		"--list-formats",
		"--dump-json",
		"--no-warnings",
		videoURL,
	}

	cmd := exec.Command("yt-dlp", args...)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list formats: %v", err)
	}

	var formats []FormatInfo
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		var format FormatInfo
		if err := json.Unmarshal([]byte(line), &format); err == nil {
			formats = append(formats, format)
		}
	}

	return formats, nil
}

// buildFormatSelector builds the format selector string for yt-dlp
func (d *Downloader) buildFormatSelector() string {
	if d.options.AudioOnly {
		// Force audio-only download
		return "bestaudio[ext=m4a]/bestaudio[ext=webm]/bestaudio"
	}
	if d.options.VideoOnly {
		// Force video-only download (no audio)
		return "bestvideo[ext=mp4]/bestvideo"
	}

	// Video + Audio downloads
	switch strings.ToLower(d.options.Quality) {
	case "worst":
		return "worst[ext=mp4]/worst"
	case "best", "":
		return "best[ext=mp4]/best"
	default:
		// For specific resolutions like "720p", "1080p"
		height := strings.TrimSuffix(strings.ToLower(d.options.Quality), "p")
		return fmt.Sprintf("best[height<=%s][ext=mp4]/best[height<=%s]", height, height)
	}
}

// downloadWithFallback downloads using a fallback method when yt-dlp is not available
func (d *Downloader) downloadWithFallback(videoURL string) error {
	fmt.Println("⚠️  yt-dlp not found. Using fallback method...")
	fmt.Println("📥 For best experience, install yt-dlp:")
	showInstallationInstructions()

	// Extract video ID from URL
	videoID, err := d.extractVideoID(videoURL)
	if err != nil {
		return err
	}

	// Use a simple fallback method (this is a simplified example)
	return d.downloadWithSimpleMethod(videoID)
}

// extractVideoID extracts the video ID from a YouTube URL
func (d *Downloader) extractVideoID(videoURL string) (string, error) {
	// Regular expressions for different YouTube URL formats
	patterns := []string{
		`(?:youtube\.com/watch\?v=|youtu\.be/|youtube\.com/embed/)([a-zA-Z0-9_-]{11})`,
		`youtube\.com/watch\?.*v=([a-zA-Z0-9_-]{11})`,
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(videoURL)
		if len(matches) > 1 {
			return matches[1], nil
		}
	}

	return "", fmt.Errorf("could not extract video ID from URL: %s", videoURL)
}

// downloadWithSimpleMethod implements a basic download method
func (d *Downloader) downloadWithSimpleMethod(videoID string) error {
	fmt.Printf("🎬 Attempting to download video ID: %s\n", videoID)
	fmt.Println("⚠️  Note: Fallback method has limited functionality")
	fmt.Println("   Install yt-dlp for full features and better quality")

	return fmt.Errorf("fallback download method not fully implemented. Please install yt-dlp")
}

// getVideoInfoWithAPI gets basic video info using YouTube's API or web scraping
func (d *Downloader) getVideoInfoWithAPI(videoURL string) (*VideoInfo, error) {
	videoID, err := d.extractVideoID(videoURL)
	if err != nil {
		return nil, err
	}

	return &VideoInfo{
		ID:          videoID,
		Title:       "Video Title (install yt-dlp for full info)",
		Duration:    "Unknown",
		Author:      "Unknown",
		ViewCount:   "Unknown",
		UploadDate:  "Unknown",
		Description: "Install yt-dlp for complete video information",
	}, nil
}

// showInstallationInstructions shows how to install yt-dlp
func showInstallationInstructions() {
	fmt.Println("\n📦 How to install yt-dlp:")

	switch runtime.GOOS {
	case "windows":
		fmt.Println("Windows:")
		fmt.Println("  Option 1 - Using pip:")
		fmt.Println("    pip install yt-dlp")
		fmt.Println("  Option 2 - Using Chocolatey:")
		fmt.Println("    choco install yt-dlp")
		fmt.Println("  Option 3 - Using Scoop:")
		fmt.Println("    scoop install yt-dlp")

	case "darwin":
		fmt.Println("macOS:")
		fmt.Println("  Option 1 - Using Homebrew:")
		fmt.Println("    brew install yt-dlp")
		fmt.Println("  Option 2 - Using pip:")
		fmt.Println("    pip install yt-dlp")

	case "linux":
		fmt.Println("Linux:")
		fmt.Println("  Option 1 - Using pip:")
		fmt.Println("    pip install yt-dlp")
		fmt.Println("  Option 2 - Using package manager:")
		fmt.Println("    sudo apt install yt-dlp  # Ubuntu/Debian")
		fmt.Println("    sudo pacman -S yt-dlp    # Arch Linux")
		fmt.Println("    sudo dnf install yt-dlp  # Fedora")

	default:
		fmt.Println("  Using pip (universal):")
		fmt.Println("    pip install yt-dlp")
	}

	fmt.Println("\n🔗 More info: https://github.com/yt-dlp/yt-dlp#installation")
}

func parseFlags() Config {
	var config Config

	// Define flags
	flag.StringVar(&config.URL, "url", "", "YouTube video or playlist URL (required)")
	flag.StringVar(&config.URL, "u", "", "YouTube video or playlist URL (shorthand)")

	flag.StringVar(&config.OutputDir, "output", "./downloads", "Output directory for downloaded files")
	flag.StringVar(&config.OutputDir, "o", "./downloads", "Output directory (shorthand)")

	flag.StringVar(&config.Quality, "quality", "best", "Video quality (best, worst, 720p, 480p, etc.)")
	flag.StringVar(&config.Quality, "q", "best", "Video quality (shorthand)")

	flag.StringVar(&config.Format, "format", "mp4", "Output format (mp4, webm, mkv for video; mp3, m4a, wav for audio)")
	flag.StringVar(&config.Format, "f", "mp4", "Output format (shorthand)")

	flag.BoolVar(&config.AudioOnly, "audio", false, "Download audio only")
	flag.BoolVar(&config.AudioOnly, "a", false, "Download audio only (shorthand)")

	flag.BoolVar(&config.VideoOnly, "video", false, "Download video only (no audio)")
	flag.BoolVar(&config.VideoOnly, "v", false, "Download video only (shorthand)")

	flag.BoolVar(&config.Playlist, "playlist", false, "Download entire playlist")
	flag.BoolVar(&config.Playlist, "p", false, "Download playlist (shorthand)")

	flag.BoolVar(&config.Verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&config.ShowInfo, "info", false, "Show video information without downloading")
	flag.BoolVar(&config.ListFormats, "list-formats", false, "List available formats for the video")

	// Custom usage function
	flag.Usage = func() {
		fmt.Printf(banner, version)
		fmt.Println("A fast and flexible YouTube downloader - video or audio!")
		fmt.Println("\nUsage:")
		fmt.Printf("  %s [options] -url <YouTube_URL>\n\n", os.Args[0])
		fmt.Println("Examples:")
		fmt.Printf("  # Download video (best quality)\n")
		fmt.Printf("  %s -url https://youtube.com/watch?v=dQw4w9WgXcQ\n\n", os.Args[0])
		fmt.Printf("  # Download specific video quality\n")
		fmt.Printf("  %s -u https://youtube.com/watch?v=dQw4w9WgXcQ -q 720p\n\n", os.Args[0])
		fmt.Printf("  # Download audio only (M4A/WebM)\n")
		fmt.Printf("  %s -u https://youtube.com/watch?v=dQw4w9WgXcQ -a\n\n", os.Args[0])
		fmt.Printf("  # Download audio as MP3 (requires ffmpeg)\n")
		fmt.Printf("  %s -u https://youtube.com/watch?v=dQw4w9WgXcQ -a -f mp3\n\n", os.Args[0])
		fmt.Printf("  # Download playlist\n")
		fmt.Printf("  %s -u https://youtube.com/playlist?list=... -p\n\n", os.Args[0])
		fmt.Printf("  # Get video info\n")
		fmt.Printf("  %s -u https://youtube.com/watch?v=dQw4w9WgXcQ --info\n", os.Args[0])
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
		fmt.Println("\nSupported Qualities:")
		fmt.Println("  best, worst, 144p, 240p, 360p, 480p, 720p, 1080p, 1440p, 2160p")
		fmt.Println("\nSupported Formats:")
		fmt.Println("  Video: mp4, webm, mkv, avi, mov")
		fmt.Println("  Audio: mp3, m4a, wav, aac, ogg (mp3/wav require ffmpeg)")
		fmt.Println("\nFeatures:")
		fmt.Println("  ✅ Clean downloads (no thumbnails or metadata)")
		fmt.Println("  ✅ Works without ffmpeg (native formats)")
		fmt.Println("  ✅ MP3 conversion with ffmpeg installed")
		fmt.Println("  ✅ Flexible for both video and audio downloads")
	}

	// Parse flags
	flag.Parse()

	// Show banner
	fmt.Printf(banner, version)

	// Validate required URL
	if config.URL == "" {
		fmt.Println("❌ Error: YouTube URL is required")
		flag.Usage()
		os.Exit(1)
	}

	// Validate URL format
	if !isValidYouTubeURL(config.URL) {
		fmt.Println("❌ Error: Invalid YouTube URL format")
		fmt.Println("   Supported formats:")
		fmt.Println("   - https://youtube.com/watch?v=VIDEO_ID")
		fmt.Println("   - https://youtu.be/VIDEO_ID")
		fmt.Println("   - https://youtube.com/playlist?list=PLAYLIST_ID")
		os.Exit(1)
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(config.OutputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Convert to absolute path
	absPath, err := filepath.Abs(config.OutputDir)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}
	config.OutputDir = absPath

	// Validate conflicting flags
	if config.AudioOnly && config.VideoOnly {
		fmt.Println("❌ Error: Cannot specify both -audio and -video flags")
		os.Exit(1)
	}

	// Set appropriate format based on audio/video flags
	if config.AudioOnly && !isAudioFormat(config.Format) {
		config.Format = "mp3" // Default audio format
	}
	if config.VideoOnly && !isVideoFormat(config.Format) {
		config.Format = "mp4" // Default video format
	}

	return config
}

func isValidYouTubeURL(url string) bool {
	url = strings.ToLower(url)
	return strings.Contains(url, "youtube.com") ||
		   strings.Contains(url, "youtu.be") ||
		   strings.Contains(url, "m.youtube.com")
}

func isAudioFormat(format string) bool {
	audioFormats := []string{"mp3", "m4a", "wav", "aac", "ogg"}
	for _, f := range audioFormats {
		if strings.ToLower(format) == f {
			return true
		}
	}
	return false
}

func isVideoFormat(format string) bool {
	videoFormats := []string{"mp4", "webm", "mkv", "avi", "mov"}
	for _, f := range videoFormats {
		if strings.ToLower(format) == f {
			return true
		}
	}
	return false
}

func showInfo(url string) {
	fmt.Println("🔍 Fetching video information...")

	downloader := NewDownloader(Options{Verbose: false})
	info, err := downloader.GetVideoInfo(url)
	if err != nil {
		log.Fatalf("Failed to get video info: %v", err)
	}

	fmt.Printf("\n📺 Video Information:\n")
	fmt.Printf("Title: %s\n", info.Title)
	fmt.Printf("Duration: %v\n", formatValue(info.Duration))
	fmt.Printf("Author: %s\n", info.Author)
	fmt.Printf("View Count: %v\n", formatValue(info.ViewCount))
	fmt.Printf("Upload Date: %s\n", info.UploadDate)
	fmt.Printf("Description: %s\n", truncateString(info.Description, 200))
}

func formatValue(v interface{}) string {
	if v == nil {
		return "Unknown"
	}
	return fmt.Sprintf("%v", v)
}

func listFormats(url string) {
	fmt.Println("📋 Fetching available formats...")

	downloader := NewDownloader(Options{Verbose: true})
	formats, err := downloader.ListFormats(url)
	if err != nil {
		log.Fatalf("Failed to list formats: %v", err)
	}

	fmt.Printf("\n📋 Available Formats:\n")
	for _, format := range formats {
		fmt.Printf("Format: %s | Quality: %s | Size: %s\n",
			format.Format, format.Quality, format.Size)
	}
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
