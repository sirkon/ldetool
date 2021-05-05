package ldetesting

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldBeImportWithCustomType_Extract(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		wantRes ShouldBeImportWithCustomType
		want    bool
		wantErr bool
	}{
		{
			name: "only-it-needed",
			line: "1:2",
			wantRes: ShouldBeImportWithCustomType{
				Rest:   "",
				First:  time.Date(2019, 9, 4, 11, 6, 27, 0, time.UTC),
				Second: time.Date(2019, 9, 4, 11, 6, 37, 0, time.UTC),
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p ShouldBeImportWithCustomType
			got, err := p.Extract(tt.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				assert.Equal(t, tt.wantRes, p)
				t.Errorf("Extract() got = %v, want %v", got, tt.want)
			}
		})
	}
}
