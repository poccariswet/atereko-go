package main

import (
	"context"
	"fmt"
	"os/exec"
)

type ffmpeg struct {
	*exec.Cmd
}

func (f *ffmpeg) setArgs(args ...string) {
	f.Args = append(f.Args, args...)
}

func (f *ffmpeg) setDir(dir string) {
	f.Dir = dir
}

func (f *ffmpeg) run(output string) error {
	f.setArgs(output)
	return f.Run()
}

func newffmpeg(ctx context.Context) (*ffmpeg, error) {
	cmdPath, err := exec.LookPath("ffmpeg")
	if err != nil {
		return nil, err
	}

	return &ffmpeg{exec.CommandContext(
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

func main() {
	ctx := context.Background()

	inputTime := format_hhmmss(10)
	output := "sample.mp4"

	f, err := newffmpeg(ctx)
	if err != nil {
		panic(err)
	}

	f.setArgs(
		"-i", "sample/girl_shout.mp4",
		"-itsoffset", inputTime, "-i",
		"sample/shout.mp3", "-c:v",
		"copy", "-map", "0:v:0", "-map", "1:a:0",
	)

	f.run(output)
}
