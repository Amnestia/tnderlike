package ping

// Service service functionality of ping domain
type Service struct {
}

// Ping health check function
func (svc *Service) Ping() string {
	return "pong"
}
