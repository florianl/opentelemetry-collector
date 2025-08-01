// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pprofile

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpprofiles "go.opentelemetry.io/collector/pdata/internal/data/protogen/profiles/v1development"
	"go.opentelemetry.io/collector/pdata/internal/json"
)

func TestLocationSlice(t *testing.T) {
	es := NewLocationSlice()
	assert.Equal(t, 0, es.Len())
	state := internal.StateMutable
	es = newLocationSlice(&[]*otlpprofiles.Location{}, &state)
	assert.Equal(t, 0, es.Len())

	emptyVal := NewLocation()
	testVal := generateTestLocation()
	for i := 0; i < 7; i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, es.At(i))
		fillTestLocation(el)
		assert.Equal(t, testVal, es.At(i))
	}
	assert.Equal(t, 7, es.Len())
}

func TestLocationSliceReadOnly(t *testing.T) {
	sharedState := internal.StateReadOnly
	es := newLocationSlice(&[]*otlpprofiles.Location{}, &sharedState)
	assert.Equal(t, 0, es.Len())
	assert.Panics(t, func() { es.AppendEmpty() })
	assert.Panics(t, func() { es.EnsureCapacity(2) })
	es2 := NewLocationSlice()
	es.CopyTo(es2)
	assert.Panics(t, func() { es2.CopyTo(es) })
	assert.Panics(t, func() { es.MoveAndAppendTo(es2) })
	assert.Panics(t, func() { es2.MoveAndAppendTo(es) })
}

func TestLocationSlice_CopyTo(t *testing.T) {
	dest := NewLocationSlice()
	// Test CopyTo empty
	NewLocationSlice().CopyTo(dest)
	assert.Equal(t, NewLocationSlice(), dest)

	// Test CopyTo larger slice
	src := generateTestLocationSlice()
	src.CopyTo(dest)
	assert.Equal(t, generateTestLocationSlice(), dest)

	// Test CopyTo same size slice
	src.CopyTo(dest)
	assert.Equal(t, generateTestLocationSlice(), dest)

	// Test CopyTo smaller size slice
	NewLocationSlice().CopyTo(dest)
	assert.Equal(t, 0, dest.Len())

	// Test CopyTo larger slice with enough capacity
	src.CopyTo(dest)
	assert.Equal(t, generateTestLocationSlice(), dest)
}

func TestLocationSlice_EnsureCapacity(t *testing.T) {
	es := generateTestLocationSlice()

	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	assert.Equal(t, es.Len(), cap(*es.orig))
	assert.Equal(t, generateTestLocationSlice(), es)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	es.EnsureCapacity(ensureLargeLen)
	assert.Less(t, generateTestLocationSlice().Len(), ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	assert.Equal(t, generateTestLocationSlice(), es)
}

func TestLocationSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestLocationSlice()
	dest := NewLocationSlice()
	src := generateTestLocationSlice()
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestLocationSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestLocationSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestLocationSlice().MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}

	dest.MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestLocationSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewLocationSlice()
	emptySlice.RemoveIf(func(el Location) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestLocationSlice()
	pos := 0
	filtered.RemoveIf(func(el Location) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestLocationSlice_RemoveIfAll(t *testing.T) {
	got := generateTestLocationSlice()
	got.RemoveIf(func(el Location) bool {
		return true
	})
	assert.Equal(t, 0, got.Len())
}

func TestLocationSliceAll(t *testing.T) {
	ms := generateTestLocationSlice()
	assert.NotEmpty(t, ms.Len())

	var c int
	for i, v := range ms.All() {
		assert.Equal(t, ms.At(i), v, "element should match")
		c++
	}
	assert.Equal(t, ms.Len(), c, "All elements should have been visited")
}

func TestLocationSlice_MarshalAndUnmarshalJSON(t *testing.T) {
	stream := json.BorrowStream(nil)
	defer json.ReturnStream(stream)
	src := generateTestLocationSlice()
	src.marshalJSONStream(stream)
	require.NoError(t, stream.Error())

	iter := json.BorrowIterator(stream.Buffer())
	defer json.ReturnIterator(iter)
	dest := NewLocationSlice()
	dest.unmarshalJSONIter(iter)
	require.NoError(t, iter.Error())

	assert.Equal(t, src, dest)
}

func TestLocationSlice_Sort(t *testing.T) {
	es := generateTestLocationSlice()
	es.Sort(func(a, b Location) bool {
		return uintptr(unsafe.Pointer(a.orig)) < uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.Less(t, uintptr(unsafe.Pointer(es.At(i-1).orig)), uintptr(unsafe.Pointer(es.At(i).orig)))
	}
	es.Sort(func(a, b Location) bool {
		return uintptr(unsafe.Pointer(a.orig)) > uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.Greater(t, uintptr(unsafe.Pointer(es.At(i-1).orig)), uintptr(unsafe.Pointer(es.At(i).orig)))
	}
}

func generateTestLocationSlice() LocationSlice {
	es := NewLocationSlice()
	fillTestLocationSlice(es)
	return es
}

func fillTestLocationSlice(es LocationSlice) {
	*es.orig = make([]*otlpprofiles.Location, 7)
	for i := 0; i < 7; i++ {
		(*es.orig)[i] = &otlpprofiles.Location{}
		fillTestLocation(newLocation((*es.orig)[i], es.state))
	}
}
