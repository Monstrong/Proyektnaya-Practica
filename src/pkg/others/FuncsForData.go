package others

import (
	"math/rand"
	"time"
)


func CurrentTime() (now string) {
	now = time.Now().Format("15:04:05")
	return
}

func Randomnum() int {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    return r.Intn(100)
}