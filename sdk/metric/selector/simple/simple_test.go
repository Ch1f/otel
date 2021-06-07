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

package simple_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Ch1f/otel/api/metric"
	export "github.com/Ch1f/otel/sdk/export/metric"
	"github.com/Ch1f/otel/sdk/metric/aggregator/array"
	"github.com/Ch1f/otel/sdk/metric/aggregator/ddsketch"
	"github.com/Ch1f/otel/sdk/metric/aggregator/histogram"
	"github.com/Ch1f/otel/sdk/metric/aggregator/minmaxsumcount"
	"github.com/Ch1f/otel/sdk/metric/aggregator/sum"
	"github.com/Ch1f/otel/sdk/metric/selector/simple"
)

var (
	testCounterDesc       = metric.NewDescriptor("counter", metric.CounterKind, metric.Int64NumberKind)
	testValueRecorderDesc = metric.NewDescriptor("valuerecorder", metric.ValueRecorderKind, metric.Int64NumberKind)
	testValueObserverDesc = metric.NewDescriptor("valueobserver", metric.ValueObserverKind, metric.Int64NumberKind)
)

func oneAgg(sel export.AggregatorSelector, desc *metric.Descriptor) export.Aggregator {
	var agg export.Aggregator
	sel.AggregatorFor(desc, &agg)
	return agg
}

func TestInexpensiveDistribution(t *testing.T) {
	inex := simple.NewWithInexpensiveDistribution()
	require.NotPanics(t, func() { _ = oneAgg(inex, &testCounterDesc).(*sum.Aggregator) })
	require.NotPanics(t, func() { _ = oneAgg(inex, &testValueRecorderDesc).(*minmaxsumcount.Aggregator) })
	require.NotPanics(t, func() { _ = oneAgg(inex, &testValueObserverDesc).(*minmaxsumcount.Aggregator) })
}

func TestSketchDistribution(t *testing.T) {
	sk := simple.NewWithSketchDistribution(ddsketch.NewDefaultConfig())
	require.NotPanics(t, func() { _ = oneAgg(sk, &testCounterDesc).(*sum.Aggregator) })
	require.NotPanics(t, func() { _ = oneAgg(sk, &testValueRecorderDesc).(*ddsketch.Aggregator) })
	require.NotPanics(t, func() { _ = oneAgg(sk, &testValueObserverDesc).(*ddsketch.Aggregator) })
}

func TestExactDistribution(t *testing.T) {
	ex := simple.NewWithExactDistribution()
	require.NotPanics(t, func() { _ = oneAgg(ex, &testCounterDesc).(*sum.Aggregator) })
	require.NotPanics(t, func() { _ = oneAgg(ex, &testValueRecorderDesc).(*array.Aggregator) })
	require.NotPanics(t, func() { _ = oneAgg(ex, &testValueObserverDesc).(*array.Aggregator) })
}

func TestHistogramDistribution(t *testing.T) {
	ex := simple.NewWithHistogramDistribution(nil)
	require.NotPanics(t, func() { _ = oneAgg(ex, &testCounterDesc).(*sum.Aggregator) })
	require.NotPanics(t, func() { _ = oneAgg(ex, &testValueRecorderDesc).(*histogram.Aggregator) })
	require.NotPanics(t, func() { _ = oneAgg(ex, &testValueObserverDesc).(*histogram.Aggregator) })
}
