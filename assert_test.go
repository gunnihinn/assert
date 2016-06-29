package assert

import "testing"

func TestAssertInitAssertsOff(t *testing.T) {
	a := Asserter{}
	if a.on {
		t.Errorf("Assertions should be off by default")
	}
}

func TestAssertNew(t *testing.T) {
	a := New()
	if !a.on {
		t.Errorf("New assertions should be on by default")
	}
}

func TestAssertOn(t *testing.T) {
	a := Asserter{}
	a.On()
	if !a.on {
		t.Errorf("On didn't turn assertions on")
	}
}

func TestAssertOff(t *testing.T) {
	a := New()
	a.Off()
	if a.on {
		t.Errorf("Off didn't turn assertions off")
	}
}

func TestAssertOffDoesntPanic(t *testing.T) {
	a := New()
	a.Off()
	a.Assert(false, "")
}

func TestAssertPass(t *testing.T) {
	a := New()
	a.Assert(true, "")
}

func TestAssertFail(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("Should have panicked")
		}
	}()
	a := New()
	a.Assert(false, "")
}
