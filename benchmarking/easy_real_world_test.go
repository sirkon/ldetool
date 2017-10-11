package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// UAPiece user agent generator prototype
type UAPiece interface {
	UserAgentPiece() string
	Geo() (float64, float64)
}

type geo struct{}

func (g geo) Geo() (float64, float64) {
	return 30. + rand.Float64()*30, 45. + rand.Float64()*30
}

// AndroidUA android device UAPiece generator
type AndroidUA struct {
	geo
}

var androidDevs = []string{
	"Samsung S7",
	"Samsung GT-i9000",
	"Sony XT-5",
	"Huawei-P6/2016",
	"Xiaomi Mi-6",
	"Nexus 5 32Gb",
	"Pixel 2 XL 128Gb",
}
var androidVersions = []string{
	"2.3.0",
	"8.0.0",
	"6.1.1",
	"7.1.1",
	"7.0.1",
	"5.1.9",
}

func (android AndroidUA) UserAgentPiece() string {
	index := rand.Int() % len(androidDevs)
	dev := androidDevs[index]
	index = rand.Int() % len(androidVersions)
	os := androidVersions[index]
	return fmt.Sprintf("%s Android/%s", dev, os)
}

// IPhoneUA device UAPiece generator
type IPhoneUA struct {
	geo
}

var iphoneDevices = []string{
	"iPhone 8+",
	"iPhone X",
	"iPhone 7",
	"iPhone 6s+",
	"iPhone 5",
}

var iosVersions = []string{
	"11.0.2",
	"10.1.2",
	"9.1.1",
}

func (i IPhoneUA) UserAgentPiece() string {
	index := rand.Int() % len(iphoneDevices)
	dev := iphoneDevices[index]
	index = rand.Int() % len(iosVersions)
	os := iosVersions[index]
	return fmt.Sprintf("%s iOS/%s", dev, os)
}

// WindowsUA ...
type WindowsUA struct{}

var windowsVersions = []string{
	"10.0.1",
	"7.1.1",
	"XP SP3",
	"Vista SP1",
}

func (w WindowsUA) UserAgentPiece() string {
	index := rand.Int() % len(windowsVersions)
	os := windowsVersions[index]
	return fmt.Sprintf("Windows/%s", os)
}

func (w WindowsUA) Geo() (float64, float64) {
	return 0.0, 0.0
}

// FullUA is not a UAPiece
type FullUA struct{}

var countries = []string{
	"RU", "US", "BR", "BY", "DE", "UK", "AU", "CZ", "PL", "KZ", "UA", "TZ", "NZ", "FR", "TR", "RO",
}

// UserAgent blah-blah-blah
func (fua FullUA) UserAgent(piece UAPiece) string {
	index := rand.Int() % len(countries)
	country := countries[index]
	return fmt.Sprintf("App.Com %s/%s", piece.UserAgentPiece(), country)
}

var easylines [][]byte

var START = time.Date(2015, 1, 1, 1, 1, 0, 0, time.UTC)

var users = map[string]struct{}{}

func init() {
	var pieces []UAPiece
	for i := 0; i < 20; i++ {
		pieces = append(pieces, AndroidUA{geo{}})
	}
	for i := 0; i < 3; i++ {
		pieces = append(pieces, IPhoneUA{geo{}})
	}
	for i := 0; i < 5; i++ {
		pieces = append(pieces, WindowsUA{})
	}

	// Generating userlist
	var userlist []string
	for i := 0; i < 100000; i++ {
		uid := rand.Uint64()
		barrier := uint64(10000000000)
		if uid < barrier {
			uid += barrier
		}
		userlist = append(userlist, fmt.Sprintf("%d", uid))
		users[fmt.Sprintf("%d", uid)] = struct{}{}
	}

	// Generating user agents for userlist
	type Node struct {
		UA  string
		Geo func() (float64, float64)
	}
	var agentmap = map[string]Node{}
	for _, u := range userlist {
		index := rand.Int() % len(pieces)
		node := Node{}
		node.UA = FullUA{}.UserAgent(pieces[index])
		node.Geo = pieces[index].Geo
		agentmap[u] = node
	}

	start := START
	total := 0
	for i := 0; i < 1000000; i++ {
		index := rand.Int() % len(userlist)
		gens := agentmap[userlist[index]]
		pid := rand.Int() % 65536
		line := fmt.Sprintf("[%d %s] PRESENCE uid=%s ua='%s'", pid, start.Format("2006-01-02T15:04:05"), userlist[index], gens.UA)
		lat, lon := gens.Geo()
		if lat != 0 || lon != 0 {
			line += fmt.Sprintf(" Geo={Lat: %f, Lon: %f}", lat, lon)
		}
		act := rand.Int() % 6
		line += fmt.Sprintf(" Activity=%d", act)
		easylines = append(easylines, []byte(line))
		total += len(line)
		start = start.Add(time.Second)
	}
	buf := make([]byte, total)
	offset := 0
	for i, line := range easylines {
		copy(buf[offset:], line)
		easylines[i] = buf[offset : offset+len(line)]
		offset += len(line)
	}
}

func TestRagelExtraction(t *testing.T) {
	p := &Easy{}
	f := &EasyFloat{}
	e := &PresenceFloats{}
	r := &Presence{}

	start := START
	for _, line := range easylines {
		ok, err := p.Extract(line)
		if err != nil {
			t.Fatal(err)
		}
		require.True(t, ok)

		ok, err = f.Extract(line)
		if err != nil {
			t.Fatal(err)
		}
		require.True(t, ok)

		ok, err = e.Extract(line)
		if err != nil {
			t.Fatal(err)
		}
		require.True(t, ok)

		ok, err = r.Extract(line)
		if err != nil {
			t.Fatal(err)
		}
		require.True(t, ok)

		require.Equal(t, start.Format("2006-01-02T15:04:05"), string(p.Time))
		if _, ok := users[string(p.UID)]; !ok {
			t.Fatalf("Unknown user `\033[01m%s\033[0m` extracted on parsing>>\033[1m%s\033[0m", string(p.UID), string(line))
		}
		start = start.Add(time.Second)

		require.Equal(t, string(p.Time), string(r.Time))
		require.Equal(t, string(p.UID), string(r.UID))
		require.Equal(t, string(p.UA), string(r.UA))
		require.Equal(t, p.Geo.Valid, r.Geo.Valid)
		if p.Geo.Valid {
			require.Equal(t, string(p.Geo.Lat), string(r.Geo.Lat))
			require.Equal(t, string(p.Geo.Lon), string(r.Geo.Lon))
		}
		require.Equal(t, string(p.Activity), string(r.Activity))

		require.Equal(t, string(p.Time), string(e.Time))
		require.Equal(t, string(p.UID), string(e.UID))
		require.Equal(t, string(p.UA), string(e.UA))
		require.Equal(t, p.Geo.Valid, e.Geo.Valid)
		if p.Geo.Valid {
			require.Equal(t, string(p.Geo.Lat), fmt.Sprintf("%f", e.Geo.Lat))
			require.Equal(t, string(p.Geo.Lon), fmt.Sprintf("%f", e.Geo.Lon))
		}
		require.Equal(t, string(p.Activity), fmt.Sprintf("%d", e.Activity))

		require.Equal(t, string(p.Time), string(e.Time))
		require.Equal(t, string(p.UID), string(e.UID))
		require.Equal(t, string(p.UA), string(e.UA))
		require.Equal(t, p.Geo.Valid, e.Geo.Valid)
		if p.Geo.Valid {
			require.Equal(t, f.Geo.Lat, e.Geo.Lat)
			require.Equal(t, f.Geo.Lon, e.Geo.Lon)
		}
		require.Equal(t, f.Activity, e.Activity)
	}
}

func BenchmarkLDEEasyRealWorld(b *testing.B) {
	p := &Presence{}
	for i := 0; i < b.N; i++ {
		for _, line := range easylines {
			p.Extract(line)
			if _, ok := users[string(p.UID)]; !ok {
				b.Fatalf("Unknown user `\033[1m%s\033[0m` extracted on parsing>>\033[1m%s\033[0m", string(p.UID), string(line))
			}
		}
	}
}

func BenchmarkLDEEasyFloatsRealWorld(b *testing.B) {
	p := &PresenceFloats{}
	for i := 0; i < b.N; i++ {
		for _, line := range easylines {
			p.Extract(line)
			if _, ok := users[string(p.UID)]; !ok {
				b.Fatalf("Unknown user `\033[1m%s\033[0m` extracted on parsing>>\033[1m%s\033[0m", string(p.UID), string(line))
			}
		}
	}
}

func BenchmarkRagelEasyRealWorld(b *testing.B) {
	p := &Easy{}
	for i := 0; i < b.N; i++ {
		for _, line := range easylines {
			p.Extract(line)
			if _, ok := users[string(p.UID)]; !ok {
				b.Fatalf("Unknown user `\033[1m%s\033[0m` extracted on parsing>>\033[1m%s\033[0m", string(p.UID), string(line))
			}
		}
	}
}

func BenchmarkRagelEasyFloatsRealWorld(b *testing.B) {
	p := &EasyFloat{}
	for i := 0; i < b.N; i++ {
		for _, line := range easylines {
			p.Extract(line)
			if _, ok := users[string(p.UID)]; !ok {
				b.Fatalf("Unknown user `\033[1m%s\033[0m` extracted on parsing>>\033[1m%s\033[0m", string(p.UID), string(line))
			}
		}
	}
}
