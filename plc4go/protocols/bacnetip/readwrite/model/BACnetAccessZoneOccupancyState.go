/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetAccessZoneOccupancyState is an enum
type BACnetAccessZoneOccupancyState uint16

type IBACnetAccessZoneOccupancyState interface {
	fmt.Stringer
	utils.Serializable
}

const (
	BACnetAccessZoneOccupancyState_NORMAL                   BACnetAccessZoneOccupancyState = 0
	BACnetAccessZoneOccupancyState_BELOW_LOWER_LIMIT        BACnetAccessZoneOccupancyState = 1
	BACnetAccessZoneOccupancyState_AT_LOWER_LIMIT           BACnetAccessZoneOccupancyState = 2
	BACnetAccessZoneOccupancyState_AT_UPPER_LIMIT           BACnetAccessZoneOccupancyState = 3
	BACnetAccessZoneOccupancyState_ABOVE_UPPER_LIMIT        BACnetAccessZoneOccupancyState = 4
	BACnetAccessZoneOccupancyState_DISABLED                 BACnetAccessZoneOccupancyState = 5
	BACnetAccessZoneOccupancyState_NOT_SUPPORTED            BACnetAccessZoneOccupancyState = 6
	BACnetAccessZoneOccupancyState_VENDOR_PROPRIETARY_VALUE BACnetAccessZoneOccupancyState = 0xFFFF
)

var BACnetAccessZoneOccupancyStateValues []BACnetAccessZoneOccupancyState

func init() {
	_ = errors.New
	BACnetAccessZoneOccupancyStateValues = []BACnetAccessZoneOccupancyState{
		BACnetAccessZoneOccupancyState_NORMAL,
		BACnetAccessZoneOccupancyState_BELOW_LOWER_LIMIT,
		BACnetAccessZoneOccupancyState_AT_LOWER_LIMIT,
		BACnetAccessZoneOccupancyState_AT_UPPER_LIMIT,
		BACnetAccessZoneOccupancyState_ABOVE_UPPER_LIMIT,
		BACnetAccessZoneOccupancyState_DISABLED,
		BACnetAccessZoneOccupancyState_NOT_SUPPORTED,
		BACnetAccessZoneOccupancyState_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetAccessZoneOccupancyStateByValue(value uint16) (enum BACnetAccessZoneOccupancyState, ok bool) {
	switch value {
	case 0:
		return BACnetAccessZoneOccupancyState_NORMAL, true
	case 0xFFFF:
		return BACnetAccessZoneOccupancyState_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetAccessZoneOccupancyState_BELOW_LOWER_LIMIT, true
	case 2:
		return BACnetAccessZoneOccupancyState_AT_LOWER_LIMIT, true
	case 3:
		return BACnetAccessZoneOccupancyState_AT_UPPER_LIMIT, true
	case 4:
		return BACnetAccessZoneOccupancyState_ABOVE_UPPER_LIMIT, true
	case 5:
		return BACnetAccessZoneOccupancyState_DISABLED, true
	case 6:
		return BACnetAccessZoneOccupancyState_NOT_SUPPORTED, true
	}
	return 0, false
}

func BACnetAccessZoneOccupancyStateByName(value string) (enum BACnetAccessZoneOccupancyState, ok bool) {
	switch value {
	case "NORMAL":
		return BACnetAccessZoneOccupancyState_NORMAL, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetAccessZoneOccupancyState_VENDOR_PROPRIETARY_VALUE, true
	case "BELOW_LOWER_LIMIT":
		return BACnetAccessZoneOccupancyState_BELOW_LOWER_LIMIT, true
	case "AT_LOWER_LIMIT":
		return BACnetAccessZoneOccupancyState_AT_LOWER_LIMIT, true
	case "AT_UPPER_LIMIT":
		return BACnetAccessZoneOccupancyState_AT_UPPER_LIMIT, true
	case "ABOVE_UPPER_LIMIT":
		return BACnetAccessZoneOccupancyState_ABOVE_UPPER_LIMIT, true
	case "DISABLED":
		return BACnetAccessZoneOccupancyState_DISABLED, true
	case "NOT_SUPPORTED":
		return BACnetAccessZoneOccupancyState_NOT_SUPPORTED, true
	}
	return 0, false
}

func BACnetAccessZoneOccupancyStateKnows(value uint16) bool {
	for _, typeValue := range BACnetAccessZoneOccupancyStateValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetAccessZoneOccupancyState(structType interface{}) BACnetAccessZoneOccupancyState {
	castFunc := func(typ interface{}) BACnetAccessZoneOccupancyState {
		if sBACnetAccessZoneOccupancyState, ok := typ.(BACnetAccessZoneOccupancyState); ok {
			return sBACnetAccessZoneOccupancyState
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAccessZoneOccupancyState) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m BACnetAccessZoneOccupancyState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessZoneOccupancyStateParse(ctx context.Context, theBytes []byte) (BACnetAccessZoneOccupancyState, error) {
	return BACnetAccessZoneOccupancyStateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAccessZoneOccupancyStateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessZoneOccupancyState, error) {
	val, err := readBuffer.ReadUint16("BACnetAccessZoneOccupancyState", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetAccessZoneOccupancyState")
	}
	if enum, ok := BACnetAccessZoneOccupancyStateByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetAccessZoneOccupancyState(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetAccessZoneOccupancyState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetAccessZoneOccupancyState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetAccessZoneOccupancyState", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAccessZoneOccupancyState) PLC4XEnumName() string {
	switch e {
	case BACnetAccessZoneOccupancyState_NORMAL:
		return "NORMAL"
	case BACnetAccessZoneOccupancyState_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetAccessZoneOccupancyState_BELOW_LOWER_LIMIT:
		return "BELOW_LOWER_LIMIT"
	case BACnetAccessZoneOccupancyState_AT_LOWER_LIMIT:
		return "AT_LOWER_LIMIT"
	case BACnetAccessZoneOccupancyState_AT_UPPER_LIMIT:
		return "AT_UPPER_LIMIT"
	case BACnetAccessZoneOccupancyState_ABOVE_UPPER_LIMIT:
		return "ABOVE_UPPER_LIMIT"
	case BACnetAccessZoneOccupancyState_DISABLED:
		return "DISABLED"
	case BACnetAccessZoneOccupancyState_NOT_SUPPORTED:
		return "NOT_SUPPORTED"
	}
	return ""
}

func (e BACnetAccessZoneOccupancyState) String() string {
	return e.PLC4XEnumName()
}
