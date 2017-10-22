// Package dateseq creates a sequence series of daily dates
package dateseq

import (
	"sort"
	"time"
)

// Sequence represents a slice of dates.
type Sequence struct {
	seq      []time.Time
	weekends bool
}

// New returns a Sequence ready for use.
func New() Sequence {
	return Sequence{weekends: true}
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

// Steps creates a date sequence with the specified length of days ending now.
// Calls to Steps will reset the sequence slice to nil before generating an new sequence.
func (s Sequence) Steps(i int) Sequence {
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

// Asc sorts a slice of dates in ascending order, i.e. 2006-01-02 comes before 2006-01-03
func (s Sequence) Asc() Sequence {
	sort.Slice(s.seq, func(i, j int) bool {
		d1 := s.seq[i]
		d2 := s.seq[j]

		// sort by date
		return d1.Before(d2)
	})

	return s
}

// Desc sorts a slice of dates in descending order, i.e. 2006-01-01 comes after 2006-01-02
func (s Sequence) Desc() Sequence {
	sort.Slice(s.seq, func(i, j int) bool {
		d1 := s.seq[i]
		d2 := s.seq[j]

		// sort by date
		return d2.Before(d1)
	})

	return s
}

// Seq returns the sequence slice.
func (s Sequence) Seq() []time.Time {
	return s.seq
}

// String returns the sequence as slice with a simple string repesentation
// of the date in the format YYYY-MM-DD.
func (s Sequence) String() []string {
	var strings = make([]string, len(s.seq))

	for k, v := range s.seq {
		strings[k] = v.Format("2006-01-02")
	}

	return strings
}

// Format returns the sequence as slice with a string repesentation of the date in the specified layout.
func (s Sequence) Format(layout string) []string {
	var strings = make([]string, len(s.seq))

	for k, v := range s.seq {
		strings[k] = v.Format(layout)
	}

	return strings
}
