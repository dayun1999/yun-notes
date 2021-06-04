/*
定义一个计时器
*/
package stopwatch

import "time"

type StopWatch struct {
	created time.Time
}

func New() *StopWatch {
	return &StopWatch{created: time.Now()}
}

// return the elapsed time since the stopwatch created
func (w *StopWatch) Elapsed() time.Duration {
	return time.Since(w.created)
}