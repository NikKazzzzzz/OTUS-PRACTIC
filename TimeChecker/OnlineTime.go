package timeghecker

import (
	"errors"

	"github.com/beevik/ntp"
)

func GetCurrentTime(server string) (string, error) {
	time, err := ntp.Time(server)
	if err != nil {
		return "", errors.New("failed to fetch time: " + err.Error())
	}
	return time.Format("2006-01-02 15:04:05"), nil
}
