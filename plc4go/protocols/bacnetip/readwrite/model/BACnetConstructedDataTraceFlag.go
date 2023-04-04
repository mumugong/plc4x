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

// BACnetConstructedDataTraceFlag is the corresponding interface of BACnetConstructedDataTraceFlag
type BACnetConstructedDataTraceFlag interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTraceFlag returns TraceFlag (property field)
	GetTraceFlag() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataTraceFlagExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTraceFlag.
// This is useful for switch cases.
type BACnetConstructedDataTraceFlagExactly interface {
	BACnetConstructedDataTraceFlag
	isBACnetConstructedDataTraceFlag() bool
}

// _BACnetConstructedDataTraceFlag is the data-structure of this message
type _BACnetConstructedDataTraceFlag struct {
	*_BACnetConstructedData
	TraceFlag BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTraceFlag) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTraceFlag) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TRACE_FLAG
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTraceFlag) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTraceFlag) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTraceFlag) GetTraceFlag() BACnetApplicationTagBoolean {
	return m.TraceFlag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTraceFlag) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetTraceFlag())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTraceFlag factory function for _BACnetConstructedDataTraceFlag
func NewBACnetConstructedDataTraceFlag(traceFlag BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTraceFlag {
	_result := &_BACnetConstructedDataTraceFlag{
		TraceFlag:              traceFlag,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTraceFlag(structType interface{}) BACnetConstructedDataTraceFlag {
	if casted, ok := structType.(BACnetConstructedDataTraceFlag); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTraceFlag); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTraceFlag) GetTypeName() string {
	return "BACnetConstructedDataTraceFlag"
}

func (m *_BACnetConstructedDataTraceFlag) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (traceFlag)
	lengthInBits += m.TraceFlag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTraceFlag) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataTraceFlagParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTraceFlag, error) {
	return BACnetConstructedDataTraceFlagParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataTraceFlagParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTraceFlag, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTraceFlag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTraceFlag")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (traceFlag)
	if pullErr := readBuffer.PullContext("traceFlag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for traceFlag")
	}
	_traceFlag, _traceFlagErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _traceFlagErr != nil {
		return nil, errors.Wrap(_traceFlagErr, "Error parsing 'traceFlag' field of BACnetConstructedDataTraceFlag")
	}
	traceFlag := _traceFlag.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("traceFlag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for traceFlag")
	}

	// Virtual field
	_actualValue := traceFlag
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTraceFlag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTraceFlag")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTraceFlag{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		TraceFlag: traceFlag,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTraceFlag) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTraceFlag) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTraceFlag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTraceFlag")
		}

		// Simple Field (traceFlag)
		if pushErr := writeBuffer.PushContext("traceFlag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for traceFlag")
		}
		_traceFlagErr := writeBuffer.WriteSerializable(ctx, m.GetTraceFlag())
		if popErr := writeBuffer.PopContext("traceFlag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for traceFlag")
		}
		if _traceFlagErr != nil {
			return errors.Wrap(_traceFlagErr, "Error serializing 'traceFlag' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTraceFlag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTraceFlag")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTraceFlag) isBACnetConstructedDataTraceFlag() bool {
	return true
}

func (m *_BACnetConstructedDataTraceFlag) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
