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
package org.apache.plc4x.java.test.readwrite;

import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.generation.ByteOrder;
import org.apache.plc4x.java.spi.generation.EvaluationHelper;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.ReadBuffer;
import org.apache.plc4x.java.spi.generation.SerializationException;
import org.apache.plc4x.java.spi.generation.WriteBuffer;
import org.apache.plc4x.java.spi.values.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

// Code generated by code-generation. DO NOT EDIT.

public class DataIOType {

  private static final Logger LOGGER = LoggerFactory.getLogger(DataIOType.class);

  public static PlcValue staticParse(ReadBuffer readBuffer, EnumType dataType)
      throws ParseException {
    if (EvaluationHelper.equals(dataType, EnumType.BOOL)) { // BOOL

      // Simple Field (value)
      Boolean value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readBit("");

      return new PlcBOOL(value);
    } else if (EvaluationHelper.equals(dataType, EnumType.UINT)) { // USINT

      // Simple Field (value)
      Short value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedShort("", 8);

      return new PlcUSINT(value);
    } else if (EvaluationHelper.equals(dataType, EnumType.INT)) { // UINT

      // Simple Field (value)
      Integer value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedInt("", 16);

      return new PlcUINT(value);
    }
    return null;
  }

  public static void staticSerialize(WriteBuffer writeBuffer, PlcValue _value, EnumType dataType)
      throws SerializationException {
    staticSerialize(writeBuffer, _value, dataType, ByteOrder.BIG_ENDIAN);
  }

  public static void staticSerialize(
      WriteBuffer writeBuffer, PlcValue _value, EnumType dataType, ByteOrder byteOrder)
      throws SerializationException {
    if (EvaluationHelper.equals(dataType, EnumType.BOOL)) { // BOOL
      // Simple Field (value)
      boolean value = (boolean) _value.getBoolean();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeBit("", (boolean) (value));
    } else if (EvaluationHelper.equals(dataType, EnumType.UINT)) { // USINT
      // Simple Field (value)
      short value = (short) _value.getShort();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedShort("", 8, ((Number) (value)).shortValue());
    } else if (EvaluationHelper.equals(dataType, EnumType.INT)) { // UINT
      // Simple Field (value)
      int value = (int) _value.getInt();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedInt("", 16, ((Number) (value)).intValue());
    }
  }

  public static int getLengthInBytes(PlcValue _value, EnumType dataType) {
    return (int) Math.ceil((float) getLengthInBits(_value, dataType) / 8.0);
  }

  public static int getLengthInBits(PlcValue _value, EnumType dataType) {
    int sizeInBits = 0;
    if (EvaluationHelper.equals(dataType, EnumType.BOOL)) { // BOOL
      // Simple Field (value)
      sizeInBits += 1;
    } else if (EvaluationHelper.equals(dataType, EnumType.UINT)) { // USINT
      // Simple Field (value)
      sizeInBits += 8;
    } else if (EvaluationHelper.equals(dataType, EnumType.INT)) { // UINT
      // Simple Field (value)
      sizeInBits += 16;
    }
    return sizeInBits;
  }
}
