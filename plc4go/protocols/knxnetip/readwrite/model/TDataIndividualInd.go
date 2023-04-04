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

// TDataIndividualInd is the corresponding interface of TDataIndividualInd
type TDataIndividualInd interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMI
}

// TDataIndividualIndExactly can be used when we want exactly this type and not a type which fulfills TDataIndividualInd.
// This is useful for switch cases.
type TDataIndividualIndExactly interface {
	TDataIndividualInd
	isTDataIndividualInd() bool
}

// _TDataIndividualInd is the data-structure of this message
type _TDataIndividualInd struct {
	*_CEMI
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TDataIndividualInd) GetMessageCode() uint8 {
	return 0x94
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TDataIndividualInd) InitializeParent(parent CEMI) {}

func (m *_TDataIndividualInd) GetParent() CEMI {
	return m._CEMI
}

// NewTDataIndividualInd factory function for _TDataIndividualInd
func NewTDataIndividualInd(size uint16) *_TDataIndividualInd {
	_result := &_TDataIndividualInd{
		_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTDataIndividualInd(structType interface{}) TDataIndividualInd {
	if casted, ok := structType.(TDataIndividualInd); ok {
		return casted
	}
	if casted, ok := structType.(*TDataIndividualInd); ok {
		return *casted
	}
	return nil
}

func (m *_TDataIndividualInd) GetTypeName() string {
	return "TDataIndividualInd"
}

func (m *_TDataIndividualInd) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_TDataIndividualInd) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TDataIndividualIndParse(theBytes []byte, size uint16) (TDataIndividualInd, error) {
	return TDataIndividualIndParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), size)
}

func TDataIndividualIndParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, size uint16) (TDataIndividualInd, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TDataIndividualInd"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TDataIndividualInd")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TDataIndividualInd"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TDataIndividualInd")
	}

	// Create a partially initialized instance
	_child := &_TDataIndividualInd{
		_CEMI: &_CEMI{
			Size: size,
		},
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_TDataIndividualInd) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TDataIndividualInd) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TDataIndividualInd"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TDataIndividualInd")
		}

		if popErr := writeBuffer.PopContext("TDataIndividualInd"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TDataIndividualInd")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TDataIndividualInd) isTDataIndividualInd() bool {
	return true
}

func (m *_TDataIndividualInd) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
