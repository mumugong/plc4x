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

// BACnetConstructedDataProgramState is the corresponding interface of BACnetConstructedDataProgramState
type BACnetConstructedDataProgramState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetProgramState returns ProgramState (property field)
	GetProgramState() BACnetProgramStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetProgramStateTagged
}

// BACnetConstructedDataProgramStateExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataProgramState.
// This is useful for switch cases.
type BACnetConstructedDataProgramStateExactly interface {
	BACnetConstructedDataProgramState
	isBACnetConstructedDataProgramState() bool
}

// _BACnetConstructedDataProgramState is the data-structure of this message
type _BACnetConstructedDataProgramState struct {
	*_BACnetConstructedData
	ProgramState BACnetProgramStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProgramState) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProgramState) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROGRAM_STATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProgramState) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataProgramState) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProgramState) GetProgramState() BACnetProgramStateTagged {
	return m.ProgramState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProgramState) GetActualValue() BACnetProgramStateTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetProgramStateTagged(m.GetProgramState())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProgramState factory function for _BACnetConstructedDataProgramState
func NewBACnetConstructedDataProgramState(programState BACnetProgramStateTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProgramState {
	_result := &_BACnetConstructedDataProgramState{
		ProgramState:           programState,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProgramState(structType interface{}) BACnetConstructedDataProgramState {
	if casted, ok := structType.(BACnetConstructedDataProgramState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProgramState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProgramState) GetTypeName() string {
	return "BACnetConstructedDataProgramState"
}

func (m *_BACnetConstructedDataProgramState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (programState)
	lengthInBits += m.ProgramState.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProgramState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataProgramStateParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProgramState, error) {
	return BACnetConstructedDataProgramStateParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataProgramStateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProgramState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProgramState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProgramState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (programState)
	if pullErr := readBuffer.PullContext("programState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for programState")
	}
	_programState, _programStateErr := BACnetProgramStateTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _programStateErr != nil {
		return nil, errors.Wrap(_programStateErr, "Error parsing 'programState' field of BACnetConstructedDataProgramState")
	}
	programState := _programState.(BACnetProgramStateTagged)
	if closeErr := readBuffer.CloseContext("programState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for programState")
	}

	// Virtual field
	_actualValue := programState
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProgramState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProgramState")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataProgramState{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ProgramState: programState,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataProgramState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataProgramState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProgramState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProgramState")
		}

		// Simple Field (programState)
		if pushErr := writeBuffer.PushContext("programState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for programState")
		}
		_programStateErr := writeBuffer.WriteSerializable(ctx, m.GetProgramState())
		if popErr := writeBuffer.PopContext("programState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for programState")
		}
		if _programStateErr != nil {
			return errors.Wrap(_programStateErr, "Error serializing 'programState' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProgramState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProgramState")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProgramState) isBACnetConstructedDataProgramState() bool {
	return true
}

func (m *_BACnetConstructedDataProgramState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
