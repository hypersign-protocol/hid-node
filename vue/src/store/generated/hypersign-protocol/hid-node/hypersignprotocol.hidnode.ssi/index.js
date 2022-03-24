import { txClient, queryClient, MissingWalletError, registry } from './module';
// @ts-ignore
import { SpVuexError } from '@starport/vuex';
import { Did } from "./module/types/ssi/v1/did";
import { Metadata } from "./module/types/ssi/v1/did";
import { DidResolveMeta } from "./module/types/ssi/v1/did";
import { VerificationMethod } from "./module/types/ssi/v1/did";
import { Service } from "./module/types/ssi/v1/did";
import { SignInfo } from "./module/types/ssi/v1/did";
import { DidDocument } from "./module/types/ssi/v1/did";
import { Params } from "./module/types/ssi/v1/params";
import { DidResolutionResponse } from "./module/types/ssi/v1/query";
import { Schema } from "./module/types/ssi/v1/schema";
import { SchemaProperty } from "./module/types/ssi/v1/schema";
export { Did, Metadata, DidResolveMeta, VerificationMethod, Service, SignInfo, DidDocument, Params, DidResolutionResponse, Schema, SchemaProperty };
async function initTxClient(vuexGetters) {
    return await txClient(vuexGetters['common/wallet/signer'], {
        addr: vuexGetters['common/env/apiTendermint']
    });
}
async function initQueryClient(vuexGetters) {
    return await queryClient({
        addr: vuexGetters['common/env/apiCosmos']
    });
}
function mergeResults(value, next_values) {
    for (let prop of Object.keys(next_values)) {
        if (Array.isArray(next_values[prop])) {
            value[prop] = [...value[prop], ...next_values[prop]];
        }
        else {
            value[prop] = next_values[prop];
        }
    }
    return value;
}
function getStructure(template) {
    let structure = { fields: [] };
    for (const [key, value] of Object.entries(template)) {
        let field = {};
        field.name = key;
        field.type = typeof value;
        structure.fields.push(field);
    }
    return structure;
}
const getDefaultState = () => {
    return {
        Params: {},
        GetSchema: {},
        Schemas: {},
        SchemaCount: {},
        ResolveDid: {},
        DidParam: {},
        _Structure: {
            Did: getStructure(Did.fromPartial({})),
            Metadata: getStructure(Metadata.fromPartial({})),
            DidResolveMeta: getStructure(DidResolveMeta.fromPartial({})),
            VerificationMethod: getStructure(VerificationMethod.fromPartial({})),
            Service: getStructure(Service.fromPartial({})),
            SignInfo: getStructure(SignInfo.fromPartial({})),
            DidDocument: getStructure(DidDocument.fromPartial({})),
            Params: getStructure(Params.fromPartial({})),
            DidResolutionResponse: getStructure(DidResolutionResponse.fromPartial({})),
            Schema: getStructure(Schema.fromPartial({})),
            SchemaProperty: getStructure(SchemaProperty.fromPartial({})),
        },
        _Registry: registry,
        _Subscriptions: new Set(),
    };
};
// initial state
const state = getDefaultState();
export default {
    namespaced: true,
    state,
    mutations: {
        RESET_STATE(state) {
            Object.assign(state, getDefaultState());
        },
        QUERY(state, { query, key, value }) {
            state[query][JSON.stringify(key)] = value;
        },
        SUBSCRIBE(state, subscription) {
            state._Subscriptions.add(JSON.stringify(subscription));
        },
        UNSUBSCRIBE(state, subscription) {
            state._Subscriptions.delete(JSON.stringify(subscription));
        }
    },
    getters: {
        getParams: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.Params[JSON.stringify(params)] ?? {};
        },
        getGetSchema: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.GetSchema[JSON.stringify(params)] ?? {};
        },
        getSchemas: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.Schemas[JSON.stringify(params)] ?? {};
        },
        getSchemaCount: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.SchemaCount[JSON.stringify(params)] ?? {};
        },
        getResolveDid: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ResolveDid[JSON.stringify(params)] ?? {};
        },
        getDidParam: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.DidParam[JSON.stringify(params)] ?? {};
        },
        getTypeStructure: (state) => (type) => {
            return state._Structure[type].fields;
        },
        getRegistry: (state) => {
            return state._Registry;
        }
    },
    actions: {
        init({ dispatch, rootGetters }) {
            console.log('Vuex module: hypersignprotocol.hidnode.ssi initialized!');
            if (rootGetters['common/env/client']) {
                rootGetters['common/env/client'].on('newblock', () => {
                    dispatch('StoreUpdate');
                });
            }
        },
        resetState({ commit }) {
            commit('RESET_STATE');
        },
        unsubscribe({ commit }, subscription) {
            commit('UNSUBSCRIBE', subscription);
        },
        async StoreUpdate({ state, dispatch }) {
            state._Subscriptions.forEach(async (subscription) => {
                try {
                    const sub = JSON.parse(subscription);
                    await dispatch(sub.action, sub.payload);
                }
                catch (e) {
                    throw new SpVuexError('Subscriptions: ' + e.message);
                }
            });
        },
        async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryParams()).data;
                commit('QUERY', { query: 'Params', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: { ...key }, query } });
                return getters['getParams']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryParams', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryGetSchema({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryGetSchema(key.schemaId)).data;
                commit('QUERY', { query: 'GetSchema', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryGetSchema', payload: { options: { all }, params: { ...key }, query } });
                return getters['getGetSchema']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryGetSchema', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QuerySchemas({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.querySchemas(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.querySchemas({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'Schemas', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QuerySchemas', payload: { options: { all }, params: { ...key }, query } });
                return getters['getSchemas']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QuerySchemas', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QuerySchemaCount({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.querySchemaCount()).data;
                commit('QUERY', { query: 'SchemaCount', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QuerySchemaCount', payload: { options: { all }, params: { ...key }, query } });
                return getters['getSchemaCount']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QuerySchemaCount', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryResolveDid({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryResolveDid(key.didId, query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryResolveDid(key.didId, { ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'ResolveDid', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryResolveDid', payload: { options: { all }, params: { ...key }, query } });
                return getters['getResolveDid']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryResolveDid', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryDidParam({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryDidParam(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryDidParam({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'DidParam', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryDidParam', payload: { options: { all }, params: { ...key }, query } });
                return getters['getDidParam']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryDidParam', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async sendMsgDeactivateDID({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgDeactivateDID(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgDeactivateDID:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgDeactivateDID:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgCreateDID({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCreateDID(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgCreateDID:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCreateDID:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgUpdateDID({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgUpdateDID(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgUpdateDID:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgUpdateDID:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgCreateSchema({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCreateSchema(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgCreateSchema:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCreateSchema:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async MsgDeactivateDID({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgDeactivateDID(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgDeactivateDID:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgDeactivateDID:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgCreateDID({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCreateDID(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgCreateDID:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCreateDID:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgUpdateDID({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgUpdateDID(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgUpdateDID:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgUpdateDID:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgCreateSchema({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCreateSchema(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgCreateSchema:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCreateSchema:Create', 'Could not create message: ' + e.message);
                }
            }
        },
    }
};
