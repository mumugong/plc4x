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

// SecurityDataEmulatedKeypad is the corresponding interface of SecurityDataEmulatedKeypad
type SecurityDataEmulatedKeypad interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetKey returns Key (property field)
	GetKey() byte
	// GetIsAscii returns IsAscii (virtual field)
	GetIsAscii() bool
	// GetIsCustom returns IsCustom (virtual field)
	GetIsCustom() bool
	// GetIsEnter returns IsEnter (virtual field)
	GetIsEnter() bool
	// GetIsShift returns IsShift (virtual field)
	GetIsShift() bool
	// GetIsPanic returns IsPanic (virtual field)
	GetIsPanic() bool
	// GetIsFire returns IsFire (virtual field)
	GetIsFire() bool
	// GetIsARM returns IsARM (virtual field)
	GetIsARM() bool
	// GetIsAway returns IsAway (virtual field)
	GetIsAway() bool
	// GetIsNight returns IsNight (virtual field)
	GetIsNight() bool
	// GetIsDay returns IsDay (virtual field)
	GetIsDay() bool
	// GetIsVacation returns IsVacation (virtual field)
	GetIsVacation() bool
}

// SecurityDataEmulatedKeypadExactly can be used when we want exactly this type and not a type which fulfills SecurityDataEmulatedKeypad.
// This is useful for switch cases.
type SecurityDataEmulatedKeypadExactly interface {
	SecurityDataEmulatedKeypad
	isSecurityDataEmulatedKeypad() bool
}

// _SecurityDataEmulatedKeypad is the data-structure of this message
type _SecurityDataEmulatedKeypad struct {
	*_SecurityData
	Key byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataEmulatedKeypad) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataEmulatedKeypad) GetParent() SecurityData {
	return m._SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataEmulatedKeypad) GetKey() byte {
	return m.Key
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_SecurityDataEmulatedKeypad) GetIsAscii() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(bool((m.GetKey()) >= (0x00))) && bool(bool((m.GetKey()) <= (0x7F))))
}

func (m *_SecurityDataEmulatedKeypad) GetIsCustom() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) >= (0x80)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsEnter() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x0D)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsShift() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x80)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsPanic() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x81)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsFire() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x82)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsARM() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x83)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsAway() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x84)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsNight() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x85)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsDay() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x86)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsVacation() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x87)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataEmulatedKeypad factory function for _SecurityDataEmulatedKeypad
func NewSecurityDataEmulatedKeypad(key byte, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataEmulatedKeypad {
	_result := &_SecurityDataEmulatedKeypad{
		Key:           key,
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataEmulatedKeypad(structType interface{}) SecurityDataEmulatedKeypad {
	if casted, ok := structType.(SecurityDataEmulatedKeypad); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataEmulatedKeypad); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataEmulatedKeypad) GetTypeName() string {
	return "SecurityDataEmulatedKeypad"
}

func (m *_SecurityDataEmulatedKeypad) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (key)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_SecurityDataEmulatedKeypad) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataEmulatedKeypadParse(theBytes []byte) (SecurityDataEmulatedKeypad, error) {
	return SecurityDataEmulatedKeypadParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataEmulatedKeypadParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataEmulatedKeypad, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataEmulatedKeypad"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataEmulatedKeypad")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (key)
	_key, _keyErr := readBuffer.ReadByte("key")
	if _keyErr != nil {
		return nil, errors.Wrap(_keyErr, "Error parsing 'key' field of SecurityDataEmulatedKeypad")
	}
	key := _key

	// Virtual field
	_isAscii := bool(bool((key) >= (0x00))) && bool(bool((key) <= (0x7F)))
	isAscii := bool(_isAscii)
	_ = isAscii

	// Virtual field
	_isCustom := bool((key) >= (0x80))
	isCustom := bool(_isCustom)
	_ = isCustom

	// Virtual field
	_isEnter := bool((key) == (0x0D))
	isEnter := bool(_isEnter)
	_ = isEnter

	// Virtual field
	_isShift := bool((key) == (0x80))
	isShift := bool(_isShift)
	_ = isShift

	// Virtual field
	_isPanic := bool((key) == (0x81))
	isPanic := bool(_isPanic)
	_ = isPanic

	// Virtual field
	_isFire := bool((key) == (0x82))
	isFire := bool(_isFire)
	_ = isFire

	// Virtual field
	_isARM := bool((key) == (0x83))
	isARM := bool(_isARM)
	_ = isARM

	// Virtual field
	_isAway := bool((key) == (0x84))
	isAway := bool(_isAway)
	_ = isAway

	// Virtual field
	_isNight := bool((key) == (0x85))
	isNight := bool(_isNight)
	_ = isNight

	// Virtual field
	_isDay := bool((key) == (0x86))
	isDay := bool(_isDay)
	_ = isDay

	// Virtual field
	_isVacation := bool((key) == (0x87))
	isVacation := bool(_isVacation)
	_ = isVacation

	if closeErr := readBuffer.CloseContext("SecurityDataEmulatedKeypad"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataEmulatedKeypad")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataEmulatedKeypad{
		_SecurityData: &_SecurityData{},
		Key:           key,
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataEmulatedKeypad) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataEmulatedKeypad) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataEmulatedKeypad"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataEmulatedKeypad")
		}

		// Simple Field (key)
		key := byte(m.GetKey())
		_keyErr := writeBuffer.WriteByte("key", (key))
		if _keyErr != nil {
			return errors.Wrap(_keyErr, "Error serializing 'key' field")
		}
		// Virtual field
		if _isAsciiErr := writeBuffer.WriteVirtual(ctx, "isAscii", m.GetIsAscii()); _isAsciiErr != nil {
			return errors.Wrap(_isAsciiErr, "Error serializing 'isAscii' field")
		}
		// Virtual field
		if _isCustomErr := writeBuffer.WriteVirtual(ctx, "isCustom", m.GetIsCustom()); _isCustomErr != nil {
			return errors.Wrap(_isCustomErr, "Error serializing 'isCustom' field")
		}
		// Virtual field
		if _isEnterErr := writeBuffer.WriteVirtual(ctx, "isEnter", m.GetIsEnter()); _isEnterErr != nil {
			return errors.Wrap(_isEnterErr, "Error serializing 'isEnter' field")
		}
		// Virtual field
		if _isShiftErr := writeBuffer.WriteVirtual(ctx, "isShift", m.GetIsShift()); _isShiftErr != nil {
			return errors.Wrap(_isShiftErr, "Error serializing 'isShift' field")
		}
		// Virtual field
		if _isPanicErr := writeBuffer.WriteVirtual(ctx, "isPanic", m.GetIsPanic()); _isPanicErr != nil {
			return errors.Wrap(_isPanicErr, "Error serializing 'isPanic' field")
		}
		// Virtual field
		if _isFireErr := writeBuffer.WriteVirtual(ctx, "isFire", m.GetIsFire()); _isFireErr != nil {
			return errors.Wrap(_isFireErr, "Error serializing 'isFire' field")
		}
		// Virtual field
		if _isARMErr := writeBuffer.WriteVirtual(ctx, "isARM", m.GetIsARM()); _isARMErr != nil {
			return errors.Wrap(_isARMErr, "Error serializing 'isARM' field")
		}
		// Virtual field
		if _isAwayErr := writeBuffer.WriteVirtual(ctx, "isAway", m.GetIsAway()); _isAwayErr != nil {
			return errors.Wrap(_isAwayErr, "Error serializing 'isAway' field")
		}
		// Virtual field
		if _isNightErr := writeBuffer.WriteVirtual(ctx, "isNight", m.GetIsNight()); _isNightErr != nil {
			return errors.Wrap(_isNightErr, "Error serializing 'isNight' field")
		}
		// Virtual field
		if _isDayErr := writeBuffer.WriteVirtual(ctx, "isDay", m.GetIsDay()); _isDayErr != nil {
			return errors.Wrap(_isDayErr, "Error serializing 'isDay' field")
		}
		// Virtual field
		if _isVacationErr := writeBuffer.WriteVirtual(ctx, "isVacation", m.GetIsVacation()); _isVacationErr != nil {
			return errors.Wrap(_isVacationErr, "Error serializing 'isVacation' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataEmulatedKeypad"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataEmulatedKeypad")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataEmulatedKeypad) isSecurityDataEmulatedKeypad() bool {
	return true
}

func (m *_SecurityDataEmulatedKeypad) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
