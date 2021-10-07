package main

type BrowserHistory struct {
	History []string
	CurIdx  int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		History: []string{homepage},
		CurIdx:  0,
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.CurIdx++
	if this.CurIdx >= len(this.History) {
		this.History = append(this.History, url)
	} else {
		this.History[this.CurIdx] = url
		this.History = this.History[:this.CurIdx+1]
	}
}

func (this *BrowserHistory) Back(steps int) string {
	this.CurIdx -= steps
	if this.CurIdx < 0 {
		this.CurIdx = 0
	}
	return this.History[this.CurIdx]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.CurIdx += steps
	if this.CurIdx >= len(this.History) {
		this.CurIdx = len(this.History) - 1
	}
	return this.History[this.CurIdx]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
