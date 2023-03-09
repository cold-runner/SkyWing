package settings

import "time"

type JwtConf struct {
	SigningKey  string        `mapstructure:"signing-key"`
	Issuer      string        `mapstructure:"issuer"`
	ExpiresTime time.Duration `mapstructure:"expiresTimeDurationMinutes"`
	BufferTime  time.Duration `mapstructure:"bufferTimeDurationMinutes"`
}
