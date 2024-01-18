package main

import (
	"fmt"
	"time"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func main() {
	// Input image file path
	imagePath := "image.png"

	// Get time since unix epoch
	timestamp := time.Now().UnixMilli()

	// Output path
	outputPath := fmt.Sprintf("output_video_%d.mp4", timestamp)

	// Duration of the video
	// https://ffmpeg.org/ffmpeg-utils.html#time-duration-syntax
	videoDuration := "30"

	// Maximum output size of the video file
	maxOutputSizeBytes := 100 * 1000 * 1000 // 100MB

	fmt.Println(timestamp)

	var ffmpegArgs []ffmpeg_go.KwArgs = []ffmpeg_go.KwArgs{
		{"t": videoDuration},
		{"format": "mp4"},
		{"fs": maxOutputSizeBytes},
	}

	// Initialize ffmpeg command
	// docs: https://ffmpeg.org/ffmpeg.html#Main-options
	ffmpeg := ffmpeg_go.Input(imagePath).Output(outputPath, ffmpegArgs...)

	err := ffmpeg.Run()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Video generated successfully!")
}
