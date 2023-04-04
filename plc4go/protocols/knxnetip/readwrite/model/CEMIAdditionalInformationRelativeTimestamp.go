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

// Constant values.
const CEMIAdditionalInformationRelativeTimestamp_LEN uint8 = uint8(2)

// CEMIAdditionalInformationRelativeTimestamp is the corresponding interface of CEMIAdditionalInformationRelativeTimestamp
type CEMIAdditionalInformationRelativeTimestamp interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMIAdditionalInformation
	// GetRelativeTimestamp returns RelativeTimestamp (property field)
	GetRelativeTimestamp() RelativeTimestamp
}

// CEMIAdditionalInformationRelativeTimestampExactly can be used when we want exactly this type and not a type which fulfills CEMIAdditionalInformationRelativeTimestamp.
// This is useful for switch cases.
type CEMIAdditionalInformationRelativeTimestampExactly interface {
	CEMIAdditionalInformationRelativeTimestamp
	isCEMIAdditionalInformationRelativeTimestamp() bool
}

// _CEMIAdditionalInformationRelativeTimestamp is the data-structure of this message
type _CEMIAdditionalInformationRelativeTimestamp struct {
	*_CEMIAdditionalInformation
	RelativeTimestamp RelativeTimestamp
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetAdditionalInformationType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CEMIAdditionalInformationRelativeTimestamp) InitializeParent(parent CEMIAdditionalInformation) {
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetParent() CEMIAdditionalInformation {
	return m._CEMIAdditionalInformation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetRelativeTimestamp() RelativeTimestamp {
	return m.RelativeTimestamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetLen() uint8 {
	return CEMIAdditionalInformationRelativeTimestamp_LEN
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCEMIAdditionalInformationRelativeTimestamp factory function for _CEMIAdditionalInformationRelativeTimestamp
func NewCEMIAdditionalInformationRelativeTimestamp(relativeTimestamp RelativeTimestamp) *_CEMIAdditionalInformationRelativeTimestamp {
	_result := &_CEMIAdditionalInformationRelativeTimestamp{
		RelativeTimestamp:          relativeTimestamp,
		_CEMIAdditionalInformation: NewCEMIAdditionalInformation(),
	}
	_result._CEMIAdditionalInformation._CEMIAdditionalInformationChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCEMIAdditionalInformationRelativeTimestamp(structType interface{}) CEMIAdditionalInformationRelativeTimestamp {
	if casted, ok := structType.(CEMIAdditionalInformationRelativeTimestamp); ok {
		return casted
	}
	if casted, ok := structType.(*CEMIAdditionalInformationRelativeTimestamp); ok {
		return *casted
	}
	return nil
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetTypeName() string {
	return "CEMIAdditionalInformationRelativeTimestamp"
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Const Field (len)
	lengthInBits += 8

	// Simple field (relativeTimestamp)
	lengthInBits += m.RelativeTimestamp.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CEMIAdditionalInformationRelativeTimestampParse(theBytes []byte) (CEMIAdditionalInformationRelativeTimestamp, error) {
	return CEMIAdditionalInformationRelativeTimestampParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func CEMIAdditionalInformationRelativeTimestampParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CEMIAdditionalInformationRelativeTimestamp, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CEMIAdditionalInformationRelativeTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CEMIAdditionalInformationRelativeTimestamp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (len)
	len, _lenErr := readBuffer.ReadUint8("len", 8)
	if _lenErr != nil {
		return nil, errors.Wrap(_lenErr, "Error parsing 'len' field of CEMIAdditionalInformationRelativeTimestamp")
	}
	if len != CEMIAdditionalInformationRelativeTimestamp_LEN {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CEMIAdditionalInformationRelativeTimestamp_LEN) + " but got " + fmt.Sprintf("%d", len))
	}

	// Simple Field (relativeTimestamp)
	if pullErr := readBuffer.PullContext("relativeTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for relativeTimestamp")
	}
	_relativeTimestamp, _relativeTimestampErr := RelativeTimestampParseWithBuffer(ctx, readBuffer)
	if _relativeTimestampErr != nil {
		return nil, errors.Wrap(_relativeTimestampErr, "Error parsing 'relativeTimestamp' field of CEMIAdditionalInformationRelativeTimestamp")
	}
	relativeTimestamp := _relativeTimestamp.(RelativeTimestamp)
	if closeErr := readBuffer.CloseContext("relativeTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for relativeTimestamp")
	}

	if closeErr := readBuffer.CloseContext("CEMIAdditionalInformationRelativeTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CEMIAdditionalInformationRelativeTimestamp")
	}

	// Create a partially initialized instance
	_child := &_CEMIAdditionalInformationRelativeTimestamp{
		_CEMIAdditionalInformation: &_CEMIAdditionalInformation{},
		RelativeTimestamp:          relativeTimestamp,
	}
	_child._CEMIAdditionalInformation._CEMIAdditionalInformationChildRequirements = _child
	return _child, nil
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CEMIAdditionalInformationRelativeTimestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CEMIAdditionalInformationRelativeTimestamp")
		}

		// Const Field (len)
		_lenErr := writeBuffer.WriteUint8("len", 8, 2)
		if _lenErr != nil {
			return errors.Wrap(_lenErr, "Error serializing 'len' field")
		}

		// Simple Field (relativeTimestamp)
		if pushErr := writeBuffer.PushContext("relativeTimestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for relativeTimestamp")
		}
		_relativeTimestampErr := writeBuffer.WriteSerializable(ctx, m.GetRelativeTimestamp())
		if popErr := writeBuffer.PopContext("relativeTimestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for relativeTimestamp")
		}
		if _relativeTimestampErr != nil {
			return errors.Wrap(_relativeTimestampErr, "Error serializing 'relativeTimestamp' field")
		}

		if popErr := writeBuffer.PopContext("CEMIAdditionalInformationRelativeTimestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CEMIAdditionalInformationRelativeTimestamp")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) isCEMIAdditionalInformationRelativeTimestamp() bool {
	return true
}

func (m *_CEMIAdditionalInformationRelativeTimestamp) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
