// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
	"go.opentelemetry.io/collector/pdata/internal/json"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ExponentialHistogramDataPointBuckets are a set of bucket counts, encoded in a contiguous array of counts.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewExponentialHistogramDataPointBuckets function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ExponentialHistogramDataPointBuckets struct {
	orig  *otlpmetrics.ExponentialHistogramDataPoint_Buckets
	state *internal.State
}

func newExponentialHistogramDataPointBuckets(orig *otlpmetrics.ExponentialHistogramDataPoint_Buckets, state *internal.State) ExponentialHistogramDataPointBuckets {
	return ExponentialHistogramDataPointBuckets{orig: orig, state: state}
}

// NewExponentialHistogramDataPointBuckets creates a new empty ExponentialHistogramDataPointBuckets.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewExponentialHistogramDataPointBuckets() ExponentialHistogramDataPointBuckets {
	state := internal.StateMutable
	return newExponentialHistogramDataPointBuckets(&otlpmetrics.ExponentialHistogramDataPoint_Buckets{}, &state)
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms ExponentialHistogramDataPointBuckets) MoveTo(dest ExponentialHistogramDataPointBuckets) {
	ms.state.AssertMutable()
	dest.state.AssertMutable()
	// If they point to the same data, they are the same, nothing to do.
	if ms.orig == dest.orig {
		return
	}
	*dest.orig = *ms.orig
	*ms.orig = otlpmetrics.ExponentialHistogramDataPoint_Buckets{}
}

// Offset returns the offset associated with this ExponentialHistogramDataPointBuckets.
func (ms ExponentialHistogramDataPointBuckets) Offset() int32 {
	return ms.orig.Offset
}

// SetOffset replaces the offset associated with this ExponentialHistogramDataPointBuckets.
func (ms ExponentialHistogramDataPointBuckets) SetOffset(v int32) {
	ms.state.AssertMutable()
	ms.orig.Offset = v
}

// BucketCounts returns the BucketCounts associated with this ExponentialHistogramDataPointBuckets.
func (ms ExponentialHistogramDataPointBuckets) BucketCounts() pcommon.UInt64Slice {
	return pcommon.UInt64Slice(internal.NewUInt64Slice(&ms.orig.BucketCounts, ms.state))
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms ExponentialHistogramDataPointBuckets) CopyTo(dest ExponentialHistogramDataPointBuckets) {
	dest.state.AssertMutable()
	copyOrigExponentialHistogramDataPointBuckets(dest.orig, ms.orig)
}

// marshalJSONStream marshals all properties from the current struct to the destination stream.
func (ms ExponentialHistogramDataPointBuckets) marshalJSONStream(dest *json.Stream) {
	dest.WriteObjectStart()
	if ms.orig.Offset != int32(0) {
		dest.WriteObjectField("offset")
		dest.WriteInt32(ms.orig.Offset)
	}
	if len(ms.orig.BucketCounts) > 0 {
		dest.WriteObjectField("bucketCounts")
		internal.MarshalJSONStreamUInt64Slice(internal.NewUInt64Slice(&ms.orig.BucketCounts, ms.state), dest)
	}
	dest.WriteObjectEnd()
}

// unmarshalJSONIter unmarshals all properties from the current struct from the source iterator.
func (ms ExponentialHistogramDataPointBuckets) unmarshalJSONIter(iter *json.Iterator) {
	iter.ReadObjectCB(func(iter *json.Iterator, f string) bool {
		switch f {
		case "offset":
			ms.orig.Offset = iter.ReadInt32()
		case "bucketCounts", "bucket_counts":
			internal.UnmarshalJSONIterUInt64Slice(internal.NewUInt64Slice(&ms.orig.BucketCounts, ms.state), iter)
		default:
			iter.Skip()
		}
		return true
	})
}

func copyOrigExponentialHistogramDataPointBuckets(dest, src *otlpmetrics.ExponentialHistogramDataPoint_Buckets) {
	dest.Offset = src.Offset
	dest.BucketCounts = internal.CopyOrigUInt64Slice(dest.BucketCounts, src.BucketCounts)
}
