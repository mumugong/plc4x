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

// Operation is an enum
type Operation uint32

type IOperation interface {
	fmt.Stringer
	utils.Serializable
}

const (
	Operation_DISCOVERY_REQUEST            Operation = 0x00000001
	Operation_DISCOVERY_RESPONSE           Operation = 0x80000001
	Operation_ADD_OR_UPDATE_ROUTE_REQUEST  Operation = 0x00000006
	Operation_ADD_OR_UPDATE_ROUTE_RESPONSE Operation = 0x80000006
	Operation_DEL_ROUTE_REQUEST            Operation = 0x00000007
	Operation_DEL_ROUTE_RESPONSE           Operation = 0x80000007
	Operation_UNKNOWN_REQUEST              Operation = 0x00000008
	Operation_UNKNOWN_RESPONSE             Operation = 0x80000008
)

var OperationValues []Operation

func init() {
	_ = errors.New
	OperationValues = []Operation{
		Operation_DISCOVERY_REQUEST,
		Operation_DISCOVERY_RESPONSE,
		Operation_ADD_OR_UPDATE_ROUTE_REQUEST,
		Operation_ADD_OR_UPDATE_ROUTE_RESPONSE,
		Operation_DEL_ROUTE_REQUEST,
		Operation_DEL_ROUTE_RESPONSE,
		Operation_UNKNOWN_REQUEST,
		Operation_UNKNOWN_RESPONSE,
	}
}

func OperationByValue(value uint32) (enum Operation, ok bool) {
	switch value {
	case 0x00000001:
		return Operation_DISCOVERY_REQUEST, true
	case 0x00000006:
		return Operation_ADD_OR_UPDATE_ROUTE_REQUEST, true
	case 0x00000007:
		return Operation_DEL_ROUTE_REQUEST, true
	case 0x00000008:
		return Operation_UNKNOWN_REQUEST, true
	case 0x80000001:
		return Operation_DISCOVERY_RESPONSE, true
	case 0x80000006:
		return Operation_ADD_OR_UPDATE_ROUTE_RESPONSE, true
	case 0x80000007:
		return Operation_DEL_ROUTE_RESPONSE, true
	case 0x80000008:
		return Operation_UNKNOWN_RESPONSE, true
	}
	return 0, false
}

func OperationByName(value string) (enum Operation, ok bool) {
	switch value {
	case "DISCOVERY_REQUEST":
		return Operation_DISCOVERY_REQUEST, true
	case "ADD_OR_UPDATE_ROUTE_REQUEST":
		return Operation_ADD_OR_UPDATE_ROUTE_REQUEST, true
	case "DEL_ROUTE_REQUEST":
		return Operation_DEL_ROUTE_REQUEST, true
	case "UNKNOWN_REQUEST":
		return Operation_UNKNOWN_REQUEST, true
	case "DISCOVERY_RESPONSE":
		return Operation_DISCOVERY_RESPONSE, true
	case "ADD_OR_UPDATE_ROUTE_RESPONSE":
		return Operation_ADD_OR_UPDATE_ROUTE_RESPONSE, true
	case "DEL_ROUTE_RESPONSE":
		return Operation_DEL_ROUTE_RESPONSE, true
	case "UNKNOWN_RESPONSE":
		return Operation_UNKNOWN_RESPONSE, true
	}
	return 0, false
}

func OperationKnows(value uint32) bool {
	for _, typeValue := range OperationValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOperation(structType interface{}) Operation {
	castFunc := func(typ interface{}) Operation {
		if sOperation, ok := typ.(Operation); ok {
			return sOperation
		}
		return 0
	}
	return castFunc(structType)
}

func (m Operation) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m Operation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OperationParse(ctx context.Context, theBytes []byte) (Operation, error) {
	return OperationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OperationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Operation, error) {
	val, err := readBuffer.ReadUint32("Operation", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading Operation")
	}
	if enum, ok := OperationByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return Operation(val), nil
	} else {
		return enum, nil
	}
}

func (e Operation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e Operation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint32("Operation", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e Operation) PLC4XEnumName() string {
	switch e {
	case Operation_DISCOVERY_REQUEST:
		return "DISCOVERY_REQUEST"
	case Operation_ADD_OR_UPDATE_ROUTE_REQUEST:
		return "ADD_OR_UPDATE_ROUTE_REQUEST"
	case Operation_DEL_ROUTE_REQUEST:
		return "DEL_ROUTE_REQUEST"
	case Operation_UNKNOWN_REQUEST:
		return "UNKNOWN_REQUEST"
	case Operation_DISCOVERY_RESPONSE:
		return "DISCOVERY_RESPONSE"
	case Operation_ADD_OR_UPDATE_ROUTE_RESPONSE:
		return "ADD_OR_UPDATE_ROUTE_RESPONSE"
	case Operation_DEL_ROUTE_RESPONSE:
		return "DEL_ROUTE_RESPONSE"
	case Operation_UNKNOWN_RESPONSE:
		return "UNKNOWN_RESPONSE"
	}
	return ""
}

func (e Operation) String() string {
	return e.PLC4XEnumName()
}
