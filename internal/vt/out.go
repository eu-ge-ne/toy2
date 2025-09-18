package vt

type out interface {
	Write(chunks ...[]byte)
}
