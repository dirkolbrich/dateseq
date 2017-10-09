// Package dateseq creates a sequence series of daily dates
package dateseq

import (
	"time"
)

// Sequence represents a slice of dates.
type Sequence struct {
	seq      []time.Time
	weekends bool
}

// New returns a Sequence ready for use.
func New() Sequence {
	return Sequence{}
}

// InclWeekends includes Saturday and Sunday into the sequence.
func (s Sequence) InclWeekends() Sequence {
	s.weekends = true
	return s
}

// ExclWeekends excludes Saturday and Sunday from the sequence.
// This is the default setting for weekends.
func (s Sequence) ExclWeekends() Sequence {
	s.weekends = false
	return s
}

// Duration creates a date sequence with the specified length of days ending now.
// Calls to Duration will reset the sequence slice to nil before generating an new sequence.
func (s Sequence) Duration(i int) Sequence {
	// get current date
	t, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))

	// reset the seqeuence slice
	if len(s.seq) != 0 {
		s.seq = []time.Time{}
	}

	for k := 0; k < i; k++ {
		if !s.weekends {
			if (t.Weekday() == 0) || (t.Weekday() == 6) {
				t = t.AddDate(0, 0, -1)
				k--
				continue
			}
		}

		s.seq = append(s.seq, t)
		t = t.AddDate(0, 0, -1)
	}

	return s
}

// Seq returns the sequence slice.
func (s Sequence) Seq() []time.Time {
	return s.seq
}
