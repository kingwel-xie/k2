import dayjs from 'dayjs';
import { set } from 'lodash-es';
import { tryParseJson } from '/@/utils/formatUtil';

export const kobhDateTimeValueFormat = 'YYYY-MM-DDTHH:mm:ssZ';

export const kobhFullDay = {
  defaultValue: [dayjs('00:00:00', 'HH:mm:ss'), dayjs('23:59:59', 'HH:mm:ss')],
};

export function kobhRegularOptions(): any {
  const present = dayjs();
  return {
    近一周: [present.subtract(7, 'day'), present],
    近一月: [present.subtract(1, 'month'), present],
    近三月: [present.subtract(3, 'month'), present],
    本月: [present.startOf('month'), present],
  };
}

// split a string with [, or \n] to a string array
// used to split delimited serial numbers
export function stringToArray(str: string, separator = /[,;\n]/) {
  if (!str || str.length === 0) {
    return undefined;
  }
  return str
    .split(separator)
    .filter((x) => x !== '')
    .map((x) => x.trim());
}

/**
 * @param {Array} arr
 * @returns {Array}
 */
export function uniqueArray<T>(arr: Array<T>) {
  return Array.from(new Set(arr));
}

export function toNestedObject(v: Recordable): Recordable {
  const newObj = {};
  Object.keys(v).forEach((key) => {
    const pathArray = key.split('.');
    set(newObj, pathArray, v[key]);
  });
  return newObj;
}

/**
 * @param {number} value, 传入
 * @returns {string}
 */
export function formatFileSize(value) {
  if (value > 1024 * 1024) {
    return (value / 1024 / 1024).toFixed(2) + ' MB';
  } else if (value > 1024) {
    return (value / 1024).toFixed(2) + ' KB';
  } else {
    return value + ' B';
  }
}

/**
 * @param {string} value, 传入
 * @returns {string}
 */
export function formatCurrencyString(value: string): string {
  const v = tryParseJson(value) || {};
  return v.value + ', ' + v.currencyType;
}

/**
 * @param {Object} value, 传入
 * @returns {string}
 */
export function asExcelString(value, elseAs: any | undefined = undefined): string | undefined {
  return value !== undefined ? String(value).trim() : elseAs;
}

/**
 * two floats deemed as exactly equal
 * @param {number} x, 传入
 * @param {number} y, 传入
 * @returns {boolean}
 */
export function floatEq(x, y): boolean {
  return Math.abs(x - y) < Number.EPSILON;
}

/**
 * two floats close enough to be deemed as the same
 * @param {number} x, 传入
 * @param {number} y, 传入
 * @returns {boolean}
 */
export function floatClose(x, y): boolean {
  return Math.abs(x - y) < 0.0001;
}

/**
 * @param {number} num, 传入的数字
 * @param {Object} n, 需要返回的字符长度
 * @returns {string}
 */
function prefixInteger(num, n) {
  return (Array(n).join('0') + num).slice(-n);
}

/**
 * @param {mark} string, FBA Mark
 * @param {index} num, 序号
 * @param {width} num, 位数
 * @param {Object} n, 需要返回的字符长度
 * @returns {string}
 */
export function expandFbaBoxNo(mark: string, index: number, width = 6): string {
  return mark + 'U' + prefixInteger(index, width);
}
