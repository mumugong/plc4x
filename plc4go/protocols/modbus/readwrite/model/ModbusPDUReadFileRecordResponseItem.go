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

// ModbusPDUReadFileRecordResponseItem is the corresponding interface of ModbusPDUReadFileRecordResponseItem
type ModbusPDUReadFileRecordResponseItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetReferenceType returns ReferenceType (property field)
	GetReferenceType() uint8
	// GetData returns Data (property field)
	GetData() []byte
}

// ModbusPDUReadFileRecordResponseItemExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadFileRecordResponseItem.
// This is useful for switch cases.
type ModbusPDUReadFileRecordResponseItemExactly interface {
	ModbusPDUReadFileRecordResponseItem
	isModbusPDUReadFileRecordResponseItem() bool
}

// _ModbusPDUReadFileRecordResponseItem is the data-structure of this message
type _ModbusPDUReadFileRecordResponseItem struct {
	ReferenceType uint8
	Data          []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadFileRecordResponseItem) GetReferenceType() uint8 {
	return m.ReferenceType
}

func (m *_ModbusPDUReadFileRecordResponseItem) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadFileRecordResponseItem factory function for _ModbusPDUReadFileRecordResponseItem
func NewModbusPDUReadFileRecordResponseItem(referenceType uint8, data []byte) *_ModbusPDUReadFileRecordResponseItem {
	return &_ModbusPDUReadFileRecordResponseItem{ReferenceType: referenceType, Data: data}
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadFileRecordResponseItem(structType interface{}) ModbusPDUReadFileRecordResponseItem {
	if casted, ok := structType.(ModbusPDUReadFileRecordResponseItem); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadFileRecordResponseItem); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadFileRecordResponseItem) GetTypeName() string {
	return "ModbusPDUReadFileRecordResponseItem"
}

func (m *_ModbusPDUReadFileRecordResponseItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (dataLength)
	lengthInBits += 8

	// Simple field (referenceType)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ModbusPDUReadFileRecordResponseItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUReadFileRecordResponseItemParse(theBytes []byte) (ModbusPDUReadFileRecordResponseItem, error) {
	return ModbusPDUReadFileRecordResponseItemParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func ModbusPDUReadFileRecordResponseItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusPDUReadFileRecordResponseItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadFileRecordResponseItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadFileRecordResponseItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (dataLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	dataLength, _dataLengthErr := readBuffer.ReadUint8("dataLength", 8)
	_ = dataLength
	if _dataLengthErr != nil {
		return nil, errors.Wrap(_dataLengthErr, "Error parsing 'dataLength' field of ModbusPDUReadFileRecordResponseItem")
	}

	// Simple Field (referenceType)
	_referenceType, _referenceTypeErr := readBuffer.ReadUint8("referenceType", 8)
	if _referenceTypeErr != nil {
		return nil, errors.Wrap(_referenceTypeErr, "Error parsing 'referenceType' field of ModbusPDUReadFileRecordResponseItem")
	}
	referenceType := _referenceType
	// Byte Array field (data)
	numberOfBytesdata := int(uint16(dataLength) - uint16(uint16(1)))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of ModbusPDUReadFileRecordResponseItem")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadFileRecordResponseItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadFileRecordResponseItem")
	}

	// Create the instance
	return &_ModbusPDUReadFileRecordResponseItem{
		ReferenceType: referenceType,
		Data:          data,
	}, nil
}

func (m *_ModbusPDUReadFileRecordResponseItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadFileRecordResponseItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ModbusPDUReadFileRecordResponseItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadFileRecordResponseItem")
	}

	// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	dataLength := uint8(uint8(uint8(len(m.GetData()))) + uint8(uint8(1)))
	_dataLengthErr := writeBuffer.WriteUint8("dataLength", 8, (dataLength))
	if _dataLengthErr != nil {
		return errors.Wrap(_dataLengthErr, "Error serializing 'dataLength' field")
	}

	// Simple Field (referenceType)
	referenceType := uint8(m.GetReferenceType())
	_referenceTypeErr := writeBuffer.WriteUint8("referenceType", 8, (referenceType))
	if _referenceTypeErr != nil {
		return errors.Wrap(_referenceTypeErr, "Error serializing 'referenceType' field")
	}

	// Array Field (data)
	// Byte Array field (data)
	if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

	if popErr := writeBuffer.PopContext("ModbusPDUReadFileRecordResponseItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ModbusPDUReadFileRecordResponseItem")
	}
	return nil
}

func (m *_ModbusPDUReadFileRecordResponseItem) isModbusPDUReadFileRecordResponseItem() bool {
	return true
}

func (m *_ModbusPDUReadFileRecordResponseItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
