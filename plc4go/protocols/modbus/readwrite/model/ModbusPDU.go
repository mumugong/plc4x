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

// ModbusPDU is the corresponding interface of ModbusPDU
type ModbusPDU interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetErrorFlag returns ErrorFlag (discriminator field)
	GetErrorFlag() bool
	// GetFunctionFlag returns FunctionFlag (discriminator field)
	GetFunctionFlag() uint8
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
}

// ModbusPDUExactly can be used when we want exactly this type and not a type which fulfills ModbusPDU.
// This is useful for switch cases.
type ModbusPDUExactly interface {
	ModbusPDU
	isModbusPDU() bool
}

// _ModbusPDU is the data-structure of this message
type _ModbusPDU struct {
	_ModbusPDUChildRequirements
}

type _ModbusPDUChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetErrorFlag() bool
	GetFunctionFlag() uint8
	GetResponse() bool
}

type ModbusPDUParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ModbusPDU, serializeChildFunction func() error) error
	GetTypeName() string
}

type ModbusPDUChild interface {
	utils.Serializable
	InitializeParent(parent ModbusPDU)
	GetParent() *ModbusPDU

	GetTypeName() string
	ModbusPDU
}

// NewModbusPDU factory function for _ModbusPDU
func NewModbusPDU() *_ModbusPDU {
	return &_ModbusPDU{}
}

// Deprecated: use the interface for direct cast
func CastModbusPDU(structType interface{}) ModbusPDU {
	if casted, ok := structType.(ModbusPDU); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDU) GetTypeName() string {
	return "ModbusPDU"
}

func (m *_ModbusPDU) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (errorFlag)
	lengthInBits += 1
	// Discriminator Field (functionFlag)
	lengthInBits += 7

	return lengthInBits
}

func (m *_ModbusPDU) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUParse(theBytes []byte, response bool) (ModbusPDU, error) {
	return ModbusPDUParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDU, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (errorFlag) (Used as input to a switch field)
	errorFlag, _errorFlagErr := readBuffer.ReadBit("errorFlag")
	if _errorFlagErr != nil {
		return nil, errors.Wrap(_errorFlagErr, "Error parsing 'errorFlag' field of ModbusPDU")
	}

	// Discriminator Field (functionFlag) (Used as input to a switch field)
	functionFlag, _functionFlagErr := readBuffer.ReadUint8("functionFlag", 7)
	if _functionFlagErr != nil {
		return nil, errors.Wrap(_functionFlagErr, "Error parsing 'functionFlag' field of ModbusPDU")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ModbusPDUChildSerializeRequirement interface {
		ModbusPDU
		InitializeParent(ModbusPDU)
		GetParent() ModbusPDU
	}
	var _childTemp interface{}
	var _child ModbusPDUChildSerializeRequirement
	var typeSwitchError error
	switch {
	case errorFlag == bool(true): // ModbusPDUError
		_childTemp, typeSwitchError = ModbusPDUErrorParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x02 && response == bool(false): // ModbusPDUReadDiscreteInputsRequest
		_childTemp, typeSwitchError = ModbusPDUReadDiscreteInputsRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x02 && response == bool(true): // ModbusPDUReadDiscreteInputsResponse
		_childTemp, typeSwitchError = ModbusPDUReadDiscreteInputsResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x01 && response == bool(false): // ModbusPDUReadCoilsRequest
		_childTemp, typeSwitchError = ModbusPDUReadCoilsRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x01 && response == bool(true): // ModbusPDUReadCoilsResponse
		_childTemp, typeSwitchError = ModbusPDUReadCoilsResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x05 && response == bool(false): // ModbusPDUWriteSingleCoilRequest
		_childTemp, typeSwitchError = ModbusPDUWriteSingleCoilRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x05 && response == bool(true): // ModbusPDUWriteSingleCoilResponse
		_childTemp, typeSwitchError = ModbusPDUWriteSingleCoilResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x0F && response == bool(false): // ModbusPDUWriteMultipleCoilsRequest
		_childTemp, typeSwitchError = ModbusPDUWriteMultipleCoilsRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x0F && response == bool(true): // ModbusPDUWriteMultipleCoilsResponse
		_childTemp, typeSwitchError = ModbusPDUWriteMultipleCoilsResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x04 && response == bool(false): // ModbusPDUReadInputRegistersRequest
		_childTemp, typeSwitchError = ModbusPDUReadInputRegistersRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x04 && response == bool(true): // ModbusPDUReadInputRegistersResponse
		_childTemp, typeSwitchError = ModbusPDUReadInputRegistersResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x03 && response == bool(false): // ModbusPDUReadHoldingRegistersRequest
		_childTemp, typeSwitchError = ModbusPDUReadHoldingRegistersRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x03 && response == bool(true): // ModbusPDUReadHoldingRegistersResponse
		_childTemp, typeSwitchError = ModbusPDUReadHoldingRegistersResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x06 && response == bool(false): // ModbusPDUWriteSingleRegisterRequest
		_childTemp, typeSwitchError = ModbusPDUWriteSingleRegisterRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x06 && response == bool(true): // ModbusPDUWriteSingleRegisterResponse
		_childTemp, typeSwitchError = ModbusPDUWriteSingleRegisterResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x10 && response == bool(false): // ModbusPDUWriteMultipleHoldingRegistersRequest
		_childTemp, typeSwitchError = ModbusPDUWriteMultipleHoldingRegistersRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x10 && response == bool(true): // ModbusPDUWriteMultipleHoldingRegistersResponse
		_childTemp, typeSwitchError = ModbusPDUWriteMultipleHoldingRegistersResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x17 && response == bool(false): // ModbusPDUReadWriteMultipleHoldingRegistersRequest
		_childTemp, typeSwitchError = ModbusPDUReadWriteMultipleHoldingRegistersRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x17 && response == bool(true): // ModbusPDUReadWriteMultipleHoldingRegistersResponse
		_childTemp, typeSwitchError = ModbusPDUReadWriteMultipleHoldingRegistersResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x16 && response == bool(false): // ModbusPDUMaskWriteHoldingRegisterRequest
		_childTemp, typeSwitchError = ModbusPDUMaskWriteHoldingRegisterRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x16 && response == bool(true): // ModbusPDUMaskWriteHoldingRegisterResponse
		_childTemp, typeSwitchError = ModbusPDUMaskWriteHoldingRegisterResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x18 && response == bool(false): // ModbusPDUReadFifoQueueRequest
		_childTemp, typeSwitchError = ModbusPDUReadFifoQueueRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x18 && response == bool(true): // ModbusPDUReadFifoQueueResponse
		_childTemp, typeSwitchError = ModbusPDUReadFifoQueueResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x14 && response == bool(false): // ModbusPDUReadFileRecordRequest
		_childTemp, typeSwitchError = ModbusPDUReadFileRecordRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x14 && response == bool(true): // ModbusPDUReadFileRecordResponse
		_childTemp, typeSwitchError = ModbusPDUReadFileRecordResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x15 && response == bool(false): // ModbusPDUWriteFileRecordRequest
		_childTemp, typeSwitchError = ModbusPDUWriteFileRecordRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x15 && response == bool(true): // ModbusPDUWriteFileRecordResponse
		_childTemp, typeSwitchError = ModbusPDUWriteFileRecordResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x07 && response == bool(false): // ModbusPDUReadExceptionStatusRequest
		_childTemp, typeSwitchError = ModbusPDUReadExceptionStatusRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x07 && response == bool(true): // ModbusPDUReadExceptionStatusResponse
		_childTemp, typeSwitchError = ModbusPDUReadExceptionStatusResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x08 && response == bool(false): // ModbusPDUDiagnosticRequest
		_childTemp, typeSwitchError = ModbusPDUDiagnosticRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x08 && response == bool(true): // ModbusPDUDiagnosticResponse
		_childTemp, typeSwitchError = ModbusPDUDiagnosticResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x0B && response == bool(false): // ModbusPDUGetComEventCounterRequest
		_childTemp, typeSwitchError = ModbusPDUGetComEventCounterRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x0B && response == bool(true): // ModbusPDUGetComEventCounterResponse
		_childTemp, typeSwitchError = ModbusPDUGetComEventCounterResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x0C && response == bool(false): // ModbusPDUGetComEventLogRequest
		_childTemp, typeSwitchError = ModbusPDUGetComEventLogRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x0C && response == bool(true): // ModbusPDUGetComEventLogResponse
		_childTemp, typeSwitchError = ModbusPDUGetComEventLogResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x11 && response == bool(false): // ModbusPDUReportServerIdRequest
		_childTemp, typeSwitchError = ModbusPDUReportServerIdRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x11 && response == bool(true): // ModbusPDUReportServerIdResponse
		_childTemp, typeSwitchError = ModbusPDUReportServerIdResponseParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x2B && response == bool(false): // ModbusPDUReadDeviceIdentificationRequest
		_childTemp, typeSwitchError = ModbusPDUReadDeviceIdentificationRequestParseWithBuffer(ctx, readBuffer, response)
	case errorFlag == bool(false) && functionFlag == 0x2B && response == bool(true): // ModbusPDUReadDeviceIdentificationResponse
		_childTemp, typeSwitchError = ModbusPDUReadDeviceIdentificationResponseParseWithBuffer(ctx, readBuffer, response)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [errorFlag=%v, functionFlag=%v, response=%v]", errorFlag, functionFlag, response)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of ModbusPDU")
	}
	_child = _childTemp.(ModbusPDUChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("ModbusPDU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDU")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_ModbusPDU) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ModbusPDU, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ModbusPDU"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ModbusPDU")
	}

	// Discriminator Field (errorFlag) (Used as input to a switch field)
	errorFlag := bool(child.GetErrorFlag())
	_errorFlagErr := writeBuffer.WriteBit("errorFlag", (errorFlag))

	if _errorFlagErr != nil {
		return errors.Wrap(_errorFlagErr, "Error serializing 'errorFlag' field")
	}

	// Discriminator Field (functionFlag) (Used as input to a switch field)
	functionFlag := uint8(child.GetFunctionFlag())
	_functionFlagErr := writeBuffer.WriteUint8("functionFlag", 7, (functionFlag))

	if _functionFlagErr != nil {
		return errors.Wrap(_functionFlagErr, "Error serializing 'functionFlag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ModbusPDU"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ModbusPDU")
	}
	return nil
}

func (m *_ModbusPDU) isModbusPDU() bool {
	return true
}

func (m *_ModbusPDU) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
