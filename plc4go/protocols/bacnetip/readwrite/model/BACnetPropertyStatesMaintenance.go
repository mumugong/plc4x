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

// BACnetPropertyStatesMaintenance is the corresponding interface of BACnetPropertyStatesMaintenance
type BACnetPropertyStatesMaintenance interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetMaintenance returns Maintenance (property field)
	GetMaintenance() BACnetMaintenanceTagged
}

// BACnetPropertyStatesMaintenanceExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesMaintenance.
// This is useful for switch cases.
type BACnetPropertyStatesMaintenanceExactly interface {
	BACnetPropertyStatesMaintenance
	isBACnetPropertyStatesMaintenance() bool
}

// _BACnetPropertyStatesMaintenance is the data-structure of this message
type _BACnetPropertyStatesMaintenance struct {
	*_BACnetPropertyStates
	Maintenance BACnetMaintenanceTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesMaintenance) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesMaintenance) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesMaintenance) GetMaintenance() BACnetMaintenanceTagged {
	return m.Maintenance
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesMaintenance factory function for _BACnetPropertyStatesMaintenance
func NewBACnetPropertyStatesMaintenance(maintenance BACnetMaintenanceTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesMaintenance {
	_result := &_BACnetPropertyStatesMaintenance{
		Maintenance:           maintenance,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesMaintenance(structType interface{}) BACnetPropertyStatesMaintenance {
	if casted, ok := structType.(BACnetPropertyStatesMaintenance); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesMaintenance); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesMaintenance) GetTypeName() string {
	return "BACnetPropertyStatesMaintenance"
}

func (m *_BACnetPropertyStatesMaintenance) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (maintenance)
	lengthInBits += m.Maintenance.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesMaintenance) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesMaintenanceParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesMaintenance, error) {
	return BACnetPropertyStatesMaintenanceParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesMaintenanceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesMaintenance, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesMaintenance"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesMaintenance")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maintenance)
	if pullErr := readBuffer.PullContext("maintenance"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maintenance")
	}
	_maintenance, _maintenanceErr := BACnetMaintenanceTaggedParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _maintenanceErr != nil {
		return nil, errors.Wrap(_maintenanceErr, "Error parsing 'maintenance' field of BACnetPropertyStatesMaintenance")
	}
	maintenance := _maintenance.(BACnetMaintenanceTagged)
	if closeErr := readBuffer.CloseContext("maintenance"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maintenance")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesMaintenance"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesMaintenance")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesMaintenance{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		Maintenance:           maintenance,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesMaintenance) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesMaintenance) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesMaintenance"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesMaintenance")
		}

		// Simple Field (maintenance)
		if pushErr := writeBuffer.PushContext("maintenance"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maintenance")
		}
		_maintenanceErr := writeBuffer.WriteSerializable(ctx, m.GetMaintenance())
		if popErr := writeBuffer.PopContext("maintenance"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maintenance")
		}
		if _maintenanceErr != nil {
			return errors.Wrap(_maintenanceErr, "Error serializing 'maintenance' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesMaintenance"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesMaintenance")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesMaintenance) isBACnetPropertyStatesMaintenance() bool {
	return true
}

func (m *_BACnetPropertyStatesMaintenance) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
