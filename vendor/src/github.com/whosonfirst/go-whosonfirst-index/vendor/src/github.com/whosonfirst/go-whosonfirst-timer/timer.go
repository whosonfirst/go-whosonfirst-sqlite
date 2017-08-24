package timer

import (
	"fmt"
	"log"
	"time"
)

type TimingCallback func(t Timing)

type TimingConstructor func(d time.Duration) Timing

type Timing interface {
	String() string
	Duration() time.Duration
}

type DefaultTiming struct {
	Timing
	duration time.Duration
}

func NewDefaultTiming(d time.Duration) Timing {

	t := DefaultTiming{
		duration: d,
	}

	return &t
}

func (dt *DefaultTiming) String() string {
	return fmt.Sprintf("TIMER %v", dt.duration)
}

func (dt *DefaultTiming) Duration() time.Duration {
	return dt.duration
}

func NewDefaultCallback(t Timing) {
     log.Println(t.String())
}

type Timer struct {
	Start       time.Time
	Timings     chan Timing
	Constructor TimingConstructor
	Callback    TimingCallback
	Done        chan bool
}

func NewDefaultTimer() (*Timer, error) {

	ch := make(chan Timing)
	con := NewDefaultTiming
	cb := NewDefaultCallback

	return NewTimer(ch, con, cb)
}

func NewTimer(ch chan Timing, con TimingConstructor, cb TimingCallback) (*Timer, error) {

	t1 := time.Now()
	done := make(chan bool)

	t := Timer{
		Start:       t1,
		Timings:     ch,
		Callback:    cb,
		Constructor: con,
		Done:        done,
	}

	go t.poll()

	return &t, nil
}

func (tm *Timer) Stop() {
	t2 := time.Since(tm.Start)
	tm.Timings <- tm.Constructor(t2)
	tm.Done <- true
}

func (tm *Timer) poll() {

	for {

		select {
		case t := <-tm.Timings:
			go tm.Callback(t)
		case <-tm.Done:
			return
		default:
			// pass
		}
	}

}