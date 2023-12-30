import { defineStore } from 'pinia';
import { store } from '../index';
import { getAllDicts } from '/@/api/sys/dict';
import { SysDictEntryModel } from '/@/api/sys/model/dictModel';
import { KobhWorksheet, excelColumnIndex2Name } from '/@/utils/excelUtils';

export const LABEL_KEY = '__kobh_label';
export const VALUE_KEY = '__kobh_value';
export const STATUS_KEY = '__kobh_status';
export const HIDDEN_DICT_SHEET = '__dict__';

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
  { kind: 'TbxCustomer', field: 'customerList', codeKey: 'customerCode', noValidation: true },
  { kind: 'TbxLabel', field: 'labelList' },
  { kind: 'TbxPort', field: 'portList' },
  { kind: 'TbxRoute', field: 'routeList' },
  { kind: 'TbxShippingVendor', field: 'shippingVendorList' },
  { kind: 'TbxShippingLine', field: 'shippingLineList', codeKey: 'lineCode', labelKey: 'lineName' },
  { kind: 'TbxSupplyAgency', field: 'supplyAgencyList' },
  { kind: 'TbxTrailer', field: 'trailerList' },
  { kind: 'TbxCarrier', field: 'carrierList' },
  { kind: 'TbxImporter', field: 'importerList' },
  { kind: 'TbxWarehouse', field: 'warehouseList', noValidation: true },
  { kind: 'TbxShippingVendor', field: 'shippingVendorList' },
];
/* eslint-enable no-multi-spaces */

const alias = [
  {
    kind: 'TbxContainerChannel',
    factory: (kind, dictRegistry, listRegistry, listValidation) => {
      const newDict = dictRegistry['TbxChannel'];
      const newList = listRegistry['TbxChannel'].filter((ch) => ch.expressType === 'container');
      dictRegistry[kind] = newDict;
      listRegistry[kind] = newList;
      listValidation.push({ kind, list: newList });
    },
  },
  {
    kind: 'TbxCargoOnlyChannel',
    factory: (kind, dictRegistry, listRegistry, listValidation) => {
      const newDict = dictRegistry['TbxChannel'];
      const newList = listRegistry['TbxChannel'].filter((ch) => ch.expressType === 'cargo');
      dictRegistry[kind] = newDict;
      listRegistry[kind] = newList;
      listValidation.push({ kind, list: newList });
    },
  },
  {
    kind: 'TbxCargoChannel',
    factory: (kind, dictRegistry, listRegistry, listValidation) => {
      const newDict = dictRegistry['TbxChannel'];
      const newList = listRegistry['TbxChannel'].filter(
        (ch) =>
          ch.expressType === 'cargo' ||
          ch.expressType === 'container' ||
          ch.expressType === 'relocation',
      );
      dictRegistry[kind] = newDict;
      listRegistry[kind] = newList;
      listValidation.push({ kind, list: newList });
    },
  },
  {
    kind: 'TbxParcelChannel',
    factory: (kind, dictRegistry, listRegistry, listValidation) => {
      const newDict = dictRegistry['TbxChannel'];
      const newList = listRegistry['TbxChannel'].filter((ch) => ch.expressType === 'parcel');
      dictRegistry[kind] = newDict;
      listRegistry[kind] = newList;
      listValidation.push({ kind, list: newList });
    },
  },
  {
    kind: 'TbxContainerChannel',
    factory: (kind, dictRegistry, listRegistry, listValidation) => {
      const newDict = dictRegistry['TbxChannel'];
      const newList = listRegistry['TbxChannel'].filter((ch) => ch.expressType === 'container');
      dictRegistry[kind] = newDict;
      listRegistry[kind] = newList;
      listValidation.push({ kind, list: newList });
    },
  },
  {
    kind: 'TbxRegion',
    factory: (kind, dictRegistry, listRegistry, listValidation) => {
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

      dictRegistry[kind] = d2;
      listRegistry[kind] = l2;
      listValidation.push({ kind, list: l2 });
    },
  },
  {
    kind: 'TbxSupplierTypeEx',
    factory: (kind, dictRegistry, listRegistry, listValidation) => {
      const makeData = (kind, option) => {
        const list = [] as any[];
        option.forEach((x) => {
          list.push({
            [LABEL_KEY]: x[LABEL_KEY],
            [VALUE_KEY]: x[VALUE_KEY],
            [STATUS_KEY]: x[STATUS_KEY],
          });
        });
        const dict = {};
        list.forEach((i) => {
          deleteUselessFields(i);
          dict[i[VALUE_KEY]] = i;
        });
        dictRegistry[kind] = dict;
        listRegistry[kind] = list;
        listValidation.push({ kind, list });
      };

      // get the raw supplier options
      const rawOptions = rawSupplierOptions(dictRegistry, listRegistry);
      makeData(kind, rawOptions);

      // construct TbxSupplyAgency:xxxx
      rawOptions.forEach((x) => {
        const value = x[VALUE_KEY];
        if (value.startsWith('TbxSupplyAgency')) {
          makeData(value, x.children);
        }
      });
    },
  },
];

const uselessFields = ['createBy', 'createdAt', 'remark', 'updateBy', 'updatedAt'];

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
  validationSheet: { dictSheet?: KobhWorksheet; indexer?: { [key: string]: Recordable } };
  // Page loading status
  dictLoaded: boolean;
}

export const useDictStore = defineStore({
  id: 'dict',
  state: (): DictState =>
    <DictState>{
      dictRegistry: {},
      listRegistry: {},
      validationSheet: {},
      dictLoaded: false,
    },
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
    getValidationSheet(): any {
      return this.validationSheet;
    },
  },
  actions: {
    resetState() {
      this.listRegistry = {};
      this.dictRegistry = {};
      this.validationSheet = {};
      this.dictLoaded = false;
    },

    async buildAllDicts() {
      const dict = await getAllDicts();

      const dictRegistry = {};
      const listRegistry = {};
      const listValidation: any[] = [];

      extDicts.forEach(
        ({
          kind,
          field,
          codeKey = 'code',
          labelKey = 'name',
          statusKey = 'status',
          noValidation = false,
        }) => {
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
          if (!noValidation) {
            listValidation.push({ kind, list });
          }
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
        listValidation.push({ kind, list });
      });

      alias.forEach(({ kind, factory }) => {
        factory(kind, dictRegistry, listRegistry, listValidation);
      });

      // construct validation sheet and its indexer
      const validationSheet = {
        dictSheet: {
          name: HIDDEN_DICT_SHEET,
          hidden: true,
          rows: [] as Recordable[],
        } as KobhWorksheet,
        indexer: {},
      };
      let rowIndex = 0;
      listValidation.forEach(({ kind, list }) => {
        rowIndex += 1;
        const availableList = list.filter((x) => x[STATUS_KEY] === '2').map((x) => x[LABEL_KEY]);
        validationSheet.dictSheet.rows.push([kind, ...availableList]);
        validationSheet.indexer[kind] = { rowIndex, count: availableList.length };
      });
      // console.log(listRegistry, dictRegistry, validationSheet);

      this.dictRegistry = dictRegistry;
      this.listRegistry = listRegistry;
      this.validationSheet = validationSheet;
      this.dictLoaded = true;
    },
  },
});

// Need to be used outside the setup
export function useDictStoreWithOut() {
  return useDictStore(store);
}

export function toOptions(dict: Recordable[], ignoreDisabled = false): DictOptionEntry[] {
  // 2023.8.30, don't show 'disabled' entry
  return dict
    .filter((e) => {
      return ignoreDisabled || e[STATUS_KEY] === '2';
    })
    .map((e) => ({
      label: e[LABEL_KEY],
      value: e[VALUE_KEY],
      children: e.children ? toOptions(e.children) : undefined,
    }));
}

// construct feeTypeOptions for FeeType cascader
export function getFeeTypeOptions(fixedTypes?: any[]) {
  const dictStore = useDictStoreWithOut();

  const fts = dictStore.listRegistry['tbx_fee_type'];
  const feeGroups: DictOptionEntry[] = [];
  const map = {};
  fts.forEach(({ [LABEL_KEY]: label, [VALUE_KEY]: value }) => {
    // special '*' 计重价
    if (value === '*') {
      return;
    }
    let [group, fee] = label.split('-', 2);
    if (!fee) {
      fee = group;
      group = '其他';
    }
    // check if we have fixedTypes set
    if (fixedTypes && fixedTypes.indexOf(group) === -1) {
      // not found, skip this group
      return;
    }

    let idx = map[group];
    // console.log('[xxx]', { label, value, group, fee, idx })
    if (typeof idx === 'undefined') {
      idx = map[group] = feeGroups.length;
      feeGroups.push({ label: group, value: group, children: [] });
    }
    feeGroups[idx].children?.push({ label: fee, value: value });
    // console.log('[zzz]', feeGroups);
  });
  return feeGroups;
}

function rawSupplierOptions(dictRegistry, listRegistry) {
  const getSupplierTypeName = (typ) => {
    return dictRegistry['tbx_supplier_type'][typ].label;
  };
  const getAgencyTypeName = (typ) => {
    return dictRegistry['tbx_agency_type'][typ].label;
  };

  const agents = listRegistry['TbxSupplyAgency'];
  const groupByType: any = {};
  agents.forEach((agency) => {
    const group = (groupByType[agency.type] ||= []);
    group.push(agency);
  });
  const opts: any[] = [];
  Object.entries(groupByType).forEach(([typ, children]) => {
    opts.push({
      [LABEL_KEY]: getAgencyTypeName(typ),
      [VALUE_KEY]: 'TbxSupplyAgency:' + typ,
      [STATUS_KEY]: '2',
      children,
    });
  });
  // console.log(agents, groupByType, opts);
  const rawOptions = [
    // { [LABEL_KEY]: '代理商', [VALUE_KEY]: 'TbxSupplyAgency', children: dictStore.listRegistry['TbxSupplyAgency'] },
    ...opts,
    {
      [LABEL_KEY]: getSupplierTypeName('TbxBroker'),
      [VALUE_KEY]: 'TbxBroker',
      [STATUS_KEY]: '2',
      children: listRegistry['TbxBroker'],
    },
    {
      [LABEL_KEY]: getSupplierTypeName('TbxTrailer'),
      [VALUE_KEY]: 'TbxTrailer',
      [STATUS_KEY]: '2',
      children: listRegistry['TbxTrailer'],
    },
    {
      [LABEL_KEY]: getSupplierTypeName('TbxCarrier'),
      [VALUE_KEY]: 'TbxCarrier',
      [STATUS_KEY]: '2',
      children: listRegistry['TbxCarrier'],
    },
  ];
  return rawOptions;
}

// construct supplierOptions for Supplier cascader
export function getSupplierOptions() {
  const dictStore = useDictStoreWithOut();
  const rawOptions = rawSupplierOptions(dictStore.dictRegistry, dictStore.listRegistry);

  // console.log('rawSupplierOptions', rawOptions, toOptions(rawOptions));
  return toOptions(rawOptions);
}

export function getShippingLineOptions(): DictOptionEntry[] {
  const dictStore = useDictStoreWithOut();

  // prepare lineCodeOptions for the 'Cascader'
  const vendorOptions = dictStore.listRegistry['TbxShippingVendor'];
  // build lineOptions from vendorOptions
  const shippingLines = dictStore.listRegistry['TbxShippingLine'];
  const vv = vendorOptions.map((v) => {
    return { ...v, children: shippingLines.filter((line) => line.vendor === v.code) };
  });
  return toOptions(vv);
}

export function dictValidationRef(dictName: string) {
  const dictStore = useDictStoreWithOut();
  const dictIndex = dictStore.validationSheet.indexer![dictName];

  const colName = excelColumnIndex2Name(dictIndex.count + 1);
  // validation list format: dict!$A$5:$C$5
  return [`${HIDDEN_DICT_SHEET}!$B$${dictIndex.rowIndex}:$${colName}$${dictIndex.rowIndex}`];
}
