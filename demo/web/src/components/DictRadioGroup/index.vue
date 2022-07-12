<template>
  <el-radio-group v-bind="$attrs" v-on="$listeners">
    <slot v-for="(opt, i) in options" :index="i" :option="opt">
      <el-radio-button v-if="button" :key="i" :label="opt[valueKey]">{{ opt[labelKey] }}</el-radio-button>
      <el-radio v-else :key="i" :label="opt[valueKey]">{{ opt[labelKey] }}</el-radio>
    </slot>
  </el-radio-group>
</template>

<script>
import { LABEL_KEY, VALUE_KEY } from '@/store/modules/dictionary'

export default {
  name: 'DictRadioGroup',
  props: {
    dict: {
      type: String,
      default: undefined
    },
    labelKey: {
      type: String,
      default: LABEL_KEY
    },
    valueKey: {
      type: String,
      default: VALUE_KEY
    },
    filter: {
      type: Function,
      default: null
    },
    button: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      rawOptions: []
    }
  },
  computed: {
    options() {
      return this.filter ? this.rawOptions.filter(this.filter) : this.rawOptions
    }
  },
  created() {
    this.rawOptions = this.$store.getters.listRegistry(this.dict)
  }
}
</script>

<style scoped>

</style>
