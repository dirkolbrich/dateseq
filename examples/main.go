package main

import (
	"fmt"

	"github.com/dirkolbrich/dateseq"
)

func main() {
	seq := dateseq.New()

	fmt.Println("Sequence excluding weekends:")
	seq.Duration(10)
	for k, v := range seq.Seq() {
		fmt.Printf("%v. %v\n", k, v.Format("2006-01-02 Mon"))
	}

	fmt.Println("Sequence including weekends:")
	seq.InclWeekends()
	seq.Duration(10)
	for k, v := range seq.Seq() {
		fmt.Printf("%v. %v\n", k, v.Format("2006-01-02 Mon"))
	}
}
