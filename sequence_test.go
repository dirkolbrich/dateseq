package dateseq

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	var testCases = []struct {
		msg    string
		expSeq Sequence
	}{
		{"testing create sequence",
			Sequence{weekends: true},
		},
	}

	for _, tc := range testCases {
		seq := New()
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v New()\nexpected %#v\nactual   %#v", tc.msg, tc.expSeq, seq)
		}
	}
}

func TestInclWeekends(t *testing.T) {
	var testCases = []struct {
		msg    string
		seq    Sequence
		expSeq Sequence
	}{
		{"testing setting weekends from false to true",
			Sequence{},
			Sequence{weekends: true},
		},
		{"testing setting weekends from already true to true",
			Sequence{weekends: true},
			Sequence{weekends: true},
		},
	}

	for _, tc := range testCases {
		seq := tc.seq.InclWeekends()
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v InclWeekends()\nexpected %#v\nactual   %#v", tc.msg, tc.expSeq, seq)
		}
	}
}

func TestExclWeekends(t *testing.T) {
	var testCases = []struct {
		msg    string
		seq    Sequence
		expSeq Sequence
	}{
		{"testing setting weekends from true to false",
			Sequence{weekends: true},
			Sequence{weekends: false},
		},
		{"testing setting weekends from already false to false",
			Sequence{weekends: false},
			Sequence{weekends: false},
		},
	}

	for _, tc := range testCases {
		seq := tc.seq.ExclWeekends()
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v ExclWeekends()\nexpected %#v\nactual   %#v", tc.msg, tc.expSeq, seq)
		}
	}
}

func TestAsc(t *testing.T) {
	time1, _ := time.Parse("2006-01-02", "2006-01-01")
	time2, _ := time.Parse("2006-01-02", "2006-01-02")
	time3, _ := time.Parse("2006-01-02", "2006-01-03")

	var testCases = []struct {
		msg    string
		seq    Sequence
		expSeq Sequence
	}{
		{"testing ascending sorting with asc sorted entries",
			Sequence{
				seq: []time.Time{time1, time2, time3},
			},
			Sequence{
				seq: []time.Time{time1, time2, time3},
			},
		},
		{"testing ascending sorting with desc sorted entries",
			Sequence{
				seq: []time.Time{time3, time2, time1},
			},
			Sequence{
				seq: []time.Time{time1, time2, time3},
			},
		},
		{"testing ascending sorting with nil entries",
			Sequence{
				seq: []time.Time{},
			},
			Sequence{
				seq: []time.Time{},
			},
		},
	}

	for _, tc := range testCases {
		seq := tc.seq.Asc()
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v Asc()\nexpected %#v\nactual   %#v", tc.msg, tc.expSeq, seq)
		}
	}
}

func TestDesc(t *testing.T) {
	time1, _ := time.Parse("2006-01-02", "2006-01-01")
	time2, _ := time.Parse("2006-01-02", "2006-01-02")
	time3, _ := time.Parse("2006-01-02", "2006-01-03")

	var testCases = []struct {
		msg    string
		seq    Sequence
		expSeq Sequence
	}{
		{"testing descending sorting with desc sorted entries",
			Sequence{
				seq: []time.Time{time3, time2, time1},
			},
			Sequence{
				seq: []time.Time{time3, time2, time1},
			},
		},
		{"testing descending sorting with asc sorted entries",
			Sequence{
				seq: []time.Time{time1, time2, time3},
			},
			Sequence{
				seq: []time.Time{time3, time2, time1},
			},
		},
		{"testing descending sorting with nil entries",
			Sequence{
				seq: []time.Time{},
			},
			Sequence{
				seq: []time.Time{},
			},
		},
	}

	for _, tc := range testCases {
		seq := tc.seq.Desc()
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v Desc()\nexpected %#v\nactual   %#v", tc.msg, tc.expSeq, seq)
		}
	}
}

func TestSeq(t *testing.T) {
	time1, _ := time.Parse("2006-01-02", "2006-01-01")
	time2, _ := time.Parse("2006-01-02", "2006-01-02")
	time3, _ := time.Parse("2006-01-02", "2006-01-03")

	var testCases = []struct {
		msg    string
		seq    Sequence
		expSeq []time.Time
	}{
		{"testing return sequenz with single entry",
			Sequence{
				seq: []time.Time{time1},
			},
			[]time.Time{time1},
		},
		{"testing return sequenz with multiple entries",
			Sequence{
				seq: []time.Time{time1, time2, time3},
			},
			[]time.Time{time1, time2, time3},
		},
		{"testing return sequenz with nil entries",
			Sequence{
				seq: []time.Time{},
			},
			[]time.Time{},
		},
	}

	for _, tc := range testCases {
		seq := tc.seq.Seq()
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v Seq()\nexpected %#v\nactual   %#v", tc.msg, tc.expSeq, seq)
		}
	}
}

func TestString(t *testing.T) {
	time1, _ := time.Parse("2006-01-02", "2006-01-01")
	time2, _ := time.Parse("2006-01-02", "2006-01-02")
	time3, _ := time.Parse("2006-01-02", "2006-01-03")

	var testCases = []struct {
		msg        string
		seq        Sequence
		expStrings []string
	}{
		{"testing return sequence with single entry",
			Sequence{
				seq: []time.Time{time1},
			},
			[]string{"2006-01-01"},
		},
		{"testing return sequence with multiple entries",
			Sequence{
				seq: []time.Time{time1, time2, time3},
			},
			[]string{"2006-01-01", "2006-01-02", "2006-01-03"},
		},
		{"testing return sequence with nil entries",
			Sequence{
				seq: []time.Time{},
			},
			[]string{},
		},
	}

	for _, tc := range testCases {
		seq := tc.seq.String()
		if !reflect.DeepEqual(seq, tc.expStrings) {
			t.Errorf("%v String()\nexpected %#v\nactual   %#v", tc.msg, tc.expStrings, seq)
		}
	}
}

func TestFormat(t *testing.T) {
	time1, _ := time.Parse("2006-01-02", "2006-01-01")

	var testCases = []struct {
		msg        string
		seq        Sequence
		format     string
		expStrings []string
	}{
		{"testing format sequence in \"\" format: ",
			Sequence{
				seq: []time.Time{time1},
			},
			"Mon 01.02.2006",
			[]string{"Sun 01.01.2006"},
		},
		{"testing format sequence in \"\" format: ",
			Sequence{
				seq: []time.Time{time1},
			},
			"January 01. 2006",
			[]string{"January 01. 2006"},
		},
	}

	for _, tc := range testCases {
		seq := tc.seq.Format(tc.format)
		if !reflect.DeepEqual(seq, tc.expStrings) {
			t.Errorf("%v Format(%v)\nexpected %#v\nactual   %#v", tc.msg, tc.format, tc.expStrings, seq)
		}
	}
}

func TestStandardSequenz(t *testing.T) {
	current := time.Now().Format("2006-01-02")
	currentDate, _ := time.Parse("2006-01-02", current)

	var testCases = []struct {
		msg    string
		steps  int
		seq    Sequence
		expSeq Sequence
	}{
		{"testing standard sequence:",
			5,
			Sequence{weekends: true},
			Sequence{
				weekends: true,
				seq: []time.Time{
					currentDate,
					currentDate.AddDate(0, 0, -1),
					currentDate.AddDate(0, 0, -2),
					currentDate.AddDate(0, 0, -3),
					currentDate.AddDate(0, 0, -4),
				},
			},
		},
		{"testing standard sequence with already set seq:",
			5,
			Sequence{
				weekends: true,
				seq: []time.Time{
					currentDate,
				},
			},
			Sequence{
				weekends: true,
				seq: []time.Time{
					currentDate,
					currentDate.AddDate(0, 0, -1),
					currentDate.AddDate(0, 0, -2),
					currentDate.AddDate(0, 0, -3),
					currentDate.AddDate(0, 0, -4),
				},
			},
		},
	}

	for _, tc := range testCases {
		seq := tc.seq.Steps(tc.steps)
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v\nexpected %#v\nactual   %#v", tc.msg, tc.expSeq, seq)
		}
	}
}
