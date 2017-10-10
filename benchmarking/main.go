package main

//go:generate ldetool generate --little-endian --package main rule.lde
//go:generate ragel -Z -G2 template.ragel
//go:generate ragel -Z -G2 easy.ragel
//go:generate ragel -Z -G2 easy_floats.ragel

func main() {
}
