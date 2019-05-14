package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "net"
    "os"
)

func server(url string) {
    log.Printf("serving on %s\n", url)
    addr, err := net.ResolveUDPAddr("udp", url)
    errcheck(err)

    conn, err := net.ListenUDP("udp", addr)
    errcheck(err)
    defer close(conn)

    msg := make([]byte, 1024)

    for {
        n, retAddr, err := conn.ReadFromUDP(msg)
        errcheck(err)

        log.Printf("received %v bytes, ret addr %v, msg %s", n, retAddr, string(msg[:n]))

        reply := []byte(fmt.Sprintf("received from you: %v bytes", n))
        n, err = conn.WriteToUDP(reply, retAddr)
        errcheck(err)

        log.Printf("sent reply %v bytes\n", n)
    }
}

func client(url string) {
    log.Printf("client for server url: %s\n", url)

    addr, err := net.ResolveUDPAddr("udp", url)
    errcheck(err)

    conn, err := net.DialUDP("udp", nil, addr)
    errcheck(err)
    defer close(conn)

    msg := make([]byte, 512)
    scanner := bufio.NewScanner(os.Stdin)
    for {
        if scanner.Scan() {
            line := scanner.Text()

            n, err := conn.Write([]byte(line))
            errcheck(err)
            log.Printf("sent %d bytes \n", n)

            n, err = conn.Read(msg)
            errcheck(err)
            log.Printf("server replied with: %s \n", string(msg[:n]))
        }
    }
}

func main() {
    args := os.Args[1:]

    if len(args) < 2 {
        log.Fatal("expecting 2 arguments client or server first followed by the address:port")
    }

    switch args[0] {
    case "server":
        server(args[1])

    case "client":
        client(args[1])

    default:
        log.Fatal("first argument must be server or client")
    }

}

func errcheck(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func close(c io.Closer) {
    err := c.Close()
    if err != nil {
        log.Fatal()
    }
}