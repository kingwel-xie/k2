<template>
  <PageWrapper v-loading="loading" title="生成任意分货标签" contentBackground contentClass="p-4">
    <div class="md:w-2/3 lg:w-1/2 w-full">
      <BasicForm @register="register" />
    </div>
    <Button v-loading="loading" type="primary" @click="handleSubmit"> 生成 </Button>
  </PageWrapper>
</template>
<script lang="ts" setup>
  import { ref } from 'vue';
  import { Button } from 'ant-design-vue';
  import { BasicForm, FormSchema, useForm } from '/@/components/Form/index';
  import { PageWrapper } from '/@/components/Page';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { genTbxLabelArbitraryCargo } from '/@/api/kobh/tbx-label';
  import { downloadByUrl } from '/@/utils/file/download';
  import { limitedDownloadUrl } from '/@/api/kobh/misc';

  const schemas: FormSchema[] = [
    {
      field: 'labelCode',
      label: '标签模板',
      component: 'DictSelect',
      componentProps: {
        dictName: 'TbxLabel',
        filter: (l) => l.type === 'cargo',
        showSearch: false,
      },
      required: true,
    },
    {
      field: 'customerOrderNo',
      label: '客户单号',
      component: 'Input',
      required: true,
    },
    {
      field: 'dstWarehouse',
      label: '目的仓',
      component: 'Input',
      required: true,
    },
    {
      field: 'numPackages',
      label: '总箱数',
      component: 'InputNumber',
      componentProps: {
        min: 1,
        max: 5000,
        precision: 0,
      },
      required: true,
    },
  ];

  const loading = ref(false);
  const { createMessage } = useMessage();

  const [register, { validate }] = useForm({
    labelWidth: 120,
    baseColProps: { span: 24 },
    schemas: schemas,
    showActionButtonGroup: false,
  });

  async function handleSubmit() {
    try {
      const values = await validate();
      Object.assign(values, {
        labelType: 'cargo',
      });

      loading.value = true;
      const data = await genTbxLabelArbitraryCargo(values);
      createMessage.success('请求的标签已生成');
      downloadByUrl({ url: limitedDownloadUrl(data) });
    } finally {
      loading.value = false;
    }
  }
</script>
