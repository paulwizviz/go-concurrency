package traffic

import (
	"context"
	"fmt"
	"time"
)

const (
	LightRed = iota
	LightRedAmber
	LightGreen
	LightAmber
)

type LightState int

func (l LightState) String() string {
	switch l {
	case LightRedAmber:
		return fmt.Sprintf("%d %s", l, "RED AMBER")
	case LightGreen:
		return fmt.Sprintf("%d %s", l, "GREEN")
	case LightAmber:
		return fmt.Sprintf("%d %s", l, "AMBER")
	default:
		return fmt.Sprintf("%d %s", l, "RED")
	}
}

// Light represents a traffic light system
type Light struct {
	state LightState
}

func (l *Light) Transition() {
	switch l.state {
	case LightRed:
		l.state = LightRedAmber
	case LightRedAmber:
		l.state = LightGreen
	case LightGreen:
		l.state = LightAmber
	case LightAmber:
		l.state = LightRed
	}
}

func (l Light) String() string {
	return l.state.String()
}

func (l Light) State() LightState {
	return l.state
}

func (l Light) Run(ctx context.Context, interval time.Duration) <-chan LightState {
	c := make(chan LightState)
	go func() {
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-ticker.C:
				l.Transition()
				c <- l.state
			case <-ctx.Done():
				return
			}
		}
	}()
	return c
}

// InitLight initialize and instantiate a Light
func InitLight() Light {
	return Light{
		state: LightRed,
	}
}
