<template>
  <div>
    <a-button type="link" @click="() => setVisible(true)"><slot>点击查看</slot></a-button>
    <Image
      v-bind="restAttrs"
      :style="{ display: 'none' }"
      :preview="{
        visible,
        onVisibleChange: setVisible,
      }"
      :src="imgSrc"
    />
  </div>
</template>
<script lang="ts" setup>
  import { ref, useAttrs } from 'vue';
  import { Image } from 'ant-design-vue';

  defineProps({
    visible: { type: Boolean },
  });

  const { src, ...restAttrs } = useAttrs();

  const imgSrc = ref('');
  const visible = ref<boolean>(false);
  const setVisible = (value): void => {
    if (!imgSrc.value) {
      imgSrc.value = src as string;
    }
    visible.value = value;
  };
</script>
