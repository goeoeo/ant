package validation

// Error show the error
type ValidError struct {
	Message, Name, Field string
}
// String Returns the Message.
func (e *ValidError) String() string {
	if e == nil {
		return ""
	}
	return e.Message
}

// Implement Error interface.
// Return e.String()
func (e *ValidError) Error() string { return e.String() }