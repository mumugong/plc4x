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
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BVLCDistributeBroadcastToNetwork is the corresponding interface of BVLCDistributeBroadcastToNetwork
type BVLCDistributeBroadcastToNetwork interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BVLC
	// GetNpdu returns Npdu (property field)
	GetNpdu() NPDU
}

// BVLCDistributeBroadcastToNetworkExactly can be used when we want exactly this type and not a type which fulfills BVLCDistributeBroadcastToNetwork.
// This is useful for switch cases.
type BVLCDistributeBroadcastToNetworkExactly interface {
	BVLCDistributeBroadcastToNetwork
	isBVLCDistributeBroadcastToNetwork() bool
}

// _BVLCDistributeBroadcastToNetwork is the data-structure of this message
type _BVLCDistributeBroadcastToNetwork struct {
	*_BVLC
	Npdu NPDU

	// Arguments.
	BvlcPayloadLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCDistributeBroadcastToNetwork) GetBvlcFunction() uint8 {
	return 0x09
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCDistributeBroadcastToNetwork) InitializeParent(parent BVLC) {}

func (m *_BVLCDistributeBroadcastToNetwork) GetParent() BVLC {
	return m._BVLC
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCDistributeBroadcastToNetwork) GetNpdu() NPDU {
	return m.Npdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCDistributeBroadcastToNetwork factory function for _BVLCDistributeBroadcastToNetwork
func NewBVLCDistributeBroadcastToNetwork(npdu NPDU, bvlcPayloadLength uint16) *_BVLCDistributeBroadcastToNetwork {
	_result := &_BVLCDistributeBroadcastToNetwork{
		Npdu:  npdu,
		_BVLC: NewBVLC(),
	}
	_result._BVLC._BVLCChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBVLCDistributeBroadcastToNetwork(structType interface{}) BVLCDistributeBroadcastToNetwork {
	if casted, ok := structType.(BVLCDistributeBroadcastToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCDistributeBroadcastToNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCDistributeBroadcastToNetwork) GetTypeName() string {
	return "BVLCDistributeBroadcastToNetwork"
}

func (m *_BVLCDistributeBroadcastToNetwork) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (npdu)
	lengthInBits += m.Npdu.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BVLCDistributeBroadcastToNetwork) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BVLCDistributeBroadcastToNetworkParse(theBytes []byte, bvlcPayloadLength uint16) (BVLCDistributeBroadcastToNetwork, error) {
	return BVLCDistributeBroadcastToNetworkParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), bvlcPayloadLength)
}

func BVLCDistributeBroadcastToNetworkParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, bvlcPayloadLength uint16) (BVLCDistributeBroadcastToNetwork, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCDistributeBroadcastToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCDistributeBroadcastToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (npdu)
	if pullErr := readBuffer.PullContext("npdu"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for npdu")
	}
	_npdu, _npduErr := NPDUParseWithBuffer(ctx, readBuffer, uint16(bvlcPayloadLength))
	if _npduErr != nil {
		return nil, errors.Wrap(_npduErr, "Error parsing 'npdu' field of BVLCDistributeBroadcastToNetwork")
	}
	npdu := _npdu.(NPDU)
	if closeErr := readBuffer.CloseContext("npdu"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for npdu")
	}

	if closeErr := readBuffer.CloseContext("BVLCDistributeBroadcastToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCDistributeBroadcastToNetwork")
	}

	// Create a partially initialized instance
	_child := &_BVLCDistributeBroadcastToNetwork{
		_BVLC: &_BVLC{},
		Npdu:  npdu,
	}
	_child._BVLC._BVLCChildRequirements = _child
	return _child, nil
}

func (m *_BVLCDistributeBroadcastToNetwork) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCDistributeBroadcastToNetwork) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCDistributeBroadcastToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCDistributeBroadcastToNetwork")
		}

		// Simple Field (npdu)
		if pushErr := writeBuffer.PushContext("npdu"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for npdu")
		}
		_npduErr := writeBuffer.WriteSerializable(ctx, m.GetNpdu())
		if popErr := writeBuffer.PopContext("npdu"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for npdu")
		}
		if _npduErr != nil {
			return errors.Wrap(_npduErr, "Error serializing 'npdu' field")
		}

		if popErr := writeBuffer.PopContext("BVLCDistributeBroadcastToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCDistributeBroadcastToNetwork")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BVLCDistributeBroadcastToNetwork) GetBvlcPayloadLength() uint16 {
	return m.BvlcPayloadLength
}

//
////

func (m *_BVLCDistributeBroadcastToNetwork) isBVLCDistributeBroadcastToNetwork() bool {
	return true
}

func (m *_BVLCDistributeBroadcastToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
