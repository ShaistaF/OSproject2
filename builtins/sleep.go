package sleep

import "time"

type Sleeper interface {
    Sleep(duration time.Duration)
}

type RealSleeper struct{}

func (s RealSleeper) Sleep(duration time.Duration) {
    time.Sleep(duration)
}

var sleeper Sleeper = RealSleeper{}

func SleepFor(duration time.Duration) {
    sleeper.Sleep(duration)
}
