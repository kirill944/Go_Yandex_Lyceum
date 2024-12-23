package calculation

import (
	"errors"
	"testing"
)

func TestCalc(t *testing.T) {
	type args struct {
		expression string
		answer     float64
		err        error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "OK",
			args: args{
				expression: "2+2",
				answer:     4,
				err:        nil,
			},
		},
		{
			name: "Invalid syntax",
			args: args{
				expression: "22+-",
				answer:     0,
				err:        errors.New("invalid syntax"),
			},
		},
		{
			name: "Division by zero",
			args: args{
				expression: "2/0",
				answer:     0,
				err:        errors.New("division by zero"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := Calc(tt.args.expression)
			if err != nil && tt.args.err == nil {
				t.Errorf("Calc() error = %v, wantErr %v", err, tt.args.err)
			}
			if ans != tt.args.answer {
				t.Errorf("Calc() = %v, want %v", ans, tt.args.answer)
			}
		})
	}
}
