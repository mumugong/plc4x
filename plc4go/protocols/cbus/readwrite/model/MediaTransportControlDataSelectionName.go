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

// MediaTransportControlDataSelectionName is the corresponding interface of MediaTransportControlDataSelectionName
type MediaTransportControlDataSelectionName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetSelectionName returns SelectionName (property field)
	GetSelectionName() string
}

// MediaTransportControlDataSelectionNameExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataSelectionName.
// This is useful for switch cases.
type MediaTransportControlDataSelectionNameExactly interface {
	MediaTransportControlDataSelectionName
	isMediaTransportControlDataSelectionName() bool
}

// _MediaTransportControlDataSelectionName is the data-structure of this message
type _MediaTransportControlDataSelectionName struct {
	*_MediaTransportControlData
	SelectionName string
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataSelectionName) InitializeParent(parent MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.MediaLinkGroup = mediaLinkGroup
}

func (m *_MediaTransportControlDataSelectionName) GetParent() MediaTransportControlData {
	return m._MediaTransportControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataSelectionName) GetSelectionName() string {
	return m.SelectionName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataSelectionName factory function for _MediaTransportControlDataSelectionName
func NewMediaTransportControlDataSelectionName(selectionName string, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataSelectionName {
	_result := &_MediaTransportControlDataSelectionName{
		SelectionName:              selectionName,
		_MediaTransportControlData: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result._MediaTransportControlData._MediaTransportControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataSelectionName(structType interface{}) MediaTransportControlDataSelectionName {
	if casted, ok := structType.(MediaTransportControlDataSelectionName); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataSelectionName); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataSelectionName) GetTypeName() string {
	return "MediaTransportControlDataSelectionName"
}

func (m *_MediaTransportControlDataSelectionName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (selectionName)
	lengthInBits += uint16(int32((int32(m.GetCommandTypeContainer().NumBytes()) - int32(int32(1)))) * int32(int32(8)))

	return lengthInBits
}

func (m *_MediaTransportControlDataSelectionName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MediaTransportControlDataSelectionNameParse(theBytes []byte, commandTypeContainer MediaTransportControlCommandTypeContainer) (MediaTransportControlDataSelectionName, error) {
	return MediaTransportControlDataSelectionNameParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), commandTypeContainer)
}

func MediaTransportControlDataSelectionNameParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, commandTypeContainer MediaTransportControlCommandTypeContainer) (MediaTransportControlDataSelectionName, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataSelectionName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataSelectionName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (selectionName)
	_selectionName, _selectionNameErr := readBuffer.ReadString("selectionName", uint32(((commandTypeContainer.NumBytes())-(1))*(8)), "UTF-8")
	if _selectionNameErr != nil {
		return nil, errors.Wrap(_selectionNameErr, "Error parsing 'selectionName' field of MediaTransportControlDataSelectionName")
	}
	selectionName := _selectionName

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataSelectionName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataSelectionName")
	}

	// Create a partially initialized instance
	_child := &_MediaTransportControlDataSelectionName{
		_MediaTransportControlData: &_MediaTransportControlData{},
		SelectionName:              selectionName,
	}
	_child._MediaTransportControlData._MediaTransportControlDataChildRequirements = _child
	return _child, nil
}

func (m *_MediaTransportControlDataSelectionName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataSelectionName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataSelectionName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataSelectionName")
		}

		// Simple Field (selectionName)
		selectionName := string(m.GetSelectionName())
		_selectionNameErr := writeBuffer.WriteString("selectionName", uint32(((m.GetCommandTypeContainer().NumBytes())-(1))*(8)), "UTF-8", (selectionName))
		if _selectionNameErr != nil {
			return errors.Wrap(_selectionNameErr, "Error serializing 'selectionName' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataSelectionName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataSelectionName")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataSelectionName) isMediaTransportControlDataSelectionName() bool {
	return true
}

func (m *_MediaTransportControlDataSelectionName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
