package databases

import (
	"flag"
)

var (
	redisHost     = flag.String("redisHost", "localhost", "Redis host")
	redisPort     = flag.Int("redisPort", 6379, "Redis port")
	redisPassword = flag.String("redisPass", "", "Redis password")
	redisUsername = flag.String("redisUsername", "", "Redis username")
	redisDB       = flag.Int("redisDB", 0, "Redis db")
)
