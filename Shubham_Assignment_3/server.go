package main

import "net"
import "fmt"
import "bufio"
import "strings"
import "os"
import "runtime"
// import "sync"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

var dictionary = make(map[string]string)

func main() {

  runtime.GOMAXPROCS(4)
  //count := 5
  //collect := make(chan string, count)
  
  fmt.Println("Launching server... getting the dictionary ready...")
  make_dictionary()

  ln, _ := net.Listen("tcp", ":8081")
  defer ln.Close()
  threadNumber := 1
  for {
    conn, err := ln.Accept()
    check(err)
    go handleConnections(conn, threadNumber)
    threadNumber++
  }

}

func handleConnections(conn net.Conn, threadNumber int){

  defer conn.Close()
  word, err := bufio.NewReader(conn).ReadString('\n')
  check(err)
  fmt.Println("Thread : ",threadNumber," is searching the word : ",word)
  fmt.Fprintf(conn, get_meaning(word)+"\n")
  

}

func get_meaning(w string) string{ //collect chan<- string, 

  word := strings.ToLower(w[:len(w) - 1])
  if dictionary[word] != "" {
    return word+" : "+dictionary[word]
  } else{
    return word+" : "+"Please check the spelling again."
    //collect <- fmt.Sprintf(word," : word incorrect")
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