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

// ModbusPDUMaskWriteHoldingRegisterRequest is the corresponding interface of ModbusPDUMaskWriteHoldingRegisterRequest
type ModbusPDUMaskWriteHoldingRegisterRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetReferenceAddress returns ReferenceAddress (property field)
	GetReferenceAddress() uint16
	// GetAndMask returns AndMask (property field)
	GetAndMask() uint16
	// GetOrMask returns OrMask (property field)
	GetOrMask() uint16
}

// ModbusPDUMaskWriteHoldingRegisterRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUMaskWriteHoldingRegisterRequest.
// This is useful for switch cases.
type ModbusPDUMaskWriteHoldingRegisterRequestExactly interface {
	ModbusPDUMaskWriteHoldingRegisterRequest
	isModbusPDUMaskWriteHoldingRegisterRequest() bool
}

// _ModbusPDUMaskWriteHoldingRegisterRequest is the data-structure of this message
type _ModbusPDUMaskWriteHoldingRegisterRequest struct {
	*_ModbusPDU
	ReferenceAddress uint16
	AndMask          uint16
	OrMask           uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetFunctionFlag() uint8 {
	return 0x16
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetReferenceAddress() uint16 {
	return m.ReferenceAddress
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetAndMask() uint16 {
	return m.AndMask
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetOrMask() uint16 {
	return m.OrMask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUMaskWriteHoldingRegisterRequest factory function for _ModbusPDUMaskWriteHoldingRegisterRequest
func NewModbusPDUMaskWriteHoldingRegisterRequest(referenceAddress uint16, andMask uint16, orMask uint16) *_ModbusPDUMaskWriteHoldingRegisterRequest {
	_result := &_ModbusPDUMaskWriteHoldingRegisterRequest{
		ReferenceAddress: referenceAddress,
		AndMask:          andMask,
		OrMask:           orMask,
		_ModbusPDU:       NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUMaskWriteHoldingRegisterRequest(structType interface{}) ModbusPDUMaskWriteHoldingRegisterRequest {
	if casted, ok := structType.(ModbusPDUMaskWriteHoldingRegisterRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUMaskWriteHoldingRegisterRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetTypeName() string {
	return "ModbusPDUMaskWriteHoldingRegisterRequest"
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (referenceAddress)
	lengthInBits += 16

	// Simple field (andMask)
	lengthInBits += 16

	// Simple field (orMask)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUMaskWriteHoldingRegisterRequestParse(theBytes []byte, response bool) (ModbusPDUMaskWriteHoldingRegisterRequest, error) {
	return ModbusPDUMaskWriteHoldingRegisterRequestParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUMaskWriteHoldingRegisterRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUMaskWriteHoldingRegisterRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUMaskWriteHoldingRegisterRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUMaskWriteHoldingRegisterRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (referenceAddress)
	_referenceAddress, _referenceAddressErr := readBuffer.ReadUint16("referenceAddress", 16)
	if _referenceAddressErr != nil {
		return nil, errors.Wrap(_referenceAddressErr, "Error parsing 'referenceAddress' field of ModbusPDUMaskWriteHoldingRegisterRequest")
	}
	referenceAddress := _referenceAddress

	// Simple Field (andMask)
	_andMask, _andMaskErr := readBuffer.ReadUint16("andMask", 16)
	if _andMaskErr != nil {
		return nil, errors.Wrap(_andMaskErr, "Error parsing 'andMask' field of ModbusPDUMaskWriteHoldingRegisterRequest")
	}
	andMask := _andMask

	// Simple Field (orMask)
	_orMask, _orMaskErr := readBuffer.ReadUint16("orMask", 16)
	if _orMaskErr != nil {
		return nil, errors.Wrap(_orMaskErr, "Error parsing 'orMask' field of ModbusPDUMaskWriteHoldingRegisterRequest")
	}
	orMask := _orMask

	if closeErr := readBuffer.CloseContext("ModbusPDUMaskWriteHoldingRegisterRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUMaskWriteHoldingRegisterRequest")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUMaskWriteHoldingRegisterRequest{
		_ModbusPDU:       &_ModbusPDU{},
		ReferenceAddress: referenceAddress,
		AndMask:          andMask,
		OrMask:           orMask,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUMaskWriteHoldingRegisterRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUMaskWriteHoldingRegisterRequest")
		}

		// Simple Field (referenceAddress)
		referenceAddress := uint16(m.GetReferenceAddress())
		_referenceAddressErr := writeBuffer.WriteUint16("referenceAddress", 16, (referenceAddress))
		if _referenceAddressErr != nil {
			return errors.Wrap(_referenceAddressErr, "Error serializing 'referenceAddress' field")
		}

		// Simple Field (andMask)
		andMask := uint16(m.GetAndMask())
		_andMaskErr := writeBuffer.WriteUint16("andMask", 16, (andMask))
		if _andMaskErr != nil {
			return errors.Wrap(_andMaskErr, "Error serializing 'andMask' field")
		}

		// Simple Field (orMask)
		orMask := uint16(m.GetOrMask())
		_orMaskErr := writeBuffer.WriteUint16("orMask", 16, (orMask))
		if _orMaskErr != nil {
			return errors.Wrap(_orMaskErr, "Error serializing 'orMask' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUMaskWriteHoldingRegisterRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUMaskWriteHoldingRegisterRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) isModbusPDUMaskWriteHoldingRegisterRequest() bool {
	return true
}

func (m *_ModbusPDUMaskWriteHoldingRegisterRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
