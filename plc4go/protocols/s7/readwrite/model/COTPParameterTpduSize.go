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

// COTPParameterTpduSize is the corresponding interface of COTPParameterTpduSize
type COTPParameterTpduSize interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	COTPParameter
	// GetTpduSize returns TpduSize (property field)
	GetTpduSize() COTPTpduSize
}

// COTPParameterTpduSizeExactly can be used when we want exactly this type and not a type which fulfills COTPParameterTpduSize.
// This is useful for switch cases.
type COTPParameterTpduSizeExactly interface {
	COTPParameterTpduSize
	isCOTPParameterTpduSize() bool
}

// _COTPParameterTpduSize is the data-structure of this message
type _COTPParameterTpduSize struct {
	*_COTPParameter
	TpduSize COTPTpduSize
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPParameterTpduSize) GetParameterType() uint8 {
	return 0xC0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPParameterTpduSize) InitializeParent(parent COTPParameter) {}

func (m *_COTPParameterTpduSize) GetParent() COTPParameter {
	return m._COTPParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPParameterTpduSize) GetTpduSize() COTPTpduSize {
	return m.TpduSize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCOTPParameterTpduSize factory function for _COTPParameterTpduSize
func NewCOTPParameterTpduSize(tpduSize COTPTpduSize, rest uint8) *_COTPParameterTpduSize {
	_result := &_COTPParameterTpduSize{
		TpduSize:       tpduSize,
		_COTPParameter: NewCOTPParameter(rest),
	}
	_result._COTPParameter._COTPParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCOTPParameterTpduSize(structType interface{}) COTPParameterTpduSize {
	if casted, ok := structType.(COTPParameterTpduSize); ok {
		return casted
	}
	if casted, ok := structType.(*COTPParameterTpduSize); ok {
		return *casted
	}
	return nil
}

func (m *_COTPParameterTpduSize) GetTypeName() string {
	return "COTPParameterTpduSize"
}

func (m *_COTPParameterTpduSize) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (tpduSize)
	lengthInBits += 8

	return lengthInBits
}

func (m *_COTPParameterTpduSize) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func COTPParameterTpduSizeParse(theBytes []byte, rest uint8) (COTPParameterTpduSize, error) {
	return COTPParameterTpduSizeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), rest)
}

func COTPParameterTpduSizeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, rest uint8) (COTPParameterTpduSize, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPParameterTpduSize"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPParameterTpduSize")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (tpduSize)
	if pullErr := readBuffer.PullContext("tpduSize"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for tpduSize")
	}
	_tpduSize, _tpduSizeErr := COTPTpduSizeParseWithBuffer(ctx, readBuffer)
	if _tpduSizeErr != nil {
		return nil, errors.Wrap(_tpduSizeErr, "Error parsing 'tpduSize' field of COTPParameterTpduSize")
	}
	tpduSize := _tpduSize
	if closeErr := readBuffer.CloseContext("tpduSize"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for tpduSize")
	}

	if closeErr := readBuffer.CloseContext("COTPParameterTpduSize"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPParameterTpduSize")
	}

	// Create a partially initialized instance
	_child := &_COTPParameterTpduSize{
		_COTPParameter: &_COTPParameter{
			Rest: rest,
		},
		TpduSize: tpduSize,
	}
	_child._COTPParameter._COTPParameterChildRequirements = _child
	return _child, nil
}

func (m *_COTPParameterTpduSize) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPParameterTpduSize) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPParameterTpduSize"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPParameterTpduSize")
		}

		// Simple Field (tpduSize)
		if pushErr := writeBuffer.PushContext("tpduSize"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for tpduSize")
		}
		_tpduSizeErr := writeBuffer.WriteSerializable(ctx, m.GetTpduSize())
		if popErr := writeBuffer.PopContext("tpduSize"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for tpduSize")
		}
		if _tpduSizeErr != nil {
			return errors.Wrap(_tpduSizeErr, "Error serializing 'tpduSize' field")
		}

		if popErr := writeBuffer.PopContext("COTPParameterTpduSize"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPParameterTpduSize")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_COTPParameterTpduSize) isCOTPParameterTpduSize() bool {
	return true
}

func (m *_COTPParameterTpduSize) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
