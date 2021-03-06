package main

import (
	"github.com/glendc/cgreader"
	"strings"
)

func PrintCNGroup(is_zero bool, counter int) (em string) {
	em += "0"
	if is_zero {
		em += "0"
	}

	em += " " + strings.Repeat("0", counter)
	return
}

func main() {
	cgreader.RunStaticPrograms(
		cgreader.GetFileList("../../input/chuck_norris_%d.txt", 4),
		cgreader.GetFileList("../../output/chuck_norris_%d.txt", 4),
		true,
		func(input <-chan string, output chan string) {
			input_message, message := <-input, ""

			for i := range input_message {
				character := int(input_message[i])
				if character != '\n' {
					var msg string
					for u := 0; u < 7; u++ {
						if character%2 == 0 {
							msg = "0" + msg
						} else {
							msg = "1" + msg
						}
						character /= 2
					}
					message += msg
				}
			}

			is_zero, counter, lc := message[0] == '0', 0, len(message)-1

			for i := range message {
				if iz, il := message[i] == '0', i == lc; il || is_zero != iz {
					if il && iz == is_zero {
						counter++
					}

					output <- PrintCNGroup(is_zero, counter)

					if il && iz != is_zero {
						output <- " " + PrintCNGroup(iz, 1)
					} else if !il {
						output <- " "
					}

					is_zero, counter = !is_zero, 0
				}
				counter++
			}
		})
}
