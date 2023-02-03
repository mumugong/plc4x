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
package org.apache.plc4x.java.profinet.readwrite;

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

public class TlvTimeToLive extends LldpUnit implements Message {

  // Accessors for discriminator values.
  public TlvType getTlvId() {
    return TlvType.TIME_TO_LIVE;
  }

  // Properties.
  protected final int tlvTimeToLiveUnit;

  public TlvTimeToLive(int tlvIdLength, int tlvTimeToLiveUnit) {
    super(tlvIdLength);
    this.tlvTimeToLiveUnit = tlvTimeToLiveUnit;
  }

  public int getTlvTimeToLiveUnit() {
    return tlvTimeToLiveUnit;
  }

  @Override
  protected void serializeLldpUnitChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("TlvTimeToLive");

    // Simple Field (tlvTimeToLiveUnit)
    writeSimpleField("tlvTimeToLiveUnit", tlvTimeToLiveUnit, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("TlvTimeToLive");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    TlvTimeToLive _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (tlvTimeToLiveUnit)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static LldpUnitBuilder staticParseLldpUnitBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("TlvTimeToLive");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int tlvTimeToLiveUnit = readSimpleField("tlvTimeToLiveUnit", readUnsignedInt(readBuffer, 16));

    readBuffer.closeContext("TlvTimeToLive");
    // Create the instance
    return new TlvTimeToLiveBuilderImpl(tlvTimeToLiveUnit);
  }

  public static class TlvTimeToLiveBuilderImpl implements LldpUnit.LldpUnitBuilder {
    private final int tlvTimeToLiveUnit;

    public TlvTimeToLiveBuilderImpl(int tlvTimeToLiveUnit) {
      this.tlvTimeToLiveUnit = tlvTimeToLiveUnit;
    }

    public TlvTimeToLive build(int tlvIdLength) {
      TlvTimeToLive tlvTimeToLive = new TlvTimeToLive(tlvIdLength, tlvTimeToLiveUnit);
      return tlvTimeToLive;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TlvTimeToLive)) {
      return false;
    }
    TlvTimeToLive that = (TlvTimeToLive) o;
    return (getTlvTimeToLiveUnit() == that.getTlvTimeToLiveUnit()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getTlvTimeToLiveUnit());
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
