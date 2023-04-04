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

// Constant values.
const SendUnitData_INTERFACEHANDLE uint32 = 0x00000000

// SendUnitData is the corresponding interface of SendUnitData
type SendUnitData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	EipPacket
	// GetTimeout returns Timeout (property field)
	GetTimeout() uint16
	// GetTypeIds returns TypeIds (property field)
	GetTypeIds() []TypeId
}

// SendUnitDataExactly can be used when we want exactly this type and not a type which fulfills SendUnitData.
// This is useful for switch cases.
type SendUnitDataExactly interface {
	SendUnitData
	isSendUnitData() bool
}

// _SendUnitData is the data-structure of this message
type _SendUnitData struct {
	*_EipPacket
	Timeout uint16
	TypeIds []TypeId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SendUnitData) GetCommand() uint16 {
	return 0x0070
}

func (m *_SendUnitData) GetResponse() bool {
	return false
}

func (m *_SendUnitData) GetPacketLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SendUnitData) InitializeParent(parent EipPacket, sessionHandle uint32, status uint32, senderContext []byte, options uint32) {
	m.SessionHandle = sessionHandle
	m.Status = status
	m.SenderContext = senderContext
	m.Options = options
}

func (m *_SendUnitData) GetParent() EipPacket {
	return m._EipPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SendUnitData) GetTimeout() uint16 {
	return m.Timeout
}

func (m *_SendUnitData) GetTypeIds() []TypeId {
	return m.TypeIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_SendUnitData) GetInterfaceHandle() uint32 {
	return SendUnitData_INTERFACEHANDLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSendUnitData factory function for _SendUnitData
func NewSendUnitData(timeout uint16, typeIds []TypeId, sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_SendUnitData {
	_result := &_SendUnitData{
		Timeout:    timeout,
		TypeIds:    typeIds,
		_EipPacket: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	_result._EipPacket._EipPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSendUnitData(structType interface{}) SendUnitData {
	if casted, ok := structType.(SendUnitData); ok {
		return casted
	}
	if casted, ok := structType.(*SendUnitData); ok {
		return *casted
	}
	return nil
}

func (m *_SendUnitData) GetTypeName() string {
	return "SendUnitData"
}

func (m *_SendUnitData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Const Field (interfaceHandle)
	lengthInBits += 32

	// Simple field (timeout)
	lengthInBits += 16

	// Implicit Field (typeIdCount)
	lengthInBits += 16

	// Array field
	if len(m.TypeIds) > 0 {
		for _curItem, element := range m.TypeIds {
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.TypeIds), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_SendUnitData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SendUnitDataParse(theBytes []byte, response bool) (SendUnitData, error) {
	return SendUnitDataParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), response)
}

func SendUnitDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (SendUnitData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SendUnitData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SendUnitData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (interfaceHandle)
	interfaceHandle, _interfaceHandleErr := readBuffer.ReadUint32("interfaceHandle", 32)
	if _interfaceHandleErr != nil {
		return nil, errors.Wrap(_interfaceHandleErr, "Error parsing 'interfaceHandle' field of SendUnitData")
	}
	if interfaceHandle != SendUnitData_INTERFACEHANDLE {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", SendUnitData_INTERFACEHANDLE) + " but got " + fmt.Sprintf("%d", interfaceHandle))
	}

	// Simple Field (timeout)
	_timeout, _timeoutErr := readBuffer.ReadUint16("timeout", 16)
	if _timeoutErr != nil {
		return nil, errors.Wrap(_timeoutErr, "Error parsing 'timeout' field of SendUnitData")
	}
	timeout := _timeout

	// Implicit Field (typeIdCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	typeIdCount, _typeIdCountErr := readBuffer.ReadUint16("typeIdCount", 16)
	_ = typeIdCount
	if _typeIdCountErr != nil {
		return nil, errors.Wrap(_typeIdCountErr, "Error parsing 'typeIdCount' field of SendUnitData")
	}

	// Array field (typeIds)
	if pullErr := readBuffer.PullContext("typeIds", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for typeIds")
	}
	// Count array
	typeIds := make([]TypeId, typeIdCount)
	// This happens when the size is set conditional to 0
	if len(typeIds) == 0 {
		typeIds = nil
	}
	{
		_numItems := uint16(typeIdCount)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := spiContext.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := TypeIdParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'typeIds' field of SendUnitData")
			}
			typeIds[_curItem] = _item.(TypeId)
		}
	}
	if closeErr := readBuffer.CloseContext("typeIds", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for typeIds")
	}

	if closeErr := readBuffer.CloseContext("SendUnitData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SendUnitData")
	}

	// Create a partially initialized instance
	_child := &_SendUnitData{
		_EipPacket: &_EipPacket{},
		Timeout:    timeout,
		TypeIds:    typeIds,
	}
	_child._EipPacket._EipPacketChildRequirements = _child
	return _child, nil
}

func (m *_SendUnitData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SendUnitData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SendUnitData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SendUnitData")
		}

		// Const Field (interfaceHandle)
		_interfaceHandleErr := writeBuffer.WriteUint32("interfaceHandle", 32, 0x00000000)
		if _interfaceHandleErr != nil {
			return errors.Wrap(_interfaceHandleErr, "Error serializing 'interfaceHandle' field")
		}

		// Simple Field (timeout)
		timeout := uint16(m.GetTimeout())
		_timeoutErr := writeBuffer.WriteUint16("timeout", 16, (timeout))
		if _timeoutErr != nil {
			return errors.Wrap(_timeoutErr, "Error serializing 'timeout' field")
		}

		// Implicit Field (typeIdCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		typeIdCount := uint16(uint16(len(m.GetTypeIds())))
		_typeIdCountErr := writeBuffer.WriteUint16("typeIdCount", 16, (typeIdCount))
		if _typeIdCountErr != nil {
			return errors.Wrap(_typeIdCountErr, "Error serializing 'typeIdCount' field")
		}

		// Array Field (typeIds)
		if pushErr := writeBuffer.PushContext("typeIds", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for typeIds")
		}
		for _curItem, _element := range m.GetTypeIds() {
			_ = _curItem
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetTypeIds()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'typeIds' field")
			}
		}
		if popErr := writeBuffer.PopContext("typeIds", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for typeIds")
		}

		if popErr := writeBuffer.PopContext("SendUnitData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SendUnitData")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SendUnitData) isSendUnitData() bool {
	return true
}

func (m *_SendUnitData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
