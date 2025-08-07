// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package changefeedpb

import (
	"encoding/json"

	"github.com/cockroachdb/errors"
)

// MarshalJSON provides clean, unwrapped JSON output for changefeedpb.Value.
func (v *Value) MarshalJSON() ([]byte, error) {
	if v == nil || v.Value == nil {
		return []byte("null"), nil
	}

	var out any
	switch val := v.Value.(type) {
	case *Value_BoolValue:
		out = val.BoolValue
	case *Value_Int64Value:
		out = val.Int64Value
	case *Value_DoubleValue:
		out = val.DoubleValue
	case *Value_StringValue:
		out = val.StringValue
	case *Value_DecimalValue:
		out = val.DecimalValue.Value
	case *Value_TimestampValue:
		out = val.TimestampValue
	case *Value_DateValue:
		out = val.DateValue
	case *Value_IntervalValue:
		out = val.IntervalValue
	case *Value_UuidValue:
		out = val.UuidValue
	case *Value_BytesValue:
		out = string(val.BytesValue)
	case *Value_TimeValue:
		out = val.TimeValue
	case *Value_ArrayValue:
		arr := make([]any, len(val.ArrayValue.Values))
		for i, elem := range val.ArrayValue.Values {
			arr[i] = elem
		}
		out = arr
	case *Value_TupleValue:
		m := make(map[string]any, len(val.TupleValue.Values))
		for k, v := range val.TupleValue.Values {
			m[k] = v
		}
		out = m
	default:
		return nil, errors.AssertionFailedf("unexpected protobuf value type: %T", v.Value)
	}
	return json.Marshal(out)
}

func (k *Key) MarshalJSON() ([]byte, error) {
	if k == nil || k.Key == nil {
		return []byte("null"), nil
	}
	return json.Marshal(k.Key)
}
