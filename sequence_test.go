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
			Sequence{},
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

func TestSeq(t *testing.T) {
	time1, _ := time.Parse("2006-01-02", "2017-10-09")
	time2, _ := time.Parse("2006-01-02", "2017-10-08")
	time3, _ := time.Parse("2006-01-02", "2017-10-07")

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
