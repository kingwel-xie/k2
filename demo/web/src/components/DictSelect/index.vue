<template>
  <el-select v-bind="$attrs" v-on="$listeners">
    <slot name="pre-static-options" />
    <slot v-for="(opt, i) in options" :index="i" :option="opt">
      <el-option :key="i" :label="opt[labelKey]" :value="opt[valueKey]" />
    </slot>
    <slot name="post-static-options" />
  </el-select>
</template>

<script>
import { LABEL_KEY, VALUE_KEY } from '@/store/modules/dictionary'

export default {
  name: 'DictSelect',
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
    debug: {
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
    if (this.debug) {
      console.log(`[DictSelect] found list of '${this.dict}':`, this.rawOptions)
    }
  }
}
</script>

<style scoped>

</style>
