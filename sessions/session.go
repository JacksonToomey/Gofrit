package sessions

// Session objects hold individual session data
type Session struct {
	sessionid string
	store     map[string]interface{}
}

// NewSession creates a new session object with a session id
func NewSession(id string) *Session {
	sess := Session{}
	sess.sessionid = id
	sess.store = make(map[string]interface{})
	return &sess
}

// Set sets a session value
func (s *Session) Set(key string, value interface{}) {
	s.store[key] = value
}

// Get gets the session value stored at key, otherwise nil
func (s *Session) Get(key string) interface{} {
	val, ok := s.store[key]
	if !ok {
		return nil
	}
	return val
}

// Delete deletes the value stored at key
func (s *Session) Delete(key string) {
	delete(s.store, key)
}

// GetID returns the session id for this session
func (s *Session) GetID() string {
	return s.sessionid
}
