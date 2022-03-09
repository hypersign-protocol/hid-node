/* eslint-disable */
const { Reader, util, configure, Writer } = require("protobufjs/minimal");
const Long = require("long");
const { Did, SignInfo } = require("./did");

const protobufPackage = "hypersignprotocol.hidnode.ssi";

const baseMsgCreateDID = { creator: "" };

const MsgCreateDID = {
    encode(message, writer = Writer.create()) {
        if (message.didDocString !== undefined) {
            Did.encode(message.didDocString, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.signatures) {
            SignInfo.encode(v, writer.uint32(18).fork()).ldelim();
        }
        if (message.creator !== "") {
            writer.uint32(26).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateDID };
        message.signatures = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.didDocString = Did.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.signatures.push(SignInfo.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.creator = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCreateDID };
        message.signatures = [];
        if (object.didDocString !== undefined && object.didDocString !== null) {
            message.didDocString = Did.fromJSON(object.didDocString);
        }
        else {
            message.didDocString = undefined;
        }
        if (object.signatures !== undefined && object.signatures !== null) {
            for (const e of object.signatures) {
                message.signatures.push(SignInfo.fromJSON(e));
            }
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.didDocString !== undefined &&
            (obj.didDocString = message.didDocString
                ? Did.toJSON(message.didDocString)
                : undefined);
        if (message.signatures) {
            obj.signatures = message.signatures.map((e) => e ? SignInfo.toJSON(e) : undefined);
        }
        else {
            obj.signatures = [];
        }
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateDID };
        message.signatures = [];
        if (object.didDocString !== undefined && object.didDocString !== null) {
            message.didDocString = Did.fromPartial(object.didDocString);
        }
        else {
            message.didDocString = undefined;
        }
        if (object.signatures !== undefined && object.signatures !== null) {
            for (const e of object.signatures) {
                message.signatures.push(SignInfo.fromPartial(e));
            }
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        return message;
    },
};
const baseMsgCreateDIDResponse = { id: 0 };
const MsgCreateDIDResponse = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateDIDResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCreateDIDResponse };
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateDIDResponse };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    },
};
const baseMsgUpdateDID = { versionId: "", creator: "" };
const MsgUpdateDID = {
    encode(message, writer = Writer.create()) {
        if (message.didDocString !== undefined) {
            Did.encode(message.didDocString, writer.uint32(10).fork()).ldelim();
        }
        if (message.versionId !== "") {
            writer.uint32(18).string(message.versionId);
        }
        for (const v of message.signatures) {
            SignInfo.encode(v, writer.uint32(26).fork()).ldelim();
        }
        if (message.creator !== "") {
            writer.uint32(34).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateDID };
        message.signatures = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.didDocString = Did.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.versionId = reader.string();
                    break;
                case 3:
                    message.signatures.push(SignInfo.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.creator = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgUpdateDID };
        message.signatures = [];
        if (object.didDocString !== undefined && object.didDocString !== null) {
            message.didDocString = Did.fromJSON(object.didDocString);
        }
        else {
            message.didDocString = undefined;
        }
        if (object.versionId !== undefined && object.versionId !== null) {
            message.versionId = String(object.versionId);
        }
        else {
            message.versionId = "";
        }
        if (object.signatures !== undefined && object.signatures !== null) {
            for (const e of object.signatures) {
                message.signatures.push(SignInfo.fromJSON(e));
            }
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.didDocString !== undefined &&
            (obj.didDocString = message.didDocString
                ? Did.toJSON(message.didDocString)
                : undefined);
        message.versionId !== undefined && (obj.versionId = message.versionId);
        if (message.signatures) {
            obj.signatures = message.signatures.map((e) => e ? SignInfo.toJSON(e) : undefined);
        }
        else {
            obj.signatures = [];
        }
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateDID };
        message.signatures = [];
        if (object.didDocString !== undefined && object.didDocString !== null) {
            message.didDocString = Did.fromPartial(object.didDocString);
        }
        else {
            message.didDocString = undefined;
        }
        if (object.versionId !== undefined && object.versionId !== null) {
            message.versionId = object.versionId;
        }
        else {
            message.versionId = "";
        }
        if (object.signatures !== undefined && object.signatures !== null) {
            for (const e of object.signatures) {
                message.signatures.push(SignInfo.fromPartial(e));
            }
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        return message;
    },
};
const baseMsgUpdateDIDResponse = { updateId: "" };
const MsgUpdateDIDResponse = {
    encode(message, writer = Writer.create()) {
        if (message.updateId !== "") {
            writer.uint32(10).string(message.updateId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateDIDResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.updateId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgUpdateDIDResponse };
        if (object.updateId !== undefined && object.updateId !== null) {
            message.updateId = String(object.updateId);
        }
        else {
            message.updateId = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.updateId !== undefined && (obj.updateId = message.updateId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateDIDResponse };
        if (object.updateId !== undefined && object.updateId !== null) {
            message.updateId = object.updateId;
        }
        else {
            message.updateId = "";
        }
        return message;
    },
};

const baseMsgDeactivateDID = { creator: "", versionId: "" };
const MsgDeactivateDID = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.didDocString !== undefined) {
            Did.encode(message.didDocString, writer.uint32(18).fork()).ldelim();
        }
        if (message.versionId !== "") {
            writer.uint32(26).string(message.versionId);
        }
        for (const v of message.signatures) {
            SignInfo.encode(v, writer.uint32(34).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDeactivateDID };
        message.signatures = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.didDocString = Did.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.versionId = reader.string();
                    break;
                case 4:
                    message.signatures.push(SignInfo.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgDeactivateDID };
        message.signatures = [];
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.didDocString !== undefined && object.didDocString !== null) {
            message.didDocString = Did.fromJSON(object.didDocString);
        }
        else {
            message.didDocString = undefined;
        }
        if (object.versionId !== undefined && object.versionId !== null) {
            message.versionId = String(object.versionId);
        }
        else {
            message.versionId = "";
        }
        if (object.signatures !== undefined && object.signatures !== null) {
            for (const e of object.signatures) {
                message.signatures.push(SignInfo.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.didDocString !== undefined &&
            (obj.didDocString = message.didDocString
                ? Did.toJSON(message.didDocString)
                : undefined);
        message.versionId !== undefined && (obj.versionId = message.versionId);
        if (message.signatures) {
            obj.signatures = message.signatures.map((e) => e ? SignInfo.toJSON(e) : undefined);
        }
        else {
            obj.signatures = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgDeactivateDID };
        message.signatures = [];
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.didDocString !== undefined && object.didDocString !== null) {
            message.didDocString = Did.fromPartial(object.didDocString);
        }
        else {
            message.didDocString = undefined;
        }
        if (object.versionId !== undefined && object.versionId !== null) {
            message.versionId = object.versionId;
        }
        else {
            message.versionId = "";
        }
        if (object.signatures !== undefined && object.signatures !== null) {
            for (const e of object.signatures) {
                message.signatures.push(SignInfo.fromPartial(e));
            }
        }
        return message;
    },
};
const baseMsgDeactivateDIDResponse = { id: 0 };
const MsgDeactivateDIDResponse = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgDeactivateDIDResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseMsgDeactivateDIDResponse,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseMsgDeactivateDIDResponse,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    },
};
class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateDID(request) {
        const data = MsgCreateDID.encode(request).finish();
        const promise = this.rpc.request("hypersignprotocol.hidnode.ssi.Msg", "CreateDID", data);
        return promise.then((data) => MsgCreateDIDResponse.decode(new Reader(data)));
    }
    UpdateDID(request) {
        const data = MsgUpdateDID.encode(request).finish();
        const promise = this.rpc.request("hypersignprotocol.hidnode.ssi.Msg", "UpdateDID", data);
        return promise.then((data) => MsgUpdateDIDResponse.decode(new Reader(data)));
    }
    DeactivateDID(request) {
        const data = MsgDeactivateDID.encode(request).finish();
        const promise = this.rpc.request("hypersignprotocol.hidnode.ssi.Msg", "DeactivateDID", data);
        return promise.then((data) => MsgDeactivateDIDResponse.decode(new Reader(data)));
    }
}
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}

module.exports = {
    protobufPackage,
    MsgCreateDID,
    MsgCreateDIDResponse,
    MsgUpdateDID,
    MsgUpdateDIDResponse,
    MsgDeactivateDID,
    MsgDeactivateDIDResponse,
    MsgClientImpl
}