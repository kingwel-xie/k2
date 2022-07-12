
<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="88px">
          <el-form-item label="代码" prop="code">
            <el-input v-model="queryParams.code" placeholder="请输入代码" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="三字符代码" prop="code2">
            <el-input v-model="queryParams.code2" placeholder="请输入三字符代码" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="数字代码" prop="digitCode">
            <el-input v-model="queryParams.digitCode" placeholder="请输入数字代码" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="电话代码" prop="teleCode">
            <el-input v-model="queryParams.teleCode" placeholder="请输入电话代码" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="分组" prop="group">
            <el-input v-model="queryParams.group" placeholder="请输入分组" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="中文简称" prop="nameCN">
            <el-input v-model="queryParams.nameCN" placeholder="请输入中文简称" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="英文简称" prop="nameEN">
            <el-input v-model="queryParams.nameEN" placeholder="请输入英文简称" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item label="描述" prop="remark">
            <el-input v-model="queryParams.remark" placeholder="请输入描述" clearable size="small" @keyup.enter.native="handleQuery" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
            <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
          </el-form-item>
        </el-form>

        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button v-permisaction="['kobh:tbxCountry:add']" type="primary" icon="el-icon-plus" size="mini" @click="handleAdd">新增</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button v-permisaction="['kobh:tbxCountry:edit']" type="success" icon="el-icon-edit" size="mini" :disabled="single" @click="handleUpdate">修改</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button v-permisaction="['kobh:tbxCountry:remove']" type="danger" icon="el-icon-delete" size="mini" :disabled="multiple" @click="handleDelete">删除</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-dropdown v-permisaction="['kobh:tbxCountry:list']" size="mini" @command="handleExport">
              <el-button type="warning" icon="el-icon-download" size="mini">
                导出...<i class="el-icon-arrow-down el-icon--right" />
              </el-button>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item v-if="!multiple" command="0">当前选中</el-dropdown-item>
                <el-dropdown-item command="1">当前页</el-dropdown-item>
                <el-dropdown-item command="2">按查询条件</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </el-col>
        </el-row>

        <el-table
          ref="mainTable"
          v-loading="loading"
          element-loading-text="加载中..."
          element-loading-spinner="el-icon-loading"
          :data="tbxCountryList"
          row-key="code"
          :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
          stripe
          border
          highlight-current-row
          @selection-change="handleSelectionChange"
          @sort-change="handleSortChange"
        >
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column label="代码" align="center" prop="code" :show-overflow-tooltip="true" />
          <el-table-column label="中文简称" align="center" prop="nameCN" :show-overflow-tooltip="true" />
          <el-table-column label="英文简称" align="center" prop="nameEN" :show-overflow-tooltip="true" />
          <el-table-column label="三字符代码" align="center" prop="code2" :show-overflow-tooltip="true" />
          <el-table-column label="分组" align="center" prop="group" :show-overflow-tooltip="true" />
          <el-table-column v-fmt.dict="'TbxCountry'" label="从属" align="center" prop="belongTo" :show-overflow-tooltip="true" />
          <el-table-column label="数字代码" align="center" prop="digitCode" :show-overflow-tooltip="true" />
          <el-table-column label="电话代码" align="center" prop="teleCode" :show-overflow-tooltip="true" />
          <el-table-column label="描述" align="center" prop="remark" :show-overflow-tooltip="true" />
          <el-table-column label="显示排序" align="center" prop="displaySort" :show-overflow-tooltip="true" />
          <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="110px">
            <template slot-scope="scope">
              <el-button v-permisaction="['kobh:tbxCountry:edit']" size="mini" type="text" icon="el-icon-edit" @click="handleUpdate(scope.row)">修改</el-button>
              <el-button v-permisaction="['kobh:tbxCountry:remove']" size="mini" type="text" icon="el-icon-delete" @click="handleDelete(scope.row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>

        <pagination v-show="total>0" :total="total" :page.sync="queryParams.pageIndex" :limit.sync="queryParams.pageSize" @pagination="getList" />

        <!-- 添加或修改对话框 -->
        <el-dialog :title="title" :visible.sync="open" width="500px">
          <el-form ref="form" v-ffiov :model="form" :rules="rules" label-width="100px">
            <el-row :gutter="10" class="mb8">
              <el-col :span="24">
                <el-form-item label="代码" prop="code">
                  <el-input v-model="form.code" placeholder="代码" :disabled="isEdit" />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="三字符代码" prop="code2">
                  <el-input v-model="form.code2" placeholder="三字符代码" />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="中文简称" prop="nameCN">
                  <el-input v-model="form.nameCN" placeholder="中文简称" />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item prop="nameEN">
                  <span slot="label">英文简称<el-tooltip content="Country name in English" placement="top"><i class="el-icon-info" /></el-tooltip></span>
                  <el-input v-model="form.nameEN" placeholder="英文简称" />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="分组" prop="group">
                  <el-input v-model="form.group" placeholder="分组" />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="从属" prop="belongTo">
                  <DictSelect v-model="form.belongTo" dict="TbxCountry" placeholder="从属" filterable />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="数字代码" prop="digitCode">
                  <el-input v-model="form.digitCode" placeholder="数字代码" />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="电话代码" prop="teleCode">
                  <el-input v-model="form.teleCode" placeholder="电话代码" />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="显示排序" prop="displaySort">
                  <el-input-number v-model="form.displaySort" :precision="0" placeholder="显示排序" />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="描述" prop="remark">
                  <el-input v-model="form.remark" type="textarea" :rows="2" placeholder="请输入内容" />
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" :loading="saving" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import { addTbxCountry, delTbxCountry, getTbxCountry, listTbxCountry, updateTbxCountry } from '@/api/kobh/tbx-country'

import { formatJson } from '@/utils'

export default {
  name: 'TbxCountry',
  components: {
  },
  data() {
    return {
      // 遮罩层
      loading: true,
      saving: false,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      isEdit: false,
      // 类型数据字典
      tbxCountryList: [],
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        code: undefined,
        code2: undefined,
        digitCode: undefined,
        teleCode: undefined,
        group: undefined,
        nameCN: undefined,
        nameEN: undefined,
        remark: undefined

      },
      // 表单参数
      form: {
      },
      // 表单校验
      rules: {
        code: [{ required: true, message: '代码不能为空', trigger: 'blur' }],
        code2: [{ required: true, message: '三字符代码不能为空', trigger: 'blur' }],
        // group: [{ required: true, message: '分组不能为空', trigger: 'blur' }],
        nameCN: [{ required: true, message: '中文简称不能为空', trigger: 'blur' }],
        nameEN: [{ required: true, message: '英文简称不能为空', trigger: 'blur' }]
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
      listTbxCountry(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.tbxCountryList = response.data.list
        this.total = response.data.count
        this.loading = false
      }).catch(() => {
        this.loading = false
      })
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.form = {
        code: undefined,
        code2: undefined,
        digitCode: undefined,
        teleCode: undefined,
        group: undefined,
        belongTo: undefined,
        nameCN: undefined,
        nameEN: undefined,
        displaySort: undefined,
        remark: undefined
      }
      this.resetForm('form')
    },
    // 文件
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = []
      this.resetForm('queryForm')
      this.handleQuery()
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset()
      this.open = true
      this.title = '添加国家编码'
      this.isEdit = false
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.code)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 排序回调函数 */
    handleSortChange({ column, prop, order }) {
      Object.keys(this.queryParams).forEach(x => {
        if (x.length > 5 && x.endsWith('Order')) {
          delete this.queryParams[x]
        }
      })
      if (order === 'descending') {
        this.queryParams[prop + 'Order'] = 'desc'
      } else if (order === 'ascending') {
        this.queryParams[prop + 'Order'] = 'asc'
      } else {
        this.queryParams[prop + 'Order'] = undefined
      }
      this.getList()
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      const code = row.code || this.ids
      this.loading = true
      getTbxCountry(code).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改国家编码'
        this.isEdit = true
        this.loading = false
      }).catch(() => {
        this.loading = false
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          this.saving = true
          if (this.isEdit) {
            updateTbxCountry(this.form).then(response => {
              this.msgSuccess(response.msg)
              this.open = false
              this.saving = false
              this.getList()
            }).catch(() => {
              this.saving = false
            })
          } else {
            addTbxCountry(this.form).then(response => {
              this.msgSuccess(response.msg)
              this.open = false
              this.saving = false
              this.getList()
            }).catch(() => {
              this.saving = false
            })
          }
        }
      })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const Ids = (row.code && [row.code]) || this.ids

      this.$confirm('是否确认删除编号为"' + Ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return delTbxCountry({ 'ids': Ids })
      }).then((response) => {
        this.msgSuccess(response.msg)
        this.open = false
        this.getList()
      }).catch(() => {})
    },
    export2Excel(data) {
      const tHeader = ['代码','三字符代码','数字代码','电话代码','分组','从属','中文简称','英文简称','显示排序','描述',]
      const filterVal = ['code','code2','digitCode','teleCode','group','belongTo','nameCN','nameEN','displaySort','remark',]
      const filename = '国家编码'
      const filtered = formatJson(filterVal, data)
      import('@/vendor/Export2Excel').then(excel => {
        excel.export_json_to_excel({
          header: tHeader,
          data: filtered,
          filename: filename,
          autoWidth: true, // Optional
          bookType: 'xlsx' // Optional
        })
      })
    },
    /** 导出按钮操作 */
    handleExport(choice) {
      switch (choice) {
        case '0':
          this.export2Excel(JSON.parse(JSON.stringify(this.$refs.mainTable.selection)))
          break
        case '1':
          this.export2Excel(JSON.parse(JSON.stringify(this.tbxCountryList)))
          break
        case '2':
          this.$confirm('是否导出满足当前查询条件的数据项（最多10000项）?', '请确认', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            const params = Object.assign({}, this.queryParams)
            params.pageIndex = 1
            params.pageSize = 10000
            this.loading = true
            listTbxCountry(this.addDateRange(params, this.dateRange)).then(response => {
              this.loading = false
              this.export2Excel(response.data.list)
            }).catch(_ => {
              this.loading = false
            })
          })
          break
      }
    }
  }
}
</script>
