package initialization

import (
	"log"
	"testing"
)

func TestName(t *testing.T) {
	cfg, err := InitCfg("../config/config.yml")
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}

	log.Printf("%#v", cfg)

	if cfg.Http.Port == ""{
		t.Error("port wasnt found")
	}
}
