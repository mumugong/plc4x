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

// AirConditioningDataSetHvacSetbackLimit is the corresponding interface of AirConditioningDataSetHvacSetbackLimit
type AirConditioningDataSetHvacSetbackLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetLimit returns Limit (property field)
	GetLimit() HVACTemperature
	// GetHvacModeAndFlags returns HvacModeAndFlags (property field)
	GetHvacModeAndFlags() HVACModeAndFlags
}

// AirConditioningDataSetHvacSetbackLimitExactly can be used when we want exactly this type and not a type which fulfills AirConditioningDataSetHvacSetbackLimit.
// This is useful for switch cases.
type AirConditioningDataSetHvacSetbackLimitExactly interface {
	AirConditioningDataSetHvacSetbackLimit
	isAirConditioningDataSetHvacSetbackLimit() bool
}

// _AirConditioningDataSetHvacSetbackLimit is the data-structure of this message
type _AirConditioningDataSetHvacSetbackLimit struct {
	*_AirConditioningData
	ZoneGroup        byte
	ZoneList         HVACZoneList
	Limit            HVACTemperature
	HvacModeAndFlags HVACModeAndFlags
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataSetHvacSetbackLimit) InitializeParent(parent AirConditioningData, commandTypeContainer AirConditioningCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_AirConditioningDataSetHvacSetbackLimit) GetParent() AirConditioningData {
	return m._AirConditioningData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetHvacSetbackLimit) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataSetHvacSetbackLimit) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataSetHvacSetbackLimit) GetLimit() HVACTemperature {
	return m.Limit
}

func (m *_AirConditioningDataSetHvacSetbackLimit) GetHvacModeAndFlags() HVACModeAndFlags {
	return m.HvacModeAndFlags
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataSetHvacSetbackLimit factory function for _AirConditioningDataSetHvacSetbackLimit
func NewAirConditioningDataSetHvacSetbackLimit(zoneGroup byte, zoneList HVACZoneList, limit HVACTemperature, hvacModeAndFlags HVACModeAndFlags, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataSetHvacSetbackLimit {
	_result := &_AirConditioningDataSetHvacSetbackLimit{
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		Limit:                limit,
		HvacModeAndFlags:     hvacModeAndFlags,
		_AirConditioningData: NewAirConditioningData(commandTypeContainer),
	}
	_result._AirConditioningData._AirConditioningDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetHvacSetbackLimit(structType interface{}) AirConditioningDataSetHvacSetbackLimit {
	if casted, ok := structType.(AirConditioningDataSetHvacSetbackLimit); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetHvacSetbackLimit); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetHvacSetbackLimit) GetTypeName() string {
	return "AirConditioningDataSetHvacSetbackLimit"
}

func (m *_AirConditioningDataSetHvacSetbackLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (limit)
	lengthInBits += m.Limit.GetLengthInBits(ctx)

	// Simple field (hvacModeAndFlags)
	lengthInBits += m.HvacModeAndFlags.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AirConditioningDataSetHvacSetbackLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AirConditioningDataSetHvacSetbackLimitParse(theBytes []byte) (AirConditioningDataSetHvacSetbackLimit, error) {
	return AirConditioningDataSetHvacSetbackLimitParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func AirConditioningDataSetHvacSetbackLimitParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AirConditioningDataSetHvacSetbackLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataSetHvacSetbackLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetHvacSetbackLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneGroup)
	_zoneGroup, _zoneGroupErr := readBuffer.ReadByte("zoneGroup")
	if _zoneGroupErr != nil {
		return nil, errors.Wrap(_zoneGroupErr, "Error parsing 'zoneGroup' field of AirConditioningDataSetHvacSetbackLimit")
	}
	zoneGroup := _zoneGroup

	// Simple Field (zoneList)
	if pullErr := readBuffer.PullContext("zoneList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneList")
	}
	_zoneList, _zoneListErr := HVACZoneListParseWithBuffer(ctx, readBuffer)
	if _zoneListErr != nil {
		return nil, errors.Wrap(_zoneListErr, "Error parsing 'zoneList' field of AirConditioningDataSetHvacSetbackLimit")
	}
	zoneList := _zoneList.(HVACZoneList)
	if closeErr := readBuffer.CloseContext("zoneList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneList")
	}

	// Simple Field (limit)
	if pullErr := readBuffer.PullContext("limit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for limit")
	}
	_limit, _limitErr := HVACTemperatureParseWithBuffer(ctx, readBuffer)
	if _limitErr != nil {
		return nil, errors.Wrap(_limitErr, "Error parsing 'limit' field of AirConditioningDataSetHvacSetbackLimit")
	}
	limit := _limit.(HVACTemperature)
	if closeErr := readBuffer.CloseContext("limit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for limit")
	}

	// Simple Field (hvacModeAndFlags)
	if pullErr := readBuffer.PullContext("hvacModeAndFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hvacModeAndFlags")
	}
	_hvacModeAndFlags, _hvacModeAndFlagsErr := HVACModeAndFlagsParseWithBuffer(ctx, readBuffer)
	if _hvacModeAndFlagsErr != nil {
		return nil, errors.Wrap(_hvacModeAndFlagsErr, "Error parsing 'hvacModeAndFlags' field of AirConditioningDataSetHvacSetbackLimit")
	}
	hvacModeAndFlags := _hvacModeAndFlags.(HVACModeAndFlags)
	if closeErr := readBuffer.CloseContext("hvacModeAndFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hvacModeAndFlags")
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetHvacSetbackLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetHvacSetbackLimit")
	}

	// Create a partially initialized instance
	_child := &_AirConditioningDataSetHvacSetbackLimit{
		_AirConditioningData: &_AirConditioningData{},
		ZoneGroup:            zoneGroup,
		ZoneList:             zoneList,
		Limit:                limit,
		HvacModeAndFlags:     hvacModeAndFlags,
	}
	_child._AirConditioningData._AirConditioningDataChildRequirements = _child
	return _child, nil
}

func (m *_AirConditioningDataSetHvacSetbackLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataSetHvacSetbackLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetHvacSetbackLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetHvacSetbackLimit")
		}

		// Simple Field (zoneGroup)
		zoneGroup := byte(m.GetZoneGroup())
		_zoneGroupErr := writeBuffer.WriteByte("zoneGroup", (zoneGroup))
		if _zoneGroupErr != nil {
			return errors.Wrap(_zoneGroupErr, "Error serializing 'zoneGroup' field")
		}

		// Simple Field (zoneList)
		if pushErr := writeBuffer.PushContext("zoneList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for zoneList")
		}
		_zoneListErr := writeBuffer.WriteSerializable(ctx, m.GetZoneList())
		if popErr := writeBuffer.PopContext("zoneList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for zoneList")
		}
		if _zoneListErr != nil {
			return errors.Wrap(_zoneListErr, "Error serializing 'zoneList' field")
		}

		// Simple Field (limit)
		if pushErr := writeBuffer.PushContext("limit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for limit")
		}
		_limitErr := writeBuffer.WriteSerializable(ctx, m.GetLimit())
		if popErr := writeBuffer.PopContext("limit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for limit")
		}
		if _limitErr != nil {
			return errors.Wrap(_limitErr, "Error serializing 'limit' field")
		}

		// Simple Field (hvacModeAndFlags)
		if pushErr := writeBuffer.PushContext("hvacModeAndFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for hvacModeAndFlags")
		}
		_hvacModeAndFlagsErr := writeBuffer.WriteSerializable(ctx, m.GetHvacModeAndFlags())
		if popErr := writeBuffer.PopContext("hvacModeAndFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for hvacModeAndFlags")
		}
		if _hvacModeAndFlagsErr != nil {
			return errors.Wrap(_hvacModeAndFlagsErr, "Error serializing 'hvacModeAndFlags' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetHvacSetbackLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetHvacSetbackLimit")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetHvacSetbackLimit) isAirConditioningDataSetHvacSetbackLimit() bool {
	return true
}

func (m *_AirConditioningDataSetHvacSetbackLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
