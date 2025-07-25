// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package internal

type UInt64Slice struct {
	orig  *[]uint64
	state *State
}

func GetOrigUInt64Slice(ms UInt64Slice) *[]uint64 {
	return ms.orig
}

func GetUInt64SliceState(ms UInt64Slice) *State {
	return ms.state
}

func NewUInt64Slice(orig *[]uint64, state *State) UInt64Slice {
	return UInt64Slice{orig: orig, state: state}
}

func CopyOrigUInt64Slice(dst, src []uint64) []uint64 {
	dst = dst[:0]
	return append(dst, src...)
}

func FillTestUInt64Slice(ms UInt64Slice) {
	*ms.orig = []uint64{1, 2, 3}
}

func GenerateTestUInt64Slice() UInt64Slice {
	orig := []uint64(nil)
	state := StateMutable
	ms := NewUInt64Slice(&orig, &state)
	FillTestUInt64Slice(ms)
	return ms
}
