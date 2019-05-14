package main

import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing
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
  addr := net.UDPAddr{
    Port: 8081,
    IP: net.ParseIP("127.0.0.1"),
  }
  ln, err := net.ListenUDP("udp", &addr)
  check(err)
  
}

func get_meaning(word string) string {

  if dictionary[word] != "" {
    return dictionary[word]
  } else{
    return "incorrect"
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