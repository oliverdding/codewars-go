package kata

import "fmt"

func HumanReadableTime(seconds int) string {
	return fmt.Sprintf("%02d:%02d:%02d", seconds/3600, (seconds%3600)/60, (seconds%3600)%60)
}
