// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"fmt"
	"strings"
	"time"

	"github.com/Ch1f/otel/api/label"
	"github.com/Ch1f/otel/api/metric"
	export "github.com/Ch1f/otel/sdk/export/metric"
	"github.com/Ch1f/otel/sdk/export/metric/aggregation"
	"github.com/Ch1f/otel/sdk/metric/aggregator/array"
	"github.com/Ch1f/otel/sdk/metric/aggregator/ddsketch"
	"github.com/Ch1f/otel/sdk/metric/aggregator/histogram"
	"github.com/Ch1f/otel/sdk/metric/aggregator/lastvalue"
	"github.com/Ch1f/otel/sdk/metric/aggregator/minmaxsumcount"
	"github.com/Ch1f/otel/sdk/metric/aggregator/sum"
)

type (
	// Output collects distinct metric/label set outputs.
	Output struct {
		Map          map[string]float64
		labelEncoder label.Encoder
	}

	// testAggregatorSelector returns aggregators consistent with
	// the test variables below, needed for testing stateful
	// processors, which clone Aggregators using AggregatorFor(desc).
	testAggregatorSelector struct{}
)

func NewOutput(labelEncoder label.Encoder) Output {
	return Output{
		Map:          make(map[string]float64),
		labelEncoder: labelEncoder,
	}
}

// AggregatorSelector returns a policy that is consistent with the
// test descriptors above.  I.e., it returns sum.New() for counter
// instruments and lastvalue.New() for lastValue instruments.
func AggregatorSelector() export.AggregatorSelector {
	return testAggregatorSelector{}
}

func (testAggregatorSelector) AggregatorFor(desc *metric.Descriptor, aggPtrs ...*export.Aggregator) {

	switch {
	case strings.HasSuffix(desc.Name(), ".disabled"):
		for i := range aggPtrs {
			*aggPtrs[i] = nil
		}
	case strings.HasSuffix(desc.Name(), ".sum"):
		aggs := sum.New(len(aggPtrs))
		for i := range aggPtrs {
			*aggPtrs[i] = &aggs[i]
		}
	case strings.HasSuffix(desc.Name(), ".minmaxsumcount"):
		aggs := minmaxsumcount.New(len(aggPtrs), desc)
		for i := range aggPtrs {
			*aggPtrs[i] = &aggs[i]
		}
	case strings.HasSuffix(desc.Name(), ".lastvalue"):
		aggs := lastvalue.New(len(aggPtrs))
		for i := range aggPtrs {
			*aggPtrs[i] = &aggs[i]
		}
	case strings.HasSuffix(desc.Name(), ".sketch"):
		aggs := ddsketch.New(len(aggPtrs), desc, ddsketch.NewDefaultConfig())
		for i := range aggPtrs {
			*aggPtrs[i] = &aggs[i]
		}
	case strings.HasSuffix(desc.Name(), ".histogram"):
		aggs := histogram.New(len(aggPtrs), desc, nil)
		for i := range aggPtrs {
			*aggPtrs[i] = &aggs[i]
		}
	case strings.HasSuffix(desc.Name(), ".exact"):
		aggs := array.New(len(aggPtrs))
		for i := range aggPtrs {
			*aggPtrs[i] = &aggs[i]
		}
	default:
		panic(fmt.Sprint("Invalid instrument name for test AggregatorSelector: ", desc.Name()))
	}
}

// AddRecord adds a string representation of the exported metric data
// to a map for use in testing.  The value taken from the record is
// either the Sum() or the LastValue() of its Aggregation(), whichever
// is defined.  Record timestamps are ignored.
func (o Output) AddRecord(rec export.Record) error {
	encoded := rec.Labels().Encoded(o.labelEncoder)
	rencoded := rec.Resource().Encoded(o.labelEncoder)
	key := fmt.Sprint(rec.Descriptor().Name(), "/", encoded, "/", rencoded)
	var value float64

	if s, ok := rec.Aggregation().(aggregation.Sum); ok {
		sum, _ := s.Sum()
		value = sum.CoerceToFloat64(rec.Descriptor().NumberKind())
	} else if l, ok := rec.Aggregation().(aggregation.LastValue); ok {
		last, _, _ := l.LastValue()
		value = last.CoerceToFloat64(rec.Descriptor().NumberKind())
	} else {
		panic(fmt.Sprintf("Unhandled aggregator type: %T", rec.Aggregation()))
	}
	o.Map[key] = value
	return nil
}

// AddAccumulation adds a string representation of the exported metric
// data to a map for use in testing.  The value taken from the
// accumulation is either the Sum() or the LastValue() of its
// Aggregator().Aggregation(), whichever is defined.
func (o Output) AddAccumulation(acc export.Accumulation) error {
	return o.AddRecord(
		export.NewRecord(
			acc.Descriptor(),
			acc.Labels(),
			acc.Resource(),
			acc.Aggregator().Aggregation(),
			time.Time{},
			time.Time{},
		),
	)
}
