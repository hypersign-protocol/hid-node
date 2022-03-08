import { Reader, Writer } from "protobufjs/minimal";
import { Did, SignInfo } from "../../ssi/v1/did";
import { Schema } from "../../ssi/v1/schema";
export declare const protobufPackage = "hypersignprotocol.hidnode.ssi";
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
export declare const MsgCreateDID: {
    encode(message: MsgCreateDID, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateDID;
    fromJSON(object: any): MsgCreateDID;
    toJSON(message: MsgCreateDID): unknown;
    fromPartial(object: DeepPartial<MsgCreateDID>): MsgCreateDID;
};
export declare const MsgCreateDIDResponse: {
    encode(message: MsgCreateDIDResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateDIDResponse;
    fromJSON(object: any): MsgCreateDIDResponse;
    toJSON(message: MsgCreateDIDResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateDIDResponse>): MsgCreateDIDResponse;
};
export declare const MsgUpdateDID: {
    encode(message: MsgUpdateDID, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateDID;
    fromJSON(object: any): MsgUpdateDID;
    toJSON(message: MsgUpdateDID): unknown;
    fromPartial(object: DeepPartial<MsgUpdateDID>): MsgUpdateDID;
};
export declare const MsgUpdateDIDResponse: {
    encode(message: MsgUpdateDIDResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateDIDResponse;
    fromJSON(object: any): MsgUpdateDIDResponse;
    toJSON(message: MsgUpdateDIDResponse): unknown;
    fromPartial(object: DeepPartial<MsgUpdateDIDResponse>): MsgUpdateDIDResponse;
};
export declare const MsgCreateSchema: {
    encode(message: MsgCreateSchema, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateSchema;
    fromJSON(object: any): MsgCreateSchema;
    toJSON(message: MsgCreateSchema): unknown;
    fromPartial(object: DeepPartial<MsgCreateSchema>): MsgCreateSchema;
};
export declare const MsgCreateSchemaResponse: {
    encode(message: MsgCreateSchemaResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateSchemaResponse;
    fromJSON(object: any): MsgCreateSchemaResponse;
    toJSON(message: MsgCreateSchemaResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateSchemaResponse>): MsgCreateSchemaResponse;
};
export declare const MsgDeactivateDID: {
    encode(message: MsgDeactivateDID, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeactivateDID;
    fromJSON(object: any): MsgDeactivateDID;
    toJSON(message: MsgDeactivateDID): unknown;
    fromPartial(object: DeepPartial<MsgDeactivateDID>): MsgDeactivateDID;
};
export declare const MsgDeactivateDIDResponse: {
    encode(message: MsgDeactivateDIDResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeactivateDIDResponse;
    fromJSON(object: any): MsgDeactivateDIDResponse;
    toJSON(message: MsgDeactivateDIDResponse): unknown;
    fromPartial(object: DeepPartial<MsgDeactivateDIDResponse>): MsgDeactivateDIDResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateDID(request: MsgCreateDID): Promise<MsgCreateDIDResponse>;
    UpdateDID(request: MsgUpdateDID): Promise<MsgUpdateDIDResponse>;
    CreateSchema(request: MsgCreateSchema): Promise<MsgCreateSchemaResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeactivateDID(request: MsgDeactivateDID): Promise<MsgDeactivateDIDResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateDID(request: MsgCreateDID): Promise<MsgCreateDIDResponse>;
    UpdateDID(request: MsgUpdateDID): Promise<MsgUpdateDIDResponse>;
    CreateSchema(request: MsgCreateSchema): Promise<MsgCreateSchemaResponse>;
    DeactivateDID(request: MsgDeactivateDID): Promise<MsgDeactivateDIDResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
