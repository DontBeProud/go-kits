// 集成协程池，便于管理EG内部的协程使用

package error_ex

import (
	"context"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
)

type token struct{}

// A ErrorGroupEx is a collection of goroutines working on subtasks that are part of
// the same overall task.
//
// A zero ErrorGroupEx is valid, has no limit on the number of active goroutines,
// and does not cancel on error.
type ErrorGroupEx struct {
	cancel  func(error)
	wg      sync.WaitGroup
	sem     chan token
	errOnce sync.Once
	err     error
	pool    *ants.Pool
}

func (g *ErrorGroupEx) done() {
	if g.sem != nil {
		<-g.sem
	}
	g.wg.Done()
}

// NewErrorGroupWithContext returns a new ErrorGroupEx and an associated Context derived from ctx.
//
// The derived Context is canceled the first time a function passed to Go
// returns a non-nil error or the first time Wait returns, whichever occurs
// first.
//
// 如果传入有效的pool，则EG使用的协程从协程池中获取
func NewErrorGroupWithContext(ctx context.Context, pool *ants.Pool) (*ErrorGroupEx, context.Context) {
	ctx, cancel := withCancelCause(ctx)
	return &ErrorGroupEx{cancel: cancel, pool: pool}, ctx
}

func withCancelCause(parent context.Context) (context.Context, func(error)) {
	return context.WithCancelCause(parent)
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (g *ErrorGroupEx) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel(g.err)
	}
	return g.err
}

// Go calls the given function in a new goroutine.
// It blocks until the new goroutine can be added without the number of
// active goroutines in the group exceeding the configured limit.
//
// The first call to return a non-nil error cancels the group's context, if the
// group was created by calling WithContext. The error will be returned by Wait.
//
// pool模式下，若pool处于关闭状态，会返回error
func (g *ErrorGroupEx) Go(f func() error) error {
	if g.sem != nil {
		g.sem <- token{}
	}

	g.wg.Add(1)
	fn := func() {
		defer g.done()
		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel(g.err)
				}
			})
		}
	}

	if g.pool == nil {
		go fn()
		return nil
	} else {
		return g.pool.Submit(fn)
	}
}

// TryGo calls the given function in a new goroutine only if the number of
// active goroutines in the group is currently below the configured limit.
//
// The return value reports whether the goroutine was started.
// // pool模式下，若pool处于关闭状态，会返回error
func (g *ErrorGroupEx) TryGo(f func() error) (bool, error) {
	if g.sem != nil {
		select {
		case g.sem <- token{}:
			// Note: this allows barging iff channels in general allow barging.
		default:
			return false, nil
		}
	}

	g.wg.Add(1)

	fn := func() {
		defer g.done()
		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel(g.err)
				}
			})
		}
	}

	if g.pool == nil {
		go fn()
		return true, nil
	} else {
		err := g.pool.Submit(fn)
		return err == nil, err
	}
}

// SetLimit limits the number of active goroutines in this group to at most n.
// A negative value indicates no limit.
//
// Any subsequent call to the Go method will block until it can add an active
// goroutine without exceeding the configured limit.
//
// The limit must not be modified while any goroutines in the group are active.
func (g *ErrorGroupEx) SetLimit(n int) {
	if n < 0 {
		g.sem = nil
		return
	}
	if len(g.sem) != 0 {
		panic(fmt.Errorf("errgroup: modify limit while %v goroutines in the group are still active", len(g.sem)))
	}
	g.sem = make(chan token, n)
}
