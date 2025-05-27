package common

import "testing"

func TestTimeAfter(t *testing.T) {
	start := "2020-01-01 00:00:00"
	end := "2020-01-01 00:00:01"
	if !TimeAfter(start, end) {
		t.Error("TimeAfter failed1")
	}else{
		t.Log("TimeAfter success1")
	}
	if TimeAfter(end, start) {
		t.Error("TimeAfter failed2")
	}else{
		t.Log("TimeAfter success2")
	}
}