<template>
  <div class="upload-container">
    <el-upload
      :action="url"
      :headers="headers"
      :data="dataObj"
      :multiple="false"
      :show-file-list="false"
      :before-upload="beforeUpload"
      :on-success="handleImageSuccess"
      accept="image/*"
      class="image-uploader"
      drag
    >
      <i class="el-icon-upload" />
      <div class="el-upload__text">
        将文件拖到此处，或<em>点击上传</em>
      </div>
    </el-upload>
    <div class="image-preview">
      <div v-show="imageUrl.length>1" class="image-preview-wrapper">
        <img :src="imageUrl">
        <div class="image-preview-action">
          <i class="el-icon-zoom-in" @click="previewImage" />
          <i class="el-icon-delete" @click="rmImage" />
        </div>
      </div>
    </div>
    <K2Dialog :visible.sync="dialogVisible" title="预览" append-to-body>
      <img width="100%" :src="imageUrl" alt="">
    </K2Dialog>
  </div>
</template>

<script>

import { getToken } from '@/utils/auth'

export default {
  name: 'SingleImageUpload',
  props: {
    value: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      url: process.env.VUE_APP_BASE_API + '/api/v1/public/uploadFile',
      headers: { 'Authorization': 'Bearer ' + getToken() },
      dataObj: { type: '1', category: 'image' },
      dialogVisible: false
    }
  },
  computed: {
    imageUrl() {
      return this.value
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
    rmImage() {
      this.emitInput('')
    },
    previewImage() {
      this.dialogVisible = true
    },
    emitInput(val) {
      this.$emit('input', val)
    },
    handleImageSuccess(res) {
      this.emitInput(res.data.full_path)
    }
  }
}
</script>

<style lang="scss" scoped>
    @import "~@/styles/mixin.scss";
    .upload-container {
        width: 100%;
        position: relative;
        @include clearfix;
        .image-uploader {
            width: 60%;
            float: left;
        }
        .image-preview {
            width: 200px;
            height: 200px;
            position: relative;
            border: 1px dashed #d9d9d9;
            float: left;
            margin-left: 50px;
            .image-preview-wrapper {
                position: relative;
                width: 100%;
                height: 100%;
                img {
                    width: 100%;
                    height: 100%;
                }
            }
            .image-preview-action {
                position: absolute;
                width: 100%;
                height: 100%;
                left: 0;
                top: 0;
                cursor: default;
                text-align: center;
                color: #fff;
                opacity: 0;
                font-size: 20px;
                background-color: rgba(0, 0, 0, .5);
                transition: opacity .3s;
                cursor: pointer;
                text-align: center;
                line-height: 200px;
                .el-icon-zoom-in {
                  font-size: 36px;
                };
                .el-icon-delete {
                    font-size: 36px;
                }
            }
            &:hover {
                .image-preview-action {
                    opacity: 1;
                }
            }
        }
    }

</style>
