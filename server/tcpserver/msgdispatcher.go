package tcpserver

import (
	"github.com/name5566/leaf/log"
	"sync"
)

type MsgDispatcher struct {
	sync.RWMutex
	// id -> handler
	handlers map[interface{}]Handler
}

type Handler func(conn *Conn, msg interface{})

func (dispatcher *MsgDispatcher) RegHandler(id interface{}, handler Handler) {
	dispatcher.Lock()
	defer dispatcher.Unlock()

	if dispatcher.handlers == nil {
		dispatcher.handlers = make(map[interface{}]Handler)
	}
	if dispatcher.handlers[id] != nil {
		// TODO: file and line
		log.Error("handler %v already registered", id)
		return
	}
	dispatcher.handlers[id] = handler
}

func (dispatcher *MsgDispatcher) Handler(id interface{}) Handler {
	dispatcher.RLock()
	defer dispatcher.RUnlock()

	if dispatcher.handlers == nil {
		log.Debug("handler %v not found", id)
		return nil
	}
	handler := dispatcher.handlers[id]
	if handler == nil {
		log.Debug("handler %v not found", id)
		return nil
	}
	return handler
}
