package main

import "net"
import "fmt"
import "bufio"
import "strings"
import "os"
//import "golang.org/x/sync/semaphore"
//import "context"
import "runtime"
// import "time"
// import "sync"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

var dictionary = make(map[string]string)

func main() {

  runtime.GOMAXPROCS(0)
  maxWorkers := runtime.GOMAXPROCS(0)
  fmt.Println(maxWorkers)
  writing := 0
  // var lock = sync.RWMutex{}
  //count := 5
  //collect := make(chan string, count)

  fmt.Println("Launching server... getting the dictionary ready...")

  ln, _ := net.Listen("tcp", ":8081")
  ln_add, _ := net.Listen("tcp", ":8082")
  //defer ln.Close()
  defer ln_add.Close()
  threadNumber := 1
  ReadthreadNumber := 1
  make_dictionary()
  make_dictionary()
  //var sem = semaphore.NewWeighted(1)
  var sem_add = make(chan int, 1)
  for {
    //conn, err := ln.Accept()
    conn_add, _ := ln_add.Accept()
    //defer conn.Close()
    defer conn_add.Close()

    sem_add <- 1
    fmt.Println("waiting for semaphone. Thread :", threadNumber)
    go func(){
      fmt.Println("Changing dictionary. Thread :", threadNumber)
      writing = 1
      word, _ := bufio.NewReader(conn_add).ReadString('\n')
      dictionary[word] = "Word added."
      <- sem_add
      writing = 0
      fmt.Println("Releasing semaphore. Thread :", threadNumber)
      fmt.Fprintf(conn_add, dictionary[word]+"\n")
    }()

    conn, _ := ln.Accept()
    defer conn.Close()
    go handleConnections(conn, ReadthreadNumber)
    // check(err)
    // ctx := context.Background()
    // fmt.Print(err)
    // word, err := bufio.NewReader(conn).ReadString('\n')
    // check(err)
    // if err = sem.Acquire(ctx, 1); err != nil {
    //   fmt.Print("semaphore is full, hence thread is blocked until vacancy for accessing resource is created.")
    // }
    // go func(){
    //   defer sem.Release(1)
    //   if writing == 1{
    //     time.Sleep(10000000000)
    //   }else {
    //     word := strings.ToLower(word[:len(word) - 1])
    //     // dictionary[word] = "yolo"
    //     // lock.RLock()
    //     // defer lock.Unlock()
    //     // handleConnections(conn, threadNumber)
    //     fmt.Fprintf(conn, dictionary[word]+"\n")
    //   }
    //
    //   //<-sem
    // }()

      //go handleConnections(conn, threadNumber)
      //sem.Release(1)
    threadNumber++
    ReadthreadNumber++
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
    f.Close()
}
