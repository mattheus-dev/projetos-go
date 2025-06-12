package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d%04d", time.Now().Unix(), rand.Intn(10000))
}