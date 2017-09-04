package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"regexp"
	"runtime"
	"testing"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Extract struct {
	Name  []byte
	Count []byte
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func name() string {
	length := rand.Int()%1024 + 1
	return RandStringRunes(length)
}

var samples [][]byte
var extracts []Extract

func sampleData(length int) {
	buf := &bytes.Buffer{}
	totalLength := 0
	for i := 0; i < length; i++ {
		buf.Reset()
		n := name()
		buf.WriteString(n)
		buf.WriteByte('|')
		io.WriteString(buf, fmt.Sprintf("%d", rand.Uint64()))
		buf.WriteByte('|')
		io.WriteString(buf, fmt.Sprintf("%d", rand.Uint64()))
		buf.WriteByte('|')
		io.WriteString(buf, fmt.Sprintf("%d", rand.Uint64()))
		buf.WriteByte('|')
		c := fmt.Sprintf("%d", rand.Uint64())
		io.WriteString(buf, c)
		buf.WriteByte('|')
		io.WriteString(buf, fmt.Sprintf("%d", rand.Uint64()))
		buf.WriteByte('|')
		samples = append(samples, []byte(buf.String()))
		totalLength += buf.Len()
		extracts = append(extracts, Extract{
			Name:  []byte(n),
			Count: []byte(c),
		})
	}
	field := make([]byte, totalLength)
	left := 0
	for i, value := range samples {
		copy(field[left:left+len(value)], value)
		samples[i] = field[left : left+len(value)]
		left += len(value)
	}
	runtime.GC()
}

func init() {
	sampleData(1000)
}

func BenchmarkLDE(b *testing.B) {
	l := &Line{}
	var ok bool
	var err error
	for n := 0; n < b.N; n++ {
		for i, line := range samples {
			ok, err = l.Extract(line)
			if !ok && err == nil {
				panic("All lines must be parseable")
			}
			if err != nil {
				b.Fatal(err)
			}
			if !bytes.Equal(l.Name, extracts[i].Name) {
				b.Fatalf(
					"Failed to extract \033[1m%s\033[0m: got \033[31m%s\033[0m instead", string(extracts[i].Name), string(l.Name),
				)
			}
			if !bytes.Equal(l.Count, extracts[i].Count) {
				b.Fatalf(
					"Failed to extract \033[1m%s\033[0m: got \033[31m%s\033[0m instead",
					string(extracts[i].Count),
					string(l.Count),
				)
			}
		}
	}
}

func BenchmarkRagel(b *testing.B) {
	r := &Ragel{}
	var ok bool
	var err error
	for n := 0; n < b.N; n++ {
		for i, line := range samples {
			ok, err = r.Extract(line)
			if !ok && err == nil {
				panic("All lines must be parseable")
			}
			if err != nil {
				b.Fatal(err)
			}
			if !bytes.Equal(r.Name, extracts[i].Name) {
				b.Fatalf(
					"Failed to extract \033[1m%s\033[0m: got \033[31m%s\033[0m instead", string(extracts[i].Name), string(r.Name),
				)
			}
			if !bytes.Equal(r.Count, extracts[i].Count) {
				b.Fatalf(
					"Failed to extract \033[1m%s\033[0m: got \033[31m%s\033[0m instead",
					string(extracts[i].Count),
					string(r.Count),
				)
			}
		}
	}
}

var r = regexp.MustCompile(`^(.*?)\|.*?\|.*?\|.*?\|(.*?)\|.*$`)

func BenchmarkRegex(b *testing.B) {
	var name []byte
	var count []byte
	for n := 0; n < b.N; n++ {
		for i, line := range samples {
			data := r.FindSubmatch(line)
			if len(data) != 3 {
				b.Fatalf("Failed to parse >>%s", string(line))
			}
			name = data[1]
			count = data[2]
			if !bytes.Equal(name, extracts[i].Name) {
				b.Fatalf(
					"Failed to extract \033[1m%s\033[0m: got \033[31m%s\033[0m instead", string(extracts[i].Name), string(name),
				)
			}
			if !bytes.Equal(count, extracts[i].Count) {
				b.Fatalf(
					"Failed to extract \033[1m%s\033[0m: got \033[31m%s\033[0m instead",
					string(extracts[i].Count),
					string(count),
				)
			}
		}
	}
}
