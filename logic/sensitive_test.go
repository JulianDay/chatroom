package logic

import (
	"chatroom/dfa"
	"testing"
)

func TestFilterSensitive(t *testing.T) {
	dfa.Init([]string{"5hit", "4r5e"})
	content := FilterSensitive("5hitzbc")
	if content != "****zbc" {
		t.Error("FilterSensitive failed")
	}
}
