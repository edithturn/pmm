// Code generated by go-swagger; DO NOT EDIT.

package object_details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetMetricsReader is a Reader for the GetMetrics structure.
type GetMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMetricsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMetricsOK creates a GetMetricsOK with default headers values
func NewGetMetricsOK() *GetMetricsOK {
	return &GetMetricsOK{}
}

/*GetMetricsOK handles this case with default header values.

A successful response.
*/
type GetMetricsOK struct {
	Payload *GetMetricsOKBody
}

func (o *GetMetricsOK) Error() string {
	return fmt.Sprintf("[POST /v0/qan/ObjectDetails/GetMetrics][%d] getMetricsOk  %+v", 200, o.Payload)
}

func (o *GetMetricsOK) GetPayload() *GetMetricsOKBody {
	return o.Payload
}

func (o *GetMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetMetricsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetricsDefault creates a GetMetricsDefault with default headers values
func NewGetMetricsDefault(code int) *GetMetricsDefault {
	return &GetMetricsDefault{
		_statusCode: code,
	}
}

/*GetMetricsDefault handles this case with default header values.

An unexpected error response.
*/
type GetMetricsDefault struct {
	_statusCode int

	Payload *GetMetricsDefaultBody
}

// Code gets the status code for the get metrics default response
func (o *GetMetricsDefault) Code() int {
	return o._statusCode
}

func (o *GetMetricsDefault) Error() string {
	return fmt.Sprintf("[POST /v0/qan/ObjectDetails/GetMetrics][%d] GetMetrics default  %+v", o._statusCode, o.Payload)
}

func (o *GetMetricsDefault) GetPayload() *GetMetricsDefaultBody {
	return o.Payload
}

func (o *GetMetricsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetMetricsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetMetricsBody MetricsRequest defines filtering of metrics for specific value of dimension (ex.: host=hostname1 or queryid=1D410B4BE5060972.
swagger:model GetMetricsBody
*/
type GetMetricsBody struct {

	// period start from
	// Format: date-time
	PeriodStartFrom strfmt.DateTime `json:"period_start_from,omitempty"`

	// period start to
	// Format: date-time
	PeriodStartTo strfmt.DateTime `json:"period_start_to,omitempty"`

	// dimension value: ex: queryid - 1D410B4BE5060972.
	FilterBy string `json:"filter_by,omitempty"`

	// one of dimension: queryid | host ...
	GroupBy string `json:"group_by,omitempty"`

	// labels
	Labels []*LabelsItems0 `json:"labels"`

	// include only fields
	IncludeOnlyFields []string `json:"include_only_fields"`

	// retrieve only values for totals, excluding N/A values
	Totals bool `json:"totals,omitempty"`
}

// Validate validates this get metrics body
func (o *GetMetricsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePeriodStartFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePeriodStartTo(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMetricsBody) validatePeriodStartFrom(formats strfmt.Registry) error {

	if swag.IsZero(o.PeriodStartFrom) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"period_start_from", "body", "date-time", o.PeriodStartFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetMetricsBody) validatePeriodStartTo(formats strfmt.Registry) error {

	if swag.IsZero(o.PeriodStartTo) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"period_start_to", "body", "date-time", o.PeriodStartTo.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetMetricsBody) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(o.Labels) { // not required
		return nil
	}

	for i := 0; i < len(o.Labels); i++ {
		if swag.IsZero(o.Labels[i]) { // not required
			continue
		}

		if o.Labels[i] != nil {
			if err := o.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetMetricsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMetricsBody) UnmarshalBinary(b []byte) error {
	var res GetMetricsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetMetricsDefaultBody get metrics default body
swagger:model GetMetricsDefaultBody
*/
type GetMetricsDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this get metrics default body
func (o *GetMetricsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMetricsDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetMetrics default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetMetricsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMetricsDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetMetricsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetMetricsOKBody MetricsReply defines metrics for specific value of dimension (ex.: host=hostname1 or queryid=1D410B4BE5060972.
swagger:model GetMetricsOKBody
*/
type GetMetricsOKBody struct {

	// metrics
	Metrics map[string]MetricsAnon `json:"metrics,omitempty"`

	// text metrics
	TextMetrics map[string]string `json:"text_metrics,omitempty"`

	// sparkline
	Sparkline []*SparklineItems0 `json:"sparkline"`

	// totals
	Totals map[string]TotalsAnon `json:"totals,omitempty"`

	// fingerprint
	Fingerprint string `json:"fingerprint,omitempty"`
}

// Validate validates this get metrics OK body
func (o *GetMetricsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSparkline(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMetricsOKBody) validateMetrics(formats strfmt.Registry) error {

	if swag.IsZero(o.Metrics) { // not required
		return nil
	}

	for k := range o.Metrics {

		if swag.IsZero(o.Metrics[k]) { // not required
			continue
		}
		if val, ok := o.Metrics[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (o *GetMetricsOKBody) validateSparkline(formats strfmt.Registry) error {

	if swag.IsZero(o.Sparkline) { // not required
		return nil
	}

	for i := 0; i < len(o.Sparkline); i++ {
		if swag.IsZero(o.Sparkline[i]) { // not required
			continue
		}

		if o.Sparkline[i] != nil {
			if err := o.Sparkline[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getMetricsOk" + "." + "sparkline" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetMetricsOKBody) validateTotals(formats strfmt.Registry) error {

	if swag.IsZero(o.Totals) { // not required
		return nil
	}

	for k := range o.Totals {

		if swag.IsZero(o.Totals[k]) { // not required
			continue
		}
		if val, ok := o.Totals[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetMetricsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMetricsOKBody) UnmarshalBinary(b []byte) error {
	var res GetMetricsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LabelsItems0 MapFieldEntry allows to pass labels/dimensions in form like {"server": ["db1", "db2"...]}.
swagger:model LabelsItems0
*/
type LabelsItems0 struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value []string `json:"value"`
}

// Validate validates this labels items0
func (o *LabelsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LabelsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LabelsItems0) UnmarshalBinary(b []byte) error {
	var res LabelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MetricsAnon MetricValues is statistics of specific metric.
swagger:model MetricsAnon
*/
type MetricsAnon struct {

	// rate
	Rate float32 `json:"rate,omitempty"`

	// cnt
	Cnt float32 `json:"cnt,omitempty"`

	// sum
	Sum float32 `json:"sum,omitempty"`

	// min
	Min float32 `json:"min,omitempty"`

	// max
	Max float32 `json:"max,omitempty"`

	// avg
	Avg float32 `json:"avg,omitempty"`

	// p99
	P99 float32 `json:"p99,omitempty"`

	// percent of total
	PercentOfTotal float32 `json:"percent_of_total,omitempty"`
}

// Validate validates this metrics anon
func (o *MetricsAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MetricsAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MetricsAnon) UnmarshalBinary(b []byte) error {
	var res MetricsAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SparklineItems0 Point contains values that represents abscissa (time) and ordinate (volume etc.)
// of every point in a coordinate system of Sparklines.
swagger:model SparklineItems0
*/
type SparklineItems0 struct {

	// The serial number of the chart point from the largest time in the time interval to the lowest time in the time range.
	Point int64 `json:"point,omitempty"`

	// Duration beetween two points.
	TimeFrame int64 `json:"time_frame,omitempty"`

	// Time of point in format RFC3339.
	Timestamp string `json:"timestamp,omitempty"`

	// load is query_time / time_range.
	Load float32 `json:"load,omitempty"`

	// number of queries in bucket.
	NumQueriesPerSec float32 `json:"num_queries_per_sec,omitempty"`

	// number of queries with errors.
	NumQueriesWithErrorsPerSec float32 `json:"num_queries_with_errors_per_sec,omitempty"`

	// number of queries with warnings.
	NumQueriesWithWarningsPerSec float32 `json:"num_queries_with_warnings_per_sec,omitempty"`

	// The statement execution time in seconds.
	MQueryTimeSumPerSec float32 `json:"m_query_time_sum_per_sec,omitempty"`

	// The time to acquire locks in seconds.
	MLockTimeSumPerSec float32 `json:"m_lock_time_sum_per_sec,omitempty"`

	// The number of rows sent to the client.
	MRowsSentSumPerSec float32 `json:"m_rows_sent_sum_per_sec,omitempty"`

	// Number of rows scanned - SELECT.
	MRowsExaminedSumPerSec float32 `json:"m_rows_examined_sum_per_sec,omitempty"`

	// Number of rows changed - UPDATE, DELETE, INSERT.
	MRowsAffectedSumPerSec float32 `json:"m_rows_affected_sum_per_sec,omitempty"`

	// The number of rows read from tables.
	MRowsReadSumPerSec float32 `json:"m_rows_read_sum_per_sec,omitempty"`

	// The number of merge passes that the sort algorithm has had to do.
	MMergePassesSumPerSec float32 `json:"m_merge_passes_sum_per_sec,omitempty"`

	// Counts the number of page read operations scheduled.
	MInnodbIorOpsSumPerSec float32 `json:"m_innodb_io_r_ops_sum_per_sec,omitempty"`

	// Similar to innodb_IO_r_ops, but the unit is bytes.
	MInnodbIorBytesSumPerSec float32 `json:"m_innodb_io_r_bytes_sum_per_sec,omitempty"`

	// Shows how long (in seconds) it took InnoDB to actually read the data from storage.
	MInnodbIorWaitSumPerSec float32 `json:"m_innodb_io_r_wait_sum_per_sec,omitempty"`

	// Shows how long (in seconds) the query waited for row locks.
	MInnodbRecLockWaitSumPerSec float32 `json:"m_innodb_rec_lock_wait_sum_per_sec,omitempty"`

	// Shows how long (in seconds) the query spent either waiting to enter the InnoDB queue or inside that queue waiting for execution.
	MInnodbQueueWaitSumPerSec float32 `json:"m_innodb_queue_wait_sum_per_sec,omitempty"`

	// Counts approximately the number of unique pages the query accessed.
	MInnodbPagesDistinctSumPerSec float32 `json:"m_innodb_pages_distinct_sum_per_sec,omitempty"`

	// Shows how long the query is.
	MQueryLengthSumPerSec float32 `json:"m_query_length_sum_per_sec,omitempty"`

	// The number of bytes sent to all clients.
	MBytesSentSumPerSec float32 `json:"m_bytes_sent_sum_per_sec,omitempty"`

	// Number of temporary tables created on memory for the query.
	MTmpTablesSumPerSec float32 `json:"m_tmp_tables_sum_per_sec,omitempty"`

	// Number of temporary tables created on disk for the query.
	MTmpDiskTablesSumPerSec float32 `json:"m_tmp_disk_tables_sum_per_sec,omitempty"`

	// Total Size in bytes for all temporary tables used in the query.
	MTmpTableSizesSumPerSec float32 `json:"m_tmp_table_sizes_sum_per_sec,omitempty"`

	// Boolean metrics:
	// - *_sum_per_sec - how many times this matric was true.
	//
	// Query Cache hits.
	MQcHitSumPerSec float32 `json:"m_qc_hit_sum_per_sec,omitempty"`

	// The query performed a full table scan.
	MFullScanSumPerSec float32 `json:"m_full_scan_sum_per_sec,omitempty"`

	// The query performed a full join (a join without indexes).
	MFullJoinSumPerSec float32 `json:"m_full_join_sum_per_sec,omitempty"`

	// The query created an implicit internal temporary table.
	MTmpTableSumPerSec float32 `json:"m_tmp_table_sum_per_sec,omitempty"`

	// The querys temporary table was stored on disk.
	MTmpTableOnDiskSumPerSec float32 `json:"m_tmp_table_on_disk_sum_per_sec,omitempty"`

	// The query used a filesort.
	MFilesortSumPerSec float32 `json:"m_filesort_sum_per_sec,omitempty"`

	// The filesort was performed on disk.
	MFilesortOnDiskSumPerSec float32 `json:"m_filesort_on_disk_sum_per_sec,omitempty"`

	// The number of joins that used a range search on a reference table.
	MSelectFullRangeJoinSumPerSec float32 `json:"m_select_full_range_join_sum_per_sec,omitempty"`

	// The number of joins that used ranges on the first table.
	MSelectRangeSumPerSec float32 `json:"m_select_range_sum_per_sec,omitempty"`

	// The number of joins without keys that check for key usage after each row.
	MSelectRangeCheckSumPerSec float32 `json:"m_select_range_check_sum_per_sec,omitempty"`

	// The number of sorts that were done using ranges.
	MSortRangeSumPerSec float32 `json:"m_sort_range_sum_per_sec,omitempty"`

	// The number of sorted rows.
	MSortRowsSumPerSec float32 `json:"m_sort_rows_sum_per_sec,omitempty"`

	// The number of sorts that were done by scanning the table.
	MSortScanSumPerSec float32 `json:"m_sort_scan_sum_per_sec,omitempty"`

	// The number of queries without index.
	MNoIndexUsedSumPerSec float32 `json:"m_no_index_used_sum_per_sec,omitempty"`

	// The number of queries without good index.
	MNoGoodIndexUsedSumPerSec float32 `json:"m_no_good_index_used_sum_per_sec,omitempty"`

	// MongoDB metrics.
	//
	// The number of returned documents.
	MDocsReturnedSumPerSec float32 `json:"m_docs_returned_sum_per_sec,omitempty"`

	// The response length of the query result in bytes.
	MResponseLengthSumPerSec float32 `json:"m_response_length_sum_per_sec,omitempty"`

	// The number of scanned documents.
	MDocsScannedSumPerSec float32 `json:"m_docs_scanned_sum_per_sec,omitempty"`

	// PostgreSQL metrics.
	//
	// Total number of shared block cache hits by the statement.
	MSharedBlksHitSumPerSec float32 `json:"m_shared_blks_hit_sum_per_sec,omitempty"`

	// Total number of shared blocks read by the statement.
	MSharedBlksReadSumPerSec float32 `json:"m_shared_blks_read_sum_per_sec,omitempty"`

	// Total number of shared blocks dirtied by the statement.
	MSharedBlksDirtiedSumPerSec float32 `json:"m_shared_blks_dirtied_sum_per_sec,omitempty"`

	// Total number of shared blocks written by the statement.
	MSharedBlksWrittenSumPerSec float32 `json:"m_shared_blks_written_sum_per_sec,omitempty"`

	// Total number of local block cache hits by the statement.
	MLocalBlksHitSumPerSec float32 `json:"m_local_blks_hit_sum_per_sec,omitempty"`

	// Total number of local blocks read by the statement.
	MLocalBlksReadSumPerSec float32 `json:"m_local_blks_read_sum_per_sec,omitempty"`

	// Total number of local blocks dirtied by the statement.
	MLocalBlksDirtiedSumPerSec float32 `json:"m_local_blks_dirtied_sum_per_sec,omitempty"`

	// Total number of local blocks written by the statement.
	MLocalBlksWrittenSumPerSec float32 `json:"m_local_blks_written_sum_per_sec,omitempty"`

	// Total number of temp blocks read by the statement.
	MTempBlksReadSumPerSec float32 `json:"m_temp_blks_read_sum_per_sec,omitempty"`

	// Total number of temp blocks written by the statement.
	MTempBlksWrittenSumPerSec float32 `json:"m_temp_blks_written_sum_per_sec,omitempty"`

	// Total time the statement spent reading blocks, in milliseconds (if track_io_timing is enabled, otherwise zero).
	MBlkReadTimeSumPerSec float32 `json:"m_blk_read_time_sum_per_sec,omitempty"`

	// Total time the statement spent writing blocks, in milliseconds (if track_io_timing is enabled, otherwise zero).
	MBlkWriteTimeSumPerSec float32 `json:"m_blk_write_time_sum_per_sec,omitempty"`

	// Total time user spent in query.
	MCPUUserTimeSumPerSec float32 `json:"m_cpu_user_time_sum_per_sec,omitempty"`

	// Total time system spent in query.
	MCPUSysTimeSumPerSec float32 `json:"m_cpu_sys_time_sum_per_sec,omitempty"`

	// pg_stat_monitor 0.9 metrics
	//
	// Total number of planned calls.
	MPlansCallsSumPerSec float32 `json:"m_plans_calls_sum_per_sec,omitempty"`

	// Total number of WAL (Write-ahead logging) records.
	MWalRecordsSumPerSec float32 `json:"m_wal_records_sum_per_sec,omitempty"`

	// Total number of FPI (full page images) in WAL (Write-ahead logging) records.
	MWalFpiSumPerSec float32 `json:"m_wal_fpi_sum_per_sec,omitempty"`

	// Total bytes of WAL (Write-ahead logging) records.
	MWalBytesSumPerSec float32 `json:"m_wal_bytes_sum_per_sec,omitempty"`

	// Plan time in per seconds.
	MPlanTimeSumPerSec float32 `json:"m_plan_time_sum_per_sec,omitempty"`
<<<<<<< HEAD

	// Top Query ID.
	TopQueryid string `json:"top_queryid,omitempty"`

	// Top Query plain text.
	TopQuery string `json:"top_query,omitempty"`

	// Plan ID.
	Planid string `json:"planid,omitempty"`

	// Query Plan plain text.
	QueryPlan string `json:"query_plan,omitempty"`
=======
>>>>>>> PMM-8895-top-query-details
}

// Validate validates this sparkline items0
func (o *SparklineItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SparklineItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SparklineItems0) UnmarshalBinary(b []byte) error {
	var res SparklineItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TotalsAnon MetricValues is statistics of specific metric.
swagger:model TotalsAnon
*/
type TotalsAnon struct {

	// rate
	Rate float32 `json:"rate,omitempty"`

	// cnt
	Cnt float32 `json:"cnt,omitempty"`

	// sum
	Sum float32 `json:"sum,omitempty"`

	// min
	Min float32 `json:"min,omitempty"`

	// max
	Max float32 `json:"max,omitempty"`

	// avg
	Avg float32 `json:"avg,omitempty"`

	// p99
	P99 float32 `json:"p99,omitempty"`

	// percent of total
	PercentOfTotal float32 `json:"percent_of_total,omitempty"`
}

// Validate validates this totals anon
func (o *TotalsAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TotalsAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TotalsAnon) UnmarshalBinary(b []byte) error {
	var res TotalsAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
