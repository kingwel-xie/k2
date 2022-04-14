<template>
  <div class="upload-container">
    <el-upload
      ref="upload"
      accept="image/*"
      class="image-uploader"
      :action="url"
      :headers="headers"
      :data="dataObj"
      :multiple="false"
      :show-file-list="false"
      :before-upload="beforeUpload"
      :on-success="handleImageSuccess">
      <el-button icon="el-icon-upload" size="mini" type="primary">上传</el-button>
    </el-upload>
  </div>
</template>

<script>

import { getToken } from '@/utils/auth'

export default {
  name: 'EditorSlideUpload',
  components: {
  },
  props: {
  },
  data() {
    return {
      url: process.env.VUE_APP_BASE_API + '/api/v1/public/uploadFile',
      headers: { 'Authorization': 'Bearer ' + getToken() },
      dataObj: { type: '1', category: 'tinymce' },
      fileList: []
    }
  },
  methods: {
    beforeUpload(file) {
      const isJPG = file.type.indexOf('image/') !== -1
      const isLt2M = file.size / 1024 / 1024 < 2

      if (!isJPG) {
        this.$message.error('上传文件只能是图片格式!')
        return false
      }
      if (!isLt2M) {
        this.$message.error('上传图片大小不能超过 2MB!')
        return false
      }
      return true
    },
    handleImageSuccess(res) {
      // this.$emit('input', res.data.full_path)
      this.$emit('successCBK', [res.data.full_path])
    }
  }
}
</script>

<style lang="scss" scoped>
</style>
