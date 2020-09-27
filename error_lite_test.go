package error_lite

import "testing"

func TestNew(t *testing.T) {
	want := "Test Error"
	if got := New(want); got.Error() != want {
		t.Errorf("New() = %q, want %q", got, want)
	}
}
