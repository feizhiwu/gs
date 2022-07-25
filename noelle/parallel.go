package noelle

import (
	"sync"
	"time"
)

// parallel instance, which executes pipelines by parallel
type parallel struct {
	wg        *sync.WaitGroup
	pipes     []*pipeline
	wgChild   *sync.WaitGroup
	children  []*parallel
	exception *Handler
}

// Active NewParallel creates a new parallel instance
func Active() *parallel {
	res := new(parallel)
	res.wg = new(sync.WaitGroup)
	res.wgChild = new(sync.WaitGroup)
	res.pipes = make([]*pipeline, 0, 10)
	return res
}

// Except set the exception handling routine, when unexpected panic occur
// this routine will be executed.
func (p *parallel) Except(f interface{}, args ...interface{}) *Handler {
	h := NewHandler(f, args...)
	p.exception = h
	return h
}

// Register add a new pipeline with a single handler info parallel
func (p *parallel) Register(f interface{}, args ...interface{}) *Handler {
	return p.Pipeline().Register(f, args...)
}

// Pipeline NewPipeline create a new pipeline of parallel
func (p *parallel) Pipeline() *pipeline {
	pipe := Pipeline()
	p.Add(pipe)
	return pipe
}

// Add new pipelines to parallel
func (p *parallel) Add(pipes ...*pipeline) *parallel {
	p.wg.Add(len(pipes))
	p.pipes = append(p.pipes, pipes...)
	return p
}

// NewChild create a new child of p
func (p *parallel) NewChild() *parallel {
	child := Active()
	child.exception = p.exception
	p.AddChildren(child)
	return child
}

// AddChildren add children to parallel to handle dependency
func (p *parallel) AddChildren(children ...*parallel) *parallel {
	p.wgChild.Add(len(children))
	p.children = append(p.children, children...)
	return p
}

// Run start up all the jobs
func (p *parallel) Run() {
	for _, child := range p.children {
		// this func will never panic
		go func(ch *parallel) {
			ch.Run()
			p.wgChild.Done()
		}(child)
	}
	p.wgChild.Wait()
	p.do()
	p.wg.Wait()
}

// Do just do it
func (p *parallel) do() {
	// if only one pipeline no need go routines
	if len(p.pipes) == 1 {
		p.secure(p.pipes[0])
		return
	}
	for _, pipe := range p.pipes {
		go p.secure(pipe)
	}
}

// exec pipeline safely
func (p *parallel) secure(pipe *pipeline) {
	defer func() {
		err := recover()
		if err != nil {
			if err == ErrArgNotFunction || err == ErrInArgLenNotMatch || err == ErrOutArgLenNotMatch || err == ErrRecvArgTypeNotPtr || err == ErrRecvArgNil {
				panic(err)
			}
			if p.exception != nil {
				p.exception.OnExcept(err)
			}
		}
		p.wg.Done()
	}()
	pipe.Do()
}

// RunWithTimeOut start up all the jobs, and time out after d duration
func (p *parallel) RunWithTimeOut(d time.Duration) {
	success := make(chan struct{}, 1)
	go func() {
		p.Run()
		success <- struct{}{}
	}()
	select {
	case <-success:
	case <-time.After(d):
	}
}
