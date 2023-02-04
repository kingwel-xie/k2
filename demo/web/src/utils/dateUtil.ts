/**
 * Independent time operation tool to facilitate subsequent switch to dayjs
 */
import dayjs from 'dayjs';
import { isString } from '/@/utils/is';

const DATE_TIME_FORMAT = 'YYYY-MM-DD HH:mm:ss';
const DATE_FORMAT = 'YYYY-MM-DD';
const DATE_TIME_FORMAT_FLEX = 'flex';
const DATE_TIME_FORMAT_FULL = 'full';

export function formatToDateTime(
  date: dayjs.Dayjs | string | undefined = undefined,
  format = DATE_TIME_FORMAT,
): string {
  // kingwel hardcode here
  if (!date || (isString(date) && date.indexOf('1-01-01') > -1)) {
    return '-';
  }
  const _this = dayjs(date);
  if (format == DATE_TIME_FORMAT_FLEX) {
    format = DATE_TIME_FORMAT;
    const now = Date.now();
    const diff = -_this.diff(now) / 1000;
    if (diff < 0) {
      // show regular time for future
    } else if (diff < 30) {
      return '刚刚';
    } else if (diff < 3600) {
      // less 1 hour
      return Math.ceil(diff / 60) + '分钟前';
    } else if (diff < 3600 * 24) {
      return Math.ceil(diff / 3600) + '小时前';
    } else if (diff < 3600 * 24 * 2) {
      return '1天前';
    }
  } else if (format == DATE_TIME_FORMAT_FULL) {
    format = DATE_TIME_FORMAT;
  }
  return _this.format(format);
}

export function formatToDate(
  date: dayjs.Dayjs | string | undefined = undefined,
  format = DATE_FORMAT,
): string {
  // kingwel hardcode here
  if (!date || (isString(date) && date.indexOf('1-01-01') > -1)) {
    return '-';
  }
  return dayjs(date).format(format);
}

export const dateUtil = dayjs;
