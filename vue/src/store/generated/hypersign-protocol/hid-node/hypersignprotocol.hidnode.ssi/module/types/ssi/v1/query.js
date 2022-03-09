/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../../ssi/v1/params";
import { Schema } from "../../ssi/v1/schema";
import { PageRequest, PageResponse, } from "../../cosmos/base/query/v1beta1/pagination";
import { Did, Metadata, DidResolveMeta } from "../../ssi/v1/did";
export const protobufPackage = "hypersignprotocol.hidnode.ssi";
const baseQueryParamsRequest = {};
export const QueryParamsRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryParamsRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseQueryParamsRequest };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryParamsRequest };
        return message;
    },
};
const baseQueryParamsResponse = {};
export const QueryParamsResponse = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryParamsResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryParamsResponse };
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryParamsResponse };
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        return message;
    },
};
const baseQueryGetSchemaRequest = { schemaId: "" };
export const QueryGetSchemaRequest = {
    encode(message, writer = Writer.create()) {
        if (message.schemaId !== "") {
            writer.uint32(10).string(message.schemaId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetSchemaRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.schemaId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetSchemaRequest };
        if (object.schemaId !== undefined && object.schemaId !== null) {
            message.schemaId = String(object.schemaId);
        }
        else {
            message.schemaId = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.schemaId !== undefined && (obj.schemaId = message.schemaId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetSchemaRequest };
        if (object.schemaId !== undefined && object.schemaId !== null) {
            message.schemaId = object.schemaId;
        }
        else {
            message.schemaId = "";
        }
        return message;
    },
};
const baseQueryGetSchemaResponse = {};
export const QueryGetSchemaResponse = {
    encode(message, writer = Writer.create()) {
        if (message.schema !== undefined) {
            Schema.encode(message.schema, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetSchemaResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.schema = Schema.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetSchemaResponse };
        if (object.schema !== undefined && object.schema !== null) {
            message.schema = Schema.fromJSON(object.schema);
        }
        else {
            message.schema = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.schema !== undefined &&
            (obj.schema = message.schema ? Schema.toJSON(message.schema) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetSchemaResponse };
        if (object.schema !== undefined && object.schema !== null) {
            message.schema = Schema.fromPartial(object.schema);
        }
        else {
            message.schema = undefined;
        }
        return message;
    },
};
const baseQuerySchemasRequest = {};
export const QuerySchemasRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQuerySchemasRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQuerySchemasRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQuerySchemasRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQuerySchemasResponse = {};
export const QuerySchemasResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.schemaList) {
            Schema.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQuerySchemasResponse };
        message.schemaList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.schemaList.push(Schema.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQuerySchemasResponse };
        message.schemaList = [];
        if (object.schemaList !== undefined && object.schemaList !== null) {
            for (const e of object.schemaList) {
                message.schemaList.push(Schema.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.schemaList) {
            obj.schemaList = message.schemaList.map((e) => e ? Schema.toJSON(e) : undefined);
        }
        else {
            obj.schemaList = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQuerySchemasResponse };
        message.schemaList = [];
        if (object.schemaList !== undefined && object.schemaList !== null) {
            for (const e of object.schemaList) {
                message.schemaList.push(Schema.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQuerySchemaCountRequest = {};
export const QuerySchemaCountRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQuerySchemaCountRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = {
            ...baseQuerySchemaCountRequest,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseQuerySchemaCountRequest,
        };
        return message;
    },
};
const baseQuerySchemaCountResponse = { count: 0 };
export const QuerySchemaCountResponse = {
    encode(message, writer = Writer.create()) {
        if (message.count !== 0) {
            writer.uint32(8).uint64(message.count);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQuerySchemaCountResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.count = longToNumber(reader.uint64());
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
            ...baseQuerySchemaCountResponse,
        };
        if (object.count !== undefined && object.count !== null) {
            message.count = Number(object.count);
        }
        else {
            message.count = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.count !== undefined && (obj.count = message.count);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQuerySchemaCountResponse,
        };
        if (object.count !== undefined && object.count !== null) {
            message.count = object.count;
        }
        else {
            message.count = 0;
        }
        return message;
    },
};
const baseQueryGetDidDocByIdRequest = { didId: "", versionId: "" };
export const QueryGetDidDocByIdRequest = {
    encode(message, writer = Writer.create()) {
        if (message.didId !== "") {
            writer.uint32(10).string(message.didId);
        }
        if (message.versionId !== "") {
            writer.uint32(18).string(message.versionId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetDidDocByIdRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.didId = reader.string();
                    break;
                case 2:
                    message.versionId = reader.string();
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
            ...baseQueryGetDidDocByIdRequest,
        };
        if (object.didId !== undefined && object.didId !== null) {
            message.didId = String(object.didId);
        }
        else {
            message.didId = "";
        }
        if (object.versionId !== undefined && object.versionId !== null) {
            message.versionId = String(object.versionId);
        }
        else {
            message.versionId = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.didId !== undefined && (obj.didId = message.didId);
        message.versionId !== undefined && (obj.versionId = message.versionId);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetDidDocByIdRequest,
        };
        if (object.didId !== undefined && object.didId !== null) {
            message.didId = object.didId;
        }
        else {
            message.didId = "";
        }
        if (object.versionId !== undefined && object.versionId !== null) {
            message.versionId = object.versionId;
        }
        else {
            message.versionId = "";
        }
        return message;
    },
};
const baseQueryGetDidDocByIdResponse = { AtContext: "" };
export const QueryGetDidDocByIdResponse = {
    encode(message, writer = Writer.create()) {
        if (message.AtContext !== "") {
            writer.uint32(10).string(message.AtContext);
        }
        if (message.didDocument !== undefined) {
            Did.encode(message.didDocument, writer.uint32(18).fork()).ldelim();
        }
        if (message.didDocumentMetadata !== undefined) {
            Metadata.encode(message.didDocumentMetadata, writer.uint32(26).fork()).ldelim();
        }
        if (message.didResolutionMetadata !== undefined) {
            DidResolveMeta.encode(message.didResolutionMetadata, writer.uint32(34).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetDidDocByIdResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.AtContext = reader.string();
                    break;
                case 2:
                    message.didDocument = Did.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.didDocumentMetadata = Metadata.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.didResolutionMetadata = DidResolveMeta.decode(reader, reader.uint32());
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
            ...baseQueryGetDidDocByIdResponse,
        };
        if (object.AtContext !== undefined && object.AtContext !== null) {
            message.AtContext = String(object.AtContext);
        }
        else {
            message.AtContext = "";
        }
        if (object.didDocument !== undefined && object.didDocument !== null) {
            message.didDocument = Did.fromJSON(object.didDocument);
        }
        else {
            message.didDocument = undefined;
        }
        if (object.didDocumentMetadata !== undefined &&
            object.didDocumentMetadata !== null) {
            message.didDocumentMetadata = Metadata.fromJSON(object.didDocumentMetadata);
        }
        else {
            message.didDocumentMetadata = undefined;
        }
        if (object.didResolutionMetadata !== undefined &&
            object.didResolutionMetadata !== null) {
            message.didResolutionMetadata = DidResolveMeta.fromJSON(object.didResolutionMetadata);
        }
        else {
            message.didResolutionMetadata = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.AtContext !== undefined && (obj.AtContext = message.AtContext);
        message.didDocument !== undefined &&
            (obj.didDocument = message.didDocument
                ? Did.toJSON(message.didDocument)
                : undefined);
        message.didDocumentMetadata !== undefined &&
            (obj.didDocumentMetadata = message.didDocumentMetadata
                ? Metadata.toJSON(message.didDocumentMetadata)
                : undefined);
        message.didResolutionMetadata !== undefined &&
            (obj.didResolutionMetadata = message.didResolutionMetadata
                ? DidResolveMeta.toJSON(message.didResolutionMetadata)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetDidDocByIdResponse,
        };
        if (object.AtContext !== undefined && object.AtContext !== null) {
            message.AtContext = object.AtContext;
        }
        else {
            message.AtContext = "";
        }
        if (object.didDocument !== undefined && object.didDocument !== null) {
            message.didDocument = Did.fromPartial(object.didDocument);
        }
        else {
            message.didDocument = undefined;
        }
        if (object.didDocumentMetadata !== undefined &&
            object.didDocumentMetadata !== null) {
            message.didDocumentMetadata = Metadata.fromPartial(object.didDocumentMetadata);
        }
        else {
            message.didDocumentMetadata = undefined;
        }
        if (object.didResolutionMetadata !== undefined &&
            object.didResolutionMetadata !== null) {
            message.didResolutionMetadata = DidResolveMeta.fromPartial(object.didResolutionMetadata);
        }
        else {
            message.didResolutionMetadata = undefined;
        }
        return message;
    },
};
const baseQueryDidParamRequest = { count: false };
export const QueryDidParamRequest = {
    encode(message, writer = Writer.create()) {
        if (message.count === true) {
            writer.uint32(8).bool(message.count);
        }
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryDidParamRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.count = reader.bool();
                    break;
                case 2:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryDidParamRequest };
        if (object.count !== undefined && object.count !== null) {
            message.count = Boolean(object.count);
        }
        else {
            message.count = false;
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.count !== undefined && (obj.count = message.count);
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryDidParamRequest };
        if (object.count !== undefined && object.count !== null) {
            message.count = object.count;
        }
        else {
            message.count = false;
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryDidParamResponse = { totalDidCount: 0 };
export const QueryDidParamResponse = {
    encode(message, writer = Writer.create()) {
        if (message.totalDidCount !== 0) {
            writer.uint32(8).uint64(message.totalDidCount);
        }
        for (const v of message.didDocList) {
            DidResolutionResponse.encode(v, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryDidParamResponse };
        message.didDocList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.totalDidCount = longToNumber(reader.uint64());
                    break;
                case 2:
                    message.didDocList.push(DidResolutionResponse.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryDidParamResponse };
        message.didDocList = [];
        if (object.totalDidCount !== undefined && object.totalDidCount !== null) {
            message.totalDidCount = Number(object.totalDidCount);
        }
        else {
            message.totalDidCount = 0;
        }
        if (object.didDocList !== undefined && object.didDocList !== null) {
            for (const e of object.didDocList) {
                message.didDocList.push(DidResolutionResponse.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.totalDidCount !== undefined &&
            (obj.totalDidCount = message.totalDidCount);
        if (message.didDocList) {
            obj.didDocList = message.didDocList.map((e) => e ? DidResolutionResponse.toJSON(e) : undefined);
        }
        else {
            obj.didDocList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryDidParamResponse };
        message.didDocList = [];
        if (object.totalDidCount !== undefined && object.totalDidCount !== null) {
            message.totalDidCount = object.totalDidCount;
        }
        else {
            message.totalDidCount = 0;
        }
        if (object.didDocList !== undefined && object.didDocList !== null) {
            for (const e of object.didDocList) {
                message.didDocList.push(DidResolutionResponse.fromPartial(e));
            }
        }
        return message;
    },
};
const baseDidResolutionResponse = { AtContext: "" };
export const DidResolutionResponse = {
    encode(message, writer = Writer.create()) {
        if (message.AtContext !== "") {
            writer.uint32(10).string(message.AtContext);
        }
        if (message.didDocument !== undefined) {
            Did.encode(message.didDocument, writer.uint32(18).fork()).ldelim();
        }
        if (message.didDocumentMetadata !== undefined) {
            Metadata.encode(message.didDocumentMetadata, writer.uint32(26).fork()).ldelim();
        }
        if (message.didResolutionMetadata !== undefined) {
            DidResolveMeta.encode(message.didResolutionMetadata, writer.uint32(34).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseDidResolutionResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.AtContext = reader.string();
                    break;
                case 2:
                    message.didDocument = Did.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.didDocumentMetadata = Metadata.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.didResolutionMetadata = DidResolveMeta.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseDidResolutionResponse };
        if (object.AtContext !== undefined && object.AtContext !== null) {
            message.AtContext = String(object.AtContext);
        }
        else {
            message.AtContext = "";
        }
        if (object.didDocument !== undefined && object.didDocument !== null) {
            message.didDocument = Did.fromJSON(object.didDocument);
        }
        else {
            message.didDocument = undefined;
        }
        if (object.didDocumentMetadata !== undefined &&
            object.didDocumentMetadata !== null) {
            message.didDocumentMetadata = Metadata.fromJSON(object.didDocumentMetadata);
        }
        else {
            message.didDocumentMetadata = undefined;
        }
        if (object.didResolutionMetadata !== undefined &&
            object.didResolutionMetadata !== null) {
            message.didResolutionMetadata = DidResolveMeta.fromJSON(object.didResolutionMetadata);
        }
        else {
            message.didResolutionMetadata = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.AtContext !== undefined && (obj.AtContext = message.AtContext);
        message.didDocument !== undefined &&
            (obj.didDocument = message.didDocument
                ? Did.toJSON(message.didDocument)
                : undefined);
        message.didDocumentMetadata !== undefined &&
            (obj.didDocumentMetadata = message.didDocumentMetadata
                ? Metadata.toJSON(message.didDocumentMetadata)
                : undefined);
        message.didResolutionMetadata !== undefined &&
            (obj.didResolutionMetadata = message.didResolutionMetadata
                ? DidResolveMeta.toJSON(message.didResolutionMetadata)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseDidResolutionResponse };
        if (object.AtContext !== undefined && object.AtContext !== null) {
            message.AtContext = object.AtContext;
        }
        else {
            message.AtContext = "";
        }
        if (object.didDocument !== undefined && object.didDocument !== null) {
            message.didDocument = Did.fromPartial(object.didDocument);
        }
        else {
            message.didDocument = undefined;
        }
        if (object.didDocumentMetadata !== undefined &&
            object.didDocumentMetadata !== null) {
            message.didDocumentMetadata = Metadata.fromPartial(object.didDocumentMetadata);
        }
        else {
            message.didDocumentMetadata = undefined;
        }
        if (object.didResolutionMetadata !== undefined &&
            object.didResolutionMetadata !== null) {
            message.didResolutionMetadata = DidResolveMeta.fromPartial(object.didResolutionMetadata);
        }
        else {
            message.didResolutionMetadata = undefined;
        }
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Params(request) {
        const data = QueryParamsRequest.encode(request).finish();
        const promise = this.rpc.request("hypersignprotocol.hidnode.ssi.Query", "Params", data);
        return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
    }
    GetSchema(request) {
        const data = QueryGetSchemaRequest.encode(request).finish();
        const promise = this.rpc.request("hypersignprotocol.hidnode.ssi.Query", "GetSchema", data);
        return promise.then((data) => QueryGetSchemaResponse.decode(new Reader(data)));
    }
    Schemas(request) {
        const data = QuerySchemasRequest.encode(request).finish();
        const promise = this.rpc.request("hypersignprotocol.hidnode.ssi.Query", "Schemas", data);
        return promise.then((data) => QuerySchemasResponse.decode(new Reader(data)));
    }
    SchemaCount(request) {
        const data = QuerySchemaCountRequest.encode(request).finish();
        const promise = this.rpc.request("hypersignprotocol.hidnode.ssi.Query", "SchemaCount", data);
        return promise.then((data) => QuerySchemaCountResponse.decode(new Reader(data)));
    }
    ResolveDid(request) {
        const data = QueryGetDidDocByIdRequest.encode(request).finish();
        const promise = this.rpc.request("hypersignprotocol.hidnode.ssi.Query", "ResolveDid", data);
        return promise.then((data) => QueryGetDidDocByIdResponse.decode(new Reader(data)));
    }
    DidParam(request) {
        const data = QueryDidParamRequest.encode(request).finish();
        const promise = this.rpc.request("hypersignprotocol.hidnode.ssi.Query", "DidParam", data);
        return promise.then((data) => QueryDidParamResponse.decode(new Reader(data)));
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
