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

// MediaTransportControlData is the corresponding interface of MediaTransportControlData
type MediaTransportControlData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() MediaTransportControlCommandTypeContainer
	// GetMediaLinkGroup returns MediaLinkGroup (property field)
	GetMediaLinkGroup() byte
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() MediaTransportControlCommandType
}

// MediaTransportControlDataExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlData.
// This is useful for switch cases.
type MediaTransportControlDataExactly interface {
	MediaTransportControlData
	isMediaTransportControlData() bool
}

// _MediaTransportControlData is the data-structure of this message
type _MediaTransportControlData struct {
	_MediaTransportControlDataChildRequirements
	CommandTypeContainer MediaTransportControlCommandTypeContainer
	MediaLinkGroup       byte
}

type _MediaTransportControlDataChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetCommandType() MediaTransportControlCommandType
}

type MediaTransportControlDataParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child MediaTransportControlData, serializeChildFunction func() error) error
	GetTypeName() string
}

type MediaTransportControlDataChild interface {
	utils.Serializable
	InitializeParent(parent MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte)
	GetParent() *MediaTransportControlData

	GetTypeName() string
	MediaTransportControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlData) GetCommandTypeContainer() MediaTransportControlCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_MediaTransportControlData) GetMediaLinkGroup() byte {
	return m.MediaLinkGroup
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlData) GetCommandType() MediaTransportControlCommandType {
	ctx := context.Background()
	_ = ctx
	return CastMediaTransportControlCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlData factory function for _MediaTransportControlData
func NewMediaTransportControlData(commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlData {
	return &_MediaTransportControlData{CommandTypeContainer: commandTypeContainer, MediaLinkGroup: mediaLinkGroup}
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlData(structType interface{}) MediaTransportControlData {
	if casted, ok := structType.(MediaTransportControlData); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlData); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlData) GetTypeName() string {
	return "MediaTransportControlData"
}

func (m *_MediaTransportControlData) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (mediaLinkGroup)
	lengthInBits += 8

	return lengthInBits
}

func (m *_MediaTransportControlData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MediaTransportControlDataParse(theBytes []byte) (MediaTransportControlData, error) {
	return MediaTransportControlDataParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func MediaTransportControlDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MediaTransportControlData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsMediaTransportControlCommandTypeContainer(readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{"no command type could be found"})
	}

	// Simple Field (commandTypeContainer)
	if pullErr := readBuffer.PullContext("commandTypeContainer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for commandTypeContainer")
	}
	_commandTypeContainer, _commandTypeContainerErr := MediaTransportControlCommandTypeContainerParseWithBuffer(ctx, readBuffer)
	if _commandTypeContainerErr != nil {
		return nil, errors.Wrap(_commandTypeContainerErr, "Error parsing 'commandTypeContainer' field of MediaTransportControlData")
	}
	commandTypeContainer := _commandTypeContainer
	if closeErr := readBuffer.CloseContext("commandTypeContainer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for commandTypeContainer")
	}

	// Virtual field
	_commandType := commandTypeContainer.CommandType()
	commandType := MediaTransportControlCommandType(_commandType)
	_ = commandType

	// Simple Field (mediaLinkGroup)
	_mediaLinkGroup, _mediaLinkGroupErr := readBuffer.ReadByte("mediaLinkGroup")
	if _mediaLinkGroupErr != nil {
		return nil, errors.Wrap(_mediaLinkGroupErr, "Error parsing 'mediaLinkGroup' field of MediaTransportControlData")
	}
	mediaLinkGroup := _mediaLinkGroup

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type MediaTransportControlDataChildSerializeRequirement interface {
		MediaTransportControlData
		InitializeParent(MediaTransportControlData, MediaTransportControlCommandTypeContainer, byte)
		GetParent() MediaTransportControlData
	}
	var _childTemp interface{}
	var _child MediaTransportControlDataChildSerializeRequirement
	var typeSwitchError error
	switch {
	case commandType == MediaTransportControlCommandType_STOP: // MediaTransportControlDataStop
		_childTemp, typeSwitchError = MediaTransportControlDataStopParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_PLAY: // MediaTransportControlDataPlay
		_childTemp, typeSwitchError = MediaTransportControlDataPlayParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_PAUSE_RESUME: // MediaTransportControlDataPauseResume
		_childTemp, typeSwitchError = MediaTransportControlDataPauseResumeParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_SELECT_CATEGORY: // MediaTransportControlDataSetCategory
		_childTemp, typeSwitchError = MediaTransportControlDataSetCategoryParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_SELECT_SELECTION: // MediaTransportControlDataSetSelection
		_childTemp, typeSwitchError = MediaTransportControlDataSetSelectionParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_SELECT_TRACK: // MediaTransportControlDataSetTrack
		_childTemp, typeSwitchError = MediaTransportControlDataSetTrackParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_SHUFFLE_ON_OFF: // MediaTransportControlDataShuffleOnOff
		_childTemp, typeSwitchError = MediaTransportControlDataShuffleOnOffParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_REPEAT_ON_OFF: // MediaTransportControlDataRepeatOnOff
		_childTemp, typeSwitchError = MediaTransportControlDataRepeatOnOffParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_NEXT_PREVIOUS_CATEGORY: // MediaTransportControlDataNextPreviousCategory
		_childTemp, typeSwitchError = MediaTransportControlDataNextPreviousCategoryParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_NEXT_PREVIOUS_SELECTION: // MediaTransportControlDataNextPreviousSelection
		_childTemp, typeSwitchError = MediaTransportControlDataNextPreviousSelectionParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_NEXT_PREVIOUS_TRACK: // MediaTransportControlDataNextPreviousTrack
		_childTemp, typeSwitchError = MediaTransportControlDataNextPreviousTrackParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_FAST_FORWARD: // MediaTransportControlDataFastForward
		_childTemp, typeSwitchError = MediaTransportControlDataFastForwardParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_REWIND: // MediaTransportControlDataRewind
		_childTemp, typeSwitchError = MediaTransportControlDataRewindParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_SOURCE_POWER_CONTROL: // MediaTransportControlDataSourcePowerControl
		_childTemp, typeSwitchError = MediaTransportControlDataSourcePowerControlParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_TOTAL_TRACKS: // MediaTransportControlDataTotalTracks
		_childTemp, typeSwitchError = MediaTransportControlDataTotalTracksParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_STATUS_REQUEST: // MediaTransportControlDataStatusRequest
		_childTemp, typeSwitchError = MediaTransportControlDataStatusRequestParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_ENUMERATE_CATEGORIES_SELECTIONS_TRACKS: // MediaTransportControlDataEnumerateCategoriesSelectionTracks
		_childTemp, typeSwitchError = MediaTransportControlDataEnumerateCategoriesSelectionTracksParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_ENUMERATION_SIZE: // MediaTransportControlDataEnumerationsSize
		_childTemp, typeSwitchError = MediaTransportControlDataEnumerationsSizeParseWithBuffer(ctx, readBuffer)
	case commandType == MediaTransportControlCommandType_TRACK_NAME: // MediaTransportControlDataTrackName
		_childTemp, typeSwitchError = MediaTransportControlDataTrackNameParseWithBuffer(ctx, readBuffer, commandTypeContainer)
	case commandType == MediaTransportControlCommandType_SELECTION_NAME: // MediaTransportControlDataSelectionName
		_childTemp, typeSwitchError = MediaTransportControlDataSelectionNameParseWithBuffer(ctx, readBuffer, commandTypeContainer)
	case commandType == MediaTransportControlCommandType_CATEGORY_NAME: // MediaTransportControlDataCategoryName
		_childTemp, typeSwitchError = MediaTransportControlDataCategoryNameParseWithBuffer(ctx, readBuffer, commandTypeContainer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of MediaTransportControlData")
	}
	_child = _childTemp.(MediaTransportControlDataChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("MediaTransportControlData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlData")
	}

	// Finish initializing
	_child.InitializeParent(_child, commandTypeContainer, mediaLinkGroup)
	return _child, nil
}

func (pm *_MediaTransportControlData) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child MediaTransportControlData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("MediaTransportControlData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for MediaTransportControlData")
	}

	// Simple Field (commandTypeContainer)
	if pushErr := writeBuffer.PushContext("commandTypeContainer"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for commandTypeContainer")
	}
	_commandTypeContainerErr := writeBuffer.WriteSerializable(ctx, m.GetCommandTypeContainer())
	if popErr := writeBuffer.PopContext("commandTypeContainer"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for commandTypeContainer")
	}
	if _commandTypeContainerErr != nil {
		return errors.Wrap(_commandTypeContainerErr, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Simple Field (mediaLinkGroup)
	mediaLinkGroup := byte(m.GetMediaLinkGroup())
	_mediaLinkGroupErr := writeBuffer.WriteByte("mediaLinkGroup", (mediaLinkGroup))
	if _mediaLinkGroupErr != nil {
		return errors.Wrap(_mediaLinkGroupErr, "Error serializing 'mediaLinkGroup' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("MediaTransportControlData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for MediaTransportControlData")
	}
	return nil
}

func (m *_MediaTransportControlData) isMediaTransportControlData() bool {
	return true
}

func (m *_MediaTransportControlData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
