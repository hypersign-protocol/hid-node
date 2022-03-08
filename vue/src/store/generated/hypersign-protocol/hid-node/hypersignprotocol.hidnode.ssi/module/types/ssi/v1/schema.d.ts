import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "hypersignprotocol.hidnode.ssi";
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
export declare const Schema: {
    encode(message: Schema, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Schema;
    fromJSON(object: any): Schema;
    toJSON(message: Schema): unknown;
    fromPartial(object: DeepPartial<Schema>): Schema;
};
export declare const SchemaProperty: {
    encode(message: SchemaProperty, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): SchemaProperty;
    fromJSON(object: any): SchemaProperty;
    toJSON(message: SchemaProperty): unknown;
    fromPartial(object: DeepPartial<SchemaProperty>): SchemaProperty;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
