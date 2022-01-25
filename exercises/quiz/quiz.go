package quiz

import (
	"fmt"
	"os"
	"strconv"
	"encoding/csv"
	// "path/filepath"
)

func Init(){
	val,i := 1,0
	for(val == 1){
		fmt.Println("Press 1 to continue the quiz to 2 stop")
		var input string  
		fmt.Scanln(&input)
		val,err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		if val == 2 {
			break
		}
		is := strconv.Itoa(i + 1)
		var records [][] string= parseQuestions()
		fmt.Println("Question " +is + ": " + records[i][0])
		var ansinput string
		fmt.Scanln(&ansinput)
		if ansinput == records[i][1] {
			fmt.Println("Correct Answer")
			if i == (len(records) - 1) {
				break
			} 
			i++
			continue
		} else {
			fmt.Println("InCorrect Answer")
		}
	}

	fmt.Println("Finished! Thanks for participating")
	
}

func parseQuestions() [][]string{

	 fileName := "quiz.csv"
	//  path, _ := filepath.Abs(fileName)
	f, err := os.Open(fileName)
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