package main

type OutputType uint8

const (
	OUTPUT_PLAIN OutputType = iota
	OUTPUT_PGSQL
)

type HaruhiOutput struct {
	outputType OutputType
	data interface{}
}