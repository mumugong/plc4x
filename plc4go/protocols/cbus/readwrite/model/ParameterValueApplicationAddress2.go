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

// ParameterValueApplicationAddress2 is the corresponding interface of ParameterValueApplicationAddress2
type ParameterValueApplicationAddress2 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ParameterValue
	// GetValue returns Value (property field)
	GetValue() ApplicationAddress2
	// GetData returns Data (property field)
	GetData() []byte
}

// ParameterValueApplicationAddress2Exactly can be used when we want exactly this type and not a type which fulfills ParameterValueApplicationAddress2.
// This is useful for switch cases.
type ParameterValueApplicationAddress2Exactly interface {
	ParameterValueApplicationAddress2
	isParameterValueApplicationAddress2() bool
}

// _ParameterValueApplicationAddress2 is the data-structure of this message
type _ParameterValueApplicationAddress2 struct {
	*_ParameterValue
	Value ApplicationAddress2
	Data  []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueApplicationAddress2) GetParameterType() ParameterType {
	return ParameterType_APPLICATION_ADDRESS_2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueApplicationAddress2) InitializeParent(parent ParameterValue) {}

func (m *_ParameterValueApplicationAddress2) GetParent() ParameterValue {
	return m._ParameterValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueApplicationAddress2) GetValue() ApplicationAddress2 {
	return m.Value
}

func (m *_ParameterValueApplicationAddress2) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewParameterValueApplicationAddress2 factory function for _ParameterValueApplicationAddress2
func NewParameterValueApplicationAddress2(value ApplicationAddress2, data []byte, numBytes uint8) *_ParameterValueApplicationAddress2 {
	_result := &_ParameterValueApplicationAddress2{
		Value:           value,
		Data:            data,
		_ParameterValue: NewParameterValue(numBytes),
	}
	_result._ParameterValue._ParameterValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastParameterValueApplicationAddress2(structType interface{}) ParameterValueApplicationAddress2 {
	if casted, ok := structType.(ParameterValueApplicationAddress2); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueApplicationAddress2); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueApplicationAddress2) GetTypeName() string {
	return "ParameterValueApplicationAddress2"
}

func (m *_ParameterValueApplicationAddress2) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ParameterValueApplicationAddress2) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ParameterValueApplicationAddress2Parse(theBytes []byte, parameterType ParameterType, numBytes uint8) (ParameterValueApplicationAddress2, error) {
	return ParameterValueApplicationAddress2ParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), parameterType, numBytes)
}

func ParameterValueApplicationAddress2ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, parameterType ParameterType, numBytes uint8) (ParameterValueApplicationAddress2, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValueApplicationAddress2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueApplicationAddress2")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((numBytes) >= (1))) {
		return nil, errors.WithStack(utils.ParseValidationError{"ApplicationAddress2 has exactly one byte"})
	}

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
	_value, _valueErr := ApplicationAddress2ParseWithBuffer(ctx, readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of ParameterValueApplicationAddress2")
	}
	value := _value.(ApplicationAddress2)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(uint16(numBytes) - uint16(uint16(1)))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of ParameterValueApplicationAddress2")
	}

	if closeErr := readBuffer.CloseContext("ParameterValueApplicationAddress2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueApplicationAddress2")
	}

	// Create a partially initialized instance
	_child := &_ParameterValueApplicationAddress2{
		_ParameterValue: &_ParameterValue{
			NumBytes: numBytes,
		},
		Value: value,
		Data:  data,
	}
	_child._ParameterValue._ParameterValueChildRequirements = _child
	return _child, nil
}

func (m *_ParameterValueApplicationAddress2) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterValueApplicationAddress2) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueApplicationAddress2"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueApplicationAddress2")
		}

		// Simple Field (value)
		if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for value")
		}
		_valueErr := writeBuffer.WriteSerializable(ctx, m.GetValue())
		if popErr := writeBuffer.PopContext("value"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for value")
		}
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ParameterValueApplicationAddress2"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueApplicationAddress2")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ParameterValueApplicationAddress2) isParameterValueApplicationAddress2() bool {
	return true
}

func (m *_ParameterValueApplicationAddress2) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
