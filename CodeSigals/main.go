package main

type StorageManager struct {
	warehouseQueue  []string
	inspectionStack []string
}

// StoreItem adds item to queue and stack
func (s *StorageManager) StoreItem(item string) {

	// Add item to stack
	s.inspectionStack = append(s.inspectionStack, item)

	// Add item to queue
	s.warehouseQueue = append(s.warehouseQueue, item)
}

// RetrieveItem gets FIRST item from queue (FIFO)
func (s *StorageManager) RetrieveItem() *string {

	// Queue is empty
	if len(s.warehouseQueue) == 0 {
		return nil
	}

	// Get first item
	firstItem := s.warehouseQueue[0]

	// Remove first item
	s.warehouseQueue = s.warehouseQueue[1:]

	return &firstItem
}

// InspectItem gets LAST item from stack (LIFO)
func (s *StorageManager) InspectItem() *string {

	// Stack is empty
	if len(s.inspectionStack) == 0 {
		return nil
	}

	// Get last item
	lastItem := s.inspectionStack[len(s.inspectionStack)-1]

	// Remove last item
	s.inspectionStack = s.inspectionStack[:len(s.inspectionStack)-1]

	return &lastItem
}

// GetAllItems returns copy of queue
func (s *StorageManager) GetAllItems() []string {

	// Return empty slice if queue is empty
	if len(s.warehouseQueue) == 0 {
		return []string{}
	}

	// Make new slice
	copyWarehouse := make([]string, len(s.warehouseQueue))

	// Copy data
	copy(copyWarehouse, s.warehouseQueue)

	return copyWarehouse
}
