package model

type Config struct {
	Repo string
}

// GetRepository returns the project repository.
func (c *Config) GetRepository() string {
	return c.Repo
}

// SetRepository sets the project repository.
func (c *Config) SetRepository(repository string) error {
	c.Repo = repository
	return nil
}
