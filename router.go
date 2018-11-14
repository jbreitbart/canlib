package can

const ALLIDs = 0xFFFFFFFF

type Router struct {
	table         map[uint32]([]chan<- *Frame)
	subscribeCh   chan routerTask
	unsubscribeCh chan routerTask
	stopCh        chan bool
}

type routerTask struct {
	ID uint32
	Ch chan<- *Frame
}

func (r *Router) run(canFD InterfaceDescriptor) {

	canCh := make(chan *Frame, 100)
	errCh := make(chan error)

	go CaptureCan(canFD, canCh, errCh)

	for {
		stop := false
		select {
		case frame := <-canCh:
			for _, c := range r.table[ALLIDs] {
				// TODO select and skip if full?
				c <- frame
			}
			for _, c := range r.table[frame.ID()] {
				// TODO select and skip if full?
				c <- frame
			}

		case sub := <-r.subscribeCh:
			r.table[sub.ID] = append(r.table[sub.ID], sub.Ch)

		case unsub := <-r.unsubscribeCh:
			toRemove := -1
			for n, c := range r.table[unsub.ID] {
				if c == unsub.Ch {
					toRemove = n
				}
			}
			if toRemove != -1 {
				r.table[unsub.ID] = append(r.table[unsub.ID][:toRemove], r.table[unsub.ID][toRemove:]...)
			}
			// TODO handle error?

		case stop = <-r.stopCh:
		}
		if stop {
			break
		}
	}

	// TODO drain all channels

	// TODO find a way to end capture
}

// Subscribe adds ch to the subscriber list for id
func (r *Router) Subscribe(id uint32, ch chan<- *Frame) {
	t := routerTask{id, ch}
	r.subscribeCh <- t
}

// Unsubscribe removes ch from the subscriber list for id
func (r *Router) Unsubscribe(id uint32, ch chan<- *Frame) {
	t := routerTask{id, ch}
	r.unsubscribeCh <- t
}

// Stop will stop the router and drains all of its channels
func (r *Router) Stop() {
	r.stopCh <- true
}

//NewRouter creates a new router and starts it
func NewRouter(canFD InterfaceDescriptor) Router {
	var ret Router

	ret.table = make(map[uint32]([]chan<- *Frame))
	ret.subscribeCh = make(chan routerTask)
	ret.unsubscribeCh = make(chan routerTask)
	ret.stopCh = make(chan bool)

	go ret.run(canFD)

	return ret
}
