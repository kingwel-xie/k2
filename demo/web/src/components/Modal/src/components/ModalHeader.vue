<template>
  <BasicTitle :helpMessage="helpMessage">
    <span class="modal-flag"></span>
    {{ title }}
  </BasicTitle>
</template>
<script lang="ts">
  import type { PropType } from 'vue';
  import { defineComponent } from 'vue';
  import { BasicTitle } from '/@/components/Basic';

  export default defineComponent({
    name: 'BasicModalHeader',
    components: { BasicTitle },
    props: {
      helpMessage: {
        type: [String, Array] as PropType<string | string[]>,
      },
      title: { type: String },
    },
    emits: ['dblclick'],
  });
</script>

<style lang="less">
  .modal-flag {
    display: inline-block;
    background-color: @primary-color;
    width: 3px;
    height: 20px;
    margin-right: 5px;
    float: left;
    margin-top: 2px;
  }
</style>
