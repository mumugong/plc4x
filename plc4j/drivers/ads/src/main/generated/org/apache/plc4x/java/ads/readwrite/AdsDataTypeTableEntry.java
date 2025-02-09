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
package org.apache.plc4x.java.ads.readwrite;

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

public class AdsDataTypeTableEntry implements Message {

  // Constant values.
  public static final Short DATATYPENAMETERMINATOR = 0x00;
  public static final Short SIMPLETYPENAMETERMINATOR = 0x00;
  public static final Short COMMENTTERMINATOR = 0x00;

  // Properties.
  protected final long entryLength;
  protected final long version;
  protected final long hashValue;
  protected final long typeHashValue;
  protected final long size;
  protected final long offset;
  protected final long dataType;
  protected final long flags;
  protected final int arrayDimensions;
  protected final int numChildren;
  protected final String dataTypeName;
  protected final String simpleTypeName;
  protected final String comment;
  protected final List<AdsDataTypeArrayInfo> arrayInfo;
  protected final List<AdsDataTypeTableChildEntry> children;
  protected final byte[] rest;

  public AdsDataTypeTableEntry(
      long entryLength,
      long version,
      long hashValue,
      long typeHashValue,
      long size,
      long offset,
      long dataType,
      long flags,
      int arrayDimensions,
      int numChildren,
      String dataTypeName,
      String simpleTypeName,
      String comment,
      List<AdsDataTypeArrayInfo> arrayInfo,
      List<AdsDataTypeTableChildEntry> children,
      byte[] rest) {
    super();
    this.entryLength = entryLength;
    this.version = version;
    this.hashValue = hashValue;
    this.typeHashValue = typeHashValue;
    this.size = size;
    this.offset = offset;
    this.dataType = dataType;
    this.flags = flags;
    this.arrayDimensions = arrayDimensions;
    this.numChildren = numChildren;
    this.dataTypeName = dataTypeName;
    this.simpleTypeName = simpleTypeName;
    this.comment = comment;
    this.arrayInfo = arrayInfo;
    this.children = children;
    this.rest = rest;
  }

  public long getEntryLength() {
    return entryLength;
  }

  public long getVersion() {
    return version;
  }

  public long getHashValue() {
    return hashValue;
  }

  public long getTypeHashValue() {
    return typeHashValue;
  }

  public long getSize() {
    return size;
  }

  public long getOffset() {
    return offset;
  }

  public long getDataType() {
    return dataType;
  }

  public long getFlags() {
    return flags;
  }

  public int getArrayDimensions() {
    return arrayDimensions;
  }

  public int getNumChildren() {
    return numChildren;
  }

  public String getDataTypeName() {
    return dataTypeName;
  }

  public String getSimpleTypeName() {
    return simpleTypeName;
  }

  public String getComment() {
    return comment;
  }

  public List<AdsDataTypeArrayInfo> getArrayInfo() {
    return arrayInfo;
  }

  public List<AdsDataTypeTableChildEntry> getChildren() {
    return children;
  }

  public byte[] getRest() {
    return rest;
  }

  public short getDataTypeNameTerminator() {
    return DATATYPENAMETERMINATOR;
  }

  public short getSimpleTypeNameTerminator() {
    return SIMPLETYPENAMETERMINATOR;
  }

  public short getCommentTerminator() {
    return COMMENTTERMINATOR;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("AdsDataTypeTableEntry");

    // Simple Field (entryLength)
    writeSimpleField(
        "entryLength",
        entryLength,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (version)
    writeSimpleField(
        "version",
        version,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (hashValue)
    writeSimpleField(
        "hashValue",
        hashValue,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (typeHashValue)
    writeSimpleField(
        "typeHashValue",
        typeHashValue,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (size)
    writeSimpleField(
        "size",
        size,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (offset)
    writeSimpleField(
        "offset",
        offset,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (dataType)
    writeSimpleField(
        "dataType",
        dataType,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (flags)
    writeSimpleField(
        "flags",
        flags,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Implicit Field (dataTypeNameLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int dataTypeNameLength = (int) (STR_LEN(getDataTypeName()));
    writeImplicitField(
        "dataTypeNameLength",
        dataTypeNameLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Implicit Field (simpleTypeNameLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int simpleTypeNameLength = (int) (STR_LEN(getSimpleTypeName()));
    writeImplicitField(
        "simpleTypeNameLength",
        simpleTypeNameLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Implicit Field (commentLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int commentLength = (int) (STR_LEN(getComment()));
    writeImplicitField(
        "commentLength",
        commentLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (arrayDimensions)
    writeSimpleField(
        "arrayDimensions",
        arrayDimensions,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (numChildren)
    writeSimpleField(
        "numChildren",
        numChildren,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (dataTypeName)
    writeSimpleField(
        "dataTypeName",
        dataTypeName,
        writeString(writeBuffer, (dataTypeNameLength) * (8)),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Const Field (dataTypeNameTerminator)
    writeConstField(
        "dataTypeNameTerminator",
        DATATYPENAMETERMINATOR,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (simpleTypeName)
    writeSimpleField(
        "simpleTypeName",
        simpleTypeName,
        writeString(writeBuffer, (simpleTypeNameLength) * (8)),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Const Field (simpleTypeNameTerminator)
    writeConstField(
        "simpleTypeNameTerminator",
        SIMPLETYPENAMETERMINATOR,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (comment)
    writeSimpleField(
        "comment",
        comment,
        writeString(writeBuffer, (commentLength) * (8)),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Const Field (commentTerminator)
    writeConstField(
        "commentTerminator",
        COMMENTTERMINATOR,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Array Field (arrayInfo)
    writeComplexTypeArrayField(
        "arrayInfo", arrayInfo, writeBuffer, WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Array Field (children)
    writeComplexTypeArrayField(
        "children", children, writeBuffer, WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Array Field (rest)
    writeByteArrayField(
        "rest",
        rest,
        writeByteArray(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("AdsDataTypeTableEntry");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    AdsDataTypeTableEntry _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (entryLength)
    lengthInBits += 32;

    // Simple field (version)
    lengthInBits += 32;

    // Simple field (hashValue)
    lengthInBits += 32;

    // Simple field (typeHashValue)
    lengthInBits += 32;

    // Simple field (size)
    lengthInBits += 32;

    // Simple field (offset)
    lengthInBits += 32;

    // Simple field (dataType)
    lengthInBits += 32;

    // Simple field (flags)
    lengthInBits += 32;

    // Implicit Field (dataTypeNameLength)
    lengthInBits += 16;

    // Implicit Field (simpleTypeNameLength)
    lengthInBits += 16;

    // Implicit Field (commentLength)
    lengthInBits += 16;

    // Simple field (arrayDimensions)
    lengthInBits += 16;

    // Simple field (numChildren)
    lengthInBits += 16;

    // Simple field (dataTypeName)
    lengthInBits += (STR_LEN(getDataTypeName())) * (8);

    // Const Field (dataTypeNameTerminator)
    lengthInBits += 8;

    // Simple field (simpleTypeName)
    lengthInBits += (STR_LEN(getSimpleTypeName())) * (8);

    // Const Field (simpleTypeNameTerminator)
    lengthInBits += 8;

    // Simple field (comment)
    lengthInBits += (STR_LEN(getComment())) * (8);

    // Const Field (commentTerminator)
    lengthInBits += 8;

    // Array field
    if (arrayInfo != null) {
      int i = 0;
      for (AdsDataTypeArrayInfo element : arrayInfo) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= arrayInfo.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Array field
    if (children != null) {
      int i = 0;
      for (AdsDataTypeTableChildEntry element : children) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= children.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Array field
    if (rest != null) {
      lengthInBits += 8 * rest.length;
    }

    return lengthInBits;
  }

  public static AdsDataTypeTableEntry staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static AdsDataTypeTableEntry staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("AdsDataTypeTableEntry");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long entryLength =
        readSimpleField(
            "entryLength",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    long version =
        readSimpleField(
            "version",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    long hashValue =
        readSimpleField(
            "hashValue",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    long typeHashValue =
        readSimpleField(
            "typeHashValue",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    long size =
        readSimpleField(
            "size",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    long offset =
        readSimpleField(
            "offset",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    long dataType =
        readSimpleField(
            "dataType",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    long flags =
        readSimpleField(
            "flags",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    int dataTypeNameLength =
        readImplicitField(
            "dataTypeNameLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    int simpleTypeNameLength =
        readImplicitField(
            "simpleTypeNameLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    int commentLength =
        readImplicitField(
            "commentLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    int arrayDimensions =
        readSimpleField(
            "arrayDimensions",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    int numChildren =
        readSimpleField(
            "numChildren",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    String dataTypeName =
        readSimpleField(
            "dataTypeName",
            readString(readBuffer, (dataTypeNameLength) * (8)),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    short dataTypeNameTerminator =
        readConstField(
            "dataTypeNameTerminator",
            readUnsignedShort(readBuffer, 8),
            AdsDataTypeTableEntry.DATATYPENAMETERMINATOR,
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    String simpleTypeName =
        readSimpleField(
            "simpleTypeName",
            readString(readBuffer, (simpleTypeNameLength) * (8)),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    short simpleTypeNameTerminator =
        readConstField(
            "simpleTypeNameTerminator",
            readUnsignedShort(readBuffer, 8),
            AdsDataTypeTableEntry.SIMPLETYPENAMETERMINATOR,
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    String comment =
        readSimpleField(
            "comment",
            readString(readBuffer, (commentLength) * (8)),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    short commentTerminator =
        readConstField(
            "commentTerminator",
            readUnsignedShort(readBuffer, 8),
            AdsDataTypeTableEntry.COMMENTTERMINATOR,
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    List<AdsDataTypeArrayInfo> arrayInfo =
        readCountArrayField(
            "arrayInfo",
            new DataReaderComplexDefault<>(
                () -> AdsDataTypeArrayInfo.staticParse(readBuffer), readBuffer),
            arrayDimensions,
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    List<AdsDataTypeTableChildEntry> children =
        readCountArrayField(
            "children",
            new DataReaderComplexDefault<>(
                () -> AdsDataTypeTableChildEntry.staticParse(readBuffer), readBuffer),
            numChildren,
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    byte[] rest =
        readBuffer.readByteArray(
            "rest",
            Math.toIntExact((entryLength) - ((positionAware.getPos() - startPos))),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("AdsDataTypeTableEntry");
    // Create the instance
    AdsDataTypeTableEntry _adsDataTypeTableEntry;
    _adsDataTypeTableEntry =
        new AdsDataTypeTableEntry(
            entryLength,
            version,
            hashValue,
            typeHashValue,
            size,
            offset,
            dataType,
            flags,
            arrayDimensions,
            numChildren,
            dataTypeName,
            simpleTypeName,
            comment,
            arrayInfo,
            children,
            rest);
    return _adsDataTypeTableEntry;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AdsDataTypeTableEntry)) {
      return false;
    }
    AdsDataTypeTableEntry that = (AdsDataTypeTableEntry) o;
    return (getEntryLength() == that.getEntryLength())
        && (getVersion() == that.getVersion())
        && (getHashValue() == that.getHashValue())
        && (getTypeHashValue() == that.getTypeHashValue())
        && (getSize() == that.getSize())
        && (getOffset() == that.getOffset())
        && (getDataType() == that.getDataType())
        && (getFlags() == that.getFlags())
        && (getArrayDimensions() == that.getArrayDimensions())
        && (getNumChildren() == that.getNumChildren())
        && (getDataTypeName() == that.getDataTypeName())
        && (getSimpleTypeName() == that.getSimpleTypeName())
        && (getComment() == that.getComment())
        && (getArrayInfo() == that.getArrayInfo())
        && (getChildren() == that.getChildren())
        && (getRest() == that.getRest())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getEntryLength(),
        getVersion(),
        getHashValue(),
        getTypeHashValue(),
        getSize(),
        getOffset(),
        getDataType(),
        getFlags(),
        getArrayDimensions(),
        getNumChildren(),
        getDataTypeName(),
        getSimpleTypeName(),
        getComment(),
        getArrayInfo(),
        getChildren(),
        getRest());
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
