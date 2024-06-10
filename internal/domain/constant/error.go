package constant

// LoginFailedError error on failed login
type LoginFailedError struct{}

func (LoginFailedError) Error() string {
	return "Email and Password cannot be found"
}
