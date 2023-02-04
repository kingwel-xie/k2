<template>
  <Select v-bind="$attrs" @change="handleChange" :options="getOptions" v-model:value="state">
    <template #[item]="data" v-for="item in Object.keys($slots)">
      <slot :name="item" v-bind="data || {}"></slot>
    </template>
  </Select>
</template>
<script lang="ts">
  import { computed, defineComponent, ref, watch } from 'vue';
  import { Select } from 'ant-design-vue';
  import { useRuleFormItem } from '/@/hooks/component/useFormItem';
  import { useAttrs } from '/@/hooks/core/useAttrs';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { toOptions, useDictStoreWithOut } from '/@/store/modules/dictionary';

  // type OptionsItem = { label: string; value: string; disabled?: boolean };

  export default defineComponent({
    name: 'DictSelect',
    components: {
      Select,
    },
    inheritAttrs: false,
    props: {
      dictName: String,
      value: [Array, String],
      filter: {
        type: Function as PropType<(_) => boolean>,
        default: null,
      },
    },
    emits: ['change', 'update:value'],
    setup(props, { emit }) {
      // const options = ref<OptionsItem[]>([]);
      const emitData = ref<any[]>([]);
      const attrs = useAttrs();
      const { t } = useI18n();

      // Embedded in the form, just use the hook binding to perform form verification
      const [state] = useRuleFormItem(props, 'value', 'change', emitData);

      const getOptions = computed(() => {
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

      function handleChange(_, ...args) {
        emitData.value = args;
      }

      return { state, attrs, getOptions, t, handleChange };
    },
  });
</script>
