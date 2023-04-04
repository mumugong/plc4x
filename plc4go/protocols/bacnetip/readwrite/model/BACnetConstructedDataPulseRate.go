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

// BACnetConstructedDataPulseRate is the corresponding interface of BACnetConstructedDataPulseRate
type BACnetConstructedDataPulseRate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPulseRate returns PulseRate (property field)
	GetPulseRate() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataPulseRateExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataPulseRate.
// This is useful for switch cases.
type BACnetConstructedDataPulseRateExactly interface {
	BACnetConstructedDataPulseRate
	isBACnetConstructedDataPulseRate() bool
}

// _BACnetConstructedDataPulseRate is the data-structure of this message
type _BACnetConstructedDataPulseRate struct {
	*_BACnetConstructedData
	PulseRate BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPulseRate) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataPulseRate) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PULSE_RATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPulseRate) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataPulseRate) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPulseRate) GetPulseRate() BACnetApplicationTagUnsignedInteger {
	return m.PulseRate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPulseRate) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetPulseRate())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPulseRate factory function for _BACnetConstructedDataPulseRate
func NewBACnetConstructedDataPulseRate(pulseRate BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPulseRate {
	_result := &_BACnetConstructedDataPulseRate{
		PulseRate:              pulseRate,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPulseRate(structType interface{}) BACnetConstructedDataPulseRate {
	if casted, ok := structType.(BACnetConstructedDataPulseRate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPulseRate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPulseRate) GetTypeName() string {
	return "BACnetConstructedDataPulseRate"
}

func (m *_BACnetConstructedDataPulseRate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (pulseRate)
	lengthInBits += m.PulseRate.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPulseRate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataPulseRateParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPulseRate, error) {
	return BACnetConstructedDataPulseRateParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataPulseRateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPulseRate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPulseRate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPulseRate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (pulseRate)
	if pullErr := readBuffer.PullContext("pulseRate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for pulseRate")
	}
	_pulseRate, _pulseRateErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _pulseRateErr != nil {
		return nil, errors.Wrap(_pulseRateErr, "Error parsing 'pulseRate' field of BACnetConstructedDataPulseRate")
	}
	pulseRate := _pulseRate.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("pulseRate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for pulseRate")
	}

	// Virtual field
	_actualValue := pulseRate
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPulseRate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPulseRate")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataPulseRate{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		PulseRate: pulseRate,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataPulseRate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPulseRate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPulseRate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPulseRate")
		}

		// Simple Field (pulseRate)
		if pushErr := writeBuffer.PushContext("pulseRate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for pulseRate")
		}
		_pulseRateErr := writeBuffer.WriteSerializable(ctx, m.GetPulseRate())
		if popErr := writeBuffer.PopContext("pulseRate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for pulseRate")
		}
		if _pulseRateErr != nil {
			return errors.Wrap(_pulseRateErr, "Error serializing 'pulseRate' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPulseRate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPulseRate")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPulseRate) isBACnetConstructedDataPulseRate() bool {
	return true
}

func (m *_BACnetConstructedDataPulseRate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
