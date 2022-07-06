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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequest is the corresponding interface of BACnetConfirmedServiceRequest
type BACnetConfirmedServiceRequest interface {
	utils.LengthAware
	utils.Serializable
	// GetServiceChoice returns ServiceChoice (discriminator field)
	GetServiceChoice() BACnetConfirmedServiceChoice
	// GetServiceRequestPayloadLength returns ServiceRequestPayloadLength (virtual field)
	GetServiceRequestPayloadLength() uint16
}

// BACnetConfirmedServiceRequestExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequest.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestExactly interface {
	BACnetConfirmedServiceRequest
	isBACnetConfirmedServiceRequest() bool
}

// _BACnetConfirmedServiceRequest is the data-structure of this message
type _BACnetConfirmedServiceRequest struct {
	_BACnetConfirmedServiceRequestChildRequirements

	// Arguments.
	ServiceRequestLength uint16
}

type _BACnetConfirmedServiceRequestChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetServiceChoice() BACnetConfirmedServiceChoice
}

type BACnetConfirmedServiceRequestParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child BACnetConfirmedServiceRequest, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetConfirmedServiceRequestChild interface {
	utils.Serializable
	InitializeParent(parent BACnetConfirmedServiceRequest)
	GetParent() *BACnetConfirmedServiceRequest

	GetTypeName() string
	BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequest) GetServiceRequestPayloadLength() uint16 {
	return uint16(utils.InlineIf(bool(bool((m.ServiceRequestLength) > (0))), func() interface{} { return uint16(uint16(uint16(m.ServiceRequestLength) - uint16(uint16(1)))) }, func() interface{} { return uint16(uint16(0)) }).(uint16))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequest factory function for _BACnetConfirmedServiceRequest
func NewBACnetConfirmedServiceRequest(serviceRequestLength uint16) *_BACnetConfirmedServiceRequest {
	return &_BACnetConfirmedServiceRequest{ServiceRequestLength: serviceRequestLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequest(structType interface{}) BACnetConfirmedServiceRequest {
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequest) GetTypeName() string {
	return "BACnetConfirmedServiceRequest"
}

func (m *_BACnetConfirmedServiceRequest) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetConfirmedServiceRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	if pullErr := readBuffer.PullContext("serviceChoice"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serviceChoice")
	}
	serviceChoice_temp, _serviceChoiceErr := BACnetConfirmedServiceChoiceParse(readBuffer)
	var serviceChoice BACnetConfirmedServiceChoice = serviceChoice_temp
	if closeErr := readBuffer.CloseContext("serviceChoice"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serviceChoice")
	}
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field of BACnetConfirmedServiceRequest")
	}

	// Virtual field
	_serviceRequestPayloadLength := utils.InlineIf(bool(bool((serviceRequestLength) > (0))), func() interface{} { return uint16(uint16(uint16(serviceRequestLength) - uint16(uint16(1)))) }, func() interface{} { return uint16(uint16(0)) }).(uint16)
	serviceRequestPayloadLength := uint16(_serviceRequestPayloadLength)
	_ = serviceRequestPayloadLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetConfirmedServiceRequestChildSerializeRequirement interface {
		BACnetConfirmedServiceRequest
		InitializeParent(BACnetConfirmedServiceRequest)
		GetParent() BACnetConfirmedServiceRequest
	}
	var _childTemp interface{}
	var _child BACnetConfirmedServiceRequestChildSerializeRequirement
	var typeSwitchError error
	switch {
	case serviceChoice == BACnetConfirmedServiceChoice_ACKNOWLEDGE_ALARM: // BACnetConfirmedServiceRequestAcknowledgeAlarm
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestAcknowledgeAlarmParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION: // BACnetConfirmedServiceRequestConfirmedCOVNotification
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestConfirmedCOVNotificationParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE: // BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_EVENT_NOTIFICATION: // BACnetConfirmedServiceRequestConfirmedEventNotification
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestConfirmedEventNotificationParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY: // BACnetConfirmedServiceRequestGetEnrollmentSummary
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestGetEnrollmentSummaryParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION: // BACnetConfirmedServiceRequestGetEventInformation
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestGetEventInformationParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_LIFE_SAFETY_OPERATION: // BACnetConfirmedServiceRequestLifeSafetyOperation
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestLifeSafetyOperationParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_SUBSCRIBE_COV: // BACnetConfirmedServiceRequestSubscribeCOV
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestSubscribeCOVParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY: // BACnetConfirmedServiceRequestSubscribeCOVProperty
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestSubscribeCOVPropertyParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE: // BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_ATOMIC_READ_FILE: // BACnetConfirmedServiceRequestAtomicReadFile
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestAtomicReadFileParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE: // BACnetConfirmedServiceRequestAtomicWriteFile
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestAtomicWriteFileParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT: // BACnetConfirmedServiceRequestAddListElement
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestAddListElementParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT: // BACnetConfirmedServiceRequestRemoveListElement
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestRemoveListElementParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_CREATE_OBJECT: // BACnetConfirmedServiceRequestCreateObject
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestCreateObjectParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_DELETE_OBJECT: // BACnetConfirmedServiceRequestDeleteObject
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestDeleteObjectParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_READ_PROPERTY: // BACnetConfirmedServiceRequestReadProperty
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestReadPropertyParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE: // BACnetConfirmedServiceRequestReadPropertyMultiple
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestReadPropertyMultipleParse(readBuffer, serviceRequestLength, serviceRequestPayloadLength)
	case serviceChoice == BACnetConfirmedServiceChoice_READ_RANGE: // BACnetConfirmedServiceRequestReadRange
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestReadRangeParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_WRITE_PROPERTY: // BACnetConfirmedServiceRequestWriteProperty
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestWritePropertyParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE: // BACnetConfirmedServiceRequestWritePropertyMultiple
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestWritePropertyMultipleParse(readBuffer, serviceRequestLength, serviceRequestPayloadLength)
	case serviceChoice == BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL: // BACnetConfirmedServiceRequestDeviceCommunicationControl
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestDeviceCommunicationControlParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER: // BACnetConfirmedServiceRequestConfirmedPrivateTransfer
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestConfirmedPrivateTransferParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_TEXT_MESSAGE: // BACnetConfirmedServiceRequestConfirmedTextMessage
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestConfirmedTextMessageParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE: // BACnetConfirmedServiceRequestReinitializeDevice
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestReinitializeDeviceParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_VT_OPEN: // BACnetConfirmedServiceRequestVTOpen
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestVTOpenParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_VT_CLOSE: // BACnetConfirmedServiceRequestVTClose
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestVTCloseParse(readBuffer, serviceRequestLength, serviceRequestPayloadLength)
	case serviceChoice == BACnetConfirmedServiceChoice_VT_DATA: // BACnetConfirmedServiceRequestVTData
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestVTDataParse(readBuffer, serviceRequestLength)
	case serviceChoice == BACnetConfirmedServiceChoice_AUTHENTICATE: // BACnetConfirmedServiceRequestAuthenticate
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestAuthenticateParse(readBuffer, serviceRequestLength, serviceRequestPayloadLength)
	case serviceChoice == BACnetConfirmedServiceChoice_REQUEST_KEY: // BACnetConfirmedServiceRequestRequestKey
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestRequestKeyParse(readBuffer, serviceRequestLength, serviceRequestPayloadLength)
	case serviceChoice == BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL: // BACnetConfirmedServiceRequestReadPropertyConditional
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestReadPropertyConditionalParse(readBuffer, serviceRequestLength, serviceRequestPayloadLength)
	case 0 == 0: // BACnetConfirmedServiceRequestUnknown
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestUnknownParse(readBuffer, serviceRequestLength, serviceRequestPayloadLength)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [serviceChoice=%v]", serviceChoice)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetConfirmedServiceRequest")
	}
	_child = _childTemp.(BACnetConfirmedServiceRequestChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequest")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_BACnetConfirmedServiceRequest) SerializeParent(writeBuffer utils.WriteBuffer, child BACnetConfirmedServiceRequest, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequest"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequest")
	}

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice := BACnetConfirmedServiceChoice(child.GetServiceChoice())
	if pushErr := writeBuffer.PushContext("serviceChoice"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for serviceChoice")
	}
	_serviceChoiceErr := writeBuffer.WriteSerializable(serviceChoice)
	if popErr := writeBuffer.PopContext("serviceChoice"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for serviceChoice")
	}

	if _serviceChoiceErr != nil {
		return errors.Wrap(_serviceChoiceErr, "Error serializing 'serviceChoice' field")
	}
	// Virtual field
	if _serviceRequestPayloadLengthErr := writeBuffer.WriteVirtual("serviceRequestPayloadLength", m.GetServiceRequestPayloadLength()); _serviceRequestPayloadLengthErr != nil {
		return errors.Wrap(_serviceRequestPayloadLengthErr, "Error serializing 'serviceRequestPayloadLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequest"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequest")
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequest) isBACnetConfirmedServiceRequest() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
