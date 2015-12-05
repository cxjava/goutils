package strutils

import (
	"fmt"
	"testing"
)

func TestSubStr(t *testing.T) {
	fmt.Println("ddddd")
}

func TestSubStringBetween(t *testing.T) {
	str := `ProxyFromEnvironment`
	out := ""
	if r := SubStringBetween("", "abc", "abc"); r != out {
		t.Errorf(`SubStringBetween(%v, %v, %v) = %v, want %v`, "", "abc", "abc", r, out)
	}

	if r := SubStringBetween(str, "", "nil"); r != out {
		t.Errorf(`SubStringBetween(%v, %v, %v) = %v, want %v`, str, "", "nil", r, out)
	}

	if r := SubStringBetween(str, "nil", ""); r != out {
		t.Errorf(`SubStringBetween(%v, %v, %v) = %v, want %v`, str, "nil", "", r, out)
	}

	if r := SubStringBetween(str, "FFF", "From"); r != out {
		t.Errorf(` SubStringBetween(%v, %v, %v) = %v, want %v`, str, "FFF", "From", r, out)
	}

	if r := SubStringBetween(str, "Proxy", "MMMM"); r != out {
		t.Errorf(` SubStringBetween(%v, %v, %v) = %v, want %v`, str, "Proxy", "MMMM", r, out)
	}

	if r := SubStringBetween(str, "Proxy", "Proxy"); r != out {
		t.Errorf(` SubStringBetween(%v, %v, %v) = %v, want %v`, str, "Proxy", "Proxy", r, out)
	}

	if r := SubStringBetween(str, "nment", "nment"); r != out {
		t.Errorf(` SubStringBetween(%v, %v, %v) = %v, want %v`, str, "nment", "nment", r, out)
	}

	out = "From"
	if r := SubStringBetween(str, "Proxy", "Env"); r != out {
		t.Errorf(` SubStringBetween(%v, %v, %v) = %v, want %v`, str, "Proxy", "Env", r, out)
	}

	str = "中文测试rune,;，；。×&&……很好"
	out = ""
	if r := SubStringBetween(str, "中", "中"); r != out {
		t.Errorf(` SubStringBetween(%v, %v, %v) = %v, want %v`, str, "中", "中", r, out)
	}
	if r := SubStringBetween(str, "文测", "文测"); r != out {
		t.Errorf(` SubStringBetween(%v, %v, %v) = %v, want %v`, str, "文测", "文测", r, out)
	}
	if r := SubStringBetween(str, "&&", "&&"); r != out {
		t.Errorf(` SubStringBetween(%v, %v, %v) = %v, want %v`, str, "&&", "&&", r, out)
	}

	out = "测试rune,"
	if r := SubStringBetween(str, "中文", ";"); r != out {
		t.Errorf(` SubStringBetween(%v, %v, %v) = %v, want %v`, str, "中文", ";", r, out)
	}

	out = ";，；。×"
	if r := SubStringBetween(str, "中文测试rune,", "&&……"); r != out {
		t.Errorf(` SubStringBetween(%v, %v, %v) = %v, want %v`, str, "中文测试rune,", "&&……", r, out)
	}

	out = "e,;，；。×&"
	if r := SubStringBetween(str, "试run", "&…"); r != out {
		t.Errorf(` SubStringBetween(%v, %v, %v) = %v, want %v`, str, "试run", "&…", r, out)
	}

	out = "。×&&……很"
	if r := SubStringBetween(str, "；", "好"); r != out {
		t.Errorf(` SubStringBetween(%v, %v, %v) = %v, want %v`, str, "；", "好", r, out)
	}
}
