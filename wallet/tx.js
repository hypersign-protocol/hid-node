/* eslint-disable */
const { Reader, util, configure, Writer } = require("protobufjs/minimal");
const Long = require("long");

const protobufPackage = "hypersignprotocol.hidnode.did";
const baseMsgCreateDID = {
    creator: "",
    did: "",
    didDocString: "",
    createdAt: "",
};
const MsgCreateDID = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator);
        }
        if (message.did !== "") {
            writer.uint32(18).string(message.did);
        }
        if (message.didDocString !== "") {
            writer.uint32(26).string(message.didDocString);
        }
        if (message.createdAt !== "") {
            writer.uint32(34).string(message.createdAt);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateDID };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.did = reader.string();
                    break;
                case 3:
                    message.didDocString = reader.string();
                    break;
                case 4:
                    message.createdAt = reader.string();
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
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = String(object.did);
        }
        else {
            message.did = "";
        }
        if (object.didDocString !== undefined && object.didDocString !== null) {
            message.didDocString = String(object.didDocString);
        }
        else {
            message.didDocString = "";
        }
        if (object.createdAt !== undefined && object.createdAt !== null) {
            message.createdAt = String(object.createdAt);
        }
        else {
            message.createdAt = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.did !== undefined && (obj.did = message.did);
        message.didDocString !== undefined &&
            (obj.didDocString = message.didDocString);
        message.createdAt !== undefined && (obj.createdAt = message.createdAt);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateDID };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.did !== undefined && object.did !== null) {
            message.did = object.did;
        }
        else {
            message.did = "";
        }
        if (object.didDocString !== undefined && object.didDocString !== null) {
            message.didDocString = object.didDocString;
        }
        else {
            message.didDocString = "";
        }
        if (object.createdAt !== undefined && object.createdAt !== null) {
            message.createdAt = object.createdAt;
        }
        else {
            message.createdAt = "";
        }
        return message;
    },
};
const baseMsgCDResponse = { id: 0 };
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
class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateDID(request) {
        const data = MsgCreateDID.encode(request).finish();
        const promise = this.rpc.request("hypersignprotocol.hidnode.did.Msg", "CreateDID", data);
        return promise.then((data) => MsgCreateDIDResponse.decode(new Reader(data)));
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
    MsgClientImpl
}