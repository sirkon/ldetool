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
