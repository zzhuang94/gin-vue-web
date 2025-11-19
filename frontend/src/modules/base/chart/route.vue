<template>
  <Chart china :option="option" margin="5px 0 0 0" height="400px" />
</template>

<script setup>
const coords = [
  { "name": "北京", "latitude": 39.9042, "longitude": 116.4074 },
  { "name": "天津", "latitude": 39.3434, "longitude": 117.3616 },
  { "name": "上海", "latitude": 31.2304, "longitude": 121.4737 },
  { "name": "重庆", "latitude": 29.5638, "longitude": 106.5506 },
  { "name": "河北", "capital": "石家庄", "latitude": 38.0428, "longitude": 114.5149 },
  { "name": "山西", "capital": "太原", "latitude": 37.8706, "longitude": 112.5489 },
  { "name": "辽宁", "capital": "沈阳", "latitude": 41.8057, "longitude": 123.4315 },
  { "name": "吉林", "capital": "长春", "latitude": 43.816, "longitude": 125.3235 },
  { "name": "黑龙江", "capital": "哈尔滨", "latitude": 45.8038, "longitude": 126.5349 },
  { "name": "江苏", "capital": "南京", "latitude": 32.0617, "longitude": 118.7969 },
  { "name": "浙江", "capital": "杭州", "latitude": 30.2741, "longitude": 120.1551 },
  { "name": "安徽", "capital": "合肥", "latitude": 31.8612, "longitude": 117.2857 },
  { "name": "福建", "capital": "福州", "latitude": 26.0745, "longitude": 119.2965 },
  { "name": "江西", "capital": "南昌", "latitude": 28.682, "longitude": 115.8579 },
  { "name": "山东", "capital": "济南", "latitude": 36.6512, "longitude": 117.1201 },
  { "name": "河南", "capital": "郑州", "latitude": 34.7466, "longitude": 113.6254 },
  { "name": "湖北", "capital": "武汉", "latitude": 30.5928, "longitude": 114.3055 },
  { "name": "湖南", "capital": "长沙", "latitude": 28.2278, "longitude": 112.9388 },
  { "name": "广东", "capital": "广州", "latitude": 23.1291, "longitude": 113.2644 },
  { "name": "海南", "capital": "海口", "latitude": 20.044, "longitude": 110.1999 },
  { "name": "四川", "capital": "成都", "latitude": 30.5723, "longitude": 104.0665 },
  { "name": "贵州", "capital": "贵阳", "latitude": 26.647, "longitude": 106.6302 },
  { "name": "云南", "capital": "昆明", "latitude": 25.0443, "longitude": 102.7183 },
  { "name": "陕西", "capital": "西安", "latitude": 34.3416, "longitude": 108.9398 },
  { "name": "甘肃", "capital": "兰州", "latitude": 36.0611, "longitude": 103.8343 },
  { "name": "青海", "capital": "西宁", "latitude": 36.6171, "longitude": 101.7782 },
  { "name": "台湾", "capital": "台北", "latitude": 25.033, "longitude": 121.5654 },
  { "name": "内蒙古", "capital": "呼和浩特", "latitude": 40.8415, "longitude": 111.7510 },
  { "name": "广西", "capital": "南宁", "latitude": 22.8155, "longitude": 108.3275 },
  { "name": "西藏", "capital": "拉萨", "latitude": 29.6456, "longitude": 91.1409 },
  { "name": "宁夏", "capital": "银川", "latitude": 38.4872, "longitude": 106.2309 },
  { "name": "新疆", "capital": "乌鲁木齐", "latitude": 43.8256, "longitude": 87.6168 },
  { "name": "香港", "capital": "香港", "latitude": 22.3193, "longitude": 114.1694 },
  { "name": "澳门", "capital": "澳门", "latitude": 22.1987, "longitude": 113.5439 }
]

function calcSrc() {
  const ans = []
  for (let i = 0; i < 10; i++) {
    const c = coords[i]
    ans.push({ name: c.name, value: [c.longitude, c.latitude] })
  }
  return ans
}

function calcTar() {
  const ans = []
  for (let i = 20; i < 30; i++) {
    const c = coords[i]
    ans.push({ name: c.name, value: [c.longitude, c.latitude] })
  }
  return ans
}

function calcLine() {
  const ans = []
  for (let i = 0; i < 10; i++) {
    const src = coords[i]
    const tar = coords[i + 20]
    ans.push({
      name: src.name + ' -> ' + tar.name,
      lineStyle: {normal: {curveness: 0.3, color: randColor()}},
      coords: [ [ src.longitude, src.latitude ], [ tar.longitude, tar.latitude ] ],
    })
  }
  return ans
}

function randColor() {
 const rc = Math.floor(Math.random() * 0xffffff);
  return `#${rc.toString(16).padStart(6, '0')}`;
}

const option = {
  title: {text: '路线图', left: 'center', top: '10px'},
  tooltip: {trigger: 'item'},
  grid: { top: 10, left: 10, right: 10 },
  geo: {
    map: 'china',
    zoom: 1.25,
    itemStyle: {normal: {areaColor: 'lightyellow', borderColor: 'gray'}},
  },
  series: [
    {
      type: 'lines',
      zlevel: 2,
      effect: {show: true, period: 2, trailLength: 0.1, symbol: 'arrow', symbolSize: 8},
      data: calcLine(),
    },
    {
      type: 'effectScatter',
      coordinateSystem: 'geo',
      zlevel: 2,
      symbolSize: 6,
      label: {fontSize: 13, show: true, position: 'right', formatter: '{b}'},
      tooltip: {formatter: '<b>{b} -> {a}</b>'},
      data: calcSrc(),
    },
    {
      type: 'effectScatter',
      coordinateSystem: 'geo',
      zlevel: 2,
      symbolSize: 6,
      label: {fontSize: 13, show: true, position: 'right', formatter: '{b}'},
      tooltip: {formatter: '<b>{b} -> {a}</b>'},
      data: calcTar(),
    }
  ]
}
</script>
