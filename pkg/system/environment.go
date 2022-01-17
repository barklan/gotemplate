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

func GetInternalEnv() (InternalEnv, bool) {
	v, ok := os.LookupEnv(InternalEnvKey)
	if !ok {
		log.Fatalf("%s not specified\n", InternalEnvKey)
	}
	var iEnv InternalEnv
	switch v {
	case "dev":
		iEnv = DevEnv
	case "prod":
		iEnv = ProdEnv
	default:
		log.Fatalf("%s\n %q not recognized", InternalEnvKey, v)
	}
	log.Printf("%s is set to %q\n", InternalEnvKey, v)

	var inDocker bool
	if os.Getenv("DOCKERIZED") == "true" {
		log.Println("running in container")
		inDocker = true
	}
	return iEnv, inDocker
}
