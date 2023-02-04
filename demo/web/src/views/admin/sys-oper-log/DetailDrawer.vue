<template>
  <BasicDrawer v-bind="$attrs" @register="registerDrawer" isDetail title="详情" width="50%">
    <template #title>
      <span class="mr-2">{{ '❅ 操作详情：' + operData.id }}</span>
    </template>
    <Row>
      <Col :span="24">
        <span>请求地址：{{ operData.operUrl }}</span>
      </Col>
      <Col :span="12">
        <span>登录信息：{{ operData.operName }} / {{ operData.operIp }} / {{ operData.operLocation }}</span>
      </Col>
      <Col :span="12">
        <span>请求方式：{{ operData.requestMethod }}</span>
      </Col>
      <Col :span="12">
        <span>耗时：{{ operData.latencyTime }}</span>
      </Col>
      <Col :span="24">
        <span>请求参数：{{ operData.operParam }}</span>
      </Col>
      <Col :span="24">
        <span>返回参数：{{ operData.jsonResult }}</span>
      </Col>
      <Col :span="12">
        <span>
          <span>API code：{{ operData.apiCode }}</span>
        </span>
      </Col>
      <Col :span="12">
        <span>操作时间：{{ formatToDateTime(operData.operTime) }}</span>
      </Col>
    </Row>
  </BasicDrawer>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { Row, Col, Tag, Divider } from 'ant-design-vue';
  import { BasicDrawer, useDrawerInner } from '/@/components/Drawer';
  import { formatToDateTime } from '/@/utils/dateUtil';

  const operData = ref<Recordable>({});

  const [registerDrawer, { setDrawerProps }] = useDrawerInner(async (data) => {
    setDrawerProps({ loading: true });
    operData.value = data.record;
    setDrawerProps({ loading: false });
  });
</script>
