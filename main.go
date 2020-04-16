package main

import (
	"context"
	"flag"
	"fmt"
	"os/exec"
)

type FFmpeg struct {
	*exec.Cmd
}

func (f *FFmpeg) setArgs(args ...string) {
	f.Args = append(f.Args, args...)
}

func (f *FFmpeg) setDir(dir string) {
	f.Dir = dir
}

func (f *FFmpeg) run(output string) error {
	f.setArgs(output)
	return f.Run()
}

func newFFmpeg(ctx context.Context) (*FFmpeg, error) {
	cmdPath, err := exec.LookPath("ffmpeg")
	if err != nil {
		return nil, err
	}

	return &FFmpeg{exec.CommandContext(
		ctx,
		cmdPath,
	)}, nil
}

func format_hhmmss(target_seconds int) string {
	target_minutes := target_seconds / 60
	seconds := target_seconds % 60
	hour := target_minutes / 60
	minutes := target_minutes % 60

	return fmt.Sprintf("%02d:%02d:%02d", hour, minutes, seconds)
}

var (
	inputVideo string
	inputAudio string
	inputTime  int
	output     string
)

func init() {
	flag.StringVar(&inputVideo, "iv", "sample/girl_shout.mp4", "input video path")
	flag.StringVar(&inputAudio, "ia", "sample/shout.mp3", "input audio path")
	flag.IntVar(&inputTime, "it", 8, "input play time")
	flag.StringVar(&output, "o", "output.mp4", "output video path like ` -o output.mp4`")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	f, err := newFFmpeg(ctx)
	if err != nil {
		panic(err)
	}

	f.setArgs(
		"-i", inputVideo,
		"-itsoffset", format_hhmmss(inputTime),
		"-i", inputAudio,
		"-c:v", "copy", "-map", "0:v:0", "-map", "1:a:0",
	)

	f.run(output)
}
