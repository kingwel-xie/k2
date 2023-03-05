<template>
  <Popover title="" trigger="click">
    <template #content>
      <List item-layout="horizontal" :data-source="receiverList">
        <template #renderItem="{ item }">
          <ListItem>
            {{ item }}
          </ListItem>
        </template>
      </List>
    </template>
    <span class="mx-1 underline">
      {{ receiverList.length > 1 ? receiverList[0] + '...' : receiverList[0] || '?' }}
    </span>
  </Popover>
</template>
<script lang="ts" setup>
  import { computed } from 'vue';
  import { Popover, List, ListItem } from 'ant-design-vue';
  import { tryParseJson } from '/@/utils/formatUtil';

  const props = defineProps({
    receivers: {
      type: String as PropType<string>,
      default: '',
    },
  });

  const receiverList = computed(() => {
    return tryParseJson(props.receivers) || [];
  });
</script>
