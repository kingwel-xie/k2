<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form v-ffiov ref="form" :model="form" :rules="rules" label-width="80px" style="width: 66%">
          <el-row :gutter="10">
            <el-col :span="18">
              <el-form-item label="系统名称" prop="sys_app_name">
                <el-input v-model="form.sys_app_name" placeholder="请输入系统名称" clearable size="small" />
              </el-form-item>
            </el-col>
            <el-col :span="18">
              <el-form-item label="初始密码" prop="sys_user_initPassword">
                <el-input v-model="form.sys_user_initPassword" placeholder="请输入初始密码" prefix-icon="el-icon-key" clearable size="small" />
              </el-form-item>
            </el-col>
            <el-col :span="18">
              <el-form-item label="皮肤样式" prop="sys_index_skinName">
                <el-select v-model="form.sys_index_skinName" placeholder="皮肤样式" clearable size="small">
                  <el-option v-for="skin in skinOptions" :key="skin.value" :label="skin.label" :value="skin.value">{{ skin.label }}</el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="18">
              <el-form-item label="侧栏主题" prop="sys_index_sideTheme">
                <el-select v-model="form.sys_index_sideTheme" placeholder="侧栏主题" clearable size="small">
                  <el-option v-for="theme in themeOptions" :key="theme.value" :label="theme.label" :value="theme.value">{{ theme.label }}</el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="18">
              <el-form-item label="系统logo" prop="sys_app_logo">
                <SingleImageUpload v-model="form.sys_app_logo" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row class="mt10">
            <el-button type="primary" :loading="saving" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">重 置</el-button>
          </el-row>
        </el-form>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import { getSetConfig, updateSetConfig } from '@/api/admin/sys-config'
import SingleImageUpload from '@/components/Upload/SingleImage'

export default {
  name: 'SysConfigSet',
  components: {
    SingleImageUpload
  },
  data() {
    return {
      // 遮罩层
      loading: true,
      saving: false,
      themeOptions: [{
        'label': '深色主题',
        'value': 'theme-dark'
      }, {
        'label': '浅色主题',
        'value': 'theme-light'
      }],
      skinOptions: [{
        'label': '蓝色',
        'value': 'skin-blue'
      }, {
        'label': '绿色',
        'value': 'skin-green'
      }, {
        'label': '紫色',
        'value': 'skin-purple'
      }, {
        'label': '红色',
        'value': 'skin-red'
      }, {
        'label': '黄色',
        'value': 'skin-yellow'
      }],
      oldConfig: {},
      form: {},
      rules: {
        sys_app_name: [
          { required: true, message: '系统名称不能为空', trigger: 'blur' }
        ],
        sys_user_initPassword: [
          { required: true, message: '初始密码不能为空', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      getSetConfig().then(response => {
        this.oldConfig = response.data
        this.form = Object.assign({}, this.oldConfig)
        this.loading = false
      })
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1
      this.getList()
    },
    cancel() {
      // reset data
      this.form = Object.assign({}, this.oldConfig)
    },
    submitForm() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          this.saving = true
          const list = []
          for (const key in this.form) {
            list.push({ 'configKey': key, 'configValue': this.form[key] })
          }
          updateSetConfig(list).then(response => {
            this.msgSuccess(response.msg)
            this.open = false
            this.saving = false
            this.getList()
            const data = Object.assign({}, this.form)
            this.$store.commit('system/SET_INFO', data)
          }).catch(() => {
            this.saving = false
          })
        }
      })
    }
  }
}
</script>
