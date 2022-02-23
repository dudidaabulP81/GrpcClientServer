package internal

import (
	"fmt"
	"time"
)

func DateToString(tm time.Time) string {
	//return tm.Format("2022-02-22")
	year, month, day := tm.Date()
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}
