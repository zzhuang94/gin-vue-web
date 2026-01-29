<template>
  <Chart :option="option" height="300px" />
</template>

<script setup lang="ts">
let base = +new Date(1988, 9, 3);
let oneDay = 24 * 3600 * 1000;
let data: [number, number][] = [[base, Math.random() * 300]];
for (let i = 1; i < 20000; i++) {
  let now = new Date((base += oneDay));
  const prevValue = data[i - 1]?.[1] ?? 0
  data.push([+now, Math.round((Math.random() - 0.5) * 20 + prevValue)]);
}
const option: Record<string, any> = {
  tooltip: {
    trigger: 'axis',
    position: function (pt: number[]) {
      return [pt[0], '10%'];
    }
  },
  grid: { left: 10, right: 10, top: 20 },
  title: {
    left: 'center',
    text: 'Large Ara Chart'
  },
  xAxis: {
    type: 'time',
    boundaryGap: false
  },
  yAxis: {
    type: 'value',
    boundaryGap: [0, '100%']
  },
  dataZoom: [
    {
      type: 'inside',
      start: 0,
      end: 20
    },
    {
      start: 0,
      end: 20
    }
  ],
  series: [
    {
      name: 'Fake Data',
      type: 'line',
      smooth: true,
      symbol: 'none',
      areaStyle: {},
      data: data
    }
  ]
};
</script>
