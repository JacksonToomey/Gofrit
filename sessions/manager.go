package sessions

// SessionManager manages all active sessions
type SessionManager struct {
	sessions     map[int]*Session
	sessionIndex int
}

const startIndex = 1

var manager *SessionManager

func newSessionManager() *SessionManager {
	m := &SessionManager{}
	m.sessions = make(map[int]*Session)
	m.sessionIndex = startIndex
	return m
}

// GetSessionManager returns the session manager, or creates on if none exist
func GetSessionManager() *SessionManager {
	if manager == nil {
		manager = newSessionManager()
	}
	return manager
}

// GetSession returns the session with the provided id, otherwise nil
func (m *SessionManager) GetSession(sessionID int) *Session {
	sess, ok := m.sessions[sessionID]

	if !ok {
		return nil
	}

	return sess
}

// AddSession adds a new empty session to the manager
func (m *SessionManager) AddSession() *Session {
	sess := NewSession(m.sessionIndex)
	m.sessions[m.sessionIndex] = sess
	m.sessionIndex++
	return sess
}

// DeleteSession removes the session from the manager
func (m *SessionManager) DeleteSession(sessionID int) {
	delete(m.sessions, sessionID)
}
