<template>
  <BasicDrawer v-bind="$attrs" @register="registerDrawer" isDetail width="61%">
    <template #title>
      <span class="mr-2">{{ '❅ 国家编码详情：' + entryData.code }}</span>
    </template>
    <Description :bordered="false" :schema="descSchema" :data="entryData" />
    <Divider />
    <Row>
      <Col>
        <Tag>Your code here:</Tag>
      </Col>
    </Row>
  </BasicDrawer>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { Row, Col, Tag, Divider } from 'ant-design-vue';
  import { BasicDrawer, useDrawerInner } from '/@/components/Drawer';
  import { Description } from '/@/components/Description';
  import { descSchema } from './data';

  const entryData = ref<Recordable>({});

  const [registerDrawer, { setDrawerProps }] = useDrawerInner(async (data) => {
    setDrawerProps({ loading: true });
    entryData.value = data.record;
    setDrawerProps({ loading: false });
  });
</script>
