package instance

import "testing"

func TestFloat_Read(t *testing.T) {
	pi := FLOAT.Read("3.14").(Float)
	if pi.Ne(Float(3.14)) {
		t.Errorf("Float Read Fail")
	}
}

func TestFloat_Show(t *testing.T) {
	pi := Float(3.14)
	if pi.Show() != "3.14" {
		t.Errorf("Float Show Fail")
	}
}