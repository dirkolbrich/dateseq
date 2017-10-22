[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE.md)
[![Travis](https://img.shields.io/travis/dirkolbrich/dateseq.svg?style=flat-square)](https://travis-ci.org/dirkolbrich/dateseq)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/dirkolbrich/dateseq)
[![Coverage Status](https://img.shields.io/coveralls/dirkolbrich/dateseq/master.svg?style=flat-square)](https://coveralls.io/github/dirkolbrich/dateseq?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/dirkolbrich/dateseq?style=flat-square)](https://goreportcard.com/report/github.com/dirkolbrich/dateseq)

# dateseq

A package to create a sequence series of daily dates, written in golang.

```go
seq := dateseq.New()
seq = seq.Steps(10)

for k, v := range seq.Seq() {
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
5. 2017-10-04 Wed
6. 2017-10-03 Tue
7. 2017-10-02 Mon
8. 2017-10-01 Sun
9. 2017-09-30 Sat
```

Weekends can be excluded with `ExclWeekends()` from the sequence series. `InclWeekends()` will include them again, which is also the default behavior.

```go
seq := dateseq.New()
seq = seq.ExclWeekends().Steps(10)

for k, v := range seq.Seq() {
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
5. 2017-10-02 Mon
6. 2017-09-29 Fri
7. 2017-09-28 Thu
8. 2017-09-27 Wed
9. 2017-09-26 Tue
```

The sequence methods are chainable, which allows creation and retrieving of the sequence slice in one go.

```go
seq := dateseq.New().Steps(10).Seq()
```

Need a slice with only the string representations of the dates?

```go
seq := dateseq.New().Steps(10).String()
fmt.Println(seq)
```

Returns the dates in format YYYY-MM-DD:

```bash
[2017-10-09 2017-10-08 2017-10-07 2017-10-06 2017-10-05]
```

For a custom format use the `Format(layout string)` method, which returns the string in a layout defined by the `time` package.

```go
seq := dateseq.New().Steps(10).Format("Jan 01. 2006")
```