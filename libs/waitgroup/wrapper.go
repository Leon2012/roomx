package waitgroup

import "sync"

type Wrapper struct {
	sync.WaitGroup
}

func (w *Wrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}
