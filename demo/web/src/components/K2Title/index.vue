<template>
  <div v-bind="$attrs" class="el-descriptions__title mb8">
    <template>
      <span>
        <span :class="className" />
        <slot></slot>
      </span>
    </template>
  </div>
</template>

<script>
export default {
  name: 'K2Title',
  props: {
    title: {
      type: String,
      default: undefined
    },
    red: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      className: this.red ? 'red-block' : 'blue-block'
    }
  },
  methods: {
  }
}
</script>

<style lang="scss" scoped>

.blue-block {
  display: inline-block;
  background-color: #81a3e5;
  width:2px;
  height:16px;
  margin-right:5px;
  float: left;
  margin-top:2px;
}

.red-block {
  display: inline-block;
  background-color: #e581a4;
  width:2px;
  height:16px;
  margin-right:5px;
  float: left;
  margin-top:2px;
}

</style>
