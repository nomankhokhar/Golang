package main

type BrowserHistory struct {
	historyStack []string
	forwardStack []string
}

// Visit adds url to history and clears forward stack
func (b *BrowserHistory) Visit(url string) {
	b.historyStack = append(b.historyStack, url)
	b.forwardStack = nil // clear forward history on new visit
}

// Back moves current page to forwardStack, returns the previous page or nil
func (b *BrowserHistory) Back() *string {
	if len(b.historyStack) <= 1 {
		return nil
	}
	popUrl := b.historyStack[len(b.historyStack)-1]
	b.historyStack = b.historyStack[:len(b.historyStack)-1]
	b.forwardStack = append(b.forwardStack, popUrl)
	return &popUrl
}

// Forward moves last forward page back to historyStack, returns it or nil
func (b *BrowserHistory) Forward() *string {
	if len(b.forwardStack) == 0 {
		return nil
	}
	popUrl := b.forwardStack[len(b.forwardStack)-1]
	b.forwardStack = b.forwardStack[:len(b.forwardStack)-1]
	b.historyStack = append(b.historyStack, popUrl)
	return &popUrl
}

// GetHistory returns a copy of the history stack
func (b *BrowserHistory) GetHistory() []string {
	if len(b.historyStack) == 0 {
		return nil
	}
	copyHistoryStack := make([]string, len(b.historyStack))
	copy(copyHistoryStack, b.historyStack)
	return copyHistoryStack
}
