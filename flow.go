package flow

import (
	"context"
)

type FlowFn func(context.Context) error

type flowSteps []FlowFn

func (s *flowSteps) ExecStepSeq(ctx context.Context) error {
	steps := *s
	for _, f := range steps {
		if err := f(ctx); err != nil {
			return err
		}
	}
	return nil
}

func FlowSteps(fns ...FlowFn) flowSteps {
	fn := make(flowSteps, len(fns))
	copy(fn, fns)
	return fn
}