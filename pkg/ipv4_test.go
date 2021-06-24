package pkg

import "testing"

func TestValidateIpV4(t *testing.T) {
	b := ValidateIpV4("0.0.0.0")
	t.Log(b) // true

	b = ValidateIpV4("192.168.1.1")
	t.Log(b) // true

	b = ValidateIpV4("192.168.1.0/24")
	t.Log(b) // true

	b = ValidateIpV4("192.168.1.256")
	t.Log(b) // true
}
