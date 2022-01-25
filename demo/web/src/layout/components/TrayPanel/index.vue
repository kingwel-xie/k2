<template>
  <transition :name="transitionName">
    <div v-show="visible" class="tray-panel">
      <ErrorLog />
    </div>
  </transition>
</template>

<script>

import ErrorLog from '@/layout/components/ErrorLog'

export default {
  name: 'TrayPanel',
  components: { ErrorLog },
  props: {
    transitionName: {
      type: String,
      default: 'fade'
    }
  },
  data() {
    return {
      visible: true
    }
  }
}
</script>

<style scoped>
.tray-panel {
  right: 10px;
  bottom: 10px;
  width: 40px;
  height: 40px;
  border-radius: 4px;
  line-height: 45px;
  position: fixed;
  display: inline-block;
  text-align: center;
  cursor: pointer;
  z-index: 999999;
}

</style>
