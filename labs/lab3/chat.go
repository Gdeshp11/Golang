// Demonstration of channels with a chat application
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Chat is a server that lets clients chat with each other.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	clientChan chan<- string // an outgoing message channel
	clientName string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.clientChan <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			cli.clientChan <- "Users Connected:"
			for entered_clients := range clients {
				cli.clientChan <- entered_clients.clientName
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.clientChan)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	tmp := client{}
	tmp.clientChan = ch
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	tmp.clientChan = ch
	tmp.clientName = who
	messages <- who + " has arrived"
	entering <- tmp

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- tmp
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {

	for msg := range ch {
		fmt.Println("in clientWriter..")
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
