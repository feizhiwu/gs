package noelle

// Pipeline instance, which executes jobs by serial
type pipeline struct {
	handlers []*Handler
}

// NewPipeline creates a new Pipeline instance
func Pipeline() *pipeline {
	res := new(pipeline)
	return res
}

// Register add a new function to pipeline
func (p *pipeline) Register(f interface{}, args ...interface{}) *Handler {
	h := NewHandler(f, args...)
	p.Add(h)
	return h
}

// Add add new handlers to pipeline
func (p *pipeline) Add(hs ...*Handler) *pipeline {
	p.handlers = append(p.handlers, hs...)
	return p
}

// Do calls all handlers as the sequence they are added into pipeline.
func (p *pipeline) Do() {
	for _, h := range p.handlers {
		h.Do()
	}
}
