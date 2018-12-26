package chronos

import (
	"fmt"
	"time"
)

type TimeInterval struct {
	Start time.Time
	End   time.Time
}

func New() *TimeInterval {
	return &TimeInterval{
		Start: time.Now(),
	}
}

func (t *TimeInterval) Finish() {
	t.End = time.Now()
}

func (t *TimeInterval) StringSec() string {
	if t.End.IsZero() {
		return "not_finished"
	}
	return fmt.Sprintf("%.2f", t.End.Sub(t.Start).Seconds())
}

func (t *TimeInterval) StringMilisec() string {
	if t.End.IsZero() {
		return "not_finished"
	}
	return fmt.Sprintf("%d", t.End.Sub(t.Start).Nanoseconds()/int64(time.Millisecond))
}

func (t *TimeInterval) String() string {
	if t.End.IsZero() {
		return "not_finished"
	}
	return fmt.Sprintf("%s", t.End.Sub(t.Start).String())

}
