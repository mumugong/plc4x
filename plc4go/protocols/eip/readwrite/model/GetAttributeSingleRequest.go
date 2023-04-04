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

// GetAttributeSingleRequest is the corresponding interface of GetAttributeSingleRequest
type GetAttributeSingleRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
}

// GetAttributeSingleRequestExactly can be used when we want exactly this type and not a type which fulfills GetAttributeSingleRequest.
// This is useful for switch cases.
type GetAttributeSingleRequestExactly interface {
	GetAttributeSingleRequest
	isGetAttributeSingleRequest() bool
}

// _GetAttributeSingleRequest is the data-structure of this message
type _GetAttributeSingleRequest struct {
	*_CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GetAttributeSingleRequest) GetService() uint8 {
	return 0x0E
}

func (m *_GetAttributeSingleRequest) GetResponse() bool {
	return bool(false)
}

func (m *_GetAttributeSingleRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GetAttributeSingleRequest) InitializeParent(parent CipService) {}

func (m *_GetAttributeSingleRequest) GetParent() CipService {
	return m._CipService
}

// NewGetAttributeSingleRequest factory function for _GetAttributeSingleRequest
func NewGetAttributeSingleRequest(serviceLen uint16) *_GetAttributeSingleRequest {
	_result := &_GetAttributeSingleRequest{
		_CipService: NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastGetAttributeSingleRequest(structType interface{}) GetAttributeSingleRequest {
	if casted, ok := structType.(GetAttributeSingleRequest); ok {
		return casted
	}
	if casted, ok := structType.(*GetAttributeSingleRequest); ok {
		return *casted
	}
	return nil
}

func (m *_GetAttributeSingleRequest) GetTypeName() string {
	return "GetAttributeSingleRequest"
}

func (m *_GetAttributeSingleRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_GetAttributeSingleRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func GetAttributeSingleRequestParse(theBytes []byte, connected bool, serviceLen uint16) (GetAttributeSingleRequest, error) {
	return GetAttributeSingleRequestParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func GetAttributeSingleRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (GetAttributeSingleRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GetAttributeSingleRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GetAttributeSingleRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("GetAttributeSingleRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GetAttributeSingleRequest")
	}

	// Create a partially initialized instance
	_child := &_GetAttributeSingleRequest{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_GetAttributeSingleRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GetAttributeSingleRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GetAttributeSingleRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GetAttributeSingleRequest")
		}

		if popErr := writeBuffer.PopContext("GetAttributeSingleRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GetAttributeSingleRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GetAttributeSingleRequest) isGetAttributeSingleRequest() bool {
	return true
}

func (m *_GetAttributeSingleRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
