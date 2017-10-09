package main

import (
	"fmt"

	"github.com/dirkolbrich/dateseq"
)

func main() {
	seq := dateseq.New()

	fmt.Println("Sequence excluding weekends:")
	seq = seq.Duration(10)
	for k, v := range seq.Seq() {
		fmt.Printf("%v. %v\n", k, v.Format("2006-01-02 Mon"))
	}

	fmt.Println("Sequence including weekends:")
	seq = seq.InclWeekends().Duration(10)
	for k, v := range seq.Seq() {
		fmt.Printf("%v. %v\n", k, v.Format("2006-01-02 Mon"))
	}

	fmt.Println("Creation and retrieval of the sequence in one step:")
	s := dateseq.New().InclWeekends().Duration(5).Seq()
	for k, v := range s {
		fmt.Printf("%v. %v\n", k, v.Format("2006-01-02 Mon"))
	}
}
