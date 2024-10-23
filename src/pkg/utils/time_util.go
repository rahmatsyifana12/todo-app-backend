package utils

import "time"

func GetTimeNowJakarta() time.Time {
	jakarta, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(jakarta)
}