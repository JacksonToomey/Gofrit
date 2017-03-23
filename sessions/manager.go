package sessions

// SessionManager manages all active sessions
type SessionManager struct {
	sessions     map[string]*Session
	sessionIndex int
}

const startIndex = 0

var manager *SessionManager

func newSessionManager() *SessionManager {
	m := &SessionManager{}
	m.sessions = make(map[string]*Session)
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
func (m *SessionManager) GetSession(sessionID string) *Session {
	sess, ok := m.sessions[sessionID]

	if !ok {
		return nil
	}

	return sess
}

// AddSession adds a new empty session to the manager
func (m *SessionManager) AddSession(sessionID string) *Session {
	sess, ok := m.sessions[sessionID]

	if !ok {
		sess = NewSession(sessionID)
		m.sessions[sessionID] = sess
	}

	return sess
}

// DeleteSession removes the session from the manager
func (m *SessionManager) DeleteSession(sessionID string) {
	delete(m.sessions, sessionID)
}
