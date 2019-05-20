package validation

// Error show the error
type ValidError struct {
	Field string //结构体字段名
	Name string //结构体字段对应的名称
	Message string //验证失败的消息提示
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