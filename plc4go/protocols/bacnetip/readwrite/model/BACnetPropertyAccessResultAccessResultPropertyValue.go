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

// BACnetPropertyAccessResultAccessResultPropertyValue is the corresponding interface of BACnetPropertyAccessResultAccessResultPropertyValue
type BACnetPropertyAccessResultAccessResultPropertyValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyAccessResultAccessResult
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() BACnetConstructedData
}

// BACnetPropertyAccessResultAccessResultPropertyValueExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyAccessResultAccessResultPropertyValue.
// This is useful for switch cases.
type BACnetPropertyAccessResultAccessResultPropertyValueExactly interface {
	BACnetPropertyAccessResultAccessResultPropertyValue
	isBACnetPropertyAccessResultAccessResultPropertyValue() bool
}

// _BACnetPropertyAccessResultAccessResultPropertyValue is the data-structure of this message
type _BACnetPropertyAccessResultAccessResultPropertyValue struct {
	*_BACnetPropertyAccessResultAccessResult
	PropertyValue BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) InitializeParent(parent BACnetPropertyAccessResultAccessResult, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) GetParent() BACnetPropertyAccessResultAccessResult {
	return m._BACnetPropertyAccessResultAccessResult
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) GetPropertyValue() BACnetConstructedData {
	return m.PropertyValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyAccessResultAccessResultPropertyValue factory function for _BACnetPropertyAccessResultAccessResultPropertyValue
func NewBACnetPropertyAccessResultAccessResultPropertyValue(propertyValue BACnetConstructedData, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetPropertyAccessResultAccessResultPropertyValue {
	_result := &_BACnetPropertyAccessResultAccessResultPropertyValue{
		PropertyValue:                           propertyValue,
		_BACnetPropertyAccessResultAccessResult: NewBACnetPropertyAccessResultAccessResult(peekedTagHeader, objectTypeArgument, propertyIdentifierArgument, propertyArrayIndexArgument),
	}
	_result._BACnetPropertyAccessResultAccessResult._BACnetPropertyAccessResultAccessResultChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyAccessResultAccessResultPropertyValue(structType interface{}) BACnetPropertyAccessResultAccessResultPropertyValue {
	if casted, ok := structType.(BACnetPropertyAccessResultAccessResultPropertyValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyAccessResultAccessResultPropertyValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) GetTypeName() string {
	return "BACnetPropertyAccessResultAccessResultPropertyValue"
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (propertyValue)
	lengthInBits += m.PropertyValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyAccessResultAccessResultPropertyValueParse(theBytes []byte, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetPropertyAccessResultAccessResultPropertyValue, error) {
	return BACnetPropertyAccessResultAccessResultPropertyValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), objectTypeArgument, propertyIdentifierArgument, propertyArrayIndexArgument)
}

func BACnetPropertyAccessResultAccessResultPropertyValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetPropertyAccessResultAccessResultPropertyValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyAccessResultAccessResultPropertyValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyAccessResultAccessResultPropertyValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (propertyValue)
	if pullErr := readBuffer.PullContext("propertyValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyValue")
	}
	_propertyValue, _propertyValueErr := BACnetConstructedDataParseWithBuffer(ctx, readBuffer, uint8(uint8(4)), BACnetObjectType(objectTypeArgument), BACnetPropertyIdentifier(propertyIdentifierArgument), propertyArrayIndexArgument)
	if _propertyValueErr != nil {
		return nil, errors.Wrap(_propertyValueErr, "Error parsing 'propertyValue' field of BACnetPropertyAccessResultAccessResultPropertyValue")
	}
	propertyValue := _propertyValue.(BACnetConstructedData)
	if closeErr := readBuffer.CloseContext("propertyValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyAccessResultAccessResultPropertyValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyAccessResultAccessResultPropertyValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyAccessResultAccessResultPropertyValue{
		_BACnetPropertyAccessResultAccessResult: &_BACnetPropertyAccessResultAccessResult{
			ObjectTypeArgument:         objectTypeArgument,
			PropertyIdentifierArgument: propertyIdentifierArgument,
			PropertyArrayIndexArgument: propertyArrayIndexArgument,
		},
		PropertyValue: propertyValue,
	}
	_child._BACnetPropertyAccessResultAccessResult._BACnetPropertyAccessResultAccessResultChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyAccessResultAccessResultPropertyValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyAccessResultAccessResultPropertyValue")
		}

		// Simple Field (propertyValue)
		if pushErr := writeBuffer.PushContext("propertyValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyValue")
		}
		_propertyValueErr := writeBuffer.WriteSerializable(ctx, m.GetPropertyValue())
		if popErr := writeBuffer.PopContext("propertyValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyValue")
		}
		if _propertyValueErr != nil {
			return errors.Wrap(_propertyValueErr, "Error serializing 'propertyValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyAccessResultAccessResultPropertyValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyAccessResultAccessResultPropertyValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) isBACnetPropertyAccessResultAccessResultPropertyValue() bool {
	return true
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
