<template>
  <div v-if="errorLogs.length>0 || messageLogs.length>0">
    <el-badge :is-dot="true" style="line-height: 25px;margin-top: -5px;" @click.native="dialogTableVisible=true">
      <el-button style="padding: 8px 10px;" size="small" type="danger">
        <svg-icon icon-class="bug" />
      </el-button>
    </el-badge>

    <el-dialog :visible.sync="dialogTableVisible" width="68%" append-to-body>
      <div slot="title">
        <span style="padding-right: 10px;">Log</span>
        <el-button size="mini" type="primary" icon="el-icon-delete" @click="clearAll">Clear All</el-button>
      </div>
      <el-tabs type="border-card">
        <el-tab-pane label="Message">
          <el-tag class="mb8">
            Messages: {{ messageLogs.length }}
          </el-tag>
          <el-table :data="messageLogs" border max-height="450px">
            <el-table-column label="Code">
              <template slot-scope="{row}">
                <div>
                  <span class="message-title" style="padding-right: 10px;">Code: </span>
                  <el-tag type="warning">
                    {{ row.code }}
                  </el-tag>
                </div>
                <div>
                  <span class="message-title" style="padding-right: 16px;">Url: </span>
                  <el-tag type="success">
                    {{ row.url }}
                  </el-tag>
                </div>
              </template>
            </el-table-column>
            <el-table-column label="Message">
              <template slot-scope="scope">
                {{ scope.row.msg }}
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
        <el-tab-pane label="Error">
          <el-tag class="mb8">
            Errors: {{ errorLogs.length }}
          </el-tag>
          <el-table :data="errorLogs" border max-height="450px">
            <el-table-column label="Message">
              <template slot-scope="{row}">
                <div>
                  <span class="message-title">Msg:</span>
                  <el-tag type="danger">
                    {{ row.err.message }}
                  </el-tag>
                </div>
                <div>
                  <span class="message-title" style="padding-right: 10px;">Info: </span>
                  <el-tag type="warning">
                    {{ row.vm.$vnode.tag }} error in {{ row.info }}
                  </el-tag>
                </div>
                <div>
                  <span class="message-title" style="padding-right: 16px;">Url: </span>
                  <el-tag type="success">
                    {{ row.url }}
                  </el-tag>
                </div>
              </template>
            </el-table-column>
            <el-table-column label="Stack">
              <template slot-scope="scope">
                <ul class="list-group">
                  <li v-for="(item, index) in scope.row.err.stack.split('\n')" :key="index" class="list-group-item">
                    <span>{{ item }}</span>
                  </li>
                </ul>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
      </el-tabs>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'ErrorLog',
  data() {
    return {
      dialogTableVisible: false
    }
  },
  computed: {
    errorLogs() {
      return this.$store.getters.errorLogs
    },
    messageLogs() {
      return this.$store.getters.messageLogs
    }
  },
  methods: {
    clearAll() {
      this.dialogTableVisible = false
      this.$store.dispatch('errorLog/clearErrorLog')
      this.$store.dispatch('errorLog/clearMessageLog')
    }
  }
}
</script>

<style scoped>
.message-title {
  font-size: 16px;
  color: #333;
  font-weight: bold;
  padding-right: 8px;
}
</style>
