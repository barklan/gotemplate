package system

import (
	"log"
	"os"
)

type InternalEnv byte

const (
	DevEnv InternalEnv = iota
	ProdEnv
)

const InternalEnvKey = "INTERNAL_ENVIRONMENT"

func (i InternalEnv) String() string {
	switch i {
	case DevEnv:
		return "Development"
	case ProdEnv:
		return "Production"
	default:
		return "Undefined"
	}
}

func GetInternalEnv() (InternalEnv, bool) {
	var iEnv InternalEnv
	v, _ := os.LookupEnv(InternalEnvKey)
	switch v {
	case "dev":
		iEnv = DevEnv
	case "prod":
		iEnv = ProdEnv
	default:
		iEnv = ProdEnv
	}
	log.Printf("Environment is set to %q\n", iEnv)

	var inDocker bool
	if os.Getenv("DOCKERIZED") == "true" {
		log.Println("running in container")
		inDocker = true
	}

	return iEnv, inDocker
}
