package util

import "strings"

func NormalizePhoneNumber(phone string) string {
	phone = strings.TrimSpace(phone)
	phone = strings.TrimPrefix(phone, "+")

	if strings.HasPrefix(phone, "0") {
		phone = "62" + phone[1:]
	}

	return phone
}
