// What is the average of all the numbers between 1 to 100 that divide either by 3 or 5?
package main

import "fmt"

func main() {
	count, total := 0, 0
	for n := 1; n <= 100; n++ {
		if n%3 == 0 || n%5 == 0 {
			count++
			total += n
		}
	}

	fmt.Println(float64(total) / float64(count))
}
git commit