// Package main provides ...
package main

import (
	"fmt"
)

type A struct {
	Face int
}

type Aa A

type IZ = int

func (a A) f() {
	fmt.Println("hi ", a.Face)
}


type Mutex struct

func (m *Mutex) Lock() {fmt.Println("Lock")}
func (m *Mutex) Unlock() {fmt.Println("Unlock")}
type NewMutex Mutex
type PtrMutex *Mutex

type PrintableMutex struct {
	Mutex
}
func main() {
	var s A = A{Face: 9}
	s.f()

	//	var sa Aa = Aa{Face: 10}
	//sa.f()
}
