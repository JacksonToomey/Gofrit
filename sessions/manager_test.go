package sessions

import "testing"

func TestGetSessionManager(t *testing.T) {
	m := GetSessionManager()

	if m == nil {
		t.Error("Unexpected nil value")
	}

	otherM := GetSessionManager()

	if m != otherM {
		t.Error("Managers should be equal")
	}
}

func TestGetSessionNoSessions(t *testing.T) {
	m := GetSessionManager()

	sess := m.GetSession("some_id")

	if sess != nil {
		t.Error("Session should be nil")
	}
}

func TestGetSession(t *testing.T) {
	m := GetSessionManager()

	sess := m.GetSession("some_id")

	if sess != nil {
		t.Error("Session should be nil")
	}

	sess = m.AddSession("some_id")

	if sess == nil {
		t.Error("Session should not be nil")
	}

	other := m.GetSession("some_id")

	if sess != other {
		t.Error("Sessions should be equal")
	}

	m.DeleteSession("some_id")
}

func TestDeleteSession(t *testing.T) {
	m := GetSessionManager()

	sess := m.GetSession("some_id")

	if sess != nil {
		t.Error("Session should be nil")
	}

	sess = m.AddSession("some_id")

	if sess == nil {
		t.Error("Session should not be nil")
	}

	m.DeleteSession("some_id")

	sess = m.GetSession("some_id")

	if sess != nil {
		t.Error("Session should be nil")
	}

}
