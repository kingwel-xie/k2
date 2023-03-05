import type { BasicColumn, BasicTableProps, CellFormat, GetColumnsParams } from '../types/table';
import type { PaginationProps } from '../types/pagination';
import type { ComputedRef } from 'vue';
import { computed, Ref, ref, reactive, toRaw, unref, watch, h } from 'vue';
import { renderEditCell } from '../components/editable';
import { usePermission } from '/@/hooks/web/usePermission';
import { useI18n } from '/@/hooks/web/useI18n';
import { isArray, isBoolean, isFunction, isMap, isString } from '/@/utils/is';
import { cloneDeep, isEqual } from 'lodash-es';
import { ACTION_COLUMN_FLAG, DEFAULT_ALIGN, INDEX_COLUMN_FLAG, PAGE_SIZE } from '../const';
import { formatValue, tryParseJson } from '/@/utils/formatUtil';
import { Tag } from 'ant-design-vue';
import { TABLE_COLUMN_CFG_KEY_ } from '/@/enums/cacheEnum';

function handleItem(item: BasicColumn, ellipsis: boolean) {
  const { key, dataIndex, children } = item;
  item.align = item.align || DEFAULT_ALIGN;
  if (ellipsis) {
    if (!key) {
      item.key = dataIndex;
    }
    if (!isBoolean(item.ellipsis)) {
      Object.assign(item, {
        ellipsis,
      });
    }
  }
  if (children && children.length) {
    handleChildren(children, !!ellipsis);
  }
}

function handleChildren(children: BasicColumn[] | undefined, ellipsis: boolean) {
  if (!children) return;
  children.forEach((item) => {
    const { children } = item;
    handleItem(item, ellipsis);
    handleChildren(children, ellipsis);
  });
}

function handleIndexColumn(
  propsRef: ComputedRef<BasicTableProps>,
  getPaginationRef: ComputedRef<boolean | PaginationProps>,
  columns: BasicColumn[],
) {
  const { t } = useI18n();

  const { showIndexColumn, indexColumnProps, isTreeTable } = unref(propsRef);

  let pushIndexColumns = false;
  if (unref(isTreeTable)) {
    return;
  }
  columns.forEach(() => {
    const indIndex = columns.findIndex((column) => column.flag === INDEX_COLUMN_FLAG);
    if (showIndexColumn) {
      pushIndexColumns = indIndex === -1;
    } else if (!showIndexColumn && indIndex !== -1) {
      columns.splice(indIndex, 1);
    }
  });

  if (!pushIndexColumns) return;

  const isFixedLeft = columns.some((item) => item.fixed === 'left');

  columns.unshift({
    flag: INDEX_COLUMN_FLAG,
    width: 50,
    title: t('component.table.index'),
    align: 'center',
    customRender: ({ index }) => {
      const getPagination = unref(getPaginationRef);
      if (isBoolean(getPagination)) {
        return `${index + 1}`;
      }
      const { current = 1, pageSize = PAGE_SIZE } = getPagination;
      return ((current < 1 ? 1 : current) - 1) * pageSize + index + 1;
    },
    ...(isFixedLeft
      ? {
          fixed: 'left',
        }
      : {}),
    ...indexColumnProps,
  });
}

function handleActionColumn(propsRef: ComputedRef<BasicTableProps>, columns: BasicColumn[]) {
  const { actionColumn } = unref(propsRef);
  if (!actionColumn) return;

  const hasIndex = columns.findIndex((column) => column.flag === ACTION_COLUMN_FLAG);
  if (hasIndex === -1) {
    columns.push({
      ...columns[hasIndex],
      fixed: 'right',
      ...actionColumn,
      flag: ACTION_COLUMN_FLAG,
    });
  }
}

export function useColumns(
  propsRef: ComputedRef<BasicTableProps>,
  getPaginationRef: ComputedRef<boolean | PaginationProps>,
) {
  const columnsRef = ref(unref(propsRef).columns) as unknown as Ref<BasicColumn[]>;
  let cacheColumns = unref(propsRef).columns;

  const getColumnsRef = computed(() => {
    const columns = cloneDeep(unref(columnsRef));

    // loading column configuration from local storage
    let hiddenList: string[] = [];
    if (unref(propsRef).name) {
      const cfgStr = localStorage.getItem(TABLE_COLUMN_CFG_KEY_ + unref(propsRef).name);
      if (cfgStr) {
        hiddenList = tryParseJson(cfgStr)?.hidden;
      }
    }
    columns.forEach((col) => {
      // this is quite ugly, but works...
      // column doesn't have 'adminHidden' when it is loaded first time
      // we set it according to the config or set to defaultHidden
      if (!col.hasOwnProperty('adminHidden')) {
        if (!hiddenList) {
          col.adminHidden = col.defaultHidden;
        } else if (hiddenList.includes(col.dataIndex?.toString() || (col.key as string))) {
          col.adminHidden = true;
        } else {
          col.adminHidden = false;
        }
      }
    });

    handleIndexColumn(propsRef, getPaginationRef, columns);
    handleActionColumn(propsRef, columns);
    if (!columns) {
      return [];
    }
    const { ellipsis } = unref(propsRef);

    columns.forEach((item) => {
      const { customRender, slots } = item;

      handleItem(
        item,
        Reflect.has(item, 'ellipsis') ? !!item.ellipsis : !!ellipsis && !customRender && !slots,
      );
    });
    return columns;
  });

  function isIfShow(column: BasicColumn): boolean {
    const ifShow = column.ifShow;

    let isIfShow = true;

    if (isBoolean(ifShow)) {
      isIfShow = ifShow;
    }
    if (isFunction(ifShow)) {
      isIfShow = ifShow(column);
    }
    return isIfShow;
  }
  const { hasPermission } = usePermission();

  const getViewColumns = computed(() => {
    const viewColumns = sortFixedColumn(unref(getColumnsRef));

    const columns = cloneDeep(viewColumns);
    return columns
      .filter((column) => {
        return hasPermission(column.auth) && isIfShow(column);
      })
      .map((column) => {
        const { slots, customRender, format, edit, editRow, flag } = column;

        if (!slots || !slots?.title) {
          // column.slots = { title: `header-${dataIndex}`, ...(slots || {}) };
          column.customTitle = column.title as string;
          Reflect.deleteProperty(column, 'title');
        }
        const isDefaultAction = [INDEX_COLUMN_FLAG, ACTION_COLUMN_FLAG].includes(flag!);
        if (!customRender && format && !edit && !isDefaultAction) {
          column.customRender = ({ text, record, index }) => {
            return formatCell(text, format, record, index);
          };
        }

        // edit table
        if ((edit || editRow) && !isDefaultAction && unref(propsRef).editable) {
          column.customRender = renderEditCell(column);
        }
        return reactive(column);
      });
  });

  watch(
    () => unref(propsRef).columns,
    (columns) => {
      columnsRef.value = columns;
      cacheColumns = columns?.filter((item) => !item.flag) ?? [];
    },
  );

  function setCacheColumnsByField(dataIndex: string | undefined, value: Partial<BasicColumn>) {
    if (!dataIndex || !value) {
      return;
    }
    cacheColumns.forEach((item) => {
      if (item.dataIndex === dataIndex) {
        Object.assign(item, value);
        return;
      }
    });
  }
  /**
   * set columns
   * @param columnList key｜column
   */
  function setColumns(columnList: Partial<BasicColumn>[] | (string | string[])[]) {
    const columns = cloneDeep(columnList);
    if (!isArray(columns)) return;

    if (columns.length <= 0) {
      columnsRef.value = [];
      return;
    }

    const firstColumn = columns[0];

    const cacheKeys = cacheColumns.map((item) => item.dataIndex);

    if (!isString(firstColumn) && !isArray(firstColumn)) {
      columnsRef.value = columns as BasicColumn[];
    } else {
      const columnKeys = (columns as (string | string[])[]).map((m) => m.toString());
      const newColumns: BasicColumn[] = [];
      cacheColumns.forEach((item) => {
        newColumns.push({
          ...item,
          adminHidden: !columnKeys.includes(item.dataIndex?.toString() || (item.key as string)),
        });
      });
      // Sort according to another array
      if (!isEqual(cacheKeys, columns)) {
        newColumns.sort((prev, next) => {
          return (
            columnKeys.indexOf(prev.dataIndex?.toString() as string) -
            columnKeys.indexOf(next.dataIndex?.toString() as string)
          );
        });
      }
      columnsRef.value = newColumns;
    }
    // save col configurations
    if (unref(propsRef).name) {
      const hidden = columnsRef.value
        .filter((col) => col.adminHidden)
        .map((col) => col.dataIndex?.toString() || (col.key as string));
      if (hidden) {
        const cfgStr = JSON.stringify({ hidden });
        localStorage.setItem(TABLE_COLUMN_CFG_KEY_ + unref(propsRef).name, cfgStr);
      }
    }
  }

  function getColumns(opt?: GetColumnsParams) {
    const { ignoreIndex, ignoreAction, sort } = opt || {};
    let columns = toRaw(unref(getColumnsRef));
    if (ignoreIndex) {
      columns = columns.filter((item) => item.flag !== INDEX_COLUMN_FLAG);
    }
    if (ignoreAction) {
      columns = columns.filter((item) => item.flag !== ACTION_COLUMN_FLAG);
    }

    if (sort) {
      columns = sortFixedColumn(columns);
    }

    return columns;
  }
  function getCacheColumns() {
    return cacheColumns;
  }

  return {
    getColumnsRef,
    getCacheColumns,
    getColumns,
    setColumns,
    getViewColumns,
    setCacheColumnsByField,
  };
}

function sortFixedColumn(columns: BasicColumn[]) {
  const fixedLeftColumns: BasicColumn[] = [];
  const fixedRightColumns: BasicColumn[] = [];
  const defColumns: BasicColumn[] = [];
  for (const column of columns) {
    if (column.fixed === 'left') {
      fixedLeftColumns.push(column);
      continue;
    }
    if (column.fixed === 'right') {
      fixedRightColumns.push(column);
      continue;
    }
    defColumns.push(column);
  }
  return [...fixedLeftColumns, ...defColumns, ...fixedRightColumns].filter(
    (item) => !item.adminHidden,
  );
}

// format cell
export function formatCell(text: any, format: CellFormat, record: Recordable, index: number) {
  if (!format) {
    return text;
  }

  // custom function
  if (isFunction(format)) {
    return format(text, record, index);
  }

  try {
    if (isString(format)) {
      // hard code for 'status'
      if (format === 'dict|sys_normal_disable') {
        const enable = text === '2' || text === 2;
        const color = enable ? 'success' : 'error';
        const v = formatValue(format, String(text));
        return h(Tag, { color: color }, () => v);
      }
      return formatValue(format, text);
    }

    // Map
    if (isMap(format)) {
      return format.get(text);
    }
  } catch (error) {
    return text;
  }
}
