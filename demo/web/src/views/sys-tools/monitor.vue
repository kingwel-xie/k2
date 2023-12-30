<template>
  <PageWrapper v-loading="loading">
    <div class="lg:w-7/10 w-full">
      <Row :gutter="10" class="mb10">
        <Col :sm="24" :md="8">
          <Card v-if="info.cpu" title="CPU使用率">
            <div class="monitor">
              <div class="monitor-header">
                <Progress type="circle" :percent="info.cpu.Percent" />
              </div>
              <div class="monitor-footer" v-if="info.cpu.cpuInfo">
                <Cell label="CPU主频" :value="info.cpu.cpuInfo[0].modelName.split('@ ')[1]" />
                <Cell label="核心数" :value="`${info.cpu.cpuInfo[0].cores}`" />
              </div>
            </div>
          </Card>
        </Col>
        <Col :sm="24" :md="8">
          <Card v-if="info.mem" title="内存使用率">
            <div class="monitor">
              <div class="monitor-header">
                <Progress type="circle" :percent="info.mem.usage" />
              </div>
              <div class="monitor-footer">
                <Cell label="总内存" :value="info.mem.total + 'G'" />
                <Cell label="已用内存" :value="info.mem.used + 'G'" />
              </div>
            </div>
          </Card>
        </Col>
        <Col :sm="24" :md="8">
          <Card v-if="info.disk" title="磁盘信息">
            <div class="monitor">
              <div class="monitor-header">
                <Progress
                  type="circle"
                  :percent="
                    Number(
                      (((info.disk.total - info.disk.free) / info.disk.total) * 100).toFixed(2),
                    )
                  "
                />
              </div>
              <div class="monitor-footer">
                <Cell label="总磁盘" :value="info.disk.total + 'G'" />
                <Cell label="已用磁盘" :value="info.disk.total - info.disk.free + 'G'" />
              </div>
            </div>
          </Card>
        </Col>
      </Row>

      <Card v-if="info.app" title="应用信息">
        <div class="monitor">
          <Cell label="租户" :value="info.app.tenant" />
          <Cell label="版本" :value="info.app.version" />
          <Cell label="环境" :value="info.app.env" />
          <Cell label="GitCommit" :value="info.app.gitCommit" />
          <Cell label="构建时间" :value="info.app.buildTime" />
        </div>
      </Card>

      <Card v-if="info.os" title="go运行环境">
        <div class="monitor">
          <Cell label="GO 版本" :value="info.os.version" />
          <Cell label="Goroutine" :value="`${info.os.numGoroutine}`" />
          <Cell label="项目地址" :value="info.os.projectDir" />
        </div>
      </Card>

      <Card v-if="info.os" title="服务器信息el-">
        <div class="monitor">
          <Cell label="主机名称" :value="info.os.hostName" />
          <Cell label="操作系统" :value="info.os.goOs" />
          <Cell label="服务器IP" :value="info.os.ip" />
          <Cell label="系统架构" :value="info.os.arch" />
          <Cell label="CPU" v-if="info.cpu.cpuInfo" :value="info.cpu.cpuInfo[0].modelName" />
          <Cell label="当前时间" :value="info.os.time" />
        </div>
      </Card>

      <Card title="磁盘状态">
        <Table :columns="columns" :dataSource="info.diskList" :pagination="false" />
      </Card>
    </div>
  </PageWrapper>
</template>

<script lang="tsx" setup>
  import { Card, Row, Col, Progress, Tag, Table } from 'ant-design-vue';
  import { PageWrapper } from '/@/components/Page';

  import { onMounted, ref } from 'vue';
  import { getServerInfo } from '/@/api/sys/monitor';
  import { BasicColumn } from '/@/components/Table';
  import { formatFloat } from '/@/utils/formatUtil';
  import Cell from './cell.vue';

  const columns: BasicColumn[] = [
    {
      title: '盘符路径',
      dataIndex: 'path',
    },
    {
      title: '文件系统',
      dataIndex: 'fstype',
    },
    {
      title: '总大小',
      dataIndex: 'total',
    },
    {
      title: '可用大小',
      dataIndex: 'free',
    },
    {
      title: '已用大小',
      dataIndex: 'used',
    },
    {
      title: '已用百分比',
      dataIndex: 'percentage',
      customRender: ({ record }) => {
        const percent = (record.used * 100) / record.total;
        const color = percent > 80 ? 'red' : 'green';
        return <Tag color={color}>{() => formatFloat(percent) + '%'}</Tag>;
      },
    },
  ];

  const loading = ref(true);
  const info = ref({});

  onMounted(async () => {
    info.value = await getServerInfo();
    console.log(info.value);
    loading.value = false;
  });
</script>
