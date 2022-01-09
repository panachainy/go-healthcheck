package externals

import "testing"

func TestGetHealthCheck(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "when_health_delay2sec_case_should_be_fail",
			args: args{
				url: "http://localhost:9091/healthz/delay2sec",
			},
			wantErr: true,
		},
		{
			name: "when_health_success_case_should_be_success",
			args: args{
				url: "http://localhost:9091/healthz/success",
			},
			wantErr: false,
		},
		{
			name: "when_health_not_exiting_case_should_be_fail",
			args: args{
				url: "http://localhost:9091/healthz/xxxxxxxx",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetHealthCheck(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("GetHealthCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
