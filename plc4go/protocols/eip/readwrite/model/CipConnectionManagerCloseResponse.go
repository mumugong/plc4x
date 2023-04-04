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

// CipConnectionManagerCloseResponse is the corresponding interface of CipConnectionManagerCloseResponse
type CipConnectionManagerCloseResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetAdditionalStatusWords returns AdditionalStatusWords (property field)
	GetAdditionalStatusWords() uint8
	// GetConnectionSerialNumber returns ConnectionSerialNumber (property field)
	GetConnectionSerialNumber() uint16
	// GetOriginatorVendorId returns OriginatorVendorId (property field)
	GetOriginatorVendorId() uint16
	// GetOriginatorSerialNumber returns OriginatorSerialNumber (property field)
	GetOriginatorSerialNumber() uint32
	// GetApplicationReplySize returns ApplicationReplySize (property field)
	GetApplicationReplySize() uint8
}

// CipConnectionManagerCloseResponseExactly can be used when we want exactly this type and not a type which fulfills CipConnectionManagerCloseResponse.
// This is useful for switch cases.
type CipConnectionManagerCloseResponseExactly interface {
	CipConnectionManagerCloseResponse
	isCipConnectionManagerCloseResponse() bool
}

// _CipConnectionManagerCloseResponse is the data-structure of this message
type _CipConnectionManagerCloseResponse struct {
	*_CipService
	Status                 uint8
	AdditionalStatusWords  uint8
	ConnectionSerialNumber uint16
	OriginatorVendorId     uint16
	OriginatorSerialNumber uint32
	ApplicationReplySize   uint8
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipConnectionManagerCloseResponse) GetService() uint8 {
	return 0x4E
}

func (m *_CipConnectionManagerCloseResponse) GetResponse() bool {
	return bool(true)
}

func (m *_CipConnectionManagerCloseResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipConnectionManagerCloseResponse) InitializeParent(parent CipService) {}

func (m *_CipConnectionManagerCloseResponse) GetParent() CipService {
	return m._CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipConnectionManagerCloseResponse) GetStatus() uint8 {
	return m.Status
}

func (m *_CipConnectionManagerCloseResponse) GetAdditionalStatusWords() uint8 {
	return m.AdditionalStatusWords
}

func (m *_CipConnectionManagerCloseResponse) GetConnectionSerialNumber() uint16 {
	return m.ConnectionSerialNumber
}

func (m *_CipConnectionManagerCloseResponse) GetOriginatorVendorId() uint16 {
	return m.OriginatorVendorId
}

func (m *_CipConnectionManagerCloseResponse) GetOriginatorSerialNumber() uint32 {
	return m.OriginatorSerialNumber
}

func (m *_CipConnectionManagerCloseResponse) GetApplicationReplySize() uint8 {
	return m.ApplicationReplySize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipConnectionManagerCloseResponse factory function for _CipConnectionManagerCloseResponse
func NewCipConnectionManagerCloseResponse(status uint8, additionalStatusWords uint8, connectionSerialNumber uint16, originatorVendorId uint16, originatorSerialNumber uint32, applicationReplySize uint8, serviceLen uint16) *_CipConnectionManagerCloseResponse {
	_result := &_CipConnectionManagerCloseResponse{
		Status:                 status,
		AdditionalStatusWords:  additionalStatusWords,
		ConnectionSerialNumber: connectionSerialNumber,
		OriginatorVendorId:     originatorVendorId,
		OriginatorSerialNumber: originatorSerialNumber,
		ApplicationReplySize:   applicationReplySize,
		_CipService:            NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipConnectionManagerCloseResponse(structType interface{}) CipConnectionManagerCloseResponse {
	if casted, ok := structType.(CipConnectionManagerCloseResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CipConnectionManagerCloseResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CipConnectionManagerCloseResponse) GetTypeName() string {
	return "CipConnectionManagerCloseResponse"
}

func (m *_CipConnectionManagerCloseResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (additionalStatusWords)
	lengthInBits += 8

	// Simple field (connectionSerialNumber)
	lengthInBits += 16

	// Simple field (originatorVendorId)
	lengthInBits += 16

	// Simple field (originatorSerialNumber)
	lengthInBits += 32

	// Simple field (applicationReplySize)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CipConnectionManagerCloseResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CipConnectionManagerCloseResponseParse(theBytes []byte, connected bool, serviceLen uint16) (CipConnectionManagerCloseResponse, error) {
	return CipConnectionManagerCloseResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func CipConnectionManagerCloseResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (CipConnectionManagerCloseResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipConnectionManagerCloseResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipConnectionManagerCloseResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CipConnectionManagerCloseResponse")
		}
		if reserved != uint8(0x00) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (status)
	_status, _statusErr := readBuffer.ReadUint8("status", 8)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field of CipConnectionManagerCloseResponse")
	}
	status := _status

	// Simple Field (additionalStatusWords)
	_additionalStatusWords, _additionalStatusWordsErr := readBuffer.ReadUint8("additionalStatusWords", 8)
	if _additionalStatusWordsErr != nil {
		return nil, errors.Wrap(_additionalStatusWordsErr, "Error parsing 'additionalStatusWords' field of CipConnectionManagerCloseResponse")
	}
	additionalStatusWords := _additionalStatusWords

	// Simple Field (connectionSerialNumber)
	_connectionSerialNumber, _connectionSerialNumberErr := readBuffer.ReadUint16("connectionSerialNumber", 16)
	if _connectionSerialNumberErr != nil {
		return nil, errors.Wrap(_connectionSerialNumberErr, "Error parsing 'connectionSerialNumber' field of CipConnectionManagerCloseResponse")
	}
	connectionSerialNumber := _connectionSerialNumber

	// Simple Field (originatorVendorId)
	_originatorVendorId, _originatorVendorIdErr := readBuffer.ReadUint16("originatorVendorId", 16)
	if _originatorVendorIdErr != nil {
		return nil, errors.Wrap(_originatorVendorIdErr, "Error parsing 'originatorVendorId' field of CipConnectionManagerCloseResponse")
	}
	originatorVendorId := _originatorVendorId

	// Simple Field (originatorSerialNumber)
	_originatorSerialNumber, _originatorSerialNumberErr := readBuffer.ReadUint32("originatorSerialNumber", 32)
	if _originatorSerialNumberErr != nil {
		return nil, errors.Wrap(_originatorSerialNumberErr, "Error parsing 'originatorSerialNumber' field of CipConnectionManagerCloseResponse")
	}
	originatorSerialNumber := _originatorSerialNumber

	// Simple Field (applicationReplySize)
	_applicationReplySize, _applicationReplySizeErr := readBuffer.ReadUint8("applicationReplySize", 8)
	if _applicationReplySizeErr != nil {
		return nil, errors.Wrap(_applicationReplySizeErr, "Error parsing 'applicationReplySize' field of CipConnectionManagerCloseResponse")
	}
	applicationReplySize := _applicationReplySize

	var reservedField1 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CipConnectionManagerCloseResponse")
		}
		if reserved != uint8(0x00) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	if closeErr := readBuffer.CloseContext("CipConnectionManagerCloseResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipConnectionManagerCloseResponse")
	}

	// Create a partially initialized instance
	_child := &_CipConnectionManagerCloseResponse{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
		Status:                 status,
		AdditionalStatusWords:  additionalStatusWords,
		ConnectionSerialNumber: connectionSerialNumber,
		OriginatorVendorId:     originatorVendorId,
		OriginatorSerialNumber: originatorSerialNumber,
		ApplicationReplySize:   applicationReplySize,
		reservedField0:         reservedField0,
		reservedField1:         reservedField1,
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_CipConnectionManagerCloseResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipConnectionManagerCloseResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipConnectionManagerCloseResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipConnectionManagerCloseResponse")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 8, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (status)
		status := uint8(m.GetStatus())
		_statusErr := writeBuffer.WriteUint8("status", 8, (status))
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		// Simple Field (additionalStatusWords)
		additionalStatusWords := uint8(m.GetAdditionalStatusWords())
		_additionalStatusWordsErr := writeBuffer.WriteUint8("additionalStatusWords", 8, (additionalStatusWords))
		if _additionalStatusWordsErr != nil {
			return errors.Wrap(_additionalStatusWordsErr, "Error serializing 'additionalStatusWords' field")
		}

		// Simple Field (connectionSerialNumber)
		connectionSerialNumber := uint16(m.GetConnectionSerialNumber())
		_connectionSerialNumberErr := writeBuffer.WriteUint16("connectionSerialNumber", 16, (connectionSerialNumber))
		if _connectionSerialNumberErr != nil {
			return errors.Wrap(_connectionSerialNumberErr, "Error serializing 'connectionSerialNumber' field")
		}

		// Simple Field (originatorVendorId)
		originatorVendorId := uint16(m.GetOriginatorVendorId())
		_originatorVendorIdErr := writeBuffer.WriteUint16("originatorVendorId", 16, (originatorVendorId))
		if _originatorVendorIdErr != nil {
			return errors.Wrap(_originatorVendorIdErr, "Error serializing 'originatorVendorId' field")
		}

		// Simple Field (originatorSerialNumber)
		originatorSerialNumber := uint32(m.GetOriginatorSerialNumber())
		_originatorSerialNumberErr := writeBuffer.WriteUint32("originatorSerialNumber", 32, (originatorSerialNumber))
		if _originatorSerialNumberErr != nil {
			return errors.Wrap(_originatorSerialNumberErr, "Error serializing 'originatorSerialNumber' field")
		}

		// Simple Field (applicationReplySize)
		applicationReplySize := uint8(m.GetApplicationReplySize())
		_applicationReplySizeErr := writeBuffer.WriteUint8("applicationReplySize", 8, (applicationReplySize))
		if _applicationReplySizeErr != nil {
			return errors.Wrap(_applicationReplySizeErr, "Error serializing 'applicationReplySize' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField1 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField1
			}
			_err := writeBuffer.WriteUint8("reserved", 8, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("CipConnectionManagerCloseResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipConnectionManagerCloseResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipConnectionManagerCloseResponse) isCipConnectionManagerCloseResponse() bool {
	return true
}

func (m *_CipConnectionManagerCloseResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
