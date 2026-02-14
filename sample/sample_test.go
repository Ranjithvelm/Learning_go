package sample

import "testing"

func TestDouble(t *testing.T) {
	retVal := double(2)
	if retVal != 4 {
		t.Errorf("double(2)=%v, want 4", retVal)
	}
}

func TestDoubleTable(t *testing.T) {
	table := []struct {
		input int
		want  int
	}{
		{0, 0},
		{1, 2},
		{2, 4},
		{3, 6},
	}
	for _, data := range table {
		result := double(data.input)
		if result != data.want {
			t.Errorf("double(%v)=%v, want %v", data.input, result, data.want)
		}
	}
}

func TestSample(t *testing.T) {
	t.Fail()
	t.Errorf("msg")
	t.FailNow()
	t.Fatalf("msg")
	t.Logf("msg")
}
