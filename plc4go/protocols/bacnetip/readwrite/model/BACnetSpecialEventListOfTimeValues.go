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
	spiContext "github.com/apache/plc4x/plc4go/spi/context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetSpecialEventListOfTimeValues is the corresponding interface of BACnetSpecialEventListOfTimeValues
type BACnetSpecialEventListOfTimeValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfTimeValues returns ListOfTimeValues (property field)
	GetListOfTimeValues() []BACnetTimeValue
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetSpecialEventListOfTimeValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetSpecialEventListOfTimeValues.
// This is useful for switch cases.
type BACnetSpecialEventListOfTimeValuesExactly interface {
	BACnetSpecialEventListOfTimeValues
	isBACnetSpecialEventListOfTimeValues() bool
}

// _BACnetSpecialEventListOfTimeValues is the data-structure of this message
type _BACnetSpecialEventListOfTimeValues struct {
	OpeningTag       BACnetOpeningTag
	ListOfTimeValues []BACnetTimeValue
	ClosingTag       BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSpecialEventListOfTimeValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetSpecialEventListOfTimeValues) GetListOfTimeValues() []BACnetTimeValue {
	return m.ListOfTimeValues
}

func (m *_BACnetSpecialEventListOfTimeValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSpecialEventListOfTimeValues factory function for _BACnetSpecialEventListOfTimeValues
func NewBACnetSpecialEventListOfTimeValues(openingTag BACnetOpeningTag, listOfTimeValues []BACnetTimeValue, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetSpecialEventListOfTimeValues {
	return &_BACnetSpecialEventListOfTimeValues{OpeningTag: openingTag, ListOfTimeValues: listOfTimeValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetSpecialEventListOfTimeValues(structType interface{}) BACnetSpecialEventListOfTimeValues {
	if casted, ok := structType.(BACnetSpecialEventListOfTimeValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSpecialEventListOfTimeValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSpecialEventListOfTimeValues) GetTypeName() string {
	return "BACnetSpecialEventListOfTimeValues"
}

func (m *_BACnetSpecialEventListOfTimeValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfTimeValues) > 0 {
		for _, element := range m.ListOfTimeValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetSpecialEventListOfTimeValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetSpecialEventListOfTimeValuesParse(theBytes []byte, tagNumber uint8) (BACnetSpecialEventListOfTimeValues, error) {
	return BACnetSpecialEventListOfTimeValuesParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetSpecialEventListOfTimeValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetSpecialEventListOfTimeValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSpecialEventListOfTimeValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSpecialEventListOfTimeValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetSpecialEventListOfTimeValues")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (listOfTimeValues)
	if pullErr := readBuffer.PullContext("listOfTimeValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfTimeValues")
	}
	// Terminated array
	var listOfTimeValues []BACnetTimeValue
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetTimeValueParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfTimeValues' field of BACnetSpecialEventListOfTimeValues")
			}
			listOfTimeValues = append(listOfTimeValues, _item.(BACnetTimeValue))
		}
	}
	if closeErr := readBuffer.CloseContext("listOfTimeValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfTimeValues")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetSpecialEventListOfTimeValues")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetSpecialEventListOfTimeValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSpecialEventListOfTimeValues")
	}

	// Create the instance
	return &_BACnetSpecialEventListOfTimeValues{
		TagNumber:        tagNumber,
		OpeningTag:       openingTag,
		ListOfTimeValues: listOfTimeValues,
		ClosingTag:       closingTag,
	}, nil
}

func (m *_BACnetSpecialEventListOfTimeValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetSpecialEventListOfTimeValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetSpecialEventListOfTimeValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetSpecialEventListOfTimeValues")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (listOfTimeValues)
	if pushErr := writeBuffer.PushContext("listOfTimeValues", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfTimeValues")
	}
	for _curItem, _element := range m.GetListOfTimeValues() {
		_ = _curItem
		arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetListOfTimeValues()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listOfTimeValues' field")
		}
	}
	if popErr := writeBuffer.PopContext("listOfTimeValues", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfTimeValues")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetSpecialEventListOfTimeValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetSpecialEventListOfTimeValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetSpecialEventListOfTimeValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetSpecialEventListOfTimeValues) isBACnetSpecialEventListOfTimeValues() bool {
	return true
}

func (m *_BACnetSpecialEventListOfTimeValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
