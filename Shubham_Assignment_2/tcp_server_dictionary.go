package main

import "net"
import "fmt"
import "bufio"
import "strings"
import "os"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

var dictionary = make(map[string]string)

func main() {

  fmt.Println("Launching server... getting the dictionary ready...")
  make_dictionary()
  ln, _ := net.Listen("tcp", ":8081")
  conn, _ := ln.Accept()

  var answer_from_dictionary string
  for {
    word, _ := bufio.NewReader(conn).ReadString('\n')

    fmt.Print("Word Received:", word)
    if word == "n" {
      fmt.Println("\nClient closed the connection.\n")
      conn.Close()
      ln.Close()
    }else {
      answer_from_dictionary = get_meaning(strings.ToLower(word[:len(word) - 1]))
      conn.Write([]byte(answer_from_dictionary + "\n"))
    }

  }
}

func get_meaning(word string) string {

  if dictionary[word] != "" {
    return dictionary[word]
  } else{
    return "word incorrect"
  }

}

func make_dictionary() {

    f, err := os.Open("Oxford_English_Dictionary.txt")
    check(err)
    reader := bufio.NewReader(f)
    var line string
    line, err = reader.ReadString('\n')

    for {
      line, err = reader.ReadString('\n')
      word := strings.Split(line, " ")[0]
      meaning := strings.Join(strings.Split(line, " ")[1:], " ")
      dictionary[strings.ToLower(word)] = meaning
      if err != nil {
        break
      }
    }
    fmt.Println("Dictionary is ready.")
}
