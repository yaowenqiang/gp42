// Package main provides ...
package main

import (
	"fmt"
	"sort"
)

type Peak struct {
	Name      string
	Elevation int
}

type person struct {
	Name string
	Age  int
}

type personSlice []person

func (s personSlice) Len() int           { return len(s) }
func (s personSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s personSlice) Less(i, j int) bool { return s[i].Age < s[j].Age }

func main() {

	a := []int{3, 6, 4, -1, 9, 11, -14}
	sort.Ints(a)
	fmt.Printf("a = %+v\n", a)

	ss := []string{"Surface", "ipad", "max pro", "mac air", "think pad", "idea pad"}
	sort.Strings(ss)
	fmt.Printf("ss = %+v\n", ss)

	sort.Sort(sort.Reverse(sort.StringSlice(ss)))
	fmt.Printf("ss = %+v\n", ss)

	b := make([]int, len(a))
	fmt.Printf("IntsAreSorted(b) = %+v\n", sort.IntsAreSorted(b))
	copy(b, a)

	fmt.Printf("sorted b = %+v\n", sort.IntSlice(b))

	sort.Sort(sort.Reverse(sort.IntSlice(b)))
	fmt.Printf("b = %+v\n", b)

	c := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 00}
	d := make([]float64, len(c))
	copy(d, c)
	fmt.Printf("d = %+v\n", d)

	fmt.Printf("c = %+v\n", c)

	sort.Float64s(c)

	fmt.Printf("sorted c = %+v\n", c)

	fmt.Printf("Float64sAreSorted(c) = %+v\n", sort.Float64sAreSorted(c))
	fmt.Printf("Float64sAreSorted(d) = %+v\n", sort.Float64sAreSorted(d))

	//e := [4]int{10, 100, 100, 343}

	//fmt.Printf("a IsSorted = %+v\n", sort.IsSorted(e))

	persons := personSlice{
		{
			Name: "aaa",
			Age:  10,
		},
		{
			Name: "bbb",
			Age:  1,
		},
		{
			Name: "ggg",
			Age:  20,
		},
		{
			Name: "ff2",
			Age:  99,
		},
		{
			Name: "ff2",
			Age:  99,
		},
		{
			Name: "ff3",
			Age:  99,
		},
		{
			Name: "zz4",
			Age:  99,
		},
		{
			Name: "ddd",
			Age:  8,
		},
		{
			Name: "fff",
			Age:  99,
		},
	}
	fmt.Printf("persons = %+v\n", persons)

	sort.Sort(persons)
	fmt.Printf("persons = %+v\n", persons)
	sort.Stable(persons)
	fmt.Printf("persons = %+v\n", persons)

	peaks := []Peak{
		{"Aconcagua", 22838},
		{"Denali", 20322},
		{"Kilimanjanro", 19341},
		{"Mount Elbrus", 18510},
		{"Mount Everest", 29029},
		{"Mount Kosciuszko", 7310},
		{"Mount Vinson ", 16050},
		{"Puncak Jaya", 16024},
	}

	sort.Slice(peaks, func(i, j int) bool {
		return peaks[i].Elevation >= peaks[j].Elevation
	})
	fmt.Printf("peaks = %+v\n", peaks)

}
