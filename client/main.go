package main

import (
	"errors"
	"flag"
	"log"
	"net"
	"sync"
	"time"
)

var Timeout = 100 * time.Millisecond

func client(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	buf := make([]byte, 100)
	for {
		if err := conn.SetWriteDeadline(time.Now().Add(Timeout)); err != nil {
			return err
		}

		if _, err := conn.Write([]byte("ok")); err != nil {
			return err
		}

		if err := conn.SetReadDeadline(time.Now().Add(Timeout)); err != nil {
			return err
		}
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}
		if n != 2 {
			return errors.New("unexpected length")
		}
		if buf[0] != 'o' || buf[1] != 'k' {
			return errors.New("invalid response")
		}
	}
}

func main() {
	thread := 1
	host := ""
	flag.IntVar(&thread, "t", thread, "Threads")
	flag.StringVar(&host, "s", host, "Addr")
	flag.Parse()

	var wg sync.WaitGroup
	for i := 0; i < thread; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if err := client(host); err != nil {
				log.Print(err)
			}
			log.Printf("Finish thread %d", i)
		}(i)
	}
	wg.Wait()
}
