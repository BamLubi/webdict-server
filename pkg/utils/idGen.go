package utils

import (
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"log"
	"strconv"
)

func IdGen() string {
	flake, err := snowflake.NewSnowflake(int64(0), int64(0))
	if err != nil {
		log.Fatalf("snowflake.NewSnowflake failed with %s\n", err)
	}
	id := flake.NextVal()
	return strconv.Itoa(int(id))[:15]
}
