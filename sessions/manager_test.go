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

	sess := m.GetSession(1)

	if sess != nil {
		t.Error("Session should be nil")
	}
}

func TestGetSession(t *testing.T) {
	m := GetSessionManager()

	sess := m.GetSession(1)

	if sess != nil {
		t.Error("Session should be nil")
	}

	sess = m.AddSession()

	if sess == nil {
		t.Error("Session should not be nil")
	}

	if sess.GetID() != 1 {
		t.Error("Session id should be 1, got ", sess.GetID())
	}
}

func TestDeleteSession(t *testing.T) {
	m := GetSessionManager()

	sess := m.GetSession(2)

	if sess != nil {
		t.Error("Session should be nil")
	}

	sess = m.AddSession()

	if sess == nil {
		t.Error("Session should not be nil")
	}

	m.DeleteSession(2)

	sess = m.GetSession(2)

	if sess != nil {
		t.Error("Session should be nil")
	}
}
