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

// BACnetHostAddressName is the corresponding interface of BACnetHostAddressName
type BACnetHostAddressName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetHostAddress
	// GetName returns Name (property field)
	GetName() BACnetContextTagCharacterString
}

// BACnetHostAddressNameExactly can be used when we want exactly this type and not a type which fulfills BACnetHostAddressName.
// This is useful for switch cases.
type BACnetHostAddressNameExactly interface {
	BACnetHostAddressName
	isBACnetHostAddressName() bool
}

// _BACnetHostAddressName is the data-structure of this message
type _BACnetHostAddressName struct {
	*_BACnetHostAddress
	Name BACnetContextTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetHostAddressName) InitializeParent(parent BACnetHostAddress, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetHostAddressName) GetParent() BACnetHostAddress {
	return m._BACnetHostAddress
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetHostAddressName) GetName() BACnetContextTagCharacterString {
	return m.Name
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetHostAddressName factory function for _BACnetHostAddressName
func NewBACnetHostAddressName(name BACnetContextTagCharacterString, peekedTagHeader BACnetTagHeader) *_BACnetHostAddressName {
	_result := &_BACnetHostAddressName{
		Name:               name,
		_BACnetHostAddress: NewBACnetHostAddress(peekedTagHeader),
	}
	_result._BACnetHostAddress._BACnetHostAddressChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetHostAddressName(structType interface{}) BACnetHostAddressName {
	if casted, ok := structType.(BACnetHostAddressName); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetHostAddressName); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetHostAddressName) GetTypeName() string {
	return "BACnetHostAddressName"
}

func (m *_BACnetHostAddressName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetHostAddressName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetHostAddressNameParse(theBytes []byte) (BACnetHostAddressName, error) {
	return BACnetHostAddressNameParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetHostAddressNameParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetHostAddressName, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostAddressName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostAddressName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (name)
	if pullErr := readBuffer.PullContext("name"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for name")
	}
	_name, _nameErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_CHARACTER_STRING))
	if _nameErr != nil {
		return nil, errors.Wrap(_nameErr, "Error parsing 'name' field of BACnetHostAddressName")
	}
	name := _name.(BACnetContextTagCharacterString)
	if closeErr := readBuffer.CloseContext("name"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for name")
	}

	if closeErr := readBuffer.CloseContext("BACnetHostAddressName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostAddressName")
	}

	// Create a partially initialized instance
	_child := &_BACnetHostAddressName{
		_BACnetHostAddress: &_BACnetHostAddress{},
		Name:               name,
	}
	_child._BACnetHostAddress._BACnetHostAddressChildRequirements = _child
	return _child, nil
}

func (m *_BACnetHostAddressName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetHostAddressName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetHostAddressName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetHostAddressName")
		}

		// Simple Field (name)
		if pushErr := writeBuffer.PushContext("name"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for name")
		}
		_nameErr := writeBuffer.WriteSerializable(ctx, m.GetName())
		if popErr := writeBuffer.PopContext("name"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for name")
		}
		if _nameErr != nil {
			return errors.Wrap(_nameErr, "Error serializing 'name' field")
		}

		if popErr := writeBuffer.PopContext("BACnetHostAddressName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetHostAddressName")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetHostAddressName) isBACnetHostAddressName() bool {
	return true
}

func (m *_BACnetHostAddressName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
