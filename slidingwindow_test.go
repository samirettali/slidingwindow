package slidingwindow

import "testing"

// Contains tests
func TestContainsEmpty(t *testing.T) {
	window, err := New(5)
	if err != nil {
		t.Error("Could not initialize MovingWindow")
	}

	result := window.Contains(1)
	expected := false
	if result != expected {
		t.Errorf("Contains(1) failed, expected %v, got %v", expected, result)
	}
}

func TestContains(t *testing.T) {
	window, err := New(5)
	if err != nil {
		t.Error("Could not initialize MovingWindow")
	}

	window.Add(1)
	window.Add(2)
	window.Add(3)
	window.Add(4)
	window.Add(5)

	result := window.Contains(5)
	expected := true
	if result != expected {
		t.Errorf("Contains(5) failed, expected %v, got %v", expected, result)
	}
}

func TestContainsFull(t *testing.T) {
	window, err := New(5)
	if err != nil {
		t.Error("Could not initialize MovingWindow")
	}

	window.Add(1)
	window.Add(2)
	window.Add(3)

	result := window.Contains(1)
	expected := true
	if result != expected {
		t.Errorf("Contains(1) failed, expected %v, got %v", expected, result)
	}
}

func TestContainsMoved(t *testing.T) {
	window, err := New(5)
	if err != nil {
		t.Error("Could not initialize MovingWindow")
	}

	window.Add(1)
	window.Add(2)
	window.Add(3)
	window.Add(4)
	window.Add(5)
	window.Add(6)

	result := window.Contains(1)
	expected := false
	if result != expected {
		t.Errorf("Contains(1) failed, expected %v, got %v", expected, result)
	}
}

// Len tests
func TestLenEmpty(t *testing.T) {
	window, err := New(5)
	if err != nil {
		t.Error("Could not initialize MovingWindow")
	}

	result := window.Len()
	expected := 0
	if result != expected {
		t.Errorf("Len() failed, expected %v, got %v", expected, result)
	}
}

func TestLen(t *testing.T) {
	window, err := New(5)
	if err != nil {
		t.Error("Could not initialize MovingWindow")
	}

	window.Add(1)

	result := window.Len()
	expected := 1
	if result != expected {
		t.Errorf("Len() failed, expected %v, got %v", expected, result)
	}
}

func TestLenFull(t *testing.T) {
	window, err := New(5)
	if err != nil {
		t.Error("Could not initialize MovingWindow")
	}

	window.Add(1)
	window.Add(2)
	window.Add(3)
	window.Add(4)
	window.Add(5)

	result := window.Len()
	expected := 5
	if result != expected {
		t.Errorf("Len() failed, expected %v, got %v", expected, result)
	}
}

func TestLenMoved(t *testing.T) {
	window, err := New(5)
	if err != nil {
		t.Error("Could not initialize MovingWindow")
	}

	window.Add(1)
	window.Add(2)
	window.Add(3)
	window.Add(4)
	window.Add(5)
	window.Add(6)

	result := window.Len()
	expected := 5
	if result != expected {
		t.Errorf("Len() failed, expected %v, got %v", expected, result)
	}
}

// String tests
func TestStringEmpty(t *testing.T) {
	window, err := New(5)
	if err != nil {
		t.Error("Could not initialize MovingWindow")
	}

	result := window.String()
	expected := "[]"
	if result != expected {
		t.Errorf("String() failed, expected %v, got %v", expected, result)
	}
}

func TestString(t *testing.T) {
	window, err := New(5)
	if err != nil {
		t.Error("Could not initialize MovingWindow")
	}

	window.Add(1)
	window.Add(2)
	window.Add(3)

	result := window.String()
	expected := "[1 2 3]"
	if result != expected {
		t.Errorf("String() failed, expected %v, got %v", expected, result)
	}
}

func TestStringFull(t *testing.T) {
	window, err := New(5)
	if err != nil {
		t.Error("Could not initialize MovingWindow")
	}

	window.Add(1)
	window.Add(2)
	window.Add(3)
	window.Add(4)
	window.Add(5)

	result := window.String()
	expected := "[1 2 3 4 5]"
	if result != expected {
		t.Errorf("String() failed, expected %v, got %v", expected, result)
	}
}

func TestStringMoved(t *testing.T) {
	window, err := New(5)
	if err != nil {
		t.Error("Could not initialize MovingWindow")
	}

	window.Add(1)
	window.Add(2)
	window.Add(3)
	window.Add(4)
	window.Add(5)
	window.Add(6)

	result := window.String()
	expected := "[2 3 4 5 6]"
	if result != expected {
		t.Errorf("String() failed, expected %v, got %v", expected, result)
	}
}
