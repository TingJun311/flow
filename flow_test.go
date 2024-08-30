package flow

import (
	"context"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_FlowSteps(t *testing.T) {
	testTable := []struct {
		name     string
		input      []FlowFn
		expectedLength int
	}{
		{
			name: "Should return len of 3",
			input: []FlowFn{
				func(ctx context.Context) error {
					fmt.Println("Step 1")
					return nil
				},
				func(ctx context.Context) error {
					fmt.Println("Step 2")
					return nil
				},
				func(ctx context.Context) error {
					fmt.Println("Step 3")
					return nil
				},
			},
			expectedLength: 3,
		},
		{
			name: "Should return empty flowfn",
			input: nil,
			expectedLength: 0,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			flow := FlowSteps(tc.input...)
			assert.Len(t, tc.input, len(tc.input), flow)
		})
	}
}