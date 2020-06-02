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

#include "cotp_packet_tpdu_error.h"

plc4c_return_code plc4c_s7_read_write_cotp_packet_tpdu_error_parse(plc4c_read_buffer buf, uint16_t cotpLen, plc4c_s7_read_write_cotp_packet_tpdu_error** message) {
  uint16_t start_pos = plc4c_spi_read_get_pos(buf);
  uint16_t cur_pos;

  plc4c_s7_read_write_cotp_packet_tpdu_error* msg = malloc(sizeof(plc4c_s7_read_write_cotp_packet_tpdu_error));

  // Simple Field (destinationReference)
  uint16_t destinationReference = plc4c_spi_read_unsigned_int(buf, 16);
  msg.destination_reference = destinationReference;

  // Simple Field (rejectCause)
  uint8_t rejectCause = plc4c_spi_read_unsigned_short(buf, 8);
  msg.reject_cause = rejectCause;

  return OK;
}

plc4c_return_code plc4c_s7_read_write_cotp_packet_tpdu_error_serialize(plc4c_write_buffer buf, plc4c_s7_read_write_cotp_packet_tpdu_error* message) {
  return OK;
}
