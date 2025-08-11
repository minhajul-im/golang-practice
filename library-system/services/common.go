package services

import (
	"math/rand"
	"time"
)

var Random = rand.New(rand.NewSource(time.Now().UnixNano()))
