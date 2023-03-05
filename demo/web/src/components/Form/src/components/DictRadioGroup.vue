<!--
 * @Description:It is troublesome to implement radio button group in the form. So it is extracted independently as a separate component
-->
<template>
  <RadioGroup v-bind="attrs" v-model:value="state" button-style="solid">
    <template v-for="item in getOptions" :key="`${item.value}`">
      <RadioButton v-if="props.isBtn" :value="item.value" :disabled="item.disabled">
        {{ item.label }}
      </RadioButton>
      <Radio v-else :value="item.value" :disabled="item.disabled">
        {{ item.label }}
      </Radio>
    </template>
  </RadioGroup>
</template>
<script lang="ts">
  import { defineComponent, PropType, computed } from 'vue';
  import { Radio } from 'ant-design-vue';
  import { useAttrs } from '/@/hooks/core/useAttrs';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { toOptions, useDictStoreWithOut } from '/@/store/modules/dictionary';
  // type OptionsItem = { label: string; value: string | number | boolean; disabled?: boolean };

  export default defineComponent({
    name: 'DictRadioGroup',
    components: {
      RadioGroup: Radio.Group,
      RadioButton: Radio.Button,
      Radio,
    },
    props: {
      dictName: String,
      value: String,
      filter: {
        type: Function as PropType<(_) => boolean>,
        default: null,
      },
      isBtn: {
        type: [Boolean] as PropType<boolean>,
        default: false,
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

      return { state, getOptions, attrs, t, props };
    },
  });
</script>
