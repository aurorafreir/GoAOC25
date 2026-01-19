package main

import (
	"fmt"
	"time"

	"example.com/utils"
)

func main() {
	mainTimeStart := time.Now()

	// Day 01
	funcTimeStart := time.Now()
	d1p1 := d1p1()
	fmt.Println("D1p1 result:", d1p1)
	fmt.Println("D1p1 took:", time.Since(funcTimeStart).Microseconds(), "microseconds")

	funcTimeStart = time.Now()
	d1p2 := d1p2()
	fmt.Println("D1p2 result:", d1p2)
	fmt.Println("D1p2 took:", time.Since(funcTimeStart).Microseconds(), "microseconds")

	// Day 02

	// Day 03
	funcTimeStart = time.Now()
	d3p1, err := d3(2)
	utils.Check(err)
	fmt.Println("D3p1 result:", d3p1)
	fmt.Println("D3p1 took:", time.Since(funcTimeStart).Microseconds(), "microseconds")

	funcTimeStart = time.Now()
	d3p2, err := d3(12)
	utils.Check(err)
	fmt.Println("D3p2 result:", d3p2)
	fmt.Println("D3p2 took:", time.Since(funcTimeStart).Microseconds(), "microseconds")

	// Day 04
	funcTimeStart = time.Now()
	d4p1, err := d4p1()
	utils.Check(err)
	fmt.Println("D4p1 result:", d4p1)
	fmt.Println("D4p1 took:", time.Since(funcTimeStart).Microseconds(), "microseconds")

	funcTimeStart = time.Now()
	d4p2, err := d4p2()
	utils.Check(err)
	fmt.Println("D4p2 result:", d4p2)
	fmt.Println("D4p2 took:", time.Since(funcTimeStart).Microseconds(), "microseconds")

	// Day 05
	funcTimeStart = time.Now()
	d5p1, err := d5p1()
	utils.Check(err)
	fmt.Println("D5p1 result:", d5p1)
	fmt.Println("D5p1 took:", time.Since(funcTimeStart).Microseconds(), "microseconds")

	funcTimeStart = time.Now()
	d5p2, err := d5p2()
	utils.Check(err)
	fmt.Println("D5p2 result:", d5p2)
	fmt.Println("D5p2 took:", time.Since(funcTimeStart).Microseconds(), "microseconds")

	fmt.Println("All completed AOC 2025 challenges finished in", time.Since(mainTimeStart).Microseconds(), "microseconds /", time.Since(mainTimeStart).Seconds(), "seconds")
}
