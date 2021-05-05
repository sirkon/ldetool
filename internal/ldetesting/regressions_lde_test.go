package ldetesting

import (
	"testing"
)

func TestRegression1_Extract(t *testing.T) {
	type fields struct {
		Rest  string
		Pid   int32
		Comm  string
		State uint8
		Ppid  int32
	}
	tests := []struct {
		name    string
		fields  fields
		line    string
		want    bool
		wantErr bool
	}{
		{
			name: "sample",
			fields: fields{
				Rest:  "",
				Pid:   8266,
				Comm:  "(chrome)",
				State: 'S',
				Ppid:  3165,
			},
			line:    "8266 (chrome) S 3165 ",
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Regression1{
				Rest:  tt.fields.Rest,
				Pid:   tt.fields.Pid,
				Comm:  tt.fields.Comm,
				State: tt.fields.State,
				Ppid:  tt.fields.Ppid,
			}
			got, err := p.Extract(tt.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Extract() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFooBarBaz_Extract(t *testing.T) {
	tests := []struct {
		name      string
		fooBarBaz FooBarBaz
		line      string
		want      bool
		wantErr   bool
	}{
		{
			name: "all-right-Foobarbaz",
			fooBarBaz: FooBarBaz{
				Stuff: "Data",
				Bar: struct {
					Valid bool
					ID    struct {
						Valid     bool
						Foobarbaz string
					}
				}{
					Valid: true,
					ID: struct {
						Valid     bool
						Foobarbaz string
					}{
						Valid:     true,
						Foobarbaz: "13",
					},
				},
				Baz: "14",
			},
			line: `<foo>Data</foo><Bar foobar='{13}'/><baz>14</baz>`,
			want: true,
		},
		{
			name: "all-right-no-foobarbaz",
			fooBarBaz: FooBarBaz{
				Stuff: "Data",
				Bar: struct {
					Valid bool
					ID    struct {
						Valid     bool
						Foobarbaz string
					}
				}{
					Valid: true,
				},
				Baz: "14",
			},
			line: `<foo>Data</foo><Bar/><baz>14</baz>`,
			want: true,
		},
		{
			name: "all-right-just-foo-and-baz",
			fooBarBaz: FooBarBaz{
				Stuff: "Data",
				Baz:   "14",
			},
			line: `<foo>Data</foo><baz>14</baz>`,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &tt.fooBarBaz
			got, err := p.Extract(tt.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Extract() got = %v, want %v", got, tt.want)
			}
		})
	}
}
