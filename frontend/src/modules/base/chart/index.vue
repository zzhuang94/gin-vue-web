<template>
  <a-row :gutter="8">
    <a-col :span="9"><Card icon="map" title="地图"><Map /></Card></a-col>
    <a-col :span="9"><Card icon="route" title="路线图"><Route /></Card></a-col>
    <a-col :span="6">
      <Card icon="chart-area" title="混合图">
        <template #more-op>
          <a href="https://echarts.apache.org/zh/index.html" style="color:yellow" target="_blank">
            <i class="fa fa-link"></i> <b>更多示例</b>
          </a>
        </template>
        <Mix />
      </Card>
    </a-col>
  </a-row>
  <a-row :gutter="8">
    <a-col :span="6"><Card icon="chart-line" title="曲线图"><Line /></Card></a-col>
    <a-col :span="6"><Card icon="chart-pie" title="饼图"><Pie /></Card></a-col>
    <a-col :span="6"><Card icon="chart-column" title="柱形图"><Bar /></Card></a-col>
    <a-col :span="6"><Card icon="chart-gantt" title="散点图"><Scatter /></Card></a-col>
  </a-row>
</template>

<script setup lang="ts">
import Card from '@components/card.vue'
import Line from './line.vue'
import Pie from './pie.vue'
import Bar from './bar.vue'
import Mix from './mix.vue'
import Map from './map.vue'
import Route from './route.vue'
import Scatter from './scatter.vue'
</script>
