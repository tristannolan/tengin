package tengin

type Config struct {
	Title     string
	TickRate  float64
	FrameRate float64
}

func NewDefaultConfig() *Config {
	return &Config{
		Title:     "",
		TickRate:  60,
		FrameRate: 60,
	}
}
