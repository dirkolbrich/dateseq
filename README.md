[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE.md)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/dirkolbrich/dateseq)
[![Go Report Card](https://goreportcard.com/badge/github.com/dirkolbrich/dateseq?style=flat-square)](https://goreportcard.com/report/github.com/dirkolbrich/dateseq)

# dateseq

A package to create a sequence series of daily dates, written in golang.

```go
seq := dateseq.New()
seq = seq.Duration(10)

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

Weekends can be included with `InclWeekends()` into the sequence series. `ExclWeekends()` will exclude them again, which is also the default behavior.

```go
seq := dateseq.New()
seq = seq.InclWeekends().Duration(10)

for k, v := range seq.Seq() {
    fmt.Printf("%v. %v\n", k, v.Format("2006-01-02 Mon"))
}
```

Returns:

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

The sequence methods are chainable, wich allows creation and retrieving of the sequence slice in one go.

```go
seq := dateseq.New().Duration(10).Seq()
```