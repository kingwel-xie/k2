<template>
  <CollapseContainer title="修改密码" :canExpan="false">
    <div class="py-8 bg-white flex flex-col justify-center items-center">
      <BasicForm @register="register" />
      <div class="flex justify-center">
        <a-button @click="resetFields">{{ t('common.resetText') }}</a-button>
        <a-button class="!ml-4" type="primary" @click="handleSubmit">{{
          t('common.okText')
        }}</a-button>
      </div>
    </div>
  </CollapseContainer>
</template>
<script lang="ts">
  import { defineComponent } from 'vue';
  import { CollapseContainer } from '/@/components/Container';
  import { BasicForm, useForm } from '/@/components/Form';
  import { formSchema } from './pwd.data';
  import { updatePassword } from '/@/api/sys/user';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { useI18n } from '/@/hooks/web/useI18n';

  export default defineComponent({
    name: 'ChangePassword',
    components: { BasicForm, CollapseContainer },
    setup() {
      const { t } = useI18n();
      const { createSuccessModal } = useMessage();
      const [register, { validate, resetFields }] = useForm({
        size: 'large',
        baseColProps: { span: 24 },
        labelWidth: 100,
        showActionButtonGroup: false,
        schemas: formSchema,
      });

      async function handleSubmit() {
        try {
          const values = await validate();
          const { passwordOld, passwordNew } = values;
          await updatePassword({ oldPassword: passwordOld, newPassword: passwordNew });
          // console.log(passwordOld, passwordNew);
          createSuccessModal({
            title: 'done',
            content: '修改成功后下次请使用新密码登录！',
          });
        } catch (error) {}
      }
      return { register, resetFields, handleSubmit, t };
    },
  });
</script>
