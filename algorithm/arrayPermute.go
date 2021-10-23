//perm([a,b,c]) = [a,perm(b,c)]) union [b,perm(a,c)]) union [c, perm(a,b)]
package algorithm

import (
	"fmt"
	"zhengyun.zy/algorithm/version"
)

func Test() {
	fmt.Println(version.Get())
	input := []interface{}{1, 2, 3, 4}

	fmt.Println("V1")
	output := PermuteV1(input)
	for _, item := range output {
		fmt.Println(item)
	}

	fmt.Println("V2")
	PermuteV2(input, make([]interface{}, 0))

}

// breath first version
func PermuteV1(input []interface{}) [][]interface{} {

	if len(input) == 1 {
		return [][]interface{}{input}
	} else {
		output := make([][]interface{}, 0)
		for index, value := range input {
			temp := make([]interface{}, 0)
			for i := 0; i < len(input); i++ {
				if i != index {
					temp = append(temp, input[i])
				}
			}
			onePermute := PermuteV1(temp)
			for _, item := range onePermute {
				output = append(output, append(item, value))
			}

		}
		return output
	}

}

// depth first version
func PermuteV2(input, fixedPart []interface{}) {
	if len(input) == 1 {
		fmt.Println(append(fixedPart, input[0]))
	} else {
		for index, value := range input {
			newFixed := append(fixedPart, value)
			temp := make([]interface{}, 0)
			for i := 0; i < len(input); i++ {
				if i != index {
					temp = append(temp, input[i])
				}
			}
			PermuteV2(temp, newFixed)
		}
	}

}
