package main

import (
	"fmt"
	"os"
	"time"

	// "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func main() {
	CreateVideo()
}

func CreateVideo() {
	// Input image file path
	imagePath := "input/image.png"

	// Input audio file path
	audioPath := "input/audio.wav"

	// Get time since unix epoch
	timestamp := time.Now().UnixMilli()

	// Output path
	outputPath := fmt.Sprintf("output/output_video_%d.mp4", timestamp)

	// Ensure that the output directory exists, if it does not, create it
	if _, err := os.Stat("output"); os.IsNotExist(err) {
		os.Mkdir("output", 0755)
	}

	// Duration of the video
	// https://ffmpeg.org/ffmpeg-utils.html#time-duration-syntax
	// videoDuration := "30"

	// Maximum output size of the video file
	maxOutputSizeBytes := 100 * 1000 * 1000 // 100MB

	var ffmpegInputArgs []ffmpeg_go.KwArgs = []ffmpeg_go.KwArgs{
		{"i": imagePath},    // Input image path
		{"stream_loop": -1}, // Loop the input stream (in this case, the image, an infinite amount of times)
		{"i": audioPath},    // Input audio path
	}

	var ffmpegOutputArgs []ffmpeg_go.KwArgs = []ffmpeg_go.KwArgs{
		{"format": "mp4"},          // output format
		{"fs": maxOutputSizeBytes}, // Maximum output size of the video file
		{"c:v": "libx264"},         // copy video codec
		{"c:a": "aac"},             // copy audio codec
		{"strict": "experimental"}, // experimental flag
		{"b:a": "192k"},            // audio bitrate
		{"pix_fmt": "yuv420p"},     // pixel format
		{"shortest": ""},           // finish encoding when the shortest input stream ends
	}

	// docs: https://ffmpeg.org/ffmpeg.html#Main-options
	// Create video stream
	video := ffmpeg_go.Input(imagePath, ffmpegInputArgs...).Output(outputPath, ffmpegOutputArgs...)

	err := video.Run()

	if err != nil {
		fmt.Println("Error occured:")
		fmt.Println(err)
		return
	}

	fmt.Println("Video generated successfully!")
}
