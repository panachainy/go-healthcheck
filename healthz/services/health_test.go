package services

import "testing"

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
