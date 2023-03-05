<template>
  <Cascader v-bind="$attrs" v-model:value="state" :options="getOptions" @change="handleChange">
    <template #tagRender="data">
      <Tag :key="data.value" color="blue">{{ data.label }} TBD</Tag>
    </template>
  </Cascader>
</template>
<script lang="ts" setup>
  import { computed, PropType, watch } from 'vue';
  import { Cascader, Tag } from 'ant-design-vue';
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

  const state = computed({
    get() {
      return props.value;
    },
    set(value) {
      emit('change', value);
      emit('update:value', value);
    },
  });

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
    emit('defaultChange', keys, args);
  }
</script>
