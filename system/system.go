package system

import "github.com/andust/internal/config"

type System struct {
	name string
	cfg  config.SystemConfig
}

func NewSystem(name string) (*System, error) {
	var cfg config.SystemConfig
	cfg, err := config.InitConfig()
	if err != nil {
		return nil, err
	}
	s := &System{
		name: name,
		cfg:  cfg,
	}

	return s, nil
}
