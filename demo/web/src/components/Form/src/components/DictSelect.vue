<template>
  <Select
    v-bind="$attrs"
    :options="getOptions"
    v-model:value="state"
    :showSearch="showSearch"
    optionFilterProp="label"
    :getPopupContainer="getPopupContainer"
  >
    <template #[item]="data" v-for="item in Object.keys($slots)">
      <slot :name="item" v-bind="data || {}"></slot>
    </template>
  </Select>
</template>
<script lang="ts">
  import { computed, defineComponent, PropType } from 'vue';
  import { Select } from 'ant-design-vue';
  import { useAttrs } from '/@/hooks/core/useAttrs';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { toOptions, useDictStoreWithOut } from '/@/store/modules/dictionary';

  const getPopupContainer = () => document.body;

  // type OptionsItem = { label: string; value: string; disabled?: boolean };

  export default defineComponent({
    name: 'DictSelect',
    components: {
      Select,
    },
    inheritAttrs: false,
    props: {
      dictName: String,
      showSearch: {
        type: Boolean as PropType<boolean>,
        default: true,
      },
      value: [Array, String],
      filter: {
        type: Function as PropType<(_) => boolean>,
        default: null,
      },
    },
    emits: ['change', 'update:value'],
    setup(props, { emit }) {
      const attrs = useAttrs();
      const { t } = useI18n();

      // Embedded in the form, just use the hook binding to perform form verification
      const state = computed({
        get() {
          return props.value;
        },
        set(value) {
          emit('change', value);
          emit('update:value', value);
        },
      });

      const getOptions = computed(() => {
        const { dictName, filter } = props;
        const dictStore = useDictStoreWithOut();
        let dict = (dictName && dictStore.listRegistry[dictName]) || [];
        if (filter) {
          dict = dict.filter(filter);
        }
        return toOptions(dict);
      });

      return { state, attrs, getOptions, t, getPopupContainer };
    },
  });
</script>
