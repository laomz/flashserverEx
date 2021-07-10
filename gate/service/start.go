package service

import "fmt"

func StartGate() error {
	if err := loadConfig(); err != nil {
		return fmt.Errorf("StartGate:loadConfig:err=%s", err.Error())
	}

	return nil
}
