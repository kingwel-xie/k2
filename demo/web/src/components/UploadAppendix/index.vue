<template>
  <div>
    <el-row :gutter="10">
      <el-table :data="fileList" stripe>
        <el-table-column label="文件名" align="center" prop="name" :show-overflow-tooltip="true" />
        <el-table-column v-if="showType" label="类型" align="center" prop="type" :show-overflow-tooltip="true">
          <editable-cell v-model="row.type" slot-scope="{row}" :can-edit="true" editable-component="DictSelect" dict="tbx_attachment_type" placeholder="类型" close-event="change">
            <span slot="content">{{ row.type | dict('tbx_attachment_type') }}</span>
          </editable-cell>
        </el-table-column>
        <el-table-column v-if="showSize" label="大小" align="center" prop="size" :show-overflow-tooltip="true" :formatter="formatSize" />
        <!--            <el-table-column label="URL" align="center" prop="url" :show-overflow-tooltip="true" />-->
        <el-table-column label="预览" align="center" min-width="50px">
          <template slot-scope="{row}">
            <el-button v-if="row.contentType && row.contentType.indexOf('image') !== -1" size="mini" type="text" icon="el-icon-search" @click="handlePreview(row)">预览</el-button>
          </template>
        </el-table-column>
        <el-table-column label="备注" align="center" prop="type" :show-overflow-tooltip="true" min-width="120px">
          <editable-cell v-model="row.remark" slot-scope="{row}" :can-edit="true" editable-component="el-input" type="textarea" placeholder="请输入备注">
            <span slot="content">{{ row.remark }}</span>
          </editable-cell>
        </el-table-column>
        <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
          <template slot-scope="{row}">
            <el-button size="mini" type="text" icon="el-icon-close" @click="handleRemove(row)">删除</el-button>
            <el-button size="mini" type="text" icon="el-icon-download" @click="handleDownload(row)">下载</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-upload
        class="mt20"
        multiple
        :action="url"
        :headers="headers"
        :data="dataObj"
        :show-file-list="false"
        :disabled="fileList.length===limit"
        :before-upload="beforeUpload"
        :on-success="handleSuccess"
      >
        <el-button slot="trigger" size="small" type="primary" :disabled="fileList.length===limit">点击上传</el-button>
        <div slot="tip" class="el-upload__tip">{{ '最多 ' + this.limit + '个文件，每个不超过 ' + formatSize(0, 0, this.sizeLimit) }}</div>
        <K2Dialog :visible.sync="dialogVisible" :title="imagePreviewTitle" append-to-body>
          <!--              <img width="100%" :src="imageUrl" alt="">-->
          <el-image :src="imageUrl" />
        </K2Dialog>
      </el-upload>
    </el-row>
  </div>
</template>
<script>

import { getToken } from '@/utils/auth'
import { formatFileSize, tryParseJson } from '@/utils'
import { downLoadFile } from '@/utils/zipdownload'
import EditableCell from '@/components/EditableCell'

export default {
  name: 'UploadAppendix',
  components: {
    EditableCell
  },
  props: {
    value: {
      type: String,
      default: ''
    },
    // category determines the destination path of the uploaded file
    // see k2/middleware/file for details
    category: {
      type: String,
      default: ''
    },
    showType: {
      type: Boolean,
      default: true
    },
    showSize: {
      type: Boolean,
      default: true
    },
    sizeLimit: {
      type: Number,
      default: 2 * 1024 * 1024
    },
    limit: {
      type: Number,
      default: 10
    }
  },
  data() {
    return {
      // upload action and header
      url: process.env.VUE_APP_BASE_API + '/api/v1/public/uploadFile',
      headers: { 'Authorization': 'Bearer ' + getToken() },
      dataObj: { type: '1', category: this.category },
      fileList: [],
      richText: '',
      dialogVisible: false,
      imageUrl: undefined,
      imagePreviewTitle: ''
    }
  },
  watch: {
    value: {
      immediate: true,
      handler(val) {
        this.fileList = tryParseJson(val) || []
      }
    },
    fileList: {
      deep: true,
      handler(val) {
        // console.log('re', val)
        this.$emit('input', JSON.stringify(val))
      }
    }
  },
  created() {
  },
  methods: {
    formatSize(r, c, value) {
      return formatFileSize(value)
    },
    beforeUpload(file) {
      if (this.fileList.length >= this.limit) {
        this.$message.error('最多' + this.limit + '个附件')
        return false
      }
      const isLt2M = file.size / this.sizeLimit
      if (!isLt2M) {
        this.$message.error('上传文件大小不能超过 ' + formatFileSize(this.sizeLimit))
        return false
      }
      return true
    },
    handleRemove(file) {
      // console.log('remove', file)
      this.fileList.splice(this.fileList.indexOf(file), 1)
    },
    handlePreview(file) {
      this.imageUrl = file.url
      this.imagePreviewTitle = '预览：' + file.name + '，' + formatFileSize(file.size)
      this.dialogVisible = true
    },
    handleDownload(file) {
      downLoadFile(file.url + '?as=' + file.name)
    },
    handleSuccess(res, file, fileList) {
      // console.log('success', res, file, fileList)
      const item = { name: file.name, size: file.size, type: 'generic', url: res.data.full_path, path: res.data.path, contentType: file.raw.type, remark: '' }
      this.fileList.push(item)
      this.fileList = this.fileList.slice(-this.limit)
    }
  }
}
</script>

<style scoped>
</style>
