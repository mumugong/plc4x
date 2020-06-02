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

#include "cotp_parameter.h"

plc4c_return_code plc4c_s7_read_write_cotp_parameter_parse(plc4c_read_buffer buf, uint8_t rest, plc4c_s7_read_write_cotp_parameter** message) {
  uint16_t start_pos = plc4c_spi_read_get_pos(buf);
  uint16_t cur_pos;

  plc4c_s7_read_write_cotp_parameter* msg = malloc(sizeof(plc4c_s7_read_write_cotp_parameter));

  // Discriminator Field (parameterType) (Used as input to a switch field)
  uint8_t parameterType = plc4c_spi_read_unsigned_short(buf, 8);

  // Implicit Field (parameterLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
  uint8_t parameterLength = plc4c_spi_read_unsigned_short(buf, 8);

  // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
  if(plc4c_spi_evaluation_helper_equals(parameterType, 0xC0)) {
    plc4c_s7_read_write_cotp_parameter_tpdu_size_parse(buf, msg, rest);
  } else 
  if(plc4c_spi_evaluation_helper_equals(parameterType, 0xC1)) {
    plc4c_s7_read_write_cotp_parameter_calling_tsap_parse(buf, msg, rest);
  } else 
  if(plc4c_spi_evaluation_helper_equals(parameterType, 0xC2)) {
    plc4c_s7_read_write_cotp_parameter_called_tsap_parse(buf, msg, rest);
  } else 
  if(plc4c_spi_evaluation_helper_equals(parameterType, 0xC3)) {
    plc4c_s7_read_write_cotp_parameter_checksum_parse(buf, msg, rest);
  } else 
  if(plc4c_spi_evaluation_helper_equals(parameterType, 0xE0)) {
    plc4c_s7_read_write_cotp_parameter_disconnect_additional_information_parse(buf, msg, rest);
  }

  return OK;
}

plc4c_return_code plc4c_s7_read_write_cotp_parameter_serialize(plc4c_write_buffer buf, plc4c_s7_read_write_cotp_parameter* message) {
  return OK;
}
