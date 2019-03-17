package main

//go:generate antlr4 -visitor -no-visitor -listener -o internal/parser -Dlanguage=Go LDE.g4
