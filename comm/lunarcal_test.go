package comm

import (
	"testing"
)

func TestCal_2009_02_02(t *testing.T) {
	ld := ToLunarDate("2009-02-02")
	if ld.Format("2006-01-02") != "2009-01-08" {
		t.Error("lunar calendar error")
	}
}
func TestCal_2010_02_02(t *testing.T) {
	ld := ToLunarDate("2010-02-02")
	if ld.Format("2006-01-02") != "2009-12-19" {
		t.Error("lunar calendar error")
	}
}
func TestCal_2017_02_02(t *testing.T) {
	ld := ToLunarDate("2017-02-02")
	if ld.Format("2006-01-02") != "2017-01-06" {
		t.Error("lunar calendar error")
	}
}
func TestCal_2017_09_26(t *testing.T) {
	ld := ToLunarDate("2017-09-26")
	if ld.Format("2006-01-02") != "2017-08-07" {
		t.Error("lunar calendar error")
	}
}
