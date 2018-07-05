package dateseq

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	current := time.Now().Format("2006-01-02")
	currentDate, _ := time.Parse("2006-01-02", current)

	var testCases = []struct {
		msg    string
		expSeq Sequence
	}{
		{"testing create new sequence with default values",
			Sequence{
				Now:       currentDate,
				From:      time.Time{},
				To:        time.Time{},
				Weekends:  true,
				steps:     0,
				ascending: true,
				seq:       nil,
				exclude:   nil,
			},
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
			Sequence{Weekends: true},
		},
		{"testing setting weekends from already true to true",
			Sequence{Weekends: true},
			Sequence{Weekends: true},
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
			Sequence{Weekends: true},
			Sequence{Weekends: false},
		},
		{"testing setting weekends from already false to false",
			Sequence{Weekends: false},
			Sequence{Weekends: false},
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

func TestFromDate(t *testing.T) {
	timeB3, _ := time.Parse("2006-01-02", "2005-12-29")
	timeB2, _ := time.Parse("2006-01-02", "2005-12-30")
	timeB1, _ := time.Parse("2006-01-02", "2005-12-31")
	current, _ := time.Parse("2006-01-02", "2006-01-01")
	timeA1, _ := time.Parse("2006-01-02", "2006-01-02")
	timeA2, _ := time.Parse("2006-01-02", "2006-01-03")
	timeA3, _ := time.Parse("2006-01-02", "2006-01-04")

	var testCases = []struct {
		msg    string
		from   string
		seq    Sequence
		expSeq Sequence
	}{
		{"testing with empty string",
			"",
			Sequence{
				Now: current,
			},
			Sequence{
				Now: current,
			},
		},
		{"testing with From sma as Now",
			"2006-01-01",
			Sequence{
				Now: current,
			},
			Sequence{
				Now:  current,
				From: current,
			},
		},
		{"testing with From after Now",
			"2006-01-04",
			Sequence{
				Now: current,
			},
			Sequence{
				Now:  current,
				From: timeA3,
				seq:  []time.Time{timeA3, timeA2, timeA1, current},
			},
		},
		{"testing with From before Now",
			"2005-12-29",
			Sequence{
				Now: current,
			},
			Sequence{
				Now:  current,
				From: timeB3,
				seq:  []time.Time{timeB3, timeB2, timeB1, current},
			},
		},
	}

	for _, tc := range testCases {
		seq := tc.seq.FromDate(tc.from)
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v FromDate(%+v)\nexpected %+v\nactual   %+v", tc.msg, tc.from, tc.expSeq.Sequence(), seq.Sequence())
		}
	}
}

func TestToDate(t *testing.T) {
	sun, _ := time.Parse("2006-01-02", "2006-01-01")
	mon, _ := time.Parse("2006-01-02", "2006-01-02")
	thue, _ := time.Parse("2006-01-02", "2006-01-03")
	wed, _ := time.Parse("2006-01-02", "2006-01-04")
	thur, _ := time.Parse("2006-01-02", "2006-01-05")
	fri, _ := time.Parse("2006-01-02", "2006-01-06")
	sat, _ := time.Parse("2006-01-02", "2006-01-07")

	var testCases = []struct {
		msg    string
		to     string
		seq    Sequence
		expSeq Sequence
	}{
		{"testing with empty string",
			"",
			Sequence{
				Now: sun,
			},
			Sequence{
				Now: sun,
			},
		},
		{"testing with To same as Now",
			"2006-01-01",
			Sequence{
				Now: sun,
			},
			Sequence{
				Now: sun,
				To:  sun,
			},
		},
		{"testing with To before Now",
			"2006-01-01",
			Sequence{
				Now: wed,
			},
			Sequence{
				Now: wed,
				To:  sun,
				seq: []time.Time{sun, mon, thue, wed},
			},
		},
		{"testing with To after Now",
			"2006-01-07",
			Sequence{
				Now: wed,
			},
			Sequence{
				Now: wed,
				To:  sat,
				seq: []time.Time{wed, thur, fri, sat},
			},
		},
	}

	for _, tc := range testCases {
		seq := tc.seq.ToDate(tc.to)
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v ToDate(%+v)\nexpected %+v\nactual   %+v", tc.msg, tc.to, tc.expSeq.Sequence(), seq.Sequence())
		}
	}
}

func TestSteps(t *testing.T) {
	current := time.Now().Format("2006-01-02")
	currentDate, _ := time.Parse("2006-01-02", current)
	fixedDate, _ := time.Parse("2006-01-02", "2018-01-01") // Monday

	var testCases = []struct {
		msg    string
		seq    Sequence
		steps  int
		expSeq Sequence
	}{
		{"testing zero steps",
			New(), 0,
			Sequence{
				Now:       currentDate,
				Weekends:  true,
				ascending: true,
			},
		},
		{"testing single positive step",
			New(), 1,
			Sequence{
				Now:       currentDate,
				Weekends:  true,
				seq:       []time.Time{currentDate},
				steps:     1,
				ascending: true,
			},
		},
		{"testing multiple positive steps",
			New(), 3,
			Sequence{
				Now:      currentDate,
				Weekends: true,
				seq: []time.Time{
					currentDate,
					currentDate.AddDate(0, 0, +1),
					currentDate.AddDate(0, 0, +2),
				},
				steps:     3,
				ascending: true,
			},
		},
		{"testing single negative step",
			New(), -1,
			Sequence{
				Now:       currentDate,
				Weekends:  true,
				seq:       []time.Time{currentDate},
				steps:     -1,
				ascending: true,
			},
		},
		{"testing multiple negative steps",
			New(), -3,
			Sequence{
				Now:      currentDate,
				Weekends: true,
				seq: []time.Time{
					currentDate.AddDate(0, 0, -2),
					currentDate.AddDate(0, 0, -1),
					currentDate,
				},
				steps:     -3,
				ascending: true,
			},
		},
		{"testing multiple positive steps exclude weekends",
			Sequence{
				Now:      fixedDate,
				Weekends: false,
			}, 6,
			Sequence{
				Now:      fixedDate, // Monday
				Weekends: false,
				seq: []time.Time{
					fixedDate,
					fixedDate.AddDate(0, 0, +1), // Thuesday
					fixedDate.AddDate(0, 0, +2), // Wednesday
					fixedDate.AddDate(0, 0, +3), // Thursday
					fixedDate.AddDate(0, 0, +4), // Friday
					fixedDate.AddDate(0, 0, +7), // Monday
				},
				steps:     6,
				ascending: false,
			},
		},
		{"testing multiple negative steps exclude weekends",
			Sequence{
				Now:      fixedDate,
				Weekends: false,
			}, -5,
			Sequence{
				Now:      fixedDate,
				Weekends: false,
				seq: []time.Time{
					fixedDate.AddDate(0, 0, -6),
					fixedDate.AddDate(0, 0, -5),
					fixedDate.AddDate(0, 0, -4),
					fixedDate.AddDate(0, 0, -3),
					fixedDate,
				},
				steps:     -5,
				ascending: false,
			},
		},
		{"testing multiple steps with already set sequence",
			Sequence{
				Now:      currentDate,
				Weekends: true,
				seq: []time.Time{
					currentDate.AddDate(0, 0, +10),
				},
				ascending: true,
			}, 3,
			Sequence{
				Now:      currentDate,
				Weekends: true,
				seq: []time.Time{
					currentDate,
					currentDate.AddDate(0, 0, +1),
					currentDate.AddDate(0, 0, +2),
				},
				steps:     3,
				ascending: true,
			},
		},
	}

	for _, tc := range testCases {
		seq := tc.seq.WithSteps(tc.steps)
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v WithSteps(%v)\nexpected %v %#v\nactual   %v %#v", tc.msg, tc.steps, tc.expSeq.String(), tc.expSeq, seq.String(), seq)
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
				seq:       []time.Time{time1, time2, time3},
				ascending: true,
			},
		},
		{"testing ascending sorting with desc sorted entries",
			Sequence{
				seq: []time.Time{time3, time2, time1},
			},
			Sequence{
				seq:       []time.Time{time1, time2, time3},
				ascending: true,
			},
		},
		{"testing ascending sorting with nil entries",
			Sequence{
				seq: []time.Time{},
			},
			Sequence{
				seq:       []time.Time{},
				ascending: true,
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
				seq:       []time.Time{time3, time2, time1},
				ascending: false,
			},
		},
		{"testing descending sorting with asc sorted entries",
			Sequence{
				seq: []time.Time{time1, time2, time3},
			},
			Sequence{
				seq:       []time.Time{time3, time2, time1},
				ascending: false,
			},
		},
		{"testing descending sorting with nil entries",
			Sequence{
				seq: []time.Time{},
			},
			Sequence{
				seq:       []time.Time{},
				ascending: false,
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

func TestCreateSequence(t *testing.T) {
	sun, _ := time.Parse("2006-01-02", "2006-01-01")
	mon, _ := time.Parse("2006-01-02", "2006-01-02")
	thue, _ := time.Parse("2006-01-02", "2006-01-03")
	wed, _ := time.Parse("2006-01-02", "2006-01-04")
	thur, _ := time.Parse("2006-01-02", "2006-01-05")
	fri, _ := time.Parse("2006-01-02", "2006-01-06")
	sat, _ := time.Parse("2006-01-02", "2006-01-07")

	var testCases = []struct {
		msg    string
		t1     time.Time
		t2     time.Time
		expSeq []time.Time
	}{
		{"testing same date",
			sun, sun,
			[]time.Time{sun},
		},
		{"testing t1 before t2",
			sun, sat,
			[]time.Time{sun, mon, thue, wed, thur, fri, sat},
		},
		{"testing t2 before t1",
			sat, sun,
			[]time.Time{sun, mon, thue, wed, thur, fri, sat},
		},
	}

	for _, tc := range testCases {
		seq := createSequence(tc.t1, tc.t2)
		if !reflect.DeepEqual(seq, tc.expSeq) {
			t.Errorf("%v createSequence(%v, %v)\nexpected %v\nactual   %v", tc.msg, tc.t1, tc.t2, seq, tc.expSeq)
		}
	}
}

func TestRemoveWeekends(t *testing.T) {
	sun, _ := time.Parse("2006-01-02", "2006-01-01")
	mon, _ := time.Parse("2006-01-02", "2006-01-02")
	thue, _ := time.Parse("2006-01-02", "2006-01-03")
	wed, _ := time.Parse("2006-01-02", "2006-01-04")
	thur, _ := time.Parse("2006-01-02", "2006-01-05")
	fri, _ := time.Parse("2006-01-02", "2006-01-06")
	sat, _ := time.Parse("2006-01-02", "2006-01-07")

	var testCases = []struct {
		msg     string
		list    []time.Time
		expList []time.Time
	}{
		{"testing empty list",
			[]time.Time{},
			[]time.Time{},
		},
		{"testing list with single weekday",
			[]time.Time{mon},
			[]time.Time{mon},
		},
		{"testing list with single weekend day",
			[]time.Time{sun},
			[]time.Time{},
		},
		{"testing list with complete week",
			[]time.Time{sun, mon, thue, wed, thur, fri, sat},
			[]time.Time{mon, thue, wed, thur, fri},
		},
	}

	for _, tc := range testCases {
		list := removeWeekendFromSequence(tc.list)
		if !reflect.DeepEqual(list, tc.expList) {
			t.Errorf("%v removeWeekendFromSequence(%v)\nexpected %v\nactual   %v", tc.msg, tc.list, tc.expList, list)
		}
	}
}

func TestAddWeekends(t *testing.T) {
	thur, _ := time.Parse("2006-01-02", "2006-01-05")
	fri, _ := time.Parse("2006-01-02", "2006-01-06")
	sat, _ := time.Parse("2006-01-02", "2006-01-07")
	sun, _ := time.Parse("2006-01-02", "2006-01-08")
	mon, _ := time.Parse("2006-01-02", "2006-01-09")
	thue, _ := time.Parse("2006-01-02", "2006-01-10")
	wed, _ := time.Parse("2006-01-02", "2006-01-11")

	var testCases = []struct {
		msg     string
		input   []time.Time
		expList []time.Time
	}{
		{"testing empty list",
			[]time.Time{},
			[]time.Time{},
		},
		{"testing list with single weekday",
			[]time.Time{mon},
			[]time.Time{mon},
		},
		{"testing list with single weekend day",
			[]time.Time{sun},
			[]time.Time{sun},
		},
		{"testing list with mixed week and weekend day",
			[]time.Time{fri, sat},
			[]time.Time{fri, sat},
		},
		{"testing list with missing weekend",
			[]time.Time{thur, fri, mon, thue, wed},
			[]time.Time{thur, fri, sat, sun, mon, thue, wed},
		},
	}

	for _, tc := range testCases {
		list := addWeekendToSequence(tc.input)
		if !reflect.DeepEqual(list, tc.expList) {
			t.Errorf("%v addWeekendToSequence(%v)\nexpected %v\nactual   %v", tc.msg, tc.input, tc.expList, list)
		}
	}
}
