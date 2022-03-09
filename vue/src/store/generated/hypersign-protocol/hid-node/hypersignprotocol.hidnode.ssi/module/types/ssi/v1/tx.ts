/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Did, SignInfo } from "../../ssi/v1/did";
import { Schema } from "../../ssi/v1/schema";

export const protobufPackage = "hypersignprotocol.hidnode.ssi";

export interface MsgCreateDID {
  didDocString: Did | undefined;
  signatures: SignInfo[];
  creator: string;
}

export interface MsgCreateDIDResponse {
  id: number;
}

export interface MsgUpdateDID {
  didDocString: Did | undefined;
  versionId: string;
  signatures: SignInfo[];
  creator: string;
}

export interface MsgUpdateDIDResponse {
  updateId: string;
}

export interface MsgCreateSchema {
  creator: string;
  schema: Schema | undefined;
  signatures: SignInfo[];
}

export interface MsgCreateSchemaResponse {
  id: number;
}

export interface MsgDeactivateDID {
  creator: string;
  didDocString: Did | undefined;
  versionId: string;
  signatures: SignInfo[];
}

export interface MsgDeactivateDIDResponse {
  id: number;
}

const baseMsgCreateDID: object = { creator: "" };

export const MsgCreateDID = {
  encode(message: MsgCreateDID, writer: Writer = Writer.create()): Writer {
    if (message.didDocString !== undefined) {
      Did.encode(message.didDocString, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.signatures) {
      SignInfo.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.creator !== "") {
      writer.uint32(26).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateDID {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateDID } as MsgCreateDID;
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

  fromJSON(object: any): MsgCreateDID {
    const message = { ...baseMsgCreateDID } as MsgCreateDID;
    message.signatures = [];
    if (object.didDocString !== undefined && object.didDocString !== null) {
      message.didDocString = Did.fromJSON(object.didDocString);
    } else {
      message.didDocString = undefined;
    }
    if (object.signatures !== undefined && object.signatures !== null) {
      for (const e of object.signatures) {
        message.signatures.push(SignInfo.fromJSON(e));
      }
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgCreateDID): unknown {
    const obj: any = {};
    message.didDocString !== undefined &&
      (obj.didDocString = message.didDocString
        ? Did.toJSON(message.didDocString)
        : undefined);
    if (message.signatures) {
      obj.signatures = message.signatures.map((e) =>
        e ? SignInfo.toJSON(e) : undefined
      );
    } else {
      obj.signatures = [];
    }
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateDID>): MsgCreateDID {
    const message = { ...baseMsgCreateDID } as MsgCreateDID;
    message.signatures = [];
    if (object.didDocString !== undefined && object.didDocString !== null) {
      message.didDocString = Did.fromPartial(object.didDocString);
    } else {
      message.didDocString = undefined;
    }
    if (object.signatures !== undefined && object.signatures !== null) {
      for (const e of object.signatures) {
        message.signatures.push(SignInfo.fromPartial(e));
      }
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgCreateDIDResponse: object = { id: 0 };

export const MsgCreateDIDResponse = {
  encode(
    message: MsgCreateDIDResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateDIDResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateDIDResponse } as MsgCreateDIDResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateDIDResponse {
    const message = { ...baseMsgCreateDIDResponse } as MsgCreateDIDResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateDIDResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateDIDResponse>): MsgCreateDIDResponse {
    const message = { ...baseMsgCreateDIDResponse } as MsgCreateDIDResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgUpdateDID: object = { versionId: "", creator: "" };

export const MsgUpdateDID = {
  encode(message: MsgUpdateDID, writer: Writer = Writer.create()): Writer {
    if (message.didDocString !== undefined) {
      Did.encode(message.didDocString, writer.uint32(10).fork()).ldelim();
    }
    if (message.versionId !== "") {
      writer.uint32(18).string(message.versionId);
    }
    for (const v of message.signatures) {
      SignInfo.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.creator !== "") {
      writer.uint32(34).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateDID {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateDID } as MsgUpdateDID;
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

  fromJSON(object: any): MsgUpdateDID {
    const message = { ...baseMsgUpdateDID } as MsgUpdateDID;
    message.signatures = [];
    if (object.didDocString !== undefined && object.didDocString !== null) {
      message.didDocString = Did.fromJSON(object.didDocString);
    } else {
      message.didDocString = undefined;
    }
    if (object.versionId !== undefined && object.versionId !== null) {
      message.versionId = String(object.versionId);
    } else {
      message.versionId = "";
    }
    if (object.signatures !== undefined && object.signatures !== null) {
      for (const e of object.signatures) {
        message.signatures.push(SignInfo.fromJSON(e));
      }
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateDID): unknown {
    const obj: any = {};
    message.didDocString !== undefined &&
      (obj.didDocString = message.didDocString
        ? Did.toJSON(message.didDocString)
        : undefined);
    message.versionId !== undefined && (obj.versionId = message.versionId);
    if (message.signatures) {
      obj.signatures = message.signatures.map((e) =>
        e ? SignInfo.toJSON(e) : undefined
      );
    } else {
      obj.signatures = [];
    }
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateDID>): MsgUpdateDID {
    const message = { ...baseMsgUpdateDID } as MsgUpdateDID;
    message.signatures = [];
    if (object.didDocString !== undefined && object.didDocString !== null) {
      message.didDocString = Did.fromPartial(object.didDocString);
    } else {
      message.didDocString = undefined;
    }
    if (object.versionId !== undefined && object.versionId !== null) {
      message.versionId = object.versionId;
    } else {
      message.versionId = "";
    }
    if (object.signatures !== undefined && object.signatures !== null) {
      for (const e of object.signatures) {
        message.signatures.push(SignInfo.fromPartial(e));
      }
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgUpdateDIDResponse: object = { updateId: "" };

export const MsgUpdateDIDResponse = {
  encode(
    message: MsgUpdateDIDResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.updateId !== "") {
      writer.uint32(10).string(message.updateId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateDIDResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateDIDResponse } as MsgUpdateDIDResponse;
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

  fromJSON(object: any): MsgUpdateDIDResponse {
    const message = { ...baseMsgUpdateDIDResponse } as MsgUpdateDIDResponse;
    if (object.updateId !== undefined && object.updateId !== null) {
      message.updateId = String(object.updateId);
    } else {
      message.updateId = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateDIDResponse): unknown {
    const obj: any = {};
    message.updateId !== undefined && (obj.updateId = message.updateId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateDIDResponse>): MsgUpdateDIDResponse {
    const message = { ...baseMsgUpdateDIDResponse } as MsgUpdateDIDResponse;
    if (object.updateId !== undefined && object.updateId !== null) {
      message.updateId = object.updateId;
    } else {
      message.updateId = "";
    }
    return message;
  },
};

const baseMsgCreateSchema: object = { creator: "" };

export const MsgCreateSchema = {
  encode(message: MsgCreateSchema, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.schema !== undefined) {
      Schema.encode(message.schema, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.signatures) {
      SignInfo.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateSchema {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateSchema } as MsgCreateSchema;
    message.signatures = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.schema = Schema.decode(reader, reader.uint32());
          break;
        case 3:
          message.signatures.push(SignInfo.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateSchema {
    const message = { ...baseMsgCreateSchema } as MsgCreateSchema;
    message.signatures = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.schema !== undefined && object.schema !== null) {
      message.schema = Schema.fromJSON(object.schema);
    } else {
      message.schema = undefined;
    }
    if (object.signatures !== undefined && object.signatures !== null) {
      for (const e of object.signatures) {
        message.signatures.push(SignInfo.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: MsgCreateSchema): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.schema !== undefined &&
      (obj.schema = message.schema ? Schema.toJSON(message.schema) : undefined);
    if (message.signatures) {
      obj.signatures = message.signatures.map((e) =>
        e ? SignInfo.toJSON(e) : undefined
      );
    } else {
      obj.signatures = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateSchema>): MsgCreateSchema {
    const message = { ...baseMsgCreateSchema } as MsgCreateSchema;
    message.signatures = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.schema !== undefined && object.schema !== null) {
      message.schema = Schema.fromPartial(object.schema);
    } else {
      message.schema = undefined;
    }
    if (object.signatures !== undefined && object.signatures !== null) {
      for (const e of object.signatures) {
        message.signatures.push(SignInfo.fromPartial(e));
      }
    }
    return message;
  },
};

const baseMsgCreateSchemaResponse: object = { id: 0 };

export const MsgCreateSchemaResponse = {
  encode(
    message: MsgCreateSchemaResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateSchemaResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateSchemaResponse,
    } as MsgCreateSchemaResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateSchemaResponse {
    const message = {
      ...baseMsgCreateSchemaResponse,
    } as MsgCreateSchemaResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateSchemaResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateSchemaResponse>
  ): MsgCreateSchemaResponse {
    const message = {
      ...baseMsgCreateSchemaResponse,
    } as MsgCreateSchemaResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgDeactivateDID: object = { creator: "", versionId: "" };

export const MsgDeactivateDID = {
  encode(message: MsgDeactivateDID, writer: Writer = Writer.create()): Writer {
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
      SignInfo.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeactivateDID {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeactivateDID } as MsgDeactivateDID;
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

  fromJSON(object: any): MsgDeactivateDID {
    const message = { ...baseMsgDeactivateDID } as MsgDeactivateDID;
    message.signatures = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.didDocString !== undefined && object.didDocString !== null) {
      message.didDocString = Did.fromJSON(object.didDocString);
    } else {
      message.didDocString = undefined;
    }
    if (object.versionId !== undefined && object.versionId !== null) {
      message.versionId = String(object.versionId);
    } else {
      message.versionId = "";
    }
    if (object.signatures !== undefined && object.signatures !== null) {
      for (const e of object.signatures) {
        message.signatures.push(SignInfo.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: MsgDeactivateDID): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.didDocString !== undefined &&
      (obj.didDocString = message.didDocString
        ? Did.toJSON(message.didDocString)
        : undefined);
    message.versionId !== undefined && (obj.versionId = message.versionId);
    if (message.signatures) {
      obj.signatures = message.signatures.map((e) =>
        e ? SignInfo.toJSON(e) : undefined
      );
    } else {
      obj.signatures = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeactivateDID>): MsgDeactivateDID {
    const message = { ...baseMsgDeactivateDID } as MsgDeactivateDID;
    message.signatures = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.didDocString !== undefined && object.didDocString !== null) {
      message.didDocString = Did.fromPartial(object.didDocString);
    } else {
      message.didDocString = undefined;
    }
    if (object.versionId !== undefined && object.versionId !== null) {
      message.versionId = object.versionId;
    } else {
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

const baseMsgDeactivateDIDResponse: object = { id: 0 };

export const MsgDeactivateDIDResponse = {
  encode(
    message: MsgDeactivateDIDResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeactivateDIDResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeactivateDIDResponse,
    } as MsgDeactivateDIDResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeactivateDIDResponse {
    const message = {
      ...baseMsgDeactivateDIDResponse,
    } as MsgDeactivateDIDResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgDeactivateDIDResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeactivateDIDResponse>
  ): MsgDeactivateDIDResponse {
    const message = {
      ...baseMsgDeactivateDIDResponse,
    } as MsgDeactivateDIDResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateDID(request: MsgCreateDID): Promise<MsgCreateDIDResponse>;
  UpdateDID(request: MsgUpdateDID): Promise<MsgUpdateDIDResponse>;
  CreateSchema(request: MsgCreateSchema): Promise<MsgCreateSchemaResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeactivateDID(request: MsgDeactivateDID): Promise<MsgDeactivateDIDResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateDID(request: MsgCreateDID): Promise<MsgCreateDIDResponse> {
    const data = MsgCreateDID.encode(request).finish();
    const promise = this.rpc.request(
      "hypersignprotocol.hidnode.ssi.Msg",
      "CreateDID",
      data
    );
    return promise.then((data) =>
      MsgCreateDIDResponse.decode(new Reader(data))
    );
  }

  UpdateDID(request: MsgUpdateDID): Promise<MsgUpdateDIDResponse> {
    const data = MsgUpdateDID.encode(request).finish();
    const promise = this.rpc.request(
      "hypersignprotocol.hidnode.ssi.Msg",
      "UpdateDID",
      data
    );
    return promise.then((data) =>
      MsgUpdateDIDResponse.decode(new Reader(data))
    );
  }

  CreateSchema(request: MsgCreateSchema): Promise<MsgCreateSchemaResponse> {
    const data = MsgCreateSchema.encode(request).finish();
    const promise = this.rpc.request(
      "hypersignprotocol.hidnode.ssi.Msg",
      "CreateSchema",
      data
    );
    return promise.then((data) =>
      MsgCreateSchemaResponse.decode(new Reader(data))
    );
  }

  DeactivateDID(request: MsgDeactivateDID): Promise<MsgDeactivateDIDResponse> {
    const data = MsgDeactivateDID.encode(request).finish();
    const promise = this.rpc.request(
      "hypersignprotocol.hidnode.ssi.Msg",
      "DeactivateDID",
      data
    );
    return promise.then((data) =>
      MsgDeactivateDIDResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
