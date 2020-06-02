/*
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
*/

#include <plc4c/spi/read_buffer.h>
#include <plc4c/spi/write_buffer.h>
#include <plc4c/spi/evaluation_helper.h>

#include "s7_parameter_user_data_item_cpu_functions.h"

plc4c_return_code plc4c_s7_read_write_s7_parameter_user_data_item_cpu_functions_parse(plc4c_read_buffer buf, plc4c_s7_read_write_s7_parameter_user_data_item_cpu_functions** message) {
  uint16_t start_pos = plc4c_spi_read_get_pos(buf);
  uint16_t cur_pos;

  plc4c_s7_read_write_s7_parameter_user_data_item_cpu_functions* msg = malloc(sizeof(plc4c_s7_read_write_s7_parameter_user_data_item_cpu_functions));

  // Implicit Field (itemLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
  uint8_t itemLength = plc4c_spi_read_unsigned_short(buf, 8);

  // Simple Field (method)
  uint8_t method = plc4c_spi_read_unsigned_short(buf, 8);
  msg.method = method;

  // Simple Field (cpuFunctionType)
  unsigned int cpuFunctionType = plc4c_spi_read_unsigned_byte(buf, 4);
  msg.cpu_function_type = cpuFunctionType;

  // Simple Field (cpuFunctionGroup)
  unsigned int cpuFunctionGroup = plc4c_spi_read_unsigned_byte(buf, 4);
  msg.cpu_function_group = cpuFunctionGroup;

  // Simple Field (cpuSubfunction)
  uint8_t cpuSubfunction = plc4c_spi_read_unsigned_short(buf, 8);
  msg.cpu_subfunction = cpuSubfunction;

  // Simple Field (sequenceNumber)
  uint8_t sequenceNumber = plc4c_spi_read_unsigned_short(buf, 8);
  msg.sequence_number = sequenceNumber;

  // Optional Field (dataUnitReferenceNumber) (Can be skipped, if a given expression evaluates to false)
  uint8_t dataUnitReferenceNumber = NULL;
  if((cpuFunctionType) == (8)) {
    dataUnitReferenceNumber = plc4c_spi_read_unsigned_short(buf, 8);
  }

  // Optional Field (lastDataUnit) (Can be skipped, if a given expression evaluates to false)
  uint8_t lastDataUnit = NULL;
  if((cpuFunctionType) == (8)) {
    lastDataUnit = plc4c_spi_read_unsigned_short(buf, 8);
  }

  // Optional Field (errorCode) (Can be skipped, if a given expression evaluates to false)
  uint16_t errorCode = NULL;
  if((cpuFunctionType) == (8)) {
    errorCode = plc4c_spi_read_unsigned_int(buf, 16);
  }

  return OK;
}

plc4c_return_code plc4c_s7_read_write_s7_parameter_user_data_item_cpu_functions_serialize(plc4c_write_buffer buf, plc4c_s7_read_write_s7_parameter_user_data_item_cpu_functions* message) {
  return OK;
}
