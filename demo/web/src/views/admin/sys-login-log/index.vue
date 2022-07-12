
<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
          <el-form-item label="用户名" prop="username"><el-input
            v-model="queryParams.username"
            placeholder="请输入用户名"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
          />
          </el-form-item>
          <el-form-item label="状态" prop="status">
            <DictSelect v-model="queryParams.status" dict="sys_normal_disable" placeholder="系统登录日志状态" clearable size="small" />
          </el-form-item>
          <el-form-item label="ip地址" prop="ipaddr"><el-input
            v-model="queryParams.ipaddr"
            placeholder="请输入ip地址"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
          />
          </el-form-item>
          <el-form-item label="登录时间">
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
              v-permisaction="['admin:sysLoginLog:remove']"
              type="danger"
              icon="el-icon-delete"
              size="mini"
              :disabled="multiple"
              @click="handleDelete"
            >删除
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-dropdown v-permisaction="['admin:sysLoginLog:list']" size="mini" @command="handleExport">
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

        <el-table ref="mainTable" v-loading="loading" :data="sysloginlogList" border stripe highlight-current-row @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column label="编号" width="70" prop="id" />
          <el-table-column
            label="用户名"
            align="center"
            prop="username"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="类型"
            align="center"
            prop="msg"
            :show-overflow-tooltip="true"
          />
          <el-table-column label="状态" align="center" prop="status">
            <template slot-scope="scope">
              <el-tag
                :type="scope.row.status === '2' ? 'success' : 'danger'"
                disable-transitions
              >{{ scope.row.status | dict('sys_normal_disable') }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column
            label="ip地址"
            align="center"
            prop="ipaddr"
          >
            <template slot-scope="scope">
              <el-popover trigger="hover" placement="top">
                <p>IP: {{ scope.row.ipaddr }}</p>
                <p>归属地: {{ scope.row.loginLocation }}</p>
                <p>浏览器: {{ scope.row.browser }}</p>
                <p>系统: {{ scope.row.os }}</p>
                <p>固件: {{ scope.row.platform }}</p>
                <div slot="reference" class="name-wrapper">
                  {{ scope.row.ipaddr }}
                </div>
              </el-popover>
            </template>
          </el-table-column>

          <el-table-column
            label="登录时间"
            align="center"
            prop="loginTime"
            width="180"
          >
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.loginTime) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="110px">
            <template slot-scope="scope">
              <el-button
                v-permisaction="['admin:sysLoginLog:remove']"
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleDelete(scope.row)"
              >删除
              </el-button>
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
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import { delSysLoginlog, getSysLoginlog, listSysLoginlog } from '@/api/admin/sys-login-log'
import { formatJson } from '@/utils'

export default {
  name: 'SysLoginLogManage',
  components: {
  },
  data() {
    return {
      // 遮罩层
      loading: true,
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
      fileOpen: false,
      fileIndex: undefined,
      // 类型数据字典
      typeOptions: [],
      sysloginlogList: [],
      // 日期范围
      dateRange: [],
      // 关系表类型

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        username: undefined,
        status: undefined,
        ipaddr: undefined,
        loginLocation: undefined
      },
      // 表单参数
      form: {
      },
      // 表单校验
      rules: {
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
      listSysLoginlog(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.sysloginlogList = response.data.list
        this.total = response.data.count
        this.loading = false
      }).catch(_ => {
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
        ID: undefined,
        username: undefined,
        status: undefined,
        ipaddr: undefined,
        loginLocation: undefined,
        browser: undefined,
        os: undefined,
        platform: undefined,
        loginTime: undefined,
        remark: undefined,
        msg: undefined
      }
      this.resetForm('form')
    },
    getImgList: function() {
      this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
    },
    fileClose: function() {
      this.fileOpen = false
    },
    // 关系
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
      this.title = '添加系统登录日志'
      this.isEdit = false
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },

    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      const ID =
                row.id || this.ids
      getSysLoginlog(ID).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改系统登录日志'
        this.isEdit = true
      })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      var Ids = (row.id && [row.id]) || this.ids

      this.$confirm('是否确认删除编号为"' + Ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return delSysLoginlog({ 'ids': Ids })
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
      const tHeader = ['日志编号', '用户名', 'ip', '状态', 'os', '平台', '浏览器', '信息', '备注', '登录日期']
      const filterVal = ['id', 'username', 'ipaddr', 'status', 'os', 'platform', 'browser', 'msg', 'remark', 'loginTime']
      const filename = '登录日志'
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
          this.export2Excel(JSON.parse(JSON.stringify(this.sysloginlogList)))
          break
        case '2':
          this.$confirm('请确认是否导出所有登录日志数据项（最多10000项）?', '警告', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            const params = Object.assign({}, this.queryParams)
            params.pageIndex = 1
            params.pageSize = 10000
            this.loading = true
            listSysLoginlog(this.addDateRange(params, this.dateRange)).then(response => {
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
