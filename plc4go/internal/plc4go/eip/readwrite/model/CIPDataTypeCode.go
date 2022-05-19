/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// CIPDataTypeCode is an enum
type CIPDataTypeCode uint16

type ICIPDataTypeCode interface {
	Size() uint8
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	CIPDataTypeCode_BOOL       CIPDataTypeCode = 0x00C1
	CIPDataTypeCode_SINT       CIPDataTypeCode = 0x00C2
	CIPDataTypeCode_INT        CIPDataTypeCode = 0x00C3
	CIPDataTypeCode_DINT       CIPDataTypeCode = 0x00C4
	CIPDataTypeCode_LINT       CIPDataTypeCode = 0x00C5
	CIPDataTypeCode_REAL       CIPDataTypeCode = 0x00CA
	CIPDataTypeCode_DWORD      CIPDataTypeCode = 0x00D3
	CIPDataTypeCode_STRUCTURED CIPDataTypeCode = 0x02A0
	CIPDataTypeCode_STRING     CIPDataTypeCode = 0x02A0
	CIPDataTypeCode_STRING36   CIPDataTypeCode = 0x02A0
)

var CIPDataTypeCodeValues []CIPDataTypeCode

func init() {
	_ = errors.New
	CIPDataTypeCodeValues = []CIPDataTypeCode{
		CIPDataTypeCode_BOOL,
		CIPDataTypeCode_SINT,
		CIPDataTypeCode_INT,
		CIPDataTypeCode_DINT,
		CIPDataTypeCode_LINT,
		CIPDataTypeCode_REAL,
		CIPDataTypeCode_DWORD,
		CIPDataTypeCode_STRUCTURED,
		CIPDataTypeCode_STRING,
		CIPDataTypeCode_STRING36,
	}
}

func (e CIPDataTypeCode) Size() uint8 {
	switch e {
	case 0x00C1:
		{ /* '0X00C1' */
			return 1
		}
	case 0x00C2:
		{ /* '0X00C2' */
			return 1
		}
	case 0x00C3:
		{ /* '0X00C3' */
			return 2
		}
	case 0x00C4:
		{ /* '0X00C4' */
			return 4
		}
	case 0x00C5:
		{ /* '0X00C5' */
			return 8
		}
	case 0x00CA:
		{ /* '0X00CA' */
			return 4
		}
	case 0x00D3:
		{ /* '0X00D3' */
			return 4
		}
	case 0x02A0:
		{ /* '0X02A0' */
			return 88
		}
	default:
		{
			return 0
		}
	}
}

func CIPDataTypeCodeFirstEnumForFieldSize(value uint8) (CIPDataTypeCode, error) {
	for _, sizeValue := range CIPDataTypeCodeValues {
		if sizeValue.Size() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing Size not found", value)
}
func CIPDataTypeCodeByValue(value uint16) CIPDataTypeCode {
	switch value {
	case 0x00C1:
		return CIPDataTypeCode_BOOL
	case 0x00C2:
		return CIPDataTypeCode_SINT
	case 0x00C3:
		return CIPDataTypeCode_INT
	case 0x00C4:
		return CIPDataTypeCode_DINT
	case 0x00C5:
		return CIPDataTypeCode_LINT
	case 0x00CA:
		return CIPDataTypeCode_REAL
	case 0x00D3:
		return CIPDataTypeCode_DWORD
	case 0x02A0:
		return CIPDataTypeCode_STRUCTURED
	}
	return 0
}

func CIPDataTypeCodeByName(value string) CIPDataTypeCode {
	switch value {
	case "BOOL":
		return CIPDataTypeCode_BOOL
	case "SINT":
		return CIPDataTypeCode_SINT
	case "INT":
		return CIPDataTypeCode_INT
	case "DINT":
		return CIPDataTypeCode_DINT
	case "LINT":
		return CIPDataTypeCode_LINT
	case "REAL":
		return CIPDataTypeCode_REAL
	case "DWORD":
		return CIPDataTypeCode_DWORD
	case "STRUCTURED":
		return CIPDataTypeCode_STRUCTURED
	}
	return 0
}

func CIPDataTypeCodeKnows(value uint16) bool {
	for _, typeValue := range CIPDataTypeCodeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastCIPDataTypeCode(structType interface{}) CIPDataTypeCode {
	castFunc := func(typ interface{}) CIPDataTypeCode {
		if sCIPDataTypeCode, ok := typ.(CIPDataTypeCode); ok {
			return sCIPDataTypeCode
		}
		return 0
	}
	return castFunc(structType)
}

func (m CIPDataTypeCode) GetLengthInBits() uint16 {
	return 16
}

func (m CIPDataTypeCode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CIPDataTypeCodeParse(readBuffer utils.ReadBuffer) (CIPDataTypeCode, error) {
	val, err := readBuffer.ReadUint16("CIPDataTypeCode", 16)
	if err != nil {
		return 0, nil
	}
	return CIPDataTypeCodeByValue(val), nil
}

func (e CIPDataTypeCode) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("CIPDataTypeCode", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e CIPDataTypeCode) name() string {
	switch e {
	case CIPDataTypeCode_BOOL:
		return "BOOL"
	case CIPDataTypeCode_SINT:
		return "SINT"
	case CIPDataTypeCode_INT:
		return "INT"
	case CIPDataTypeCode_DINT:
		return "DINT"
	case CIPDataTypeCode_LINT:
		return "LINT"
	case CIPDataTypeCode_REAL:
		return "REAL"
	case CIPDataTypeCode_DWORD:
		return "DWORD"
	case CIPDataTypeCode_STRUCTURED:
		return "STRUCTURED"
	}
	return ""
}

func (e CIPDataTypeCode) String() string {
	return e.name()
}
