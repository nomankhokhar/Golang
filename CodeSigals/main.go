package main

type TicketBookingSystem struct {
	priorityQueue []string
	regularQueue  []string
}

func (t *TicketBookingSystem) AddCustomer(customer string, isPriority bool) {
	// TODO: Append the customer to priorityQueue if isPriority is true, otherwise to regularQueue.
	if isPriority {
		t.priorityQueue = append(t.priorityQueue, customer)
		return
	}
	t.regularQueue = append(t.regularQueue, customer)
}

func (t *TicketBookingSystem) ServeCustomer() (string, bool) {
	// TODO: Serve first customer from priorityQueue if not empty, otherwise from regularQueue.
	// Return the customer's name and true on success. If queues are empty, return empty string and false.
	if len(t.priorityQueue) >= 1 {
		nameofCus := t.priorityQueue[0]
		t.priorityQueue = t.priorityQueue[1:]
		return nameofCus, true
	}
	if len(t.regularQueue) >= 1 {
		nameofCus := t.regularQueue[0]
		t.regularQueue = t.regularQueue[1:]
		return nameofCus, true
	}

	return "", false
}

func (t *TicketBookingSystem) NextCustomer() (string, bool) {
	// TODO: Return the name of the first customer from priorityQueue if not empty, otherwise from regularQueue.
	// If queues are empty, return empty string and false.
	if len(t.priorityQueue) >= 1 {
		return t.priorityQueue[0], true
	}
	if len(t.regularQueue) >= 1 {
		return t.regularQueue[0], true
	}
	return "", false
}
