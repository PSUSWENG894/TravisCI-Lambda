package main

func Sum(stati []string) int {
	total := 0;
	for index := 0; index < len(stati); index++ {
		if stati[index] == "Failed" || stati[index] == "Error" {
			total++
		}
	}
	return total
}

func main() {
	test := make([]string, 10)
	test[0] = "Success"
	test[1] = "Failed"
	test[2] = "Error"
	test[3] = "Ignore"
	test[4] = "Skipped"
	test[5] = "Failed"
	test[6] = "Failed"
	test[7] = "Skipped"
	test[8] = "Error"
	test[9] = "Success"
	Sum(test)
}
