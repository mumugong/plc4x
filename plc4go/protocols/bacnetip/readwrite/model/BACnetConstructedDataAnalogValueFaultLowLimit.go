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

// BACnetConstructedDataAnalogValueFaultLowLimit is the corresponding interface of BACnetConstructedDataAnalogValueFaultLowLimit
type BACnetConstructedDataAnalogValueFaultLowLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultLowLimit returns FaultLowLimit (property field)
	GetFaultLowLimit() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataAnalogValueFaultLowLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAnalogValueFaultLowLimit.
// This is useful for switch cases.
type BACnetConstructedDataAnalogValueFaultLowLimitExactly interface {
	BACnetConstructedDataAnalogValueFaultLowLimit
	isBACnetConstructedDataAnalogValueFaultLowLimit() bool
}

// _BACnetConstructedDataAnalogValueFaultLowLimit is the data-structure of this message
type _BACnetConstructedDataAnalogValueFaultLowLimit struct {
	*_BACnetConstructedData
	FaultLowLimit BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ANALOG_VALUE
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_LOW_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) GetFaultLowLimit() BACnetApplicationTagReal {
	return m.FaultLowLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetFaultLowLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAnalogValueFaultLowLimit factory function for _BACnetConstructedDataAnalogValueFaultLowLimit
func NewBACnetConstructedDataAnalogValueFaultLowLimit(faultLowLimit BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAnalogValueFaultLowLimit {
	_result := &_BACnetConstructedDataAnalogValueFaultLowLimit{
		FaultLowLimit:          faultLowLimit,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAnalogValueFaultLowLimit(structType interface{}) BACnetConstructedDataAnalogValueFaultLowLimit {
	if casted, ok := structType.(BACnetConstructedDataAnalogValueFaultLowLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAnalogValueFaultLowLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) GetTypeName() string {
	return "BACnetConstructedDataAnalogValueFaultLowLimit"
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (faultLowLimit)
	lengthInBits += m.FaultLowLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAnalogValueFaultLowLimitParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAnalogValueFaultLowLimit, error) {
	return BACnetConstructedDataAnalogValueFaultLowLimitParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAnalogValueFaultLowLimitParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAnalogValueFaultLowLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAnalogValueFaultLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAnalogValueFaultLowLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (faultLowLimit)
	if pullErr := readBuffer.PullContext("faultLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for faultLowLimit")
	}
	_faultLowLimit, _faultLowLimitErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _faultLowLimitErr != nil {
		return nil, errors.Wrap(_faultLowLimitErr, "Error parsing 'faultLowLimit' field of BACnetConstructedDataAnalogValueFaultLowLimit")
	}
	faultLowLimit := _faultLowLimit.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("faultLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for faultLowLimit")
	}

	// Virtual field
	_actualValue := faultLowLimit
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAnalogValueFaultLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAnalogValueFaultLowLimit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAnalogValueFaultLowLimit{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FaultLowLimit: faultLowLimit,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAnalogValueFaultLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAnalogValueFaultLowLimit")
		}

		// Simple Field (faultLowLimit)
		if pushErr := writeBuffer.PushContext("faultLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for faultLowLimit")
		}
		_faultLowLimitErr := writeBuffer.WriteSerializable(ctx, m.GetFaultLowLimit())
		if popErr := writeBuffer.PopContext("faultLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for faultLowLimit")
		}
		if _faultLowLimitErr != nil {
			return errors.Wrap(_faultLowLimitErr, "Error serializing 'faultLowLimit' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAnalogValueFaultLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAnalogValueFaultLowLimit")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) isBACnetConstructedDataAnalogValueFaultLowLimit() bool {
	return true
}

func (m *_BACnetConstructedDataAnalogValueFaultLowLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
