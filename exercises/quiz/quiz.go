package quiz

import (
	"fmt"
	//"bufio"
	"os"
	"strconv"
	"encoding/csv"
	"strings"
)

func init(){
	val,i := 1,0
	for(val == 1){
		fmt.Println("Press 1 to continue the quiz to 2 stop")
		//bufio.NewScanner(os.Stdin)
		var input string  
		fmt.Scanln(&input)
		val,err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		if val == 2 {
			continue
		}
		is := strconv.Itoa(i + 1)
		var records [][] string= parseQuestions()
		res := strings.Split(records[i][0], ",")
		fmt.Println("Question " +is + ": " + res[0])
		var ansinput string
		fmt.Scanln(&ansinput)
		if ansinput == res[1] {
			fmt.Println("Correct Answer")
			continue
		} else {
			fmt.Println("InCorrect Answer")
		}
		i++
	}

	fmt.Println("Thanks for participating")
	
}

func parseQuestions() [][]string{
	filePath := "quiz.csv"
	f, err := os.Open(filePath)
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        fmt.Println(err)
    }

	
    return records
}