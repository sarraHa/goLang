// This program extends part 6 by adding a Peers type.
// The rest of the code is left as-is, so functionally there is no change.
//
// However we have added a peers_test.go file, so that running
//   go test
// from the package directory will test your implementation of the Peers type.
//
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"github.com/sbinet/whispering-gophers/util"
)

var (
	peerAddr = flag.String("peer", "", "peer host:port")
	self     string
)

type Message struct {
	Addr string
	Body string
}

func main() {
	flag.Parse()

	l, err := util.Listen()
	if err != nil {
		log.Fatal(err)
	}
	self = l.Addr().String()
	log.Println("Listening on", self)

	go dial(*peerAddr)
	go readInput()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go serve(c)
	}
}

type Peers struct {
	m  map[string]chan<- Message
	mu sync.RWMutex
}

// Add creates and returns a new channel for the given peer address.
// If an address already exists in the registry, it returns nil.


//func Add(p *Peers , addr string ) <- chan Message
func (p *Peers) Add(addr string) <-chan Message {

	p.mu.Lock()		
	defer p.mu.Unlock()
	

	
	if _, ok := p.m[addr]; ok{
		
		return nil
	}

	// TODO: Make a new channel of messages
	msgs := make( chan Message)

	// TODO: Add it to the peers map
	p.m[addr] = msgs
	// TODO: Return the newly created channel.

	return msgs
}

// Remove deletes the specified peer from the registry.
func (p *Peers) Remove(addr string) {

	// TODO: Take the write lock on p.mu. Unlock it before returning (using defer).
	p.mu.Lock()
	defer p.mu.Unlock()
	// TODO: Delete the peer from the peers map.
	delete(p.m, addr);
}

// List returns a slice of all active peer channels.
// chanell dans lequel on peut que lire
func (p *Peers) List() []chan<- Message {

	// TODO: Take the read lock on p.mu. Unlock it before returning (using defer).
	p.mu.RLock()
	defer p.mu.RUnlock()

	// TODO: Declare a slice of chan<- Message.
	//slice []chan<- Message
	slice := make([]chan<-Message, 0, len(p.m))  



	for _,ch:= range p.m/* TODO: Iterate over the map using range */ {
		// TODO: Append each channel into the slice.
		slice = append(slice,  ch)	

	}

	// TODO: Return the slice.
	return slice
}

func serve(c net.Conn) {
	defer c.Close()
	d := json.NewDecoder(c)
	for {
		var m Message
		err := d.Decode(&m)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("%#v\n", m)
	}
}

var peer = make(chan Message)

func readInput() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		m := Message{
			Addr: self,
			Body: s.Text(),
		}
		peer <- m
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func dial(addr string) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(addr, err)
		return
	}
	defer c.Close()

	e := json.NewEncoder(c)

	for m := range peer {
		err := e.Encode(m)
		if err != nil {
			log.Println(addr, err)
			return
		}
	}
}

