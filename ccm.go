package main

import "sync/atomic"

type CCM struct {
	MaxThreads int
	cThreads   int32
	recv       chan bool
}

func (c *CCM) Wait() {
	for {
		if atomic.LoadInt32(&(c.cThreads)) < int32(c.MaxThreads) {
			return
		}
		<-c.recv
	}
}

func (c *CCM) Add(i int) {
	atomic.AddInt32(&(c.cThreads), int32(i))
}

func (c *CCM) Done() {
	atomic.AddInt32(&(c.cThreads), -1)
	c.recv <- true
}

func NewCCM(threadLimit int) (c CCM) {
	c.MaxThreads = threadLimit
	c.recv = make(chan bool)
	return
}
