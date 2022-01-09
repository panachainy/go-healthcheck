package services

import (
	"fmt"
	"go-healthcheck/healthz/dto"
	"go-healthcheck/healthz/externals"
	mockExternals "go-healthcheck/healthz/externals/mocks"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetHealthSummary(t *testing.T) {
	type args struct {
		healths []dto.Health
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func(ctrl *gomock.Controller) *mockExternals.MockIHealthService
		want     *dto.Summary
	}{
		{
			name: "when_call_with_success1_fail1_should_be_success1_fail1",
			args: args{
				healths: []dto.Health{{URL: "http://localhost:8080/success"}, {URL: "http://localhost:8080/fail"}},
			},
			mockFunc: func(ctrl *gomock.Controller) *mockExternals.MockIHealthService {
				mock := mockExternals.NewMockIHealthService(ctrl)
				mock.EXPECT().GetHealthCheck("http://localhost:8080/fail").Return(fmt.Errorf("Error : %v", "error ja"))
				mock.EXPECT().GetHealthCheck("http://localhost:8080/success").Return(nil)

				return mock
			},
			want: &dto.Summary{
				TotalWebsites: 2,
				Success:       1,
				Failure:       1,
				TotalTime:     0,
			},
		},
		{
			name: "when_call_with_success1_should_be_success1",
			args: args{
				healths: []dto.Health{{URL: "http://localhost:8080/success"}},
			},
			mockFunc: func(ctrl *gomock.Controller) *mockExternals.MockIHealthService {
				mock := mockExternals.NewMockIHealthService(ctrl)
				mock.EXPECT().GetHealthCheck("http://localhost:8080/success").Return(nil)

				return mock
			},
			want: &dto.Summary{
				TotalWebsites: 1,
				Success:       1,
				Failure:       0,
				TotalTime:     0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Assert mock
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			externals.Client = tt.mockFunc(ctrl)

			got := GetHealthSummary(tt.args.healths)
			// ignore expect time duration
			tt.want.TotalTime = got.TotalTime

			if !reflect.DeepEqual(got, tt.want) {
				// TODO: check why can't compare between to obj
				t.Errorf("GetHealthSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkCSVHeader(t *testing.T) {
	type args struct {
		line []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "when_call_with_url_header_should_return_nil",
			args:    args{line: []string{"url"}},
			wantErr: false,
		},
		{
			name:    "when_call_with_URL_header_should_return_nil",
			args:    args{line: []string{"URL"}},
			wantErr: false,
		},
		{
			name:    "when_call_with_Url_header_should_return_nil",
			args:    args{line: []string{"Url"}},
			wantErr: false,
		},
		{
			name:    "when_call_with_mistake_header_should_return_error",
			args:    args{line: []string{"mistake"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkCSVHeader(tt.args.line); (err != nil) != tt.wantErr {
				t.Errorf("checkCSVHeader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
