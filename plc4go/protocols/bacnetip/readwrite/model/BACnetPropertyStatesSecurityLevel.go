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

// BACnetPropertyStatesSecurityLevel is the corresponding interface of BACnetPropertyStatesSecurityLevel
type BACnetPropertyStatesSecurityLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetSecurityLevel returns SecurityLevel (property field)
	GetSecurityLevel() BACnetSecurityLevelTagged
}

// BACnetPropertyStatesSecurityLevelExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesSecurityLevel.
// This is useful for switch cases.
type BACnetPropertyStatesSecurityLevelExactly interface {
	BACnetPropertyStatesSecurityLevel
	isBACnetPropertyStatesSecurityLevel() bool
}

// _BACnetPropertyStatesSecurityLevel is the data-structure of this message
type _BACnetPropertyStatesSecurityLevel struct {
	*_BACnetPropertyStates
	SecurityLevel BACnetSecurityLevelTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesSecurityLevel) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesSecurityLevel) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesSecurityLevel) GetSecurityLevel() BACnetSecurityLevelTagged {
	return m.SecurityLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesSecurityLevel factory function for _BACnetPropertyStatesSecurityLevel
func NewBACnetPropertyStatesSecurityLevel(securityLevel BACnetSecurityLevelTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesSecurityLevel {
	_result := &_BACnetPropertyStatesSecurityLevel{
		SecurityLevel:         securityLevel,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesSecurityLevel(structType interface{}) BACnetPropertyStatesSecurityLevel {
	if casted, ok := structType.(BACnetPropertyStatesSecurityLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesSecurityLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesSecurityLevel) GetTypeName() string {
	return "BACnetPropertyStatesSecurityLevel"
}

func (m *_BACnetPropertyStatesSecurityLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (securityLevel)
	lengthInBits += m.SecurityLevel.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesSecurityLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesSecurityLevelParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesSecurityLevel, error) {
	return BACnetPropertyStatesSecurityLevelParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesSecurityLevelParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesSecurityLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesSecurityLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesSecurityLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (securityLevel)
	if pullErr := readBuffer.PullContext("securityLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for securityLevel")
	}
	_securityLevel, _securityLevelErr := BACnetSecurityLevelTaggedParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _securityLevelErr != nil {
		return nil, errors.Wrap(_securityLevelErr, "Error parsing 'securityLevel' field of BACnetPropertyStatesSecurityLevel")
	}
	securityLevel := _securityLevel.(BACnetSecurityLevelTagged)
	if closeErr := readBuffer.CloseContext("securityLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for securityLevel")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesSecurityLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesSecurityLevel")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesSecurityLevel{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		SecurityLevel:         securityLevel,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesSecurityLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesSecurityLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesSecurityLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesSecurityLevel")
		}

		// Simple Field (securityLevel)
		if pushErr := writeBuffer.PushContext("securityLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securityLevel")
		}
		_securityLevelErr := writeBuffer.WriteSerializable(ctx, m.GetSecurityLevel())
		if popErr := writeBuffer.PopContext("securityLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securityLevel")
		}
		if _securityLevelErr != nil {
			return errors.Wrap(_securityLevelErr, "Error serializing 'securityLevel' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesSecurityLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesSecurityLevel")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesSecurityLevel) isBACnetPropertyStatesSecurityLevel() bool {
	return true
}

func (m *_BACnetPropertyStatesSecurityLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
