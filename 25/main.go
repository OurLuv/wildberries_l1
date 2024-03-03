package main

import (
	"fmt"
	"time"
)

func sleep(dur time.Duration) {
	// сигнал придёт через 3 секунды
	// до этого программа не продолжит выполнение
	<-time.After(dur)
}

func main() {
	start := time.Now()
	sleep(3 * time.Second)
	fmt.Print(time.Since(start))
}
