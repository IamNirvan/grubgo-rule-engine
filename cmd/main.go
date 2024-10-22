package main

import (
	"fmt"

	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Yeah buddy!")

	config := config.LoadConfig()
	log.Debugf("config loaded: %v", config)

}
