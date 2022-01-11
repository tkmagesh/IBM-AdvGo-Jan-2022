package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"pool-demo/pool"
	"sync"
	"time"
)

//resource
type DBConnection struct {
	ID int
}

func (c *DBConnection) Close() error {
	fmt.Printf("Closing %d\n", c.ID)
	return nil
}

//resouce factory
var IDCounter int

func DBConnectionFactory() (io.Closer, error) {
	IDCounter++
	fmt.Printf("DBConnectionFactory : Creating resource %d\n", IDCounter)
	return &DBConnection{ID: IDCounter}, nil
}

func main() {
	p, err := pool.New(DBConnectionFactory /* factory */, 3 /* size */)
	if err != nil {
		log.Fatalln(err)
	}
	wg := &sync.WaitGroup{}
	clientCount := 10
	wg.Add(clientCount)
	for client := 0; client < clientCount; client++ {
		go func(client int) {
			doWork(client, p)
			wg.Done()
		}(client)
	}
	wg.Wait()

	fmt.Println("Second batch of operations")
	var input string
	fmt.Scanln(&input)
	wg = &sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(client int) {
			doWork(client, p)
			wg.Done()
		}(i)
	}
	wg.Wait()
	p.Close()
}

func doWork(id int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Fatalln(err)
	}
	defer p.Release(conn)
	fmt.Printf("Worker : %d, Acquired %d:\n", id, conn.(*DBConnection).ID)
	//use connection
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("Worker Done : %d, Released %d:\n", id, conn.(*DBConnection).ID)
}
