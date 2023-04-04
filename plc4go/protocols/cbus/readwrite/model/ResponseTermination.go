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

// Constant values.
const ResponseTermination_CR byte = 0x0D
const ResponseTermination_LF byte = 0x0A

// ResponseTermination is the corresponding interface of ResponseTermination
type ResponseTermination interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// ResponseTerminationExactly can be used when we want exactly this type and not a type which fulfills ResponseTermination.
// This is useful for switch cases.
type ResponseTerminationExactly interface {
	ResponseTermination
	isResponseTermination() bool
}

// _ResponseTermination is the data-structure of this message
type _ResponseTermination struct {
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_ResponseTermination) GetCr() byte {
	return ResponseTermination_CR
}

func (m *_ResponseTermination) GetLf() byte {
	return ResponseTermination_LF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewResponseTermination factory function for _ResponseTermination
func NewResponseTermination() *_ResponseTermination {
	return &_ResponseTermination{}
}

// Deprecated: use the interface for direct cast
func CastResponseTermination(structType interface{}) ResponseTermination {
	if casted, ok := structType.(ResponseTermination); ok {
		return casted
	}
	if casted, ok := structType.(*ResponseTermination); ok {
		return *casted
	}
	return nil
}

func (m *_ResponseTermination) GetTypeName() string {
	return "ResponseTermination"
}

func (m *_ResponseTermination) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (cr)
	lengthInBits += 8

	// Const Field (lf)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ResponseTermination) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ResponseTerminationParse(theBytes []byte) (ResponseTermination, error) {
	return ResponseTerminationParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func ResponseTerminationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ResponseTermination, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ResponseTermination"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ResponseTermination")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (cr)
	cr, _crErr := readBuffer.ReadByte("cr")
	if _crErr != nil {
		return nil, errors.Wrap(_crErr, "Error parsing 'cr' field of ResponseTermination")
	}
	if cr != ResponseTermination_CR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", ResponseTermination_CR) + " but got " + fmt.Sprintf("%d", cr))
	}

	// Const Field (lf)
	lf, _lfErr := readBuffer.ReadByte("lf")
	if _lfErr != nil {
		return nil, errors.Wrap(_lfErr, "Error parsing 'lf' field of ResponseTermination")
	}
	if lf != ResponseTermination_LF {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", ResponseTermination_LF) + " but got " + fmt.Sprintf("%d", lf))
	}

	if closeErr := readBuffer.CloseContext("ResponseTermination"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ResponseTermination")
	}

	// Create the instance
	return &_ResponseTermination{}, nil
}

func (m *_ResponseTermination) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ResponseTermination) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ResponseTermination"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ResponseTermination")
	}

	// Const Field (cr)
	_crErr := writeBuffer.WriteByte("cr", 0x0D)
	if _crErr != nil {
		return errors.Wrap(_crErr, "Error serializing 'cr' field")
	}

	// Const Field (lf)
	_lfErr := writeBuffer.WriteByte("lf", 0x0A)
	if _lfErr != nil {
		return errors.Wrap(_lfErr, "Error serializing 'lf' field")
	}

	if popErr := writeBuffer.PopContext("ResponseTermination"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ResponseTermination")
	}
	return nil
}

func (m *_ResponseTermination) isResponseTermination() bool {
	return true
}

func (m *_ResponseTermination) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
