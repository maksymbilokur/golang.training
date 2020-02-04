package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)
func main() {
	filename:=flag.String("f","testnum.txt","Filename")
	flag.Parse()
	perfectNum(filename)
}
func perfectNum(filename *string) {

	if validNoArgs() {
			perfectNumA()
		} else if &filename == "testnum.txt"{
		perfectNumF(filename)} else{
			perfectNumNA()
	}
}


func perfectNumNA() {
	n := "string"
	for n != "no" {
		var n1 int
		fmt.Print("Input n(no for quit): ")
		fmt.Fscan(os.Stdin, &n)
		n1, _ = strconv.Atoi(n)
		if n1 <= 0 {
			fmt.Printf("invalid type of n")
			continue
		}
		for i := 1; i < n1; i++ {
			checkNum(i)
		}
	}
}

func perfectNumA(){
	n, _ := strconv.Atoi(os.Args[1])

	if n <= 0 {
		fmt.Printf("invalid type of n")
		return
	}
	for i := 1; i < n; i++ {
		checkNum(i)
	}

}
func perfectNumF(filename *string){

	var n int
	if n <= 0 {
		fmt.Printf("invalid type of n")
		return
	}
	for i := 1; i < n; i++ {
		checkNum(i)
	}
}


func checkNum(n int) {
	sum:=0
	for i:=1;i<=n/2;i++{
		if n%i==0{
			sum= sum+i
		}
	}
	if sum==n{
		print(strconv.Itoa(sum))
	}
}
func print(s string){
	fmt.Println(s)
}

func validNoArgs()bool{
	if len(os.Args)<2{
		return true
	}
	return false
}