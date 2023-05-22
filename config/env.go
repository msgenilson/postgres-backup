package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	//POSTGRES IMPORT
	PG_IMPORT_HOST     = getEnv("PGI_HOST", "")
	PG_IMPORT_PORT     = getEnv("PGI_PORT", "")
	PG_IMPORT_USER     = getEnv("PGI_USER", "")
	PG_IMPORT_PASSWORD = getEnv("PGI_PASSWORD", "")
	PG_IMPORT_DBNAME   = getEnv("PGI_DBNAME", "")
	//VARS
	PATH_DB_BK     = getEnv("PATH_DB_BK", "")
	PATH_PG_DUMP   = getEnv("PATH_PG_DUMP", "")
	LIMIT_DATABASE = getEnv("LIMIT_DATABASE", "5")
)

func getEnv(name string, fallback string) string {
	load()

	value := os.Getenv(name)
	if value != "" {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Variável de ambiente não encontrada :: %v`, name))
}

func load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}
}
