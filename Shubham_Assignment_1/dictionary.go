package main

import (
	"fmt"
	"bufio"
	"os"
  "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main(){

	f, err := os.Open("Oxford_English_Dictionary.txt")
    check(err)
    reader := bufio.NewReader(f)
    // scanner := bufio.NewScanner(os.Stdin)

    var line string
    line, err = reader.ReadString('\n')
    dictionary := make(map[string]string)

    for {
    	line, err = reader.ReadString('\n')
    	word := strings.Split(line, " ")[0]
    	meaning := strings.Join(strings.Split(line, " ")[1:], " ")
    	dictionary[strings.ToLower(word)] = meaning
    	if err != nil {
    		break
    	}
    }


    var word string
    var correct_word int
    correct_word = 0
		var i = 100000
    for {
			if i < 1 {
				fmt.Println(correct_word)
				break
			}
    	// fmt.Println("Enter a word ... ")
    	// scanner.Scan()
    	// word = scanner.Text()
			word = "zoo"
    	if dictionary[strings.ToLower(word)] != "" {
    		fmt.Println("Meaning of ", word, " is : ", dictionary[strings.ToLower(word)])
    		correct_word = 1
    	} else{
    		fmt.Println("Please check the spelling again.")
    	}
			i -= 1
    }


}
