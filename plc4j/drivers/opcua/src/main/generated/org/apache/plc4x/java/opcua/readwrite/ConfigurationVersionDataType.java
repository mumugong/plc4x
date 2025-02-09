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
package org.apache.plc4x.java.opcua.readwrite;

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

public class ConfigurationVersionDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "14595";
  }

  // Properties.
  protected final long majorVersion;
  protected final long minorVersion;

  public ConfigurationVersionDataType(long majorVersion, long minorVersion) {
    super();
    this.majorVersion = majorVersion;
    this.minorVersion = minorVersion;
  }

  public long getMajorVersion() {
    return majorVersion;
  }

  public long getMinorVersion() {
    return minorVersion;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ConfigurationVersionDataType");

    // Simple Field (majorVersion)
    writeSimpleField("majorVersion", majorVersion, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (minorVersion)
    writeSimpleField("minorVersion", minorVersion, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("ConfigurationVersionDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ConfigurationVersionDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (majorVersion)
    lengthInBits += 32;

    // Simple field (minorVersion)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("ConfigurationVersionDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long majorVersion = readSimpleField("majorVersion", readUnsignedLong(readBuffer, 32));

    long minorVersion = readSimpleField("minorVersion", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("ConfigurationVersionDataType");
    // Create the instance
    return new ConfigurationVersionDataTypeBuilderImpl(majorVersion, minorVersion);
  }

  public static class ConfigurationVersionDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long majorVersion;
    private final long minorVersion;

    public ConfigurationVersionDataTypeBuilderImpl(long majorVersion, long minorVersion) {
      this.majorVersion = majorVersion;
      this.minorVersion = minorVersion;
    }

    public ConfigurationVersionDataType build() {
      ConfigurationVersionDataType configurationVersionDataType =
          new ConfigurationVersionDataType(majorVersion, minorVersion);
      return configurationVersionDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ConfigurationVersionDataType)) {
      return false;
    }
    ConfigurationVersionDataType that = (ConfigurationVersionDataType) o;
    return (getMajorVersion() == that.getMajorVersion())
        && (getMinorVersion() == that.getMinorVersion())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getMajorVersion(), getMinorVersion());
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
