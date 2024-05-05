package sleep

import (
    "testing"
    "time"
)

type MockSleeper struct {
    SleepCalled bool
    LastDuration time.Duration
}

func (m *MockSleeper) Sleep(duration time.Duration) {
    m.SleepCalled = true
    m.LastDuration = duration
}

func TestSleepFor(t *testing.T) {
    mock := &MockSleeper{}
    sleeper = mock
    defer func() { sleeper = RealSleeper{} }()

    SleepFor(10 * time.Millisecond)
    if !mock.SleepCalled {
        t.Error("Expected Sleep to be called")
    }
    if mock.LastDuration != 10*time.Millisecond {
        t.Errorf("Expected duration to be 10ms, got %s", mock.LastDuration)
    }
}
