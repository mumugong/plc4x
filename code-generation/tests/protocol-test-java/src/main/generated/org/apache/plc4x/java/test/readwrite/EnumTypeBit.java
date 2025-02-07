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

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum EnumTypeBit {
  TRUE((boolean) true),
  FALSE((boolean) false);
  private static final Map<Boolean, EnumTypeBit> map;

  static {
    map = new HashMap<>();
    for (EnumTypeBit value : EnumTypeBit.values()) {
      map.put((boolean) value.getValue(), value);
    }
  }

  private final boolean value;

  EnumTypeBit(boolean value) {
    this.value = value;
  }

  public boolean getValue() {
    return value;
  }

  public static EnumTypeBit enumForValue(boolean value) {
    return map.get(value);
  }

  public static Boolean isDefined(boolean value) {
    return map.containsKey(value);
  }
}
