package processorsProductValidators

import (
	"learning/services/requestService"
	"testing"
)

func Test_processorValidateCategoryId_IsValid(t *testing.T) {
	type args struct {
		params *requestService.RequestParams
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Тестирование c 0",
			args: args{
				params: &requestService.RequestParams{
					CategoryID: 0,
				},
			},
			want: false,
		},
		{
			name: "Тестирование с 1",
			args: args{
				params: &requestService.RequestParams{
					CategoryID: 1,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := processorValidateCategoryId{}
			if got := p.IsValid(tt.args.params); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}