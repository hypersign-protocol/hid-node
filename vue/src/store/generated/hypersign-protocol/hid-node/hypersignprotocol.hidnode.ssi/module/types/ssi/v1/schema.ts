/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "hypersignprotocol.hidnode.ssi";

export interface Schema {
  type: string;
  modelVersion: string;
  id: string;
  name: string;
  author: string;
  authored: string;
  schema: SchemaProperty | undefined;
}

export interface SchemaProperty {
  schema: string;
  description: string;
  type: string;
  properties: string;
  required: string[];
  additionalProperties: boolean;
}

const baseSchema: object = {
  type: "",
  modelVersion: "",
  id: "",
  name: "",
  author: "",
  authored: "",
};

export const Schema = {
  encode(message: Schema, writer: Writer = Writer.create()): Writer {
    if (message.type !== "") {
      writer.uint32(10).string(message.type);
    }
    if (message.modelVersion !== "") {
      writer.uint32(18).string(message.modelVersion);
    }
    if (message.id !== "") {
      writer.uint32(26).string(message.id);
    }
    if (message.name !== "") {
      writer.uint32(34).string(message.name);
    }
    if (message.author !== "") {
      writer.uint32(42).string(message.author);
    }
    if (message.authored !== "") {
      writer.uint32(50).string(message.authored);
    }
    if (message.schema !== undefined) {
      SchemaProperty.encode(message.schema, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Schema {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSchema } as Schema;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = reader.string();
          break;
        case 2:
          message.modelVersion = reader.string();
          break;
        case 3:
          message.id = reader.string();
          break;
        case 4:
          message.name = reader.string();
          break;
        case 5:
          message.author = reader.string();
          break;
        case 6:
          message.authored = reader.string();
          break;
        case 7:
          message.schema = SchemaProperty.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Schema {
    const message = { ...baseSchema } as Schema;
    if (object.type !== undefined && object.type !== null) {
      message.type = String(object.type);
    } else {
      message.type = "";
    }
    if (object.modelVersion !== undefined && object.modelVersion !== null) {
      message.modelVersion = String(object.modelVersion);
    } else {
      message.modelVersion = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.author !== undefined && object.author !== null) {
      message.author = String(object.author);
    } else {
      message.author = "";
    }
    if (object.authored !== undefined && object.authored !== null) {
      message.authored = String(object.authored);
    } else {
      message.authored = "";
    }
    if (object.schema !== undefined && object.schema !== null) {
      message.schema = SchemaProperty.fromJSON(object.schema);
    } else {
      message.schema = undefined;
    }
    return message;
  },

  toJSON(message: Schema): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = message.type);
    message.modelVersion !== undefined &&
      (obj.modelVersion = message.modelVersion);
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.author !== undefined && (obj.author = message.author);
    message.authored !== undefined && (obj.authored = message.authored);
    message.schema !== undefined &&
      (obj.schema = message.schema
        ? SchemaProperty.toJSON(message.schema)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Schema>): Schema {
    const message = { ...baseSchema } as Schema;
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = "";
    }
    if (object.modelVersion !== undefined && object.modelVersion !== null) {
      message.modelVersion = object.modelVersion;
    } else {
      message.modelVersion = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.author !== undefined && object.author !== null) {
      message.author = object.author;
    } else {
      message.author = "";
    }
    if (object.authored !== undefined && object.authored !== null) {
      message.authored = object.authored;
    } else {
      message.authored = "";
    }
    if (object.schema !== undefined && object.schema !== null) {
      message.schema = SchemaProperty.fromPartial(object.schema);
    } else {
      message.schema = undefined;
    }
    return message;
  },
};

const baseSchemaProperty: object = {
  schema: "",
  description: "",
  type: "",
  properties: "",
  required: "",
  additionalProperties: false,
};

export const SchemaProperty = {
  encode(message: SchemaProperty, writer: Writer = Writer.create()): Writer {
    if (message.schema !== "") {
      writer.uint32(10).string(message.schema);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.type !== "") {
      writer.uint32(26).string(message.type);
    }
    if (message.properties !== "") {
      writer.uint32(34).string(message.properties);
    }
    for (const v of message.required) {
      writer.uint32(42).string(v!);
    }
    if (message.additionalProperties === true) {
      writer.uint32(48).bool(message.additionalProperties);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SchemaProperty {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSchemaProperty } as SchemaProperty;
    message.required = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.schema = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.type = reader.string();
          break;
        case 4:
          message.properties = reader.string();
          break;
        case 5:
          message.required.push(reader.string());
          break;
        case 6:
          message.additionalProperties = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SchemaProperty {
    const message = { ...baseSchemaProperty } as SchemaProperty;
    message.required = [];
    if (object.schema !== undefined && object.schema !== null) {
      message.schema = String(object.schema);
    } else {
      message.schema = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = String(object.type);
    } else {
      message.type = "";
    }
    if (object.properties !== undefined && object.properties !== null) {
      message.properties = String(object.properties);
    } else {
      message.properties = "";
    }
    if (object.required !== undefined && object.required !== null) {
      for (const e of object.required) {
        message.required.push(String(e));
      }
    }
    if (
      object.additionalProperties !== undefined &&
      object.additionalProperties !== null
    ) {
      message.additionalProperties = Boolean(object.additionalProperties);
    } else {
      message.additionalProperties = false;
    }
    return message;
  },

  toJSON(message: SchemaProperty): unknown {
    const obj: any = {};
    message.schema !== undefined && (obj.schema = message.schema);
    message.description !== undefined &&
      (obj.description = message.description);
    message.type !== undefined && (obj.type = message.type);
    message.properties !== undefined && (obj.properties = message.properties);
    if (message.required) {
      obj.required = message.required.map((e) => e);
    } else {
      obj.required = [];
    }
    message.additionalProperties !== undefined &&
      (obj.additionalProperties = message.additionalProperties);
    return obj;
  },

  fromPartial(object: DeepPartial<SchemaProperty>): SchemaProperty {
    const message = { ...baseSchemaProperty } as SchemaProperty;
    message.required = [];
    if (object.schema !== undefined && object.schema !== null) {
      message.schema = object.schema;
    } else {
      message.schema = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = "";
    }
    if (object.properties !== undefined && object.properties !== null) {
      message.properties = object.properties;
    } else {
      message.properties = "";
    }
    if (object.required !== undefined && object.required !== null) {
      for (const e of object.required) {
        message.required.push(e);
      }
    }
    if (
      object.additionalProperties !== undefined &&
      object.additionalProperties !== null
    ) {
      message.additionalProperties = object.additionalProperties;
    } else {
      message.additionalProperties = false;
    }
    return message;
  },
};

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
