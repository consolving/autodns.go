package autodns

import "time"

const (
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
)

func Time(value string) (time.Time, error) {
	return time.Parse(DDMMYYYYhhmmss, value)
}

func TimeInLocation(value string, location *time.Location) (time.Time, error) {
	return time.ParseInLocation(DDMMYYYYhhmmss, value, location)
}

func FormatTime(time time.Time) string {
	return time.Format(DDMMYYYYhhmmss)
}
