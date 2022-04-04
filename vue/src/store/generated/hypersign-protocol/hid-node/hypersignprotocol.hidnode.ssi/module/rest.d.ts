export interface ProtobufAny {
    "@type"?: string;
}
export interface RpcStatus {
    /** @format int32 */
    code?: number;
    message?: string;
    details?: ProtobufAny[];
}
export interface SsiDid {
    context?: string[];
    id?: string;
    controller?: string[];
    alsoKnownAs?: string[];
    verificationMethod?: SsiVerificationMethod[];
    authentication?: string[];
    assertionMethod?: string[];
    keyAgreement?: string[];
    capabilityInvocation?: string[];
    capabilityDelegation?: string[];
    service?: SsiService[];
}
export interface SsiDidResolutionResponse {
    AtContext?: string;
    didDocument?: SsiDid;
    didDocumentMetadata?: SsiMetadata;
    didResolutionMetadata?: SsiDidResolveMeta;
}
export interface SsiDidResolveMeta {
    retrieved?: string;
    error?: string;
}
export interface SsiMetadata {
    created?: string;
    updated?: string;
    deactivated?: boolean;
    versionId?: string;
}
export interface SsiMsgCreateDIDResponse {
    /** @format uint64 */
    id?: string;
}
export interface SsiMsgCreateSchemaResponse {
    /** @format uint64 */
    id?: string;
}
export interface SsiMsgDeactivateDIDResponse {
    /** @format uint64 */
    id?: string;
}
export interface SsiMsgUpdateDIDResponse {
    updateId?: string;
}
/**
 * Params defines the parameters for the module.
 */
export declare type SsiParams = object;
export interface SsiQueryDidParamResponse {
    /** @format uint64 */
    totalDidCount?: string;
    didDocList?: SsiDidResolutionResponse[];
}
export interface SsiQueryGetDidDocByIdResponse {
    AtContext?: string;
    didDocument?: SsiDid;
    didDocumentMetadata?: SsiMetadata;
    didResolutionMetadata?: SsiDidResolveMeta;
}
export interface SsiQueryGetSchemaResponse {
    schema?: SsiSchema[];
}
/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface SsiQueryParamsResponse {
    /** params holds all the parameters of this module. */
    params?: SsiParams;
}
export interface SsiQuerySchemaParamResponse {
    /** @format uint64 */
    totalCount?: string;
    schemaList?: SsiSchema[];
}
export interface SsiSchema {
    type?: string;
    modelVersion?: string;
    id?: string;
    name?: string;
    author?: string;
    authored?: string;
    schema?: SsiSchemaProperty;
}
export interface SsiSchemaProperty {
    schema?: string;
    description?: string;
    type?: string;
    properties?: string;
    required?: string[];
    additionalProperties?: boolean;
}
export interface SsiService {
    id?: string;
    type?: string;
    serviceEndpoint?: string;
}
export interface SsiSignInfo {
    verificationMethodId?: string;
    signature?: string;
}
export interface SsiVerificationMethod {
    id?: string;
    type?: string;
    controller?: string;
    publicKeyMultibase?: string;
}
/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
    /**
     * key is a value returned in PageResponse.next_key to begin
     * querying the next page most efficiently. Only one of offset or key
     * should be set.
     * @format byte
     */
    key?: string;
    /**
     * offset is a numeric offset that can be used when key is unavailable.
     * It is less efficient than using key. Only one of offset or key should
     * be set.
     * @format uint64
     */
    offset?: string;
    /**
     * limit is the total number of results to be returned in the result page.
     * If left empty it will default to a value to be set by each app.
     * @format uint64
     */
    limit?: string;
    /**
     * count_total is set to true  to indicate that the result set should include
     * a count of the total number of items available for pagination in UIs.
     * count_total is only respected when offset is used. It is ignored when key
     * is set.
     */
    countTotal?: boolean;
    /**
     * reverse is set to true if results are to be returned in the descending order.
     *
     * Since: cosmos-sdk 0.43
     */
    reverse?: boolean;
}
export declare type QueryParamsType = Record<string | number, any>;
export declare type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;
export interface FullRequestParams extends Omit<RequestInit, "body"> {
    /** set parameter to `true` for call `securityWorker` for this request */
    secure?: boolean;
    /** request path */
    path: string;
    /** content type of request body */
    type?: ContentType;
    /** query params */
    query?: QueryParamsType;
    /** format of response (i.e. response.json() -> format: "json") */
    format?: keyof Omit<Body, "body" | "bodyUsed">;
    /** request body */
    body?: unknown;
    /** base url */
    baseUrl?: string;
    /** request cancellation token */
    cancelToken?: CancelToken;
}
export declare type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;
export interface ApiConfig<SecurityDataType = unknown> {
    baseUrl?: string;
    baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
    securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}
export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
    data: D;
    error: E;
}
declare type CancelToken = Symbol | string | number;
export declare enum ContentType {
    Json = "application/json",
    FormData = "multipart/form-data",
    UrlEncoded = "application/x-www-form-urlencoded"
}
export declare class HttpClient<SecurityDataType = unknown> {
    baseUrl: string;
    private securityData;
    private securityWorker;
    private abortControllers;
    private baseApiParams;
    constructor(apiConfig?: ApiConfig<SecurityDataType>);
    setSecurityData: (data: SecurityDataType) => void;
    private addQueryParam;
    protected toQueryString(rawQuery?: QueryParamsType): string;
    protected addQueryParams(rawQuery?: QueryParamsType): string;
    private contentFormatters;
    private mergeRequestParams;
    private createAbortSignal;
    abortRequest: (cancelToken: CancelToken) => void;
    request: <T = any, E = any>({ body, secure, path, type, query, format, baseUrl, cancelToken, ...params }: FullRequestParams) => Promise<HttpResponse<T, E>>;
}
/**
 * @title ssi/v1/did.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryDidParam
     * @summary Did Param
     * @request GET:/hypersign-protocol/hidnode/ssi/did
     */
    queryDidParam: (query?: {
        count?: boolean;
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<SsiQueryDidParamResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryResolveDid
     * @summary Resolve DID
     * @request GET:/hypersign-protocol/hidnode/ssi/did/{didId}
     */
    queryResolveDid: (didId: string, query?: {
        versionId?: string;
    }, params?: RequestParams) => Promise<HttpResponse<SsiQueryGetDidDocByIdResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QuerySchemaParam
     * @summary Schema Param
     * @request GET:/hypersign-protocol/hidnode/ssi/schema
     */
    querySchemaParam: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<SsiQuerySchemaParamResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetSchema
     * @summary Queries a list of GetSchema items.
     * @request GET:/hypersign-protocol/hidnode/ssi/schema/{schemaId}
     */
    queryGetSchema: (schemaId: string, params?: RequestParams) => Promise<HttpResponse<SsiQueryGetSchemaResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryParams
     * @summary Parameters queries the parameters of the module.
     * @request GET:/hypersignprotocol/hidnode/ssi/params
     */
    queryParams: (params?: RequestParams) => Promise<HttpResponse<SsiQueryParamsResponse, RpcStatus>>;
}
export {};
