<template>
  <Cascader v-bind="$attrs" v-model:value="state" :options="getOptions" @change="handleChange">
    <template #tagRender="data">
      <Tag :key="data.value" color="blue">{{ data.label }} TBD</Tag>
    </template>
  </Cascader>
</template>
<script lang="ts" setup>
  import { computed, PropType, ref, unref, watch } from 'vue';
  import { Cascader, Tag } from 'ant-design-vue';
  import { useRuleFormItem } from '/@/hooks/component/useFormItem';
  import { toOptions, useDictStoreWithOut } from '/@/store/modules/dictionary';

  interface Option {
    value: string;
    label: string;
    children?: Option[];
  }

  const props = defineProps({
    value: {
      type: Array,
    },
    dictName: String,
    filter: {
      type: Function as PropType<(_) => boolean>,
      default: null,
    },
    displayRenderArray: {
      type: Array,
    },
  });
  const emit = defineEmits(['change', 'update:value', 'defaultChange']);

  const emitData = ref<any[]>([]);
  // Embedded in the form, just use the hook binding to perform form verification
  const [state] = useRuleFormItem(props, 'value', 'change', emitData);

  const getOptions = computed<Option[]>(() => {
    const { dictName, filter } = props;
    const dictStore = useDictStoreWithOut();
    let dict = (dictName && dictStore.listRegistry[dictName]) || [];
    if (filter) {
      dict = dict.filter(filter);
    }
    return toOptions(dict);
  });

  watch(
    () => state.value,
    (v) => {
      emit('update:value', v);
    },
  );

  function handleChange(keys, args) {
    emitData.value = keys;
    emit('defaultChange', keys, args);
  }

  function handleRenderDisplay({ labels, selectedOptions }) {
    if (unref(emitData).length === selectedOptions.length) {
      return labels.join(' / ');
    }
    if (props.displayRenderArray) {
      return props.displayRenderArray.join(' / ');
    }
    return '';
  }
</script>
