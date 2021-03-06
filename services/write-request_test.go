// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package services

import (
	"testing"
	"time"

	"github.com/wmnsk/gopcua/datatypes"
	"github.com/wmnsk/gopcua/utils/codectest"
)

func TestWriteRequest(t *testing.T) {
	cases := []codectest.Case{
		{
			Name: "single-writevalue",
			Struct: NewWriteRequest(
				NewRequestHeader(
					datatypes.NewOpaqueNodeID(0x00, []byte{
						0x08, 0x22, 0x87, 0x62, 0xba, 0x81, 0xe1, 0x11,
						0xa6, 0x43, 0xf8, 0x77, 0x7b, 0xc6, 0x2f, 0xc8,
					}),
					time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
					1, 0, 0, "", NewNullAdditionalHeader(), nil,
				),
				datatypes.NewWriteValue(
					datatypes.NewFourByteNodeID(0, 2256),
					datatypes.IntegerIDValue,
					"",
					datatypes.NewDataValue(
						true, false, true, false, true, false,
						datatypes.NewVariant(datatypes.NewFloat(2.50017)),
						0,
						time.Date(2018, time.September, 17, 14, 28, 29, 112000000, time.UTC),
						0,
						time.Date(2018, time.September, 17, 14, 28, 29, 112000000, time.UTC),
						0,
					),
				),
			),
			Bytes: []byte{
				// TypeID
				0x01, 0x00, 0xa1, 0x02,
				// RequestHeader
				// AuthenticationToken
				0x05, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x08,
				0x22, 0x87, 0x62, 0xba, 0x81, 0xe1, 0x11, 0xa6,
				0x43, 0xf8, 0x77, 0x7b, 0xc6, 0x2f, 0xc8,
				// Timestamp
				0x00, 0x98, 0x67, 0xdd, 0xfd, 0x30, 0xd4, 0x01,
				// RequestHandle
				0x01, 0x00, 0x00, 0x00,
				// ReturnDiagnostics
				0x00, 0x00, 0x00, 0x00,
				// AuditEntryID
				0xff, 0xff, 0xff, 0xff,
				// TimeoutHint
				0x00, 0x00, 0x00, 0x00,
				// AdditionalHeader
				0x00, 0x00, 0x00,
				// NodesToWrite
				// ArraySize
				0x01, 0x00, 0x00, 0x00,
				// NodeID
				0x01, 0x00, 0xd0, 0x08,
				// AttributeID
				0x0d, 0x00, 0x00, 0x00,
				// IndexRange
				0xff, 0xff, 0xff, 0xff,
				// Value
				0x0d, 0x0a, 0xc9, 0x02, 0x20, 0x40, 0x80, 0x3b,
				0xe8, 0xb3, 0x92, 0x4e, 0xd4, 0x01, 0x80, 0x3b,
				0xe8, 0xb3, 0x92, 0x4e, 0xd4, 0x01,
			},
		},
		{
			Name: "multiple-writevalue",
			Struct: NewWriteRequest(
				NewRequestHeader(
					datatypes.NewOpaqueNodeID(0x00, []byte{
						0x08, 0x22, 0x87, 0x62, 0xba, 0x81, 0xe1, 0x11,
						0xa6, 0x43, 0xf8, 0x77, 0x7b, 0xc6, 0x2f, 0xc8,
					}),
					time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
					1, 0, 0, "", NewNullAdditionalHeader(), nil,
				),
				datatypes.NewWriteValue(
					datatypes.NewFourByteNodeID(0, 2256),
					datatypes.IntegerIDValue,
					"",
					datatypes.NewDataValue(
						true, false, true, false, true, false,
						datatypes.NewVariant(datatypes.NewFloat(2.50017)),
						0,
						time.Date(2018, time.September, 17, 14, 28, 29, 112000000, time.UTC),
						0,
						time.Date(2018, time.September, 17, 14, 28, 29, 112000000, time.UTC),
						0,
					),
				),
				datatypes.NewWriteValue(
					datatypes.NewFourByteNodeID(0, 2256),
					datatypes.IntegerIDValue,
					"",
					datatypes.NewDataValue(
						true, false, true, false, true, false,
						datatypes.NewVariant(datatypes.NewFloat(2.50017)),
						0,
						time.Date(2018, time.September, 17, 14, 28, 29, 112000000, time.UTC),
						0,
						time.Date(2018, time.September, 17, 14, 28, 29, 112000000, time.UTC),
						0,
					),
				),
			),
			Bytes: []byte{
				// TypeID
				0x01, 0x00, 0xa1, 0x02,
				// RequestHeader
				// AuthenticationToken
				0x05, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x08,
				0x22, 0x87, 0x62, 0xba, 0x81, 0xe1, 0x11, 0xa6,
				0x43, 0xf8, 0x77, 0x7b, 0xc6, 0x2f, 0xc8,
				// Timestamp
				0x00, 0x98, 0x67, 0xdd, 0xfd, 0x30, 0xd4, 0x01,
				// RequestHandle
				0x01, 0x00, 0x00, 0x00,
				// ReturnDiagnostics
				0x00, 0x00, 0x00, 0x00,
				// AuditEntryID
				0xff, 0xff, 0xff, 0xff,
				// TimeoutHint
				0x00, 0x00, 0x00, 0x00,
				// AdditionalHeader
				0x00, 0x00, 0x00,
				// NodesToWrite
				// ArraySize
				0x02, 0x00, 0x00, 0x00,
				// NodeID
				0x01, 0x00, 0xd0, 0x08,
				// AttributeID
				0x0d, 0x00, 0x00, 0x00,
				// IndexRange
				0xff, 0xff, 0xff, 0xff,
				// Value
				0x0d, 0x0a, 0xc9, 0x02, 0x20, 0x40, 0x80, 0x3b,
				0xe8, 0xb3, 0x92, 0x4e, 0xd4, 0x01, 0x80, 0x3b,
				0xe8, 0xb3, 0x92, 0x4e, 0xd4, 0x01,
				// NodeID
				0x01, 0x00, 0xd0, 0x08,
				// AttributeID
				0x0d, 0x00, 0x00, 0x00,
				// IndexRange
				0xff, 0xff, 0xff, 0xff,
				// Value
				0x0d, 0x0a, 0xc9, 0x02, 0x20, 0x40, 0x80, 0x3b,
				0xe8, 0xb3, 0x92, 0x4e, 0xd4, 0x01, 0x80, 0x3b,
				0xe8, 0xb3, 0x92, 0x4e, 0xd4, 0x01,
			},
		},
	}
	codectest.Run(t, cases, func(b []byte) (codectest.S, error) {
		v, err := DecodeWriteRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})

	t.Run("service-id", func(t *testing.T) {
		id := new(WriteRequest).ServiceType()
		if got, want := id, uint16(ServiceTypeWriteRequest); got != want {
			t.Fatalf("got %d want %d", got, want)
		}
	})
}
