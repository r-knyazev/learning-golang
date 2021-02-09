package processorsProductValidators

import (
	"learning/services/requestService"
	"testing"
)

func Test_processorValidateName_IsValid(t *testing.T) {
	type args struct {
		params *requestService.RequestParams
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Тестирование с 1",
			args: args{
				params: &requestService.RequestParams{
					Name: "1",
				},
			},
			want: true,
		},
		{
			name: "Тестирование с пустой строкой",
			args: args{
				params: &requestService.RequestParams{
					Name: "",
				},
			},
			want: false,
		},
		{
			name: "Тестирование с кириллицей",
			args: args{
				params: &requestService.RequestParams{
					Name: "Тестирование",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := processorValidateName{}
			if got := p.IsValid(tt.args.params); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}