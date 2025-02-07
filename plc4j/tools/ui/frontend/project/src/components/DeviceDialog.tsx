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

import {Device} from "../generated/plc4j-tools-ui-frontend.ts";
import {Dialog} from "primereact/dialog";
import {Button} from "primereact/button";
import {InputText} from "primereact/inputtext";
import {DataTable} from "primereact/datatable";
import {Column} from "primereact/column";

interface ConnectionDialogProps {
    device: Device
    visible: boolean

    onSave: (device:Device) => void
    onCancel: () => void
}

interface TableEntry {
    key: string
    value: string
}

type Dictionary = { [index: string]: string }

function mapToTableEntry(map : Dictionary):TableEntry[] {
    let tableEntries = [] as TableEntry[]
    for(const key in map) {
        const value = map[key]
        tableEntries = [...tableEntries, {key: key, value: value}]
    }
    return tableEntries
}

export default function DeviceDialog({device, visible, onSave, onCancel}: ConnectionDialogProps) {
    function handleSave() {
        onSave(device)
    }
    function handleCancel() {
        onCancel()
    }

    return (
        <Dialog visible={visible} modal style={{width: '60rem'}} draggable={true} resizable={true} onHide={() => {
        }}>
            <div className="formgrid grid">
                <label htmlFor="name" className="col-12 mb-2 md:col-2 md:mb-0">Device Name</label>
                <div className="col-12 md:col-10">
                    <InputText id="name" value={device.name}/>
                </div>
            </div>
            <div className="formgrid grid">
                <label htmlFor="protocol" className="col-12 mb-2 md:col-2 md:mb-0">Protocol</label>
                <div className="col-12 md:col-10">
                    <InputText id="protocol" value={device.protocolCode}/>
                </div>
            </div>
            <div className="formgrid grid">
                <label htmlFor="transportCode" className="col-12 mb-2 md:col-2 md:mb-0">Transport Type</label>
                <div className="col-12 md:col-10">
                    <InputText id="transportCode" value={device.transportCode}/>
                </div>
            </div>
            <div className="formgrid grid">
                <label htmlFor="transportUrl" className="col-12 mb-2 md:col-2 md:mb-0">Transport URL</label>
                <div className="col-12 md:col-10">
                    <InputText id="transportUrl" value={device.transportUrl}/>
                </div>
            </div>
            <div className="formgrid grid">
                <label htmlFor="options" className="col-12 mb-2 md:col-2 md:mb-0">Options</label>
                <div className="col-12 md:col-10">
                    <DataTable id="options" value={mapToTableEntry(device.options)} tableStyle={{minWidth: '50rem'}}>
                        <Column field="key" header="Name"/>
                        <Column field="value" header="Value"/>
                    </DataTable>
                </div>
            </div>
            <div className="formgrid grid">
                <label htmlFor="attributes" className="col-12 mb-2 md:col-2 md:mb-0">Attributes</label>
                <div className="col-12 md:col-10">
                    <DataTable id="attributes" value={mapToTableEntry(device.attributes)} tableStyle={{minWidth: '50rem'}}>
                        <Column field="key" header="Name"/>
                        <Column field="value" header="Value"/>
                    </DataTable>
                </div>
            </div>
            <div className="formgrid grid">
                <div className="field col">
                    <Button label={"Save"} onClick={handleSave}/>
                    <Button label={"Cancel"} onClick={handleCancel}/>
                </div>
            </div>
        </Dialog>
    )
}