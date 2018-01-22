package core

import (
	"encoding/json"
	"log"
	"net"
	"sync"
	"time"
)

const (
	producer = "p"
	consumer = "c"
)

// App is me core
type App struct {
	// topic list
	tList map[string]BaseQueue
	// listener port
	port     string
	listener net.Listener
	mu       sync.Mutex
}

// NewApp makes instance
func NewApp(p string) *App {
	return &App{
		tList: make(map[string]BaseQueue),
		port:  p,
	}
}

// Start service
func (a *App) Start() {
	log.Println("[Me] Starting service...")
	a.handler()
}

func (a *App) handler() {
	var err error
	a.listener, err = net.Listen("tcp", "0.0.0.0:"+a.port)
	if err != nil {
		log.Fatal("[Me] Listen failed, ", err)
	}
	a.handle()
}

func (a *App) checkTopic(t string) bool {
	defer a.mu.Unlock()
	a.mu.Lock()
	if _, found := a.tList[t]; found {
		return true
	}
	return false
}

func (a *App) newQueue(t string) BaseQueue {
	defer a.mu.Unlock()
	a.mu.Lock()
	log.Println("[Me] Init a new topicQueue: ", t)
	queue := NewMeQueue()
	a.tList[t] = queue
	return queue
}

func (a *App) handle() {
	defer a.listener.Close()
	for {
		conn, err := a.listener.Accept()
		if err != nil {
			log.Println("[Me] Accept request failed, ", err)
		}
		go a.handleConn(conn)
	}
}

func (a *App) handleConn(conn net.Conn) {
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	var r Request
	err := json.Unmarshal(buf[:n], &r)
	if err != nil {
		conn.Write([]byte("Invaild Request!"))
		conn.Close()
		return
	}
	switch r.Role {
	case producer:
		go a.handleProducer(conn, r.Topic)
	case consumer:
		go a.handleConsumer(conn, r.Topic)
	}
}

func (a *App) handleProducer(conn net.Conn, t string) {
	var queue BaseQueue
	if !a.checkTopic(t) {
		queue = a.newQueue(t)
	} else {
		queue = a.tList[t]
	}
	for {
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		queue.Push(buf[:n])
	}
}

func (a *App) handleConsumer(conn net.Conn, t string) {
	if !a.checkTopic(t) {
		conn.Write([]byte("No that topic!"))
		conn.Close()
		return
	}
	msgChannel := make(chan interface{}, 1)
	queue := a.tList[t]
	go func() {
		for {
			if !queue.Empty() {
				msgChannel <- queue.Pop()
			} else {
				<-time.After(100 * time.Millisecond)
			}
		}
	}()
	for {
		select {
		case v := <-msgChannel:
			if res, ok := v.([]byte); ok {
				conn.Write(res)
			}
		}
	}
}
