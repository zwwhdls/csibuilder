package model

type Config struct {
	Repo         string
	TemplatePath string
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

// GetTemplatePath returns the template path
func (c *Config) GetTemplatePath() string {
	return c.TemplatePath
}

// SetTemplatePath sets the template path.
func (c *Config) SetTemplatePath(templatePath string) error {
	c.TemplatePath = templatePath
	return nil
}
