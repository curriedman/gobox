package instance

import "testing"

func TestInteger_Read(t *testing.T) {
	if STATIC.FuncRead("1").(Integer).Ne(Integer(1)) {
		t.Errorf("Integer Read Fail")
	}
}

func TestInteger_Show(t *testing.T) {
	i := Integer(100)
	if i.Show() != "100" {
		t.Errorf("Integer Show Fail")
	}

}
