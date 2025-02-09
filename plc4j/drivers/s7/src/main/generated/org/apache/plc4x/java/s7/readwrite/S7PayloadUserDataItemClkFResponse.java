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
package org.apache.plc4x.java.s7.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class S7PayloadUserDataItemClkFResponse extends S7PayloadUserDataItem implements Message {

  // Accessors for discriminator values.
  public Byte getCpuFunctionGroup() {
    return (byte) 0x07;
  }

  public Byte getCpuFunctionType() {
    return (byte) 0x08;
  }

  public Short getCpuSubfunction() {
    return (short) 0x03;
  }

  // Properties.
  protected final short Reserved;
  protected final short Year1;
  protected final DateAndTime TimeStamp;

  public S7PayloadUserDataItemClkFResponse(
      DataTransportErrorCode returnCode,
      DataTransportSize transportSize,
      int dataLength,
      short Reserved,
      short Year1,
      DateAndTime TimeStamp) {
    super(returnCode, transportSize, dataLength);
    this.Reserved = Reserved;
    this.Year1 = Year1;
    this.TimeStamp = TimeStamp;
  }

  public short getReserved() {
    return Reserved;
  }

  public short getYear1() {
    return Year1;
  }

  public DateAndTime getTimeStamp() {
    return TimeStamp;
  }

  @Override
  protected void serializeS7PayloadUserDataItemChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("S7PayloadUserDataItemClkFResponse");

    // Simple Field (Reserved)
    writeSimpleField("Reserved", Reserved, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (Year1)
    writeSimpleField("Year1", Year1, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (TimeStamp)
    writeSimpleField("TimeStamp", TimeStamp, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("S7PayloadUserDataItemClkFResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    S7PayloadUserDataItemClkFResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (Reserved)
    lengthInBits += 8;

    // Simple field (Year1)
    lengthInBits += 8;

    // Simple field (TimeStamp)
    lengthInBits += TimeStamp.getLengthInBits();

    return lengthInBits;
  }

  public static S7PayloadUserDataItemBuilder staticParseS7PayloadUserDataItemBuilder(
      ReadBuffer readBuffer,
      Integer dataLength,
      Byte cpuFunctionGroup,
      Byte cpuFunctionType,
      Short cpuSubfunction)
      throws ParseException {
    readBuffer.pullContext("S7PayloadUserDataItemClkFResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short Reserved = readSimpleField("Reserved", readUnsignedShort(readBuffer, 8));

    short Year1 = readSimpleField("Year1", readUnsignedShort(readBuffer, 8));

    DateAndTime TimeStamp =
        readSimpleField(
            "TimeStamp",
            new DataReaderComplexDefault<>(() -> DateAndTime.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("S7PayloadUserDataItemClkFResponse");
    // Create the instance
    return new S7PayloadUserDataItemClkFResponseBuilderImpl(Reserved, Year1, TimeStamp);
  }

  public static class S7PayloadUserDataItemClkFResponseBuilderImpl
      implements S7PayloadUserDataItem.S7PayloadUserDataItemBuilder {
    private final short Reserved;
    private final short Year1;
    private final DateAndTime TimeStamp;

    public S7PayloadUserDataItemClkFResponseBuilderImpl(
        short Reserved, short Year1, DateAndTime TimeStamp) {
      this.Reserved = Reserved;
      this.Year1 = Year1;
      this.TimeStamp = TimeStamp;
    }

    public S7PayloadUserDataItemClkFResponse build(
        DataTransportErrorCode returnCode, DataTransportSize transportSize, int dataLength) {
      S7PayloadUserDataItemClkFResponse s7PayloadUserDataItemClkFResponse =
          new S7PayloadUserDataItemClkFResponse(
              returnCode, transportSize, dataLength, Reserved, Year1, TimeStamp);
      return s7PayloadUserDataItemClkFResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7PayloadUserDataItemClkFResponse)) {
      return false;
    }
    S7PayloadUserDataItemClkFResponse that = (S7PayloadUserDataItemClkFResponse) o;
    return (getReserved() == that.getReserved())
        && (getYear1() == that.getYear1())
        && (getTimeStamp() == that.getTimeStamp())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getReserved(), getYear1(), getTimeStamp());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
