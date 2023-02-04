import { defineStore } from 'pinia';
import { store } from '../index';
import { getAllDicts } from '/@/api/sys/dict';
import { SysDictEntryModel } from '/@/api/sys/model/dictModel';

export const LABEL_KEY = '__kobh_label';
export const VALUE_KEY = '__kobh_value';
export const STATUS_KEY = '__kobh_status';

interface DictEntry {
  [LABEL_KEY]: string;
  [VALUE_KEY]: string;
  [STATUS_KEY]?: string;
  children?: DictEntry[];
}

export interface DictOptionEntry {
  label: string;
  value: string;
  disabled?: boolean;
  children?: DictOptionEntry[];
}

/* eslint-disable no-multi-spaces */
const extDicts = [
  { kind: 'TbxBroker', field: 'brokerList', statusKey: 'status' },
  { kind: 'TbxChannel', field: 'channelList' },
  { kind: 'TbxCountry', field: 'countryList', labelKey: 'nameCN' },
  { kind: 'TbxCustomer', field: 'customerList', codeKey: 'customerCode' },
  { kind: 'TbxLabel', field: 'labelList' },
  { kind: 'TbxPort', field: 'portList' },
  { kind: 'TbxRoute', field: 'routeList' },
  { kind: 'TbxShippingVendor', field: 'shippingVendorList' },
  { kind: 'TbxShippingLine', field: 'shippingLineList', codeKey: 'lineCode', labelKey: 'lineName' },
  { kind: 'TbxSupplyAgency', field: 'supplyAgencyList' },
  { kind: 'TbxTrailer', field: 'trailerList' },
  { kind: 'TbxCarrier', field: 'carrierList' },
  { kind: 'TbxWarehouse', field: 'warehouseList' },
  { kind: 'TbxShippingVendor', field: 'shippingVendorList' },
];
/* eslint-enable no-multi-spaces */

const alias = [
  {
    kind: 'TbxContainerChannel',
    factory: (dictRegistry, listRegistry) => {
      return [
        dictRegistry['TbxChannel'],
        listRegistry['TbxChannel'].filter((ch) => ch.expressType === 'container'),
      ];
    },
  },
  {
    kind: 'TbxCargoOnlyChannel',
    factory: (dictRegistry, listRegistry) => {
      return [
        dictRegistry['TbxChannel'],
        listRegistry['TbxChannel'].filter((ch) => ch.expressType === 'cargo'),
      ];
    },
  },
  {
    kind: 'TbxCargoChannel',
    factory: (dictRegistry, listRegistry) => {
      return [
        dictRegistry['TbxChannel'],
        listRegistry['TbxChannel'].filter(
          (ch) => ch.expressType === 'cargo' || ch.expressType === 'container',
        ),
      ];
    },
  },
  {
    kind: 'TbxParcelChannel',
    factory: (dictRegistry, listRegistry) => {
      return [
        dictRegistry['TbxChannel'],
        listRegistry['TbxChannel'].filter((ch) => ch.expressType === 'parcel'),
      ];
    },
  },
  {
    kind: 'TbxRegion',
    factory: (_, listRegistry) => {
      const list = listRegistry['TbxCountry'];
      const l2: SysDictEntryModel[] = [];
      list.forEach((x) => {
        l2.push(x);
        if (x.children) {
          const children = x.children.map((y) => {
            const z = Object.assign({}, y);
            z.nameCN = x.nameCN + '/' + y.nameCN;
            return z;
          });
          l2.push(...children);
        }
      });
      const d2 = {};
      l2.forEach((i) => {
        deleteUselessFields(i);
        // LABEL_KEY VALUE_KEY
        i[LABEL_KEY] = i['nameCN'];
        i[VALUE_KEY] = i['code'];
        d2[i['code']] = i;
      });
      return [d2, l2];
    },
  },
];

const uselessFields = ['createBy', 'createdAt', 'remark', 'status', 'updateBy', 'updatedAt'];

function deleteUselessFields(x) {
  uselessFields.forEach((field) => {
    delete x[field];
  });
}
//
// const state = {
//   dictRegistry: {},
//   listRegistry: {},
// };
//
// const mutations = {
//   RESET_ALL_DICT: (state, dict) => {
//   },
//   REGISTRY_MISSING_DICTS: (state, { kind, dicts }) => {
//     if (!state.dictRegistry[kind]) {
//       state.dictRegistry[kind] = dicts;
//     }
//   },
//   REGISTRY_MISSING_LIST: (state, { kind, list }) => {
//     if (!state.listRegistry[kind]) {
//       state.dictRegistry[kind] = list;
//     }
//   },
// };
//
// const actions = {
//   async getAllDict({ commit }) {
//     const res = await getAllDict();
//     commit('RESET_ALL_DICT', res.data);
//     return res;
//   },
//   registryMissingDicts({ commit }, payload) {
//     commit('REGISTRY_MISSING_DICTS', payload);
//   },
//   registryMissingList({ commit }, payload) {
//     commit('REGISTRY_MISSING_LIST', payload);
//   },
// };

interface DictState {
  dictRegistry: { [key: string]: { [key: string]: Recordable } };
  listRegistry: { [key: string]: Recordable[] };
  // Page loading status
  dictLoaded: boolean;
}

export const useDictStore = defineStore({
  id: 'dict',
  state: (): DictState => ({
    dictRegistry: {},
    listRegistry: {},
    dictLoaded: false,
  }),
  getters: {
    getDictLoaded(): boolean {
      return this.dictLoaded;
    },
    getDictRegistry(): any {
      return this.dictRegistry;
    },
    getListRegistry(): any {
      return this.listRegistry;
    },
  },
  actions: {
    resetState() {
      this.listRegistry = {};
      this.dictRegistry = {};
      this.dictLoaded = false;
    },

    async buildAllDicts() {
      const dict = await getAllDicts();

      const dictRegistry = {};
      const listRegistry = {};

      extDicts.forEach(
        ({ kind, field, codeKey = 'code', labelKey = 'name', statusKey = 'status' }) => {
          const list: DictEntry[] = dict[field] || [];
          const dicts = {};

          const processList = (list) => {
            list.forEach((i) => {
              deleteUselessFields(i);
              // LABEL_KEY VALUE_KEY for @/components/DictSelect
              i[LABEL_KEY] = i[labelKey];
              i[VALUE_KEY] = i[codeKey];
              i[STATUS_KEY] = i[statusKey] || '2';
              dicts[i[codeKey]] = i;
              i.children && processList(i.children);
            });
          };
          processList(list);
          dictRegistry[kind] = dicts;
          listRegistry[kind] = list;
        },
      );

      const systemDicts: SysDictEntryModel[] = dict.systemList || [];

      systemDicts.sort((a, b) => {
        return b.dictSort - a.dictSort;
      });

      const groupByDictType: { [key: string]: { list: DictEntry[]; dicts: {} } } = {};

      systemDicts.forEach(({ dictType, dictValue, dictLabel, status, remark }) => {
        const group = (groupByDictType[dictType] ||= { list: [], dicts: {} });
        const item = {
          [LABEL_KEY]: dictLabel,
          [VALUE_KEY]: dictValue,
          [STATUS_KEY]: status == 1 ? '1' : '2',
          label: dictLabel,
          value: dictValue,
          remark,
        };
        group.list.push(item);
        group.dicts[dictValue] = item;
      });

      Object.entries(groupByDictType).forEach(([kind, { list, dicts }]) => {
        if (dictRegistry[kind]) {
          console.warn(`dictionary index name '${kind}' is duplicated`);
          return;
        }
        dictRegistry[kind] = dicts;
        listRegistry[kind] = list;
      });

      alias.forEach(({ kind, factory }) => {
        const [dicts, list] = factory(dictRegistry, listRegistry);
        dictRegistry[kind] = dicts;
        listRegistry[kind] = list;
      });

      this.dictRegistry = dictRegistry;
      this.listRegistry = listRegistry;
      this.dictLoaded = true;
    },
  },
});

// Need to be used outside the setup
export function useDictStoreWithOut() {
  return useDictStore(store);
}

export function toOptions(dict: Recordable[]): DictOptionEntry[] {
  return dict.map((e) => ({
    label: e[LABEL_KEY],
    value: e[VALUE_KEY],
    disabled: e[STATUS_KEY] !== '2',
    children: e.children ? toOptions(e.children) : undefined,
  }));
}
