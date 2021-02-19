package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	letter := make(chan bool)
	number := make(chan bool)
	wait := sync.WaitGroup{}

	wait.Add(2)

	go func(wait *sync.WaitGroup) {
		i := 1
		for {
			select {
			case <-number:
				if i == 28 {
					fmt.Println("aaaaaa")
					wait.Done()
					return
				}
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}(&wait)
	go func(wait *sync.WaitGroup) {
		i := 0
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for {
			select {
			case <-letter:
				if i >= strings.Count(str, "")-1 {
					fmt.Println("bbb")
					wait.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				fmt.Print(str[i : i+1])
				i++
				number <- true
				break
			default:
				break
			}
		}
	}(&wait)

	number <- true
	wait.Wait()
}
