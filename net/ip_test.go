package net

import "testing"

func TestIPv4ToUint(t *testing.T) {
	res, err := IPv4ToUint("125.71.135.247")
	if err != nil {
		t.Error("转化125.71.135.247为uint64失败,原因:", err.Error())
	} else if res != uint64(2101839863) {
		t.Errorf("转化125.74.135.247为uint64失败,转化后的值为:%d,期望值为:%d", res, 2101839863)
	}
}

func TestUintToIPv4(t *testing.T) {
    if ip := UintToIPv4(uint64(2101839863));ip != "125.71.135.247" {
        t.Error("转化2101839863为IP失败,转化后的值为:%s,期望值为:",ip,"125.71.135.247")
    }
}

