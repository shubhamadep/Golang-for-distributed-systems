package main

import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing
import "os"
// import "time"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

var dictionary = make(map[string]string)

func main() {

  fmt.Println("Launching server... getting the dictionary ready...")
  make_dictionary()
  // p :=make([]byte, 2048)
  addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8081")
  check(err)
  conn, err := net.ListenUDP("udp", addr)
  check(err)
  defer conn.Close()

  var answer_from_dictionary string
  buffer := make([]byte, 1024)
  // conn.SetReadDeadline(time.Now().Add(15))

  for {
    n, retAddr, err := conn.ReadFromUDP(buffer)
    check(err)
    var word = string(buffer[:n])
    if word == "n" {
      fmt.Println("Client closed the connection.")
      conn.Close()
    }
    fmt.Println("Word recieved : ", word)
    answer_from_dictionary = get_meaning(strings.ToLower(word[:len(word) - 1]))
    meaning := []byte(fmt.Sprintf(answer_from_dictionary))
    _, err = conn.WriteToUDP(meaning, retAddr)
    check(err)
  }

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