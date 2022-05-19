/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// APDUAbort is the data-structure of this message
type APDUAbort struct {
	*APDU
	Server           bool
	OriginalInvokeId uint8
	AbortReason      *AbortReasonTagged

	// Arguments.
	ApduLength uint16
}

// IAPDUAbort is the corresponding interface of APDUAbort
type IAPDUAbort interface {
	IAPDU
	// GetServer returns Server (property field)
	GetServer() bool
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetAbortReason returns AbortReason (property field)
	GetAbortReason() *AbortReasonTagged
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *APDUAbort) GetApduType() uint8 {
	return 0x7
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *APDUAbort) InitializeParent(parent *APDU) {}

func (m *APDUAbort) GetParent() *APDU {
	return m.APDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *APDUAbort) GetServer() bool {
	return m.Server
}

func (m *APDUAbort) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *APDUAbort) GetAbortReason() *AbortReasonTagged {
	return m.AbortReason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAPDUAbort factory function for APDUAbort
func NewAPDUAbort(server bool, originalInvokeId uint8, abortReason *AbortReasonTagged, apduLength uint16) *APDUAbort {
	_result := &APDUAbort{
		Server:           server,
		OriginalInvokeId: originalInvokeId,
		AbortReason:      abortReason,
		APDU:             NewAPDU(apduLength),
	}
	_result.Child = _result
	return _result
}

func CastAPDUAbort(structType interface{}) *APDUAbort {
	if casted, ok := structType.(APDUAbort); ok {
		return &casted
	}
	if casted, ok := structType.(*APDUAbort); ok {
		return casted
	}
	if casted, ok := structType.(APDU); ok {
		return CastAPDUAbort(casted.Child)
	}
	if casted, ok := structType.(*APDU); ok {
		return CastAPDUAbort(casted.Child)
	}
	return nil
}

func (m *APDUAbort) GetTypeName() string {
	return "APDUAbort"
}

func (m *APDUAbort) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *APDUAbort) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 3

	// Simple field (server)
	lengthInBits += 1

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (abortReason)
	lengthInBits += m.AbortReason.GetLengthInBits()

	return lengthInBits
}

func (m *APDUAbort) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func APDUAbortParse(readBuffer utils.ReadBuffer, apduLength uint16) (*APDUAbort, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUAbort"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 3)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (server)
	_server, _serverErr := readBuffer.ReadBit("server")
	if _serverErr != nil {
		return nil, errors.Wrap(_serverErr, "Error parsing 'server' field")
	}
	server := _server

	// Simple Field (originalInvokeId)
	_originalInvokeId, _originalInvokeIdErr := readBuffer.ReadUint8("originalInvokeId", 8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field")
	}
	originalInvokeId := _originalInvokeId

	// Simple Field (abortReason)
	if pullErr := readBuffer.PullContext("abortReason"); pullErr != nil {
		return nil, pullErr
	}
	_abortReason, _abortReasonErr := AbortReasonTaggedParse(readBuffer, uint32(uint32(1)))
	if _abortReasonErr != nil {
		return nil, errors.Wrap(_abortReasonErr, "Error parsing 'abortReason' field")
	}
	abortReason := CastAbortReasonTagged(_abortReason)
	if closeErr := readBuffer.CloseContext("abortReason"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("APDUAbort"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &APDUAbort{
		Server:           server,
		OriginalInvokeId: originalInvokeId,
		AbortReason:      CastAbortReasonTagged(abortReason),
		APDU:             &APDU{},
	}
	_child.APDU.Child = _child
	return _child, nil
}

func (m *APDUAbort) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUAbort"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 3, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (server)
		server := bool(m.Server)
		_serverErr := writeBuffer.WriteBit("server", (server))
		if _serverErr != nil {
			return errors.Wrap(_serverErr, "Error serializing 'server' field")
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.OriginalInvokeId)
		_originalInvokeIdErr := writeBuffer.WriteUint8("originalInvokeId", 8, (originalInvokeId))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Simple Field (abortReason)
		if pushErr := writeBuffer.PushContext("abortReason"); pushErr != nil {
			return pushErr
		}
		_abortReasonErr := m.AbortReason.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("abortReason"); popErr != nil {
			return popErr
		}
		if _abortReasonErr != nil {
			return errors.Wrap(_abortReasonErr, "Error serializing 'abortReason' field")
		}

		if popErr := writeBuffer.PopContext("APDUAbort"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *APDUAbort) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
