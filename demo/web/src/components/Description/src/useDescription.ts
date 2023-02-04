import type { DescriptionProps, DescInstance, UseDescReturnType } from './typing';
import { ref, getCurrentInstance, unref } from 'vue';
import { isProdMode } from '/@/utils/env';
import { isFunction, isMap, isString } from '/@/utils/is';
import { DescItemFormat } from './typing';
import { formatValue } from '/@/utils/formatUtil';

export function useDescription(props?: Partial<DescriptionProps>): UseDescReturnType {
  if (!getCurrentInstance()) {
    throw new Error('useDescription() can only be used inside setup() or functional components!');
  }
  const desc = ref<Nullable<DescInstance>>(null);
  const loaded = ref(false);

  function register(instance: DescInstance) {
    if (unref(loaded) && isProdMode()) {
      return;
    }
    desc.value = instance;
    props && instance.setDescProps(props);
    loaded.value = true;
  }

  const methods: DescInstance = {
    setDescProps: (descProps: Partial<DescriptionProps>): void => {
      unref(desc)?.setDescProps(descProps);
    },
  };

  return [register, methods];
}

// format cell
export function formatItem(text: any, record: Recordable, format: DescItemFormat) {
  if (!format) {
    return text;
  }

  // custom function
  if (isFunction(format)) {
    return format(text, record);
  }

  try {
    if (isString(format)) {
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
