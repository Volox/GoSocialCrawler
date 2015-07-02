package social

// Credentials repesents the keys to use on the social networks
type Credentials struct {
	Key         string
	Secret      string
	Token       string
	TokenSecret string
}

// AppOnly verifies if the credentials are siutable for App Only authentication
func (c *Credentials) AppOnly() bool {
	if c.Token == "" && c.TokenSecret == "" {
		return true
	}

	return false
}
