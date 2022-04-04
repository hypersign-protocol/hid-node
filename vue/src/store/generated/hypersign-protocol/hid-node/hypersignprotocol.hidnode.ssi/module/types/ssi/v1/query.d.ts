import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../../ssi/v1/params";
import { Schema } from "../../ssi/v1/schema";
import { PageRequest } from "../../cosmos/base/query/v1beta1/pagination";
import { Did, Metadata, DidResolveMeta } from "../../ssi/v1/did";
export declare const protobufPackage = "hypersignprotocol.hidnode.ssi";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryGetSchemaRequest {
    schemaId: string;
}
export interface QueryGetSchemaResponse {
    schema: Schema[];
}
export interface QuerySchemaParamRequest {
    pagination: PageRequest | undefined;
}
export interface QuerySchemaParamResponse {
    totalCount: number;
    schemaList: Schema[];
}
export interface QueryGetDidDocByIdRequest {
    didId: string;
    versionId: string;
}
export interface QueryGetDidDocByIdResponse {
    AtContext: string;
    didDocument: Did | undefined;
    didDocumentMetadata: Metadata | undefined;
    didResolutionMetadata: DidResolveMeta | undefined;
}
export interface QueryDidParamRequest {
    count: boolean;
    pagination: PageRequest | undefined;
}
export interface QueryDidParamResponse {
    totalDidCount: number;
    didDocList: DidResolutionResponse[];
}
export interface DidResolutionResponse {
    AtContext: string;
    didDocument: Did | undefined;
    didDocumentMetadata: Metadata | undefined;
    didResolutionMetadata: DidResolveMeta | undefined;
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryGetSchemaRequest: {
    encode(message: QueryGetSchemaRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSchemaRequest;
    fromJSON(object: any): QueryGetSchemaRequest;
    toJSON(message: QueryGetSchemaRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetSchemaRequest>): QueryGetSchemaRequest;
};
export declare const QueryGetSchemaResponse: {
    encode(message: QueryGetSchemaResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSchemaResponse;
    fromJSON(object: any): QueryGetSchemaResponse;
    toJSON(message: QueryGetSchemaResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetSchemaResponse>): QueryGetSchemaResponse;
};
export declare const QuerySchemaParamRequest: {
    encode(message: QuerySchemaParamRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QuerySchemaParamRequest;
    fromJSON(object: any): QuerySchemaParamRequest;
    toJSON(message: QuerySchemaParamRequest): unknown;
    fromPartial(object: DeepPartial<QuerySchemaParamRequest>): QuerySchemaParamRequest;
};
export declare const QuerySchemaParamResponse: {
    encode(message: QuerySchemaParamResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QuerySchemaParamResponse;
    fromJSON(object: any): QuerySchemaParamResponse;
    toJSON(message: QuerySchemaParamResponse): unknown;
    fromPartial(object: DeepPartial<QuerySchemaParamResponse>): QuerySchemaParamResponse;
};
export declare const QueryGetDidDocByIdRequest: {
    encode(message: QueryGetDidDocByIdRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetDidDocByIdRequest;
    fromJSON(object: any): QueryGetDidDocByIdRequest;
    toJSON(message: QueryGetDidDocByIdRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetDidDocByIdRequest>): QueryGetDidDocByIdRequest;
};
export declare const QueryGetDidDocByIdResponse: {
    encode(message: QueryGetDidDocByIdResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetDidDocByIdResponse;
    fromJSON(object: any): QueryGetDidDocByIdResponse;
    toJSON(message: QueryGetDidDocByIdResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetDidDocByIdResponse>): QueryGetDidDocByIdResponse;
};
export declare const QueryDidParamRequest: {
    encode(message: QueryDidParamRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryDidParamRequest;
    fromJSON(object: any): QueryDidParamRequest;
    toJSON(message: QueryDidParamRequest): unknown;
    fromPartial(object: DeepPartial<QueryDidParamRequest>): QueryDidParamRequest;
};
export declare const QueryDidParamResponse: {
    encode(message: QueryDidParamResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryDidParamResponse;
    fromJSON(object: any): QueryDidParamResponse;
    toJSON(message: QueryDidParamResponse): unknown;
    fromPartial(object: DeepPartial<QueryDidParamResponse>): QueryDidParamResponse;
};
export declare const DidResolutionResponse: {
    encode(message: DidResolutionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): DidResolutionResponse;
    fromJSON(object: any): DidResolutionResponse;
    toJSON(message: DidResolutionResponse): unknown;
    fromPartial(object: DeepPartial<DidResolutionResponse>): DidResolutionResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a list of GetSchema items. */
    GetSchema(request: QueryGetSchemaRequest): Promise<QueryGetSchemaResponse>;
    /** Schema Param */
    SchemaParam(request: QuerySchemaParamRequest): Promise<QuerySchemaParamResponse>;
    /** Resolve DID */
    ResolveDid(request: QueryGetDidDocByIdRequest): Promise<QueryGetDidDocByIdResponse>;
    /** Did Param */
    DidParam(request: QueryDidParamRequest): Promise<QueryDidParamResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    GetSchema(request: QueryGetSchemaRequest): Promise<QueryGetSchemaResponse>;
    SchemaParam(request: QuerySchemaParamRequest): Promise<QuerySchemaParamResponse>;
    ResolveDid(request: QueryGetDidDocByIdRequest): Promise<QueryGetDidDocByIdResponse>;
    DidParam(request: QueryDidParamRequest): Promise<QueryDidParamResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
