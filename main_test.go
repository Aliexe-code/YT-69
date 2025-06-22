package main

import (
	"testing"
)

func TestNewDownloader(t *testing.T) {
	opts := Options{
		OutputDir: "./test-downloads",
		Quality:   "720p",
		Format:    "mp4",
		Verbose:   true,
	}

	downloader := NewDownloader(opts)

	if downloader == nil {
		t.Fatal("NewDownloader returned nil")
	}

	if downloader.options.OutputDir != opts.OutputDir {
		t.Errorf("Expected OutputDir %s, got %s", opts.OutputDir, downloader.options.OutputDir)
	}

	if downloader.options.Quality != opts.Quality {
		t.Errorf("Expected Quality %s, got %s", opts.Quality, downloader.options.Quality)
	}

	if downloader.options.Format != opts.Format {
		t.Errorf("Expected Format %s, got %s", opts.Format, downloader.options.Format)
	}
}

func TestNewDownloaderDefaults(t *testing.T) {
	downloader := NewDownloader(Options{})

	if downloader.options.OutputDir != "./downloads" {
		t.Errorf("Expected default OutputDir './downloads', got %s", downloader.options.OutputDir)
	}

	if downloader.options.Quality != "best" {
		t.Errorf("Expected default Quality 'best', got %s", downloader.options.Quality)
	}

	if downloader.options.Format != "mp4" {
		t.Errorf("Expected default Format 'mp4', got %s", downloader.options.Format)
	}
}

func TestExtractVideoID(t *testing.T) {
	downloader := NewDownloader(Options{})

	testCases := []struct {
		url      string
		expected string
		hasError bool
	}{
		{
			url:      "https://youtube.com/watch?v=dQw4w9WgXcQ",
			expected: "dQw4w9WgXcQ",
			hasError: false,
		},
		{
			url:      "https://youtu.be/dQw4w9WgXcQ",
			expected: "dQw4w9WgXcQ",
			hasError: false,
		},
		{
			url:      "https://youtube.com/watch?v=dQw4w9WgXcQ&t=30s",
			expected: "dQw4w9WgXcQ",
			hasError: false,
		},
		{
			url:      "https://m.youtube.com/watch?v=dQw4w9WgXcQ",
			expected: "dQw4w9WgXcQ",
			hasError: false,
		},
		{
			url:      "invalid-url",
			expected: "",
			hasError: true,
		},
		{
			url:      "https://example.com/video",
			expected: "",
			hasError: true,
		},
	}

	for _, tc := range testCases {
		result, err := downloader.extractVideoID(tc.url)

		if tc.hasError {
			if err == nil {
				t.Errorf("Expected error for URL %s, but got none", tc.url)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for URL %s: %v", tc.url, err)
			}
			if result != tc.expected {
				t.Errorf("For URL %s, expected %s, got %s", tc.url, tc.expected, result)
			}
		}
	}
}

func TestBuildFormatSelector(t *testing.T) {
	testCases := []struct {
		options  Options
		expected string
	}{
		{
			options:  Options{AudioOnly: true},
			expected: "bestaudio[ext=m4a]/bestaudio[ext=webm]/bestaudio",
		},
		{
			options:  Options{VideoOnly: true},
			expected: "bestvideo[ext=mp4]/bestvideo",
		},
		{
			options:  Options{Quality: "best"},
			expected: "best[ext=mp4]/best",
		},
		{
			options:  Options{Quality: "worst"},
			expected: "worst[ext=mp4]/worst",
		},
		{
			options:  Options{Quality: "720p"},
			expected: "best[height<=720][ext=mp4]/best[height<=720]",
		},
		{
			options:  Options{Quality: "1080p"},
			expected: "best[height<=1080][ext=mp4]/best[height<=1080]",
		},
	}

	for _, tc := range testCases {
		downloader := NewDownloader(tc.options)
		result := downloader.buildFormatSelector()

		if result != tc.expected {
			t.Errorf("For options %+v, expected %s, got %s", tc.options, tc.expected, result)
		}
	}
}

func TestIsValidYouTubeURL(t *testing.T) {
	testCases := []struct {
		url      string
		expected bool
	}{
		{"https://youtube.com/watch?v=dQw4w9WgXcQ", true},
		{"https://youtu.be/dQw4w9WgXcQ", true},
		{"https://m.youtube.com/watch?v=dQw4w9WgXcQ", true},
		{"https://www.youtube.com/watch?v=dQw4w9WgXcQ", true},
		{"https://example.com/video", false},
		{"invalid-url", false},
		{"", false},
	}

	for _, tc := range testCases {
		result := isValidYouTubeURL(tc.url)
		if result != tc.expected {
			t.Errorf("For URL %s, expected %v, got %v", tc.url, tc.expected, result)
		}
	}
}

func TestIsAudioFormat(t *testing.T) {
	testCases := []struct {
		format   string
		expected bool
	}{
		{"mp3", true},
		{"m4a", true},
		{"wav", true},
		{"aac", true},
		{"ogg", true},
		{"MP3", true}, // Test case insensitive
		{"mp4", false},
		{"webm", false},
		{"invalid", false},
	}

	for _, tc := range testCases {
		result := isAudioFormat(tc.format)
		if result != tc.expected {
			t.Errorf("For format %s, expected %v, got %v", tc.format, tc.expected, result)
		}
	}
}

func TestIsVideoFormat(t *testing.T) {
	testCases := []struct {
		format   string
		expected bool
	}{
		{"mp4", true},
		{"webm", true},
		{"mkv", true},
		{"avi", true},
		{"mov", true},
		{"MP4", true}, // Test case insensitive
		{"mp3", false},
		{"m4a", false},
		{"invalid", false},
	}

	for _, tc := range testCases {
		result := isVideoFormat(tc.format)
		if result != tc.expected {
			t.Errorf("For format %s, expected %v, got %v", tc.format, tc.expected, result)
		}
	}
}

func TestTruncateString(t *testing.T) {
	testCases := []struct {
		input    string
		maxLen   int
		expected string
	}{
		{"Hello World", 20, "Hello World"},
		{"Hello World", 5, "Hello..."},
		{"", 10, ""},
		{"Short", 10, "Short"},
	}

	for _, tc := range testCases {
		result := truncateString(tc.input, tc.maxLen)
		if result != tc.expected {
			t.Errorf("For input '%s' with maxLen %d, expected '%s', got '%s'", 
				tc.input, tc.maxLen, tc.expected, result)
		}
	}
}

// Benchmark tests
func BenchmarkExtractVideoID(b *testing.B) {
	downloader := NewDownloader(Options{})
	url := "https://youtube.com/watch?v=dQw4w9WgXcQ"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = downloader.extractVideoID(url)
	}
}

func BenchmarkNewDownloader(b *testing.B) {
	opts := Options{
		OutputDir: "./downloads",
		Quality:   "720p",
		Format:    "mp4",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewDownloader(opts)
	}
}
