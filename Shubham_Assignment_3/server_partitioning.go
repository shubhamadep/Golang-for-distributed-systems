package main

import "net"
import "fmt"
import "bufio"
import "strings"
import "os"
// import "sync"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

var dictionary = make(map[string]string)
var threadNumber int = 1

func main() {

  //count := 5
  //collect := make(chan string, count)
  fmt.Println("Launching server... getting the dictionary ready...")
  make_dictionary()


  ln, _ := net.Listen("tcp", ":8081")
  defer ln.Close()
  for {
    conn, err := ln.Accept()
    check(err)
    go handleConnections(conn)

  }

}

func handleConnections(conn net.Conn){

  
  word, err := bufio.NewReader(conn).ReadString('\n')
  check(err)
  fmt.Println("Goroutine : ", threadNumber, " handling connection for word : ", word)
  threadNumber++
  go get_meaning(word, conn)

}

func get_meaning(w string, conn net.Conn) { //collect chan<- string, 

  word := strings.ToLower(w[:len(w) - 1])
  fmt.Println("Goroutine : ", threadNumber, " getting meaning for word : ", word)
  threadNumber++
  if dictionary[word] != "" {
    //fmt.Println(dictionary[word])
    go send(word+" : "+dictionary[word], conn)
  } else{
    go send(word+" : "+"Please check the spelling again.", conn)
    //collect <- fmt.Sprintf(word," : word incorrect")
  }

}

func send(meaning string ,conn net.Conn){
  defer conn.Close()
  fmt.Println("Goroutine : ", threadNumber, " sending result to the client : ", conn)
  threadNumber++
  fmt.Fprintf(conn, meaning+"\n")
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