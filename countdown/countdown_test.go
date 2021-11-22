package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpyCountdoenOperations struct {
	Calls []string
}

func (s *SpyCountdoenOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdoenOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type Spytime struct {
	durationSlept time.Duration
}

func (s *Spytime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdoenOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdoenOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &Spytime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
