package main

import (
	api "1_Test/api"
	"fmt"
	"math"
	"strconv"
)

func main() {
	stop := false
	//Console app
	for !stop {
		fmt.Println("Please input 8 param of sequence with 3 unknowns (X,Y,Z)")
		fmt.Println("Example data: [1, X, 8, 17, Y, Z, 78, 113]")
		fmt.Println("************* Press Q to exit ************")
		fmt.Println("--------------------------------------------------------")

		//AX = B matrix
		var amatrix = make([]float64, 25)
		var bmatrix = []float64{}
		var postx, posty, postz int
		var strinput string
		startindex := 0
		endindex := 0

		//Create matrix A, B fot API input
		for i := 1; i <= 8 && !stop; {
			fmt.Printf(">> %d. param = ", i)
			fmt.Scanln(&strinput)
			if strinput == "Q" || strinput == "q" {
				stop = true
			} else {
				floinput, err := strconv.ParseFloat(strinput, 64)
				if err != nil {
					if strinput != "X" && strinput != "Y" && strinput != "Z" {
						fmt.Println("Invalid input! please enter only numeric or X/Y/Z")
					} else {
						switch strinput {
						case "X":
							if postx != 0 {
								fmt.Println("Already assigned param!")
							} else {
								postx = i
								i++
							}
						case "Y":
							if posty != 0 {
								fmt.Println("Already assigned param!")
							} else {
								posty = i
								i++
							}
						case "Z":
							if postz != 0 {
								fmt.Println("Already assigned param!")
							} else {
								postz = i
								i++
							}
						default: //Do nothing
						}
					}
				} else {
					//Set A matrix
					pow := 4
					endindex = startindex + 4
					for j := startindex; j <= endindex; j++ {
						amatrix[j] = math.Pow(float64(i), float64(pow))
						pow--
					}

					//Set B index
					bmatrix = append(bmatrix, floinput)
					startindex = endindex + 1
					i++
				}
			}
		}

		if !stop {
			/*For testing
			fmt.Println(amatrix)
			fmt.Println(bmatrix)
			fmt.Println(postx)
			fmt.Println(posty)
			fmt.Println(postz)
			*/

			if postx == 0 || posty == 0 || postz == 0 {
				fmt.Println("Missing input! please enter all of X/Y/Z")
			} else {
				result_X, result_Y, result_Z, err := api.Getresult(amatrix, bmatrix, postx, posty, postz)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("X = %.6F\n", result_X)
					fmt.Printf("Y = %.6F\n", result_Y)
					fmt.Printf("Z = %.6F\n", result_Z)
				}
			}
		}

		fmt.Println("--------------------------------------------------------")
		fmt.Printf("\n")
	}
	fmt.Println("Application stop!")
}
