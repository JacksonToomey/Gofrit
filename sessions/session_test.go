package sessions

import "testing"

func TestSessionGetID(t *testing.T) {
	s := NewSession("testid")
	id := s.GetID()
	if id != "testid" {
		t.Error("Expected testid, got", id)
	}
}

func TestSessionGetKeyNoKey(t *testing.T) {
	s := NewSession("testid")
	foo := s.Get("bar")

	if foo != nil {
		t.Error("Expected nil, got", foo)
	}
}

func TestSessionGetKey(t *testing.T) {
	s := NewSession("testid")
	s.Set("bar", "bat")
	foo := s.Get("bar")

	if foo != "bat" {
		t.Error("Expected 'bat', got", foo)
	}
}

func TestSessionDelete(t *testing.T) {
	s := NewSession("testid")
	s.Set("bar", "bat")
	foo := s.Get("bar")

	if foo != "bat" {
		t.Error("Expected 'bat', got", foo)
	}

	s.Delete("bar")

	foo = s.Get("bar")

	if foo != nil {
		t.Error("Expected nil, got", foo)
	}
}
