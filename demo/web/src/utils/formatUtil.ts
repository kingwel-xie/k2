/**
 * Dict tool
 */
import { LABEL_KEY, useDictStoreWithOut } from '/@/store/modules/dictionary';
import { formatToDate, formatToDateTime } from '/@/utils/dateUtil';
import { isNumber } from '/@/utils/is';

export function formatToDictLabel(value: string, dictName: string): string {
  const dictStore = useDictStoreWithOut();
  const dict = dictStore.dictRegistry[dictName] || {};
  const entry = dict[value] || {};
  return entry[LABEL_KEY] || value;
}

export function formatFromDictLabel(label: string | undefined, dictName: string): string {
  if (!label) return '';
  const dictStore = useDictStoreWithOut();
  const list = dictStore.listRegistry[dictName] || [];
  const entry = list.find((l) => l.label === label);
  return (entry && entry.value) || '';
}

/**
 * @returns {Object}
 * @param str
 */
export function tryParseJson<T = any>(str: string): T | undefined {
  try {
    return JSON.parse(str);
  } catch {}
  return undefined;
}

// dict type
const DICT_FORMAT_PREFIX = 'dict|';
// date type
const DATE_FORMAT_PREFIX = 'datetime|';
// bool type
const BOOL_FORMAT_PREFIX = 'bool|';
// float type
const FLOAT_FORMAT_PREFIX = 'float|';

export function formatValue(format: string, val: any): string {
  if (format.startsWith(DICT_FORMAT_PREFIX)) {
    const dictFormat = format.replace(DICT_FORMAT_PREFIX, '');
    if (!dictFormat) {
      return val;
    }
    return formatToDictLabel(val, dictFormat);
  } else if (format.startsWith(BOOL_FORMAT_PREFIX)) {
    const boolFormat = format.replace(BOOL_FORMAT_PREFIX, '');
    let trueOrFalse = boolFormat.split(/[,;/s]/) || [];
    if (trueOrFalse.length !== 2) {
      trueOrFalse = ['Y', 'N'];
    }
    return trueOrFalse[val ? 0 : 1];
  } else if (format.startsWith(DATE_FORMAT_PREFIX)) {
    const dateFormat = format.replace(DATE_FORMAT_PREFIX, '');
    if (!dateFormat) {
      return formatToDateTime(val);
    } else {
      switch (dateFormat) {
        case 'date':
          return formatToDate(val);
        case 'time':
          return formatToDateTime(val, 'HH:mm:ss');
        default:
          return formatToDateTime(val, dateFormat);
      }
    }
  } else if (format.startsWith(FLOAT_FORMAT_PREFIX)) {
    if (!isNumber(val)) return '-';
    const temp = format.replace(FLOAT_FORMAT_PREFIX, '');
    const decimal = Number(temp) || 2;
    return val.toFixed(decimal);
  }
  return val;
}

export function formatFloat(val: any, decimal = 2): string {
  return formatValue(FLOAT_FORMAT_PREFIX + decimal, val);
}

export function formatDict(val: any, dict: string): string {
  return formatValue(DICT_FORMAT_PREFIX + dict, val);
}

export function formatBool(val: any, dict = ''): string {
  return formatValue(BOOL_FORMAT_PREFIX + dict, val);
}

// helper function, to convert dict values in JSON string into a dict lable string, separated
// by the given separator
export function formatMultiDictLabel(data: string, dictName: string, separator = ', ') {
  return (tryParseJson(data) || []).map((x) => formatToDictLabel(x, dictName)).join(separator);
}
