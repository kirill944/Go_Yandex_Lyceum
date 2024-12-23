package application

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestCalcHandler(t *testing.T) {
	type args struct {
		expression string
		code       int
		result     float64
		err        string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "OK",
			args: args{
				expression: "2+2",
				code:       200,
				result:     4,
			},
		},
		{
			name: "Expression is not valid",
			args: args{
				expression: "2(",
				code:       422,
				err:        `Expression is not valid`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/api/v1/calculate", bytes.NewReader([]byte("{\"expression\": \""+tt.args.expression+"\"}")))
			req.Header.Set("Content-Type", "application/json")
			recorder := httptest.NewRecorder()
			CalcHandler(recorder, req)
			if recorder.Code != tt.args.code {
				t.Errorf("excepted status code %d, got %d", tt.args.code, recorder.Code)
			}
		})
	}
}
