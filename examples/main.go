package main

import (
	"fmt"

	"github.com/dirkolbrich/dateseq"
)

func main() {
	seq := dateseq.New()

	fmt.Println("Standard sequence with weekends:")
	seq = seq.WithSteps(10)
	for k, v := range seq.Sequence() {
		fmt.Printf("%v. %v\n", k, v.Format("2006-01-02 Mon"))
	}

	fmt.Println("Sequence excluding weekends:")
	seq = seq.ExcludeWeekends().WithSteps(10)
	for k, v := range seq.Sequence() {
		fmt.Printf("%v. %v\n", k, v.Format("2006-01-02 Mon"))
	}

	fmt.Println("Creation and retrieval of the sequence in one step:")
	s := dateseq.New().WithSteps(5).Sequence()
	for k, v := range s {
		fmt.Printf("%v. %v\n", k, v.Format("2006-01-02 Mon"))
	}

	fmt.Println("Slice with a simple string representation of the dates:")
	strings := dateseq.New().WithSteps(5).String()
	fmt.Println(strings)
}
