package comm

import (
	"testing"
)

func TestString(t *testing.T) {
	source := "我爱中国"
	str := Substring(source, 2, 3)
	if str != "中" {
		t.Error("substring error")
	}
}
