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

// BACnetConstructedDataUnits is the corresponding interface of BACnetConstructedDataUnits
type BACnetConstructedDataUnits interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetUnits returns Units (property field)
	GetUnits() BACnetEngineeringUnitsTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEngineeringUnitsTagged
}

// BACnetConstructedDataUnitsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataUnits.
// This is useful for switch cases.
type BACnetConstructedDataUnitsExactly interface {
	BACnetConstructedDataUnits
	isBACnetConstructedDataUnits() bool
}

// _BACnetConstructedDataUnits is the data-structure of this message
type _BACnetConstructedDataUnits struct {
	*_BACnetConstructedData
	Units BACnetEngineeringUnitsTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUnits) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUnits) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_UNITS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUnits) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataUnits) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUnits) GetUnits() BACnetEngineeringUnitsTagged {
	return m.Units
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataUnits) GetActualValue() BACnetEngineeringUnitsTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetEngineeringUnitsTagged(m.GetUnits())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataUnits factory function for _BACnetConstructedDataUnits
func NewBACnetConstructedDataUnits(units BACnetEngineeringUnitsTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUnits {
	_result := &_BACnetConstructedDataUnits{
		Units:                  units,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUnits(structType interface{}) BACnetConstructedDataUnits {
	if casted, ok := structType.(BACnetConstructedDataUnits); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUnits); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUnits) GetTypeName() string {
	return "BACnetConstructedDataUnits"
}

func (m *_BACnetConstructedDataUnits) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (units)
	lengthInBits += m.Units.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataUnits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataUnitsParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataUnits, error) {
	return BACnetConstructedDataUnitsParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataUnitsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataUnits, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUnits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUnits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (units)
	if pullErr := readBuffer.PullContext("units"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for units")
	}
	_units, _unitsErr := BACnetEngineeringUnitsTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _unitsErr != nil {
		return nil, errors.Wrap(_unitsErr, "Error parsing 'units' field of BACnetConstructedDataUnits")
	}
	units := _units.(BACnetEngineeringUnitsTagged)
	if closeErr := readBuffer.CloseContext("units"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for units")
	}

	// Virtual field
	_actualValue := units
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUnits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUnits")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataUnits{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Units: units,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataUnits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataUnits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUnits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUnits")
		}

		// Simple Field (units)
		if pushErr := writeBuffer.PushContext("units"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for units")
		}
		_unitsErr := writeBuffer.WriteSerializable(ctx, m.GetUnits())
		if popErr := writeBuffer.PopContext("units"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for units")
		}
		if _unitsErr != nil {
			return errors.Wrap(_unitsErr, "Error serializing 'units' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUnits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUnits")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUnits) isBACnetConstructedDataUnits() bool {
	return true
}

func (m *_BACnetConstructedDataUnits) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
