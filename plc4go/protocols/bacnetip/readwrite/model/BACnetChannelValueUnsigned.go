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

// BACnetChannelValueUnsigned is the corresponding interface of BACnetChannelValueUnsigned
type BACnetChannelValueUnsigned interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetApplicationTagUnsignedInteger
}

// BACnetChannelValueUnsignedExactly can be used when we want exactly this type and not a type which fulfills BACnetChannelValueUnsigned.
// This is useful for switch cases.
type BACnetChannelValueUnsignedExactly interface {
	BACnetChannelValueUnsigned
	isBACnetChannelValueUnsigned() bool
}

// _BACnetChannelValueUnsigned is the data-structure of this message
type _BACnetChannelValueUnsigned struct {
	*_BACnetChannelValue
	UnsignedValue BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueUnsigned) InitializeParent(parent BACnetChannelValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetChannelValueUnsigned) GetParent() BACnetChannelValue {
	return m._BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueUnsigned) GetUnsignedValue() BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueUnsigned factory function for _BACnetChannelValueUnsigned
func NewBACnetChannelValueUnsigned(unsignedValue BACnetApplicationTagUnsignedInteger, peekedTagHeader BACnetTagHeader) *_BACnetChannelValueUnsigned {
	_result := &_BACnetChannelValueUnsigned{
		UnsignedValue:       unsignedValue,
		_BACnetChannelValue: NewBACnetChannelValue(peekedTagHeader),
	}
	_result._BACnetChannelValue._BACnetChannelValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueUnsigned(structType interface{}) BACnetChannelValueUnsigned {
	if casted, ok := structType.(BACnetChannelValueUnsigned); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueUnsigned); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueUnsigned) GetTypeName() string {
	return "BACnetChannelValueUnsigned"
}

func (m *_BACnetChannelValueUnsigned) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueUnsigned) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetChannelValueUnsignedParse(theBytes []byte) (BACnetChannelValueUnsigned, error) {
	return BACnetChannelValueUnsignedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetChannelValueUnsignedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetChannelValueUnsigned, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueUnsigned"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueUnsigned")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unsignedValue)
	if pullErr := readBuffer.PullContext("unsignedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unsignedValue")
	}
	_unsignedValue, _unsignedValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _unsignedValueErr != nil {
		return nil, errors.Wrap(_unsignedValueErr, "Error parsing 'unsignedValue' field of BACnetChannelValueUnsigned")
	}
	unsignedValue := _unsignedValue.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("unsignedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unsignedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueUnsigned"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueUnsigned")
	}

	// Create a partially initialized instance
	_child := &_BACnetChannelValueUnsigned{
		_BACnetChannelValue: &_BACnetChannelValue{},
		UnsignedValue:       unsignedValue,
	}
	_child._BACnetChannelValue._BACnetChannelValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetChannelValueUnsigned) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueUnsigned) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueUnsigned"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueUnsigned")
		}

		// Simple Field (unsignedValue)
		if pushErr := writeBuffer.PushContext("unsignedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for unsignedValue")
		}
		_unsignedValueErr := writeBuffer.WriteSerializable(ctx, m.GetUnsignedValue())
		if popErr := writeBuffer.PopContext("unsignedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for unsignedValue")
		}
		if _unsignedValueErr != nil {
			return errors.Wrap(_unsignedValueErr, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueUnsigned"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueUnsigned")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueUnsigned) isBACnetChannelValueUnsigned() bool {
	return true
}

func (m *_BACnetChannelValueUnsigned) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
