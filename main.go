//server.go

package main

import (
    "fmt"
    "net"
    "os"
    "strconv"
    "bytes"
)

const (
    CONN_HOST = ""
    CONN_PORT = "80"
    CONN_TYPE = "tcp"
)

func main() {
    // Listen for incoming connections.
    l, err := net.Listen(CONN_TYPE, ":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }

        //logs an incoming message
        fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

func myHomeIP() string {
    ifaces, err := net.Interfaces()
    if err != nil {
        return "(can't enumerate net interfaces: " + err.Error() + ")"
    }

    for _, i := range ifaces {
        addrs, err := i.Addrs()
        if err != nil {
            continue
        }

        for _, addr := range addrs {
            var ip net.IP
            switch v := addr.(type) {
            case *net.IPNet:
                ip = v.IP
            case *net.IPAddr:
                ip = v.IP
            }

            if ip.IsLoopback() {
                continue
            }

            return ip.String()
        }
    }

    return "(no IP address found)"
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
  // Make a buffer to hold incoming data.
  buf := make([]byte, 1024)
  // Read the incoming connection into the buffer.
  reqLen, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading:", err.Error())
  }
  // Builds the message.
  message := " Upgrade Message was: "
  message += strconv.Itoa(reqLen)
  message += " bytes long and that's what it said: \""
  n := bytes.Index(buf, []byte{0})
  message += string(buf[:n-1])
  message += "My Docker IP address is: "
  message += myHomeIP()

  // Write the message in the connection channel.
  conn.Write([]byte(message));
  // Close the connection when you're done with it.
  conn.Close()
}
