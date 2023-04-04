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

// BACnetConstructedDataAccompanimentTime is the corresponding interface of BACnetConstructedDataAccompanimentTime
type BACnetConstructedDataAccompanimentTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAccompanimentTime returns AccompanimentTime (property field)
	GetAccompanimentTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataAccompanimentTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccompanimentTime.
// This is useful for switch cases.
type BACnetConstructedDataAccompanimentTimeExactly interface {
	BACnetConstructedDataAccompanimentTime
	isBACnetConstructedDataAccompanimentTime() bool
}

// _BACnetConstructedDataAccompanimentTime is the data-structure of this message
type _BACnetConstructedDataAccompanimentTime struct {
	*_BACnetConstructedData
	AccompanimentTime BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccompanimentTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccompanimentTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCOMPANIMENT_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccompanimentTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAccompanimentTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccompanimentTime) GetAccompanimentTime() BACnetApplicationTagUnsignedInteger {
	return m.AccompanimentTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccompanimentTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetAccompanimentTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccompanimentTime factory function for _BACnetConstructedDataAccompanimentTime
func NewBACnetConstructedDataAccompanimentTime(accompanimentTime BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccompanimentTime {
	_result := &_BACnetConstructedDataAccompanimentTime{
		AccompanimentTime:      accompanimentTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccompanimentTime(structType interface{}) BACnetConstructedDataAccompanimentTime {
	if casted, ok := structType.(BACnetConstructedDataAccompanimentTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccompanimentTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccompanimentTime) GetTypeName() string {
	return "BACnetConstructedDataAccompanimentTime"
}

func (m *_BACnetConstructedDataAccompanimentTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (accompanimentTime)
	lengthInBits += m.AccompanimentTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccompanimentTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAccompanimentTimeParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccompanimentTime, error) {
	return BACnetConstructedDataAccompanimentTimeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAccompanimentTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccompanimentTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccompanimentTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccompanimentTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (accompanimentTime)
	if pullErr := readBuffer.PullContext("accompanimentTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accompanimentTime")
	}
	_accompanimentTime, _accompanimentTimeErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _accompanimentTimeErr != nil {
		return nil, errors.Wrap(_accompanimentTimeErr, "Error parsing 'accompanimentTime' field of BACnetConstructedDataAccompanimentTime")
	}
	accompanimentTime := _accompanimentTime.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("accompanimentTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accompanimentTime")
	}

	// Virtual field
	_actualValue := accompanimentTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccompanimentTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccompanimentTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAccompanimentTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AccompanimentTime: accompanimentTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAccompanimentTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccompanimentTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccompanimentTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccompanimentTime")
		}

		// Simple Field (accompanimentTime)
		if pushErr := writeBuffer.PushContext("accompanimentTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accompanimentTime")
		}
		_accompanimentTimeErr := writeBuffer.WriteSerializable(ctx, m.GetAccompanimentTime())
		if popErr := writeBuffer.PopContext("accompanimentTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accompanimentTime")
		}
		if _accompanimentTimeErr != nil {
			return errors.Wrap(_accompanimentTimeErr, "Error serializing 'accompanimentTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccompanimentTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccompanimentTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccompanimentTime) isBACnetConstructedDataAccompanimentTime() bool {
	return true
}

func (m *_BACnetConstructedDataAccompanimentTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
