[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE.md)
[![Travis](https://img.shields.io/travis/dirkolbrich/dateseq.svg?style=flat-square)](https://travis-ci.org/dirkolbrich/dateseq)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/dirkolbrich/dateseq)
[![Coverage Status](https://img.shields.io/coveralls/dirkolbrich/dateseq/master.svg?style=flat-square)](https://coveralls.io/github/dirkolbrich/dateseq?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/dirkolbrich/dateseq?style=flat-square)](https://goreportcard.com/report/github.com/dirkolbrich/dateseq)

# dateseq

A package to create a sequence series of daily dates, written in golang.

Sequence is an immutable struct, each operations returns a new Sequence and leaves the original untouched.

```go
seq := dateseq.New()
seq = seq.WithSteps(5)

for k, v := range seq.Sequence() {
    fmt.Printf("%v. %v\n", k, v.Format("2006-01-02 Mon"))
}
```

Returns a sequence of dates from the current date counted backwards (*assuming, you start this this programm on 2017-10-09*):

```bash
0. 2017-10-09 Mon
1. 2017-10-08 Sun
2. 2017-10-07 Sat
3. 2017-10-06 Fri
4. 2017-10-05 Thu
```

Weekends can be excluded with `ExcludeWeekends()` from the sequence series. `IncludeWeekends()` will include them again, which is also the default behavior.

```go
seq := dateseq.New()
seq = seq.ExcludeWeekends().WithSteps(5)

for k, v := range seq.Sequence() {
    fmt.Printf("%v. %v\n", k, v.Format("2006-01-02 Mon"))
}
```

Returns:

```bash
0. 2017-10-09 Mon
1. 2017-10-06 Fri
2. 2017-10-05 Thu
3. 2017-10-04 Wed
4. 2017-10-03 Tue
```

Exclude some specific dates with `Exclude(list []string)`.

```go
seq := dateseq.New()
exclude := []string{
    "2017-12-25",
    "2017-12-26",
}

seq = seq.WithSteps(5).Exclude(exclude)

```

The sequence methods are chainable, which allows creation and retrieving of the sequence slice in one go.

```go
seq := dateseq.New().WithSteps(5).Sequence()
```

Need a slice with only the string representations of the dates?

```go
seq := dateseq.New().WithSteps(5).String()
fmt.Println(seq)
```

Returns the dates in format YYYY-MM-DD:

```bash
[2017-10-09 2017-10-08 2017-10-07 2017-10-06 2017-10-05]
```

For a custom output format use the `Format(layout string)` method, which returns the string in a layout defined by the `time` package.

```go
seq := dateseq.New().WithSteps(5).Format("Jan 01. 2006")
```