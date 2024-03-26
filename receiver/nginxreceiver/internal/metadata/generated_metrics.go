// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

// AttributeResponses specifies the a value responses attribute.
type AttributeResponses int

const (
	_ AttributeResponses = iota
	AttributeResponsesFailed
	AttributeResponsesBuffered
)

// String returns the string representation of the AttributeResponses.
func (av AttributeResponses) String() string {
	switch av {
	case AttributeResponsesFailed:
		return "failed"
	case AttributeResponsesBuffered:
		return "buffered"
	}
	return ""
}

// MapAttributeResponses is a helper map of string to AttributeResponses attribute value.
var MapAttributeResponses = map[string]AttributeResponses{
	"failed":   AttributeResponsesFailed,
	"buffered": AttributeResponsesBuffered,
}

// AttributeState specifies the a value state attribute.
type AttributeState int

const (
	_ AttributeState = iota
	AttributeStateActive
	AttributeStateReading
	AttributeStateWriting
	AttributeStateWaiting
)

// String returns the string representation of the AttributeState.
func (av AttributeState) String() string {
	switch av {
	case AttributeStateActive:
		return "active"
	case AttributeStateReading:
		return "reading"
	case AttributeStateWriting:
		return "writing"
	case AttributeStateWaiting:
		return "waiting"
	}
	return ""
}

// MapAttributeState is a helper map of string to AttributeState attribute value.
var MapAttributeState = map[string]AttributeState{
	"active":  AttributeStateActive,
	"reading": AttributeStateReading,
	"writing": AttributeStateWriting,
	"waiting": AttributeStateWaiting,
}

// AttributeStatusCode specifies the a value status_code attribute.
type AttributeStatusCode int

const (
	_ AttributeStatusCode = iota
	AttributeStatusCode1xx
	AttributeStatusCode2xx
	AttributeStatusCode4xx
	AttributeStatusCode5xx
)

// String returns the string representation of the AttributeStatusCode.
func (av AttributeStatusCode) String() string {
	switch av {
	case AttributeStatusCode1xx:
		return "1xx"
	case AttributeStatusCode2xx:
		return "2xx"
	case AttributeStatusCode4xx:
		return "4xx"
	case AttributeStatusCode5xx:
		return "5xx"
	}
	return ""
}

// MapAttributeStatusCode is a helper map of string to AttributeStatusCode attribute value.
var MapAttributeStatusCode = map[string]AttributeStatusCode{
	"1xx": AttributeStatusCode1xx,
	"2xx": AttributeStatusCode2xx,
	"4xx": AttributeStatusCode4xx,
	"5xx": AttributeStatusCode5xx,
}

type metricNginxConnectionsAccepted struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nginx.connections_accepted metric with initial data.
func (m *metricNginxConnectionsAccepted) init() {
	m.data.SetName("nginx.connections_accepted")
	m.data.SetDescription("The total number of accepted client connections")
	m.data.SetUnit("connections")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricNginxConnectionsAccepted) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNginxConnectionsAccepted) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNginxConnectionsAccepted) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNginxConnectionsAccepted(cfg MetricConfig) metricNginxConnectionsAccepted {
	m := metricNginxConnectionsAccepted{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNginxConnectionsCurrent struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nginx.connections_current metric with initial data.
func (m *metricNginxConnectionsCurrent) init() {
	m.data.SetName("nginx.connections_current")
	m.data.SetDescription("The current number of nginx connections by state")
	m.data.SetUnit("connections")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNginxConnectionsCurrent) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, stateAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("state", stateAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNginxConnectionsCurrent) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNginxConnectionsCurrent) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNginxConnectionsCurrent(cfg MetricConfig) metricNginxConnectionsCurrent {
	m := metricNginxConnectionsCurrent{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNginxConnectionsHandled struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nginx.connections_handled metric with initial data.
func (m *metricNginxConnectionsHandled) init() {
	m.data.SetName("nginx.connections_handled")
	m.data.SetDescription("The total number of handled connections. Generally, the parameter value is the same as nginx.connections_accepted unless some resource limits have been reached (for example, the worker_connections limit).")
	m.data.SetUnit("connections")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricNginxConnectionsHandled) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNginxConnectionsHandled) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNginxConnectionsHandled) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNginxConnectionsHandled(cfg MetricConfig) metricNginxConnectionsHandled {
	m := metricNginxConnectionsHandled{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNginxHTTPStatus struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nginx.http.status metric with initial data.
func (m *metricNginxHTTPStatus) init() {
	m.data.SetName("nginx.http.status")
	m.data.SetDescription("The count of HTTP response statuses")
	m.data.SetUnit("statuses")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNginxHTTPStatus) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, statusCodeAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("status_code", statusCodeAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNginxHTTPStatus) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNginxHTTPStatus) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNginxHTTPStatus(cfg MetricConfig) metricNginxHTTPStatus {
	m := metricNginxHTTPStatus{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNginxRequests struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nginx.requests metric with initial data.
func (m *metricNginxRequests) init() {
	m.data.SetName("nginx.requests")
	m.data.SetDescription("Total number of requests made to the server since it started")
	m.data.SetUnit("requests")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricNginxRequests) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNginxRequests) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNginxRequests) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNginxRequests(cfg MetricConfig) metricNginxRequests {
	m := metricNginxRequests{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNginxUpstreamsResponse struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nginx.upstreams_response metric with initial data.
func (m *metricNginxUpstreamsResponse) init() {
	m.data.SetName("nginx.upstreams_response")
	m.data.SetDescription("The total number of HTTP requests from upstream responses")
	m.data.SetUnit("upstreams")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNginxUpstreamsResponse) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, responsesAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("responses", responsesAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNginxUpstreamsResponse) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNginxUpstreamsResponse) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNginxUpstreamsResponse(cfg MetricConfig) metricNginxUpstreamsResponse {
	m := metricNginxUpstreamsResponse{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                         MetricsBuilderConfig // config of the metrics builder.
	startTime                      pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity                int                  // maximum observed number of metrics per resource.
	metricsBuffer                  pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                      component.BuildInfo  // contains version information.
	metricNginxConnectionsAccepted metricNginxConnectionsAccepted
	metricNginxConnectionsCurrent  metricNginxConnectionsCurrent
	metricNginxConnectionsHandled  metricNginxConnectionsHandled
	metricNginxHTTPStatus          metricNginxHTTPStatus
	metricNginxRequests            metricNginxRequests
	metricNginxUpstreamsResponse   metricNginxUpstreamsResponse
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:                         mbc,
		startTime:                      pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                  pmetric.NewMetrics(),
		buildInfo:                      settings.BuildInfo,
		metricNginxConnectionsAccepted: newMetricNginxConnectionsAccepted(mbc.Metrics.NginxConnectionsAccepted),
		metricNginxConnectionsCurrent:  newMetricNginxConnectionsCurrent(mbc.Metrics.NginxConnectionsCurrent),
		metricNginxConnectionsHandled:  newMetricNginxConnectionsHandled(mbc.Metrics.NginxConnectionsHandled),
		metricNginxHTTPStatus:          newMetricNginxHTTPStatus(mbc.Metrics.NginxHTTPStatus),
		metricNginxRequests:            newMetricNginxRequests(mbc.Metrics.NginxRequests),
		metricNginxUpstreamsResponse:   newMetricNginxUpstreamsResponse(mbc.Metrics.NginxUpstreamsResponse),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/nginxreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricNginxConnectionsAccepted.emit(ils.Metrics())
	mb.metricNginxConnectionsCurrent.emit(ils.Metrics())
	mb.metricNginxConnectionsHandled.emit(ils.Metrics())
	mb.metricNginxHTTPStatus.emit(ils.Metrics())
	mb.metricNginxRequests.emit(ils.Metrics())
	mb.metricNginxUpstreamsResponse.emit(ils.Metrics())

	for _, op := range rmo {
		op(rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordNginxConnectionsAcceptedDataPoint adds a data point to nginx.connections_accepted metric.
func (mb *MetricsBuilder) RecordNginxConnectionsAcceptedDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricNginxConnectionsAccepted.recordDataPoint(mb.startTime, ts, val)
}

// RecordNginxConnectionsCurrentDataPoint adds a data point to nginx.connections_current metric.
func (mb *MetricsBuilder) RecordNginxConnectionsCurrentDataPoint(ts pcommon.Timestamp, val int64, stateAttributeValue AttributeState) {
	mb.metricNginxConnectionsCurrent.recordDataPoint(mb.startTime, ts, val, stateAttributeValue.String())
}

// RecordNginxConnectionsHandledDataPoint adds a data point to nginx.connections_handled metric.
func (mb *MetricsBuilder) RecordNginxConnectionsHandledDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricNginxConnectionsHandled.recordDataPoint(mb.startTime, ts, val)
}

// RecordNginxHTTPStatusDataPoint adds a data point to nginx.http.status metric.
func (mb *MetricsBuilder) RecordNginxHTTPStatusDataPoint(ts pcommon.Timestamp, val int64, statusCodeAttributeValue AttributeStatusCode) {
	mb.metricNginxHTTPStatus.recordDataPoint(mb.startTime, ts, val, statusCodeAttributeValue.String())
}

// RecordNginxRequestsDataPoint adds a data point to nginx.requests metric.
func (mb *MetricsBuilder) RecordNginxRequestsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricNginxRequests.recordDataPoint(mb.startTime, ts, val)
}

// RecordNginxUpstreamsResponseDataPoint adds a data point to nginx.upstreams_response metric.
func (mb *MetricsBuilder) RecordNginxUpstreamsResponseDataPoint(ts pcommon.Timestamp, val int64, responsesAttributeValue AttributeResponses) {
	mb.metricNginxUpstreamsResponse.recordDataPoint(mb.startTime, ts, val, responsesAttributeValue.String())
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
