//perm([a,b,c]) = [a,perm(b,c)]) union [b,perm(a,c)]) union [c, perm(a,b)]
package main

import "fmt"
import "zhengyun.zy/algorithm"

func main() {
	input := []interface{}{1, 2, 3, 4}

	fmt.Println(input)
	algorithm.PermuteV2(input,make([]interface{},0))
	
	

}