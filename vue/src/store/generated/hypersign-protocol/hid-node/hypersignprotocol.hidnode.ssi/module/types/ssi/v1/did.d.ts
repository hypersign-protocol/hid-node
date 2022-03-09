import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "hypersignprotocol.hidnode.ssi";
export interface Did {
    context: string[];
    id: string;
    /** DID Controller Spec: https://www.w3.org/TR/did-core/#did-controller */
    controller: string[];
    alsoKnownAs: string[];
    verificationMethod: VerificationMethod[];
    authentication: string[];
    assertionMethod: string[];
    keyAgreement: string[];
    capabilityInvocation: string[];
    capabilityDelegation: string[];
    service: Service[];
}
export interface Metadata {
    created: string;
    updated: string;
    deactivated: boolean;
    versionId: string;
}
export interface DidResolveMeta {
    retrieved: string;
    error: string;
}
export interface VerificationMethod {
    id: string;
    type: string;
    controller: string;
    publicKeyMultibase: string;
}
export interface Service {
    id: string;
    type: string;
    serviceEndpoint: string;
}
export interface SignInfo {
    verificationMethodId: string;
    signature: string;
}
export interface DidDocument {
    did: Did | undefined;
    metadata: Metadata | undefined;
}
export declare const Did: {
    encode(message: Did, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Did;
    fromJSON(object: any): Did;
    toJSON(message: Did): unknown;
    fromPartial(object: DeepPartial<Did>): Did;
};
export declare const Metadata: {
    encode(message: Metadata, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Metadata;
    fromJSON(object: any): Metadata;
    toJSON(message: Metadata): unknown;
    fromPartial(object: DeepPartial<Metadata>): Metadata;
};
export declare const DidResolveMeta: {
    encode(message: DidResolveMeta, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): DidResolveMeta;
    fromJSON(object: any): DidResolveMeta;
    toJSON(message: DidResolveMeta): unknown;
    fromPartial(object: DeepPartial<DidResolveMeta>): DidResolveMeta;
};
export declare const VerificationMethod: {
    encode(message: VerificationMethod, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): VerificationMethod;
    fromJSON(object: any): VerificationMethod;
    toJSON(message: VerificationMethod): unknown;
    fromPartial(object: DeepPartial<VerificationMethod>): VerificationMethod;
};
export declare const Service: {
    encode(message: Service, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Service;
    fromJSON(object: any): Service;
    toJSON(message: Service): unknown;
    fromPartial(object: DeepPartial<Service>): Service;
};
export declare const SignInfo: {
    encode(message: SignInfo, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): SignInfo;
    fromJSON(object: any): SignInfo;
    toJSON(message: SignInfo): unknown;
    fromPartial(object: DeepPartial<SignInfo>): SignInfo;
};
export declare const DidDocument: {
    encode(message: DidDocument, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): DidDocument;
    fromJSON(object: any): DidDocument;
    toJSON(message: DidDocument): unknown;
    fromPartial(object: DeepPartial<DidDocument>): DidDocument;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
