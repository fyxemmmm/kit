package pkg

import (
	"regexp"
	"strings"
)

func ValidateIpV4(ip string) bool {
	addr := strings.Trim(ip, " ")

	if ip == "0.0.0.0" {
		return true
	}

	if match, _ := regexp.MatchString(`^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`, addr); match {
		return true
	}

	// CIDR
	if match, _ := regexp.MatchString("^(?:(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\/([1-9]|[1-2]\\d|3[0-2])$", addr);match {
		return true
	}

	return false
}
