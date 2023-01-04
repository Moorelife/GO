package system

import (
	"testing"
)

func TestNewSystemEmpty(t *testing.T) {
	want := System{"", "", nil}
	got := *NewSystem("", "", nil)
	if got.architecture != want.architecture ||
		got.operatingSystem != want.operatingSystem ||
		len(got.drives) != len(want.drives) {
		t.Errorf("NewSystem() want %v, got %v", want, got)
	}
}
