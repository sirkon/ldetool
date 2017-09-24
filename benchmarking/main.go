package main

//go:generate ldetool generate --little-endian --package main rule.lde
//go:generate ragel -Z -G2 template.ragel

func main() {
}
