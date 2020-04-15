package main

import "fmt"

func format_hhmmss(target_seconds int) string {
	target_minutes := target_seconds / 60
	seconds := target_seconds % 60
	hour := target_minutes / 60
	minutes := target_minutes % 60

	return fmt.Sprintf("%02d:%02d:%02d", hour, minutes, seconds)
}

func main() {
	fmt.Println(format_hhmmss(90))
}
