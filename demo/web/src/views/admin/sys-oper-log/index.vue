<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
          <el-form-item label="状态" prop="apiCode">
            <el-input
              v-model="queryParams.apiCode"
              size="small"
              placeholder="请输入API Code"
              clearable
              style="width: 160px"
            />
          </el-form-item>
          <el-form-item label="操作人员" prop="operName">
            <el-input
              v-model="queryParams.operName"
              size="small"
              placeholder="请输入操作人员"
              clearable
              style="width: 160px"
            />
          </el-form-item>
          <el-form-item label="创建时间">
            <DatetimeRanger v-model="dateRange" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
            <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
          </el-form-item>
        </el-form>

        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:sysOperLog:remove']"
              type="danger"
              icon="el-icon-delete"
              size="mini"
              :disabled="multiple"
              @click="handleDelete"
            >删除</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-dropdown v-permisaction="['admin:sysOperLog:list']" size="mini" @command="handleExport">
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

        <el-table v-loading="loading" :data="list" border stripe highlight-current-row @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column label="编号" width="70" prop="id" />
          <el-table-column
            label="Request info"
            prop="operUrl"
          >
            <template slot-scope="scope">
              <el-popover trigger="hover" placement="top">
                <p>Request:
                  <el-tag v-if="scope.row.requestMethod=='GET'">{{ scope.row.requestMethod }}</el-tag>
                  <el-tag v-if="scope.row.requestMethod=='POST'" type="success">{{ scope.row.requestMethod }}</el-tag>
                  <el-tag v-if="scope.row.requestMethod=='PUT'" type="warning">{{ scope.row.requestMethod }}</el-tag>
                  <el-tag v-if="scope.row.requestMethod=='DELETE'" type="danger">{{ scope.row.requestMethod }}</el-tag>
                  {{ scope.row.operUrl }}
                </p>
                <p>Host: {{ scope.row.operIp }}</p>
                <p>Location: {{ scope.row.operLocation }}</p>
                <p>耗时: {{ scope.row.latencyTime }}</p>
                <div slot="reference" class="name-wrapper">
                  <el-tag v-if="scope.row.requestMethod=='GET'">{{ scope.row.requestMethod }}</el-tag>
                  <el-tag v-if="scope.row.requestMethod=='POST'" type="success">{{ scope.row.requestMethod }}</el-tag>
                  <el-tag v-if="scope.row.requestMethod=='PUT'" type="warning">{{ scope.row.requestMethod }}</el-tag>
                  <el-tag v-if="scope.row.requestMethod=='DELETE'" type="danger">{{ scope.row.requestMethod }}</el-tag>
                  {{ scope.row.operUrl }}
                </div>
              </el-popover>
            </template>
          </el-table-column>
          <el-table-column
            label="操作人员"
            prop="operName"
            width="160"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="IP"
            prop="operIp"
            width="160"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="API code"
            prop="apiCode"
            width="100"
            :show-overflow-tooltip="true"
          >
            <template slot-scope="scope">
              <el-tag
                :type="scope.row.apiCode === 200 ? 'success' : 'danger'"
                disable-transitions
              >{{ scope.row.apiCode }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="操作日期" prop="operTime" width="160">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.operTime) }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="操作"
            width="80"
            class-name="small-padding fixed-width"
          >
            <template slot-scope="scope">
              <el-button
                size="mini"
                type="text"
                icon="el-icon-view"
                @click="handleView(scope.row,scope.index)"
              >详细</el-button>
            </template>
          </el-table-column>
        </el-table>

        <pagination
          v-show="total>0"
          :total="total"
          :page.sync="queryParams.pageIndex"
          :limit.sync="queryParams.pageSize"
          @pagination="getList"
        />

        <!-- 操作日志详细 -->
        <el-dialog title="操作日志详细" :visible.sync="open" width="700px">
          <el-form ref="form" :model="form" label-width="100px" size="mini">
            <el-row>
              <el-col :span="24">
                <el-form-item label="请求地址：">{{ form.operUrl }}</el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item
                  label="登录信息："
                >{{ form.operName }} / {{ form.operIp }} / {{ form.operLocation }}</el-form-item>
              </el-col>

              <el-col :span="12">
                <el-form-item label="请求方式：">{{ form.requestMethod }}</el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="耗时：">{{ form.latencyTime }}</el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="请求参数：">{{ form.operParam }}</el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="返回参数：">{{ form.jsonResult }}</el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="API code：">
                  <span>{{ form.apiCode }}</span>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="操作时间：">{{ parseTime(form.operTime) }}</el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item v-if="form.apiCode !== 200" label="异常信息：">{{ form.errorMsg }}</el-form-item>
              </el-col>
            </el-row>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="open = false">关 闭</el-button>
          </div>
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import { listSysOperlog, delSysOperlog, cleanOperlog } from '@/api/admin/sys-opera-log'
import { formatJson } from '@/utils'

export default {
  name: 'SysOperLogManage',
  data() {
    return {
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 表格数据
      list: [],
      // 是否显示弹出层
      open: false,
      // 日期范围
      dateRange: [],
      // 表单参数
      form: {},
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        title: undefined,
        operName: undefined,
        businessType: undefined,
        apiCode: undefined
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    /** 查询登录日志 */
    getList() {
      this.loading = true
      listSysOperlog(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.list = response.data.list
        this.total = response.data.count
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    },
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
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id)
      this.multiple = !selection.length
    },
    /** 详细按钮操作 */
    handleView(row) {
      this.open = true
      this.form = row
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const operIds = (row.id && [row.id]) || this.ids
      this.$confirm('是否确认删除日志编号为"' + operIds + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return delSysOperlog({ 'ids': operIds })
      }).then((response) => {
        if (response.code === 200) {
          this.msgSuccess(response.msg)
          this.open = false
          this.getList()
        } else {
          this.msgError(response.msg)
        }
      }).catch(function() {})
    },
    /** 清空按钮操作 */
    handleClean() {
      this.$confirm('是否确认清空所有操作日志数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return cleanOperlog()
      }).then((response) => {
        if (response.code === 200) {
          this.msgSuccess(response.msg)
          this.open = false
          this.getList()
        } else {
          this.msgError(response.msg)
        }
      }).catch(function() {})
    },
    export2Excel(data) {
      const tHeader = ['日志编号', '系统模块', '操作类型', '请求方式', '操作人员', '主机', '操作地点', 'ApiCode', '操作url', '操作日期']
      const filterVal = ['id', 'title', 'businessType', 'requestMethod', 'operName', 'operIp', 'operLocation', 'apiCode', 'operUrl', 'operTime']
      const filename = '操作日志'
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
          this.export2Excel(JSON.parse(JSON.stringify(this.list)))
          break
        case '2':
          this.$confirm('请确认是否导出所有操作日志数据项（最多10000项）?', '警告', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            const params = Object.assign({}, this.queryParams)
            params.pageIndex = 1
            params.pageSize = 10000
            this.loading = true
            listSysOperlog(this.addDateRange(params, this.dateRange)).then(response => {
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

