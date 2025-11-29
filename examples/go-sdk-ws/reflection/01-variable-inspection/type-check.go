package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"strconv"
)

func init() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func main() {
	var a uint8 = 2
	var b = 37
	var c = "3.2"
	res := sum(a, b, c)
	log.Info().Any("result", res).Msg("Sum of values")
}

func sum(v ...interface{}) float64 {
	var res = 0.0

	for _, val := range v {
		switch val.(type) {
		case int:
			res += float64(val.(int))
		case int64:
			res += float64(val.(int64))
		case uint8:
			res += float64(val.(uint8))
		case string:
			a, err := strconv.ParseFloat(val.(string), 64)
			if err != nil {
				panic(err)
			}
			res += a
		default:
			log.Info().Any("result", val).Msg("Unsupported type, cannot sum")
		}
	}
	return res
}
