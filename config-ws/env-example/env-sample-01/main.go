package main

import (
	//"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	log.Info().Msg("godotenv example starts")
}

// This example shows how to load environment variables from a .env file.
func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Err(err).Msg("Error loading .env file")
	//}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	// now do something with s3 or whatever
	log.Info().Msgf("S3 Bucket: %s", s3Bucket)
	log.Info().Msgf("Secret Key: %s", secretKey)
}
