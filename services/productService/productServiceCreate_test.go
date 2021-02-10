package productService

import (
	"errors"
	processorsProductConditions "learning/processors/productProcessors/conditions"
	processorsProductValidators "learning/processors/productProcessors/validators"
	"learning/repository/productRepository"
	"learning/services/requestService"
	"reflect"
	"testing"
)

func Test_productService_CreateProduct(t *testing.T) {
	type fields struct {
		processorsCond     []processorsProductConditions.ProcessorProductConditionInterface
		processorsValidate []processorsProductValidators.ProcessorProductValidateInterface
	}
	type args struct {
		params *requestService.RequestParams
	}
	var testRequestParams = &requestService.RequestParams{
		CategoryID	: 1,
		SKU			: "test",
		Name		: "test",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *productRepository.Product
		wantErr bool
	}{
		{
			name: "Тестирование без процессоров",
			fields: fields{
				processorsCond:     nil,
				processorsValidate: nil,
			},
			args: args{
				params: testRequestParams,
			},
			want : &productRepository.Product{
				CategoryID	: 1,
				SKU			: "test",
				Name		: "test",
			},
		},
		{
			name: "Тестирование с ошибкой процессора",
			fields: fields{
				processorsCond:     nil,
				processorsValidate: []processorsProductValidators.ProcessorProductValidateInterface{
					processorsProductValidators.ProcessorValidatorsMock{
						IsValidResult:  false,
						GetErrorResult: errors.New("test_err"),
					},
				},
			},
			args: args{
				params: &requestService.RequestParams{},
			},
			want : nil,
			wantErr: true,
		},
		{
			name: "Тестирование без ошибок процессора",
			fields: fields{
				processorsCond:     nil,
				processorsValidate: []processorsProductValidators.ProcessorProductValidateInterface{
					processorsProductValidators.ProcessorValidatorsMock{
						IsValidResult:  true,
						GetErrorResult: nil,
					},
				},
			},
			args: args{
				params: testRequestParams,
			},
			want : &productRepository.Product{
				CategoryID	: 1,
				SKU			: "test",
				Name		: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &productService{
				processorsCond:     tt.fields.processorsCond,
				processorsValidate: tt.fields.processorsValidate,
			}
			got, err := s.CreateProduct(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}