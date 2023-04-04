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

// ApduDataExtDomainAddressSerialNumberResponse is the corresponding interface of ApduDataExtDomainAddressSerialNumberResponse
type ApduDataExtDomainAddressSerialNumberResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtDomainAddressSerialNumberResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtDomainAddressSerialNumberResponse.
// This is useful for switch cases.
type ApduDataExtDomainAddressSerialNumberResponseExactly interface {
	ApduDataExtDomainAddressSerialNumberResponse
	isApduDataExtDomainAddressSerialNumberResponse() bool
}

// _ApduDataExtDomainAddressSerialNumberResponse is the data-structure of this message
type _ApduDataExtDomainAddressSerialNumberResponse struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtDomainAddressSerialNumberResponse) GetExtApciType() uint8 {
	return 0x2D
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtDomainAddressSerialNumberResponse) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtDomainAddressSerialNumberResponse) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtDomainAddressSerialNumberResponse factory function for _ApduDataExtDomainAddressSerialNumberResponse
func NewApduDataExtDomainAddressSerialNumberResponse(length uint8) *_ApduDataExtDomainAddressSerialNumberResponse {
	_result := &_ApduDataExtDomainAddressSerialNumberResponse{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtDomainAddressSerialNumberResponse(structType interface{}) ApduDataExtDomainAddressSerialNumberResponse {
	if casted, ok := structType.(ApduDataExtDomainAddressSerialNumberResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtDomainAddressSerialNumberResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtDomainAddressSerialNumberResponse) GetTypeName() string {
	return "ApduDataExtDomainAddressSerialNumberResponse"
}

func (m *_ApduDataExtDomainAddressSerialNumberResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtDomainAddressSerialNumberResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataExtDomainAddressSerialNumberResponseParse(theBytes []byte, length uint8) (ApduDataExtDomainAddressSerialNumberResponse, error) {
	return ApduDataExtDomainAddressSerialNumberResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), length)
}

func ApduDataExtDomainAddressSerialNumberResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, length uint8) (ApduDataExtDomainAddressSerialNumberResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtDomainAddressSerialNumberResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtDomainAddressSerialNumberResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtDomainAddressSerialNumberResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtDomainAddressSerialNumberResponse")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtDomainAddressSerialNumberResponse{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtDomainAddressSerialNumberResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtDomainAddressSerialNumberResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtDomainAddressSerialNumberResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtDomainAddressSerialNumberResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtDomainAddressSerialNumberResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtDomainAddressSerialNumberResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtDomainAddressSerialNumberResponse) isApduDataExtDomainAddressSerialNumberResponse() bool {
	return true
}

func (m *_ApduDataExtDomainAddressSerialNumberResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
