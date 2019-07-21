package ldetesting

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCommon(t *testing.T) {
	data := `[bugaga] -123  234 abcdef`
	e := &Rule{}
	if ok, err := e.Extract(data); !ok || err != nil {
		if err != nil {
			t.Fatal(err)
		}
		require.True(t, ok)
	}
	require.Equal(t, "bugaga", e.Data)
	require.Equal(t, -123, e.Signed)
	require.Equal(t, uint(234), e.Unsigned)
	require.Equal(t, "abcdef", e.Str)
}

func TestBeforeLookup(t *testing.T) {
	data := "     abc123"
	var e BeforeLookup
	if ok, err := e.Extract(data); !ok || err != nil {
		if err != nil {
			t.Fatal(err)
		}
		require.True(t, ok)
	}
	require.Equal(t, "abc123", e.Data)
	require.Equal(t, "", e.Rest)
}

func TestCheckPrefix(t *testing.T) {
	data := "abc123"
	var e CheckPrefix
	if ok, err := e.Extract(data); !ok || err != nil {
		if err != nil {
			t.Fatal(err)
		}
		require.True(t, ok)
	}
	require.Equal(t, "abc123", e.Data)
	require.Equal(t, "", e.Rest)
}

func TestPassHeadingStringRegression(t *testing.T) {
	data := "#################################3 123"
	var e PassHeadingStringRegression
	if ok, err := e.Extract(data); !ok {
		if err != nil {
			t.Fatal(err)
		}
		require.True(t, ok)
	}
	require.Equal(t, "3 123", e.Data)
	require.Equal(t, "", e.Rest)
}

func TestRegressionCheck1(t *testing.T) {
	var rc RegressionCheck1

	if ok, err := rc.Extract("17.965 Pump 10 State change LOCKED_PSTATE to CALLING_PSTATE [31]"); !ok {
		if err != nil {
			t.Fatal(err)
		}
		require.True(t, ok)
	}
	require.Equal(t, "17.965", rc.Time)
	require.Equal(t, int8(10), rc.Pump)
	require.Equal(t, "CALLING_PSTATE ", rc.GetPStateState())
	require.Equal(t, "", rc.GetIStateState())

	if ok, err := rc.Extract("19.996 Pump 10 change internal state AUTHORISE_ISTATE to IDLE_ISTATE"); !ok {
		if err != nil {
			t.Fatal(err)
		}
		require.True(t, ok)
	}
	require.Equal(t, "19.996", rc.Time)
	require.Equal(t, int8(10), rc.Pump)
	require.Equal(t, "", rc.GetPStateState())
	require.Equal(t, "IDLE_ISTATE", rc.GetIStateState())
}

func TestRegressionCheck2(t *testing.T) {
	var rc RegressionCheck2
	if ok, err := rc.Extract("ï»¿*** Time: 2/1/2019 12:10:17"); !ok {
		if err != nil {
			t.Fatal(err)
		}
		require.True(t, ok)
	}
	require.Equal(t, "2/1/2019 12:10:17", rc.Time)
}

func TestRegressionCheck3_Extract(t *testing.T) {
	type fields struct {
		Rest string
	}

	tests := []struct {
		name    string
		fields  fields
		line    string
		wantOK  bool
		wantErr bool
	}{
		{
			name: "match-both",
			fields: fields{
				Rest: "ab:cd",
			},
			line:    "ab:cd",
			wantOK:  true,
			wantErr: false,
		},
		{
			name: "match-first-only",
			fields: fields{
				Rest: "ab:bd",
			},
			line:    "ab:bd",
			wantOK:  false,
			wantErr: false,
		},
		{
			name: "match-last-only",
			fields: fields{
				Rest: "ac:cd",
			},
			line:    "ac:cd",
			wantOK:  false,
			wantErr: false,
		},
		{
			name: "not-match-both",
			fields: fields{
				Rest: "1234",
			},
			line:    "1234",
			wantOK:  false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &RegressionCheck3{
				Rest: tt.fields.Rest,
			}
			got, err := p.Extract(tt.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegressionCheck3.Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantOK {
				t.Errorf("RegressionCheck3.Extract() = %v, wantOK %v", got, tt.wantOK)
			}
		})
	}
}

func TestCustom(t *testing.T) {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		t.Fatal(err)
	}
	sampleTime := time.Date(2019, 7, 20, 14, 41, 04, 0, loc)
	line := sampleTime.Format(time.RFC3339) + " addr: 10.20.30.40 ze rest"

	var e Custom
	if ok, err := e.Extract(line); !ok {
		if err != nil {
			t.Fatal(err)
		}
		require.True(t, ok)
	}

	require.Equal(t, "ze rest", e.Rest)
	if !sampleTime.Equal(e.Time) {
		t.Errorf("%s != %s", sampleTime.Format(time.RFC3339), e.Time.Format(time.RFC3339))
	}
	require.Equal(t, "10.20.30.40", e.GetAddrIP().String())
}

func TestCustomBuiltin(t *testing.T) {
	var e CustomBuiltin

	if ok, err := e.Extract("12"); !ok {
		if err != nil {
			t.Fatal(err)
		}
		require.True(t, ok)
	}
	require.Equal(t, CustomBuiltin{
		Rest:  "",
		Field: 12,
	}, e)

	if ok, err := e.Extract("12ab"); ok {
		t.Errorf("should not be here")
	} else {
		require.Error(t, err)
	}
}

func TestBoolean_Extract(t *testing.T) {
	type fields struct {
		Rest  string
		Check bool
	}
	tests := []struct {
		name    string
		fields  fields
		line    string
		wantOK  bool
		wantErr bool
	}{
		{
			name: "ok-true",
			fields: fields{
				Rest:  "",
				Check: true,
			},
			line:    "1",
			wantOK:  true,
			wantErr: false,
		},
		{
			name: "ok-false",
			fields: fields{
				Rest:  "",
				Check: false,
			},
			line:    "0",
			wantOK:  true,
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				Rest:  "abc",
				Check: false,
			},
			line:    "abc",
			wantOK:  false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Boolean{
				Rest:  tt.fields.Rest,
				Check: tt.fields.Check,
			}
			got, err := p.Extract(tt.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("Boolean.Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantOK {
				t.Errorf("Boolean.Extract() = %v, wantOK %v", got, tt.wantOK)
			}
		})
	}
}

func TestSilentAreas_Extract(t *testing.T) {
	type fields struct {
		Rest string
		Alt1 struct {
			Valid  bool
			Amount int
		}
		Alt2 struct {
			Valid  bool
			Amount string
		}
	}
	tests := []struct {
		name    string
		fields  fields
		line    string
		wantOK  bool
		wantErr bool
	}{
		{
			name: "alt1",
			fields: fields{
				Rest: "",
				Alt1: struct {
					Valid  bool
					Amount int
				}{
					Valid:  true,
					Amount: 123,
				},
				Alt2: struct {
					Valid  bool
					Amount string
				}{
					Valid:  false,
					Amount: "",
				},
			},
			line:    "Amount: 123",
			wantOK:  true,
			wantErr: false,
		},
		{
			name: "alt2",
			fields: fields{
				Rest: "",
				Alt1: struct {
					Valid  bool
					Amount int
				}{
					Valid:  false,
					Amount: 0,
				},
				Alt2: struct {
					Valid  bool
					Amount string
				}{
					Valid:  true,
					Amount: "123USD",
				},
			},
			line:    "Amount: 123USD",
			wantOK:  true,
			wantErr: false,
		},
		{
			name: "no-alt",
			fields: fields{
				Rest: "Amount:Nothing",
				Alt1: struct {
					Valid  bool
					Amount int
				}{
					Valid:  false,
					Amount: 0,
				},
				Alt2: struct {
					Valid  bool
					Amount string
				}{
					Valid:  false,
					Amount: "",
				},
			},
			line:    "Amount:Nothing",
			wantOK:  true,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &SilentAreas{
				Rest: tt.fields.Rest,
				Alt1: tt.fields.Alt1,
				Alt2: tt.fields.Alt2,
			}
			got, err := p.Extract(tt.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("SilentAreas.Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantOK {
				t.Errorf("SilentAreas.Extract() = %v, wantOK %v", got, tt.wantOK)
			}
		})
	}
}

func TestTrickyDateParsing_Extract(t *testing.T) {
	type fields struct {
		Rest string
		Full struct {
			Valid     bool
			Day       int
			Month     int
			Year      int
			Hour      int
			Minute    int
			Second    int
			Microsecs int
		}
		Hour struct {
			Valid     bool
			Hour      int
			Minute    int
			Second    int
			Microsecs int
		}
		Seconds struct {
			Valid     bool
			Second    int
			Microsecs int
		}
	}
	tests := []struct {
		name    string
		fields  fields
		line    string
		wantOK  bool
		wantErr bool
	}{
		{
			name: "full-date",
			fields: fields{
				Rest: "",
				Full: struct {
					Valid     bool
					Day       int
					Month     int
					Year      int
					Hour      int
					Minute    int
					Second    int
					Microsecs int
				}{
					Valid:     true,
					Day:       15,
					Month:     5,
					Year:      2019,
					Hour:      6,
					Minute:    42,
					Second:    22,
					Microsecs: 841,
				},
			},
			line:    "15/05/2019 06:42:22.841",
			wantOK:  true,
			wantErr: false,
		},
		{
			name: "hour",
			fields: fields{
				Rest: "",
				Hour: struct {
					Valid     bool
					Hour      int
					Minute    int
					Second    int
					Microsecs int
				}{
					Valid:     true,
					Hour:      12,
					Minute:    10,
					Second:    17,
					Microsecs: 0,
				},
			},
			line:    "12:10:17.000",
			wantOK:  true,
			wantErr: false,
		},
		{
			name: "second",
			fields: fields{
				Rest: "",
				Seconds: struct {
					Valid     bool
					Second    int
					Microsecs int
				}{
					Valid:     true,
					Second:    19,
					Microsecs: 996,
				},
			},
			line:    "19.996",
			wantOK:  true,
			wantErr: false,
		},
		{
			name:    "fail",
			line:    "19:19",
			wantOK:  false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &TrickyDateParsing{
				Rest:    tt.fields.Rest,
				Full:    tt.fields.Full,
				Hour:    tt.fields.Hour,
				Seconds: tt.fields.Seconds,
			}
			got, err := p.Extract(tt.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("TrickyDateParsing.Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantOK {
				t.Errorf("TrickyDateParsing.Extract() = %v, want %v", got, tt.wantOK)
			}
		})
	}
}
