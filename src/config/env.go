package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	currentDir, _ := os.Getwd()
	envFilePath := filepath.Join(currentDir, ".env")
	godotenv.Load(envFilePath)

}
