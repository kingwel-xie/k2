<template>
  <div class="container" style="width: 90%">
    <Checkbox v-model:checked="state.checked">发送站内信</Checkbox>
    <Row v-if="state.checked" class="mt-3" :gutter="[10, 10]">
      <Col :span="3">
        <span>类别：</span>
      </Col>
      <Col :span="6">
        <DictSelect
          v-model:value="state.targetType"
          dictName="sys_notification_target_type"
          :showSearch="false"
          style="width: 90%"
          @change="state.targets = undefined"
        />
      </Col>
      <Col :span="3">
        <span>收件人：</span>
      </Col>
      <Col :span="12">
        <ApiTreeSelect
          v-if="state.targetType === 'dept'"
          v-model:value="state.targets"
          :api="getDeptTree"
          style="width: 100%"
          allowClear
          treeCheckable
          :showCheckedStrategy="TreeSelect.SHOW_ALL"
          :field-names="{ label: 'label', value: 'id', options: 'children' }"
        />
        <ApiSelect
          v-else-if="state.targetType === 'user'"
          v-model:value="state.targets"
          :api="listAccountNoCheck"
          allowClear
          showSearch
          mode="multiple"
          style="width: 100%"
          optionFilterProp="label"
          resultField="list"
          labelField="username"
          valueField="username"
        />
        <ApiSelect
          v-else-if="state.targetType === 'role'"
          v-model:value="state.targets"
          :api="listRoleNoCheck"
          allowClear
          showSearch
          mode="multiple"
          style="width: 100%"
          optionFilterProp="label"
          resultField="list"
          labelField="roleName"
          valueField="roleKey"
        />
      </Col>
      <Col :span="3">
        <span>标题：</span>
      </Col>
      <Col :span="21">
        <Input v-model:value="state.title" showCount :maxlength="128" :disabled="!state.checked" />
      </Col>
      <Col :span="3">
        <span>内容：</span>
      </Col>
      <Col :span="21">
        <Textarea
          v-model:value="state.content"
          :rows="3"
          showCount
          :maxlength="512"
          :disabled="!state.checked"
        />
      </Col>
    </Row>
  </div>
</template>
<script lang="ts" setup>
  import { computed } from 'vue';
  import { Input, Checkbox, Textarea, Row, Col, TreeSelect } from 'ant-design-vue';
  import DictSelect from '/@/components/Form/src/components/DictSelect.vue';
  import ApiSelect from '/@/components/Form/src/components/ApiSelect.vue';
  import ApiTreeSelect from '/@/components/Form/src/components/ApiTreeSelect.vue';
  import { listAccountNoCheck, listRoleNoCheck, getDeptTree } from '/@/api/admin/system';
  import { SendMessageType } from './typing';

  const emit = defineEmits(['change', 'update:value']);
  const props = defineProps({
    value: {
      type: Object as PropType<SendMessageType>,
      default: () => {},
    },
  });

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
</script>
