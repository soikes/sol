package web

// Passwords can be JSON unmarshaled (in) but not marshalled (out)
type password string
func (password) MarshalJSON() ([]byte, error) {
	return []byte(`""`), nil
}