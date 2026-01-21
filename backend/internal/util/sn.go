package util

import (
	"fmt"
	"math/rand"
	"time"
)

func GenSN() string {
	return time.Now().Format("20060102150405") + fmt.Sprintf("%04d", rand.Intn(10000))
}
