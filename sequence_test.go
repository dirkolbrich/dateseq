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

func TestIncludeWeekends(t *testing.T) {
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
		seq := tc.seq.IncludeWeekends()
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v IncludeWeekends()\nexpected %#v\nactual   %#v", tc.msg, tc.expSeq, seq)
		}
	}
}

func TestExcludeWeekends(t *testing.T) {
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
		seq := tc.seq.ExcludeWeekends()
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v ExcludeWeekends()\nexpected %#v\nactual   %#v", tc.msg, tc.expSeq, seq)
		}
	}
}

func TestExclude(t *testing.T) {
	time1, _ := time.Parse("2006-01-02", "2006-01-01")
	time2, _ := time.Parse("2006-01-02", "2006-01-02")
	time3, _ := time.Parse("2006-01-02", "2006-01-03")
	time4, _ := time.Parse("2006-01-02", "2006-01-04")

	var testCases = []struct {
		msg     string
		seq     Sequence
		exclude []string
		expSeq  Sequence
	}{
		{"testing exclude with empty exclude list",
			Sequence{
				seq: []time.Time{time1, time2, time3},
			},
			[]string{},
			Sequence{
				seq: []time.Time{time1, time2, time3},
			},
		},
		{"testing exclude from empty sequence",
			Sequence{},
			[]string{"2006-01-01"},
			Sequence{
				exclude: []time.Time{time1},
			},
		},
		{"testing exclude from single sequence with single matching date",
			Sequence{
				seq: []time.Time{time1},
			},
			[]string{"2006-01-01"},
			Sequence{
				exclude: []time.Time{time1},
				seq:     []time.Time{},
			},
		},
		{"testing exclude from sequence with single non matching date",
			Sequence{
				seq: []time.Time{time1, time2, time3},
			},
			[]string{"2006-01-04"},
			Sequence{
				exclude: []time.Time{time4},
				seq:     []time.Time{time1, time2, time3},
			},
		},
		{"testing exclude from start of sequence with single matching date",
			Sequence{
				seq: []time.Time{time1, time2, time3},
			},
			[]string{"2006-01-01"},
			Sequence{
				exclude: []time.Time{time1},
				seq:     []time.Time{time2, time3},
			},
		},
		{"testing exclude from within of sequence with single matching date",
			Sequence{
				seq: []time.Time{time1, time2, time3},
			},
			[]string{"2006-01-02"},
			Sequence{
				exclude: []time.Time{time2},
				seq:     []time.Time{time1, time3},
			},
		},
		{"testing exclude from end of sequence with single matching date",
			Sequence{
				seq: []time.Time{time1, time2, time3},
			},
			[]string{"2006-01-03"},
			Sequence{
				exclude: []time.Time{time3},
				seq:     []time.Time{time1, time2},
			},
		},
		{"testing exclude from with multiple matching date",
			Sequence{
				seq: []time.Time{time1, time2, time3, time4},
			},
			[]string{"2006-01-03", "2006-01-02"},
			Sequence{
				exclude: []time.Time{time2, time3},
				seq:     []time.Time{time1, time4},
			},
		},
	}

	for _, tc := range testCases {
		seq := tc.seq.Exclude(tc.exclude)
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v Exclude(%+v)\nexpected %#v\nactual   %#v", tc.msg, tc.exclude, tc.expSeq, seq)
		}
	}
}

func TestSteps(t *testing.T) {
	current := time.Now().Format("2006-01-02")
	currentDate, _ := time.Parse("2006-01-02", current)

	var testCases = []struct {
		msg    string
		seq    Sequence
		steps  int
		expSeq Sequence
	}{
		{"testing zero steps",
			Sequence{}, 0,
			Sequence{},
		},
		{"testing single positive step",
			Sequence{}, 1,
			Sequence{
				seq:   []time.Time{currentDate},
				steps: 1,
			},
		},
		{"testing single multiple positive steps",
			Sequence{}, 3,
			Sequence{
				seq: []time.Time{
					currentDate,
					currentDate.AddDate(0, 0, +1),
					currentDate.AddDate(0, 0, +2),
				},
				steps: 3,
			},
		},
		{"testing single negative step",
			Sequence{}, -1,
			Sequence{
				seq:   []time.Time{currentDate},
				steps: -1,
			},
		},
		{"testing single multiple negative steps",
			Sequence{}, -3,
			Sequence{
				seq: []time.Time{
					currentDate.AddDate(0, 0, -2),
					currentDate.AddDate(0, 0, -1),
					currentDate,
				},
				steps: -3,
			},
		},
	}

	for _, tc := range testCases {
		seq := tc.seq.Steps(tc.steps)
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v Steps()\nexpected %v %#v\nactual   %v %#v", tc.msg, tc.expSeq.String(), tc.expSeq, seq.String(), seq)
		}
	}
}

func TestSortAsc(t *testing.T) {
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
		seq := tc.seq.SortAsc()
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v SortAsc()\nexpected %#v\nactual   %#v", tc.msg, tc.expSeq, seq)
		}
	}
}

func TestSortDesc(t *testing.T) {
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
		seq := tc.seq.SortDesc()
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v SortDesc()\nexpected %#v\nactual   %#v", tc.msg, tc.expSeq, seq)
		}
	}
}

func TestSequence(t *testing.T) {
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
		seq := tc.seq.Sequence()
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
