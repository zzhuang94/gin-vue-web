<template>
  <Chart china :option="option" margin="5px 0 0 0" height="400px" />
</template>

<script setup lang="ts">
const data: Record<string, any> = {
  "上海": {
    "all_count": 4068,
    "city": {
      "上海": {
        "all_count": 4068,
        "sum": 1570764,
        "err_count": 3,
        "ok_count": 4065,
        "val": 386,
        "ok_rate": 99.926
      }
    },
    "sum": 1570764,
    "err_count": 3,
    "ok_count": 4065,
    "val": 386,
    "ok_rate": 99.926
  },
  "云南": {
    "all_count": 4094,
    "city": {
      "昆明": {
        "all_count": 1028,
        "sum": 504488,
        "err_count": 2,
        "ok_count": 1026,
        "val": 492,
        "ok_rate": 99.805
      },
      "曲靖": {
        "all_count": 941,
        "sum": 249553,
        "ok_count": 941,
        "err_count": 0,
        "val": 265,
        "ok_rate": 100
      },
      "昭通": {
        "all_count": 815,
        "sum": 329092,
        "ok_count": 815,
        "err_count": 0,
        "val": 404,
        "ok_rate": 100
      },
      "楚雄彝族自治州": {
        "all_count": 547,
        "sum": 457517,
        "err_count": 2,
        "ok_count": 545,
        "val": 839,
        "ok_rate": 99.634
      },
      "红河哈尼族彝族自治州": {
        "all_count": 181,
        "sum": 82942,
        "ok_count": 181,
        "err_count": 0,
        "val": 458,
        "ok_rate": 100
      },
      "玉溪": {
        "all_count": 163,
        "sum": 68920,
        "err_count": 1,
        "ok_count": 162,
        "val": 425,
        "ok_rate": 99.387
      },
      "大理白族自治州": {
        "all_count": 86,
        "sum": 39783,
        "ok_count": 86,
        "err_count": 0,
        "val": 463,
        "ok_rate": 100
      },
      "西双版纳傣族自治州": {
        "all_count": 80,
        "sum": 39522,
        "ok_count": 80,
        "err_count": 0,
        "val": 494,
        "ok_rate": 100
      },
      "普洱": {
        "all_count": 47,
        "sum": 31903,
        "ok_count": 47,
        "err_count": 0,
        "val": 679,
        "ok_rate": 100
      },
      "丽江": {
        "all_count": 41,
        "sum": 12112,
        "ok_count": 41,
        "err_count": 0,
        "val": 295,
        "ok_rate": 100
      },
      "保山": {
        "all_count": 39,
        "sum": 23867,
        "ok_count": 39,
        "err_count": 0,
        "val": 612,
        "ok_rate": 100
      },
      "德宏傣族景颇族自治州": {
        "all_count": 33,
        "sum": 16501,
        "ok_count": 33,
        "err_count": 0,
        "val": 500,
        "ok_rate": 100
      },
      "文山壮族苗族自治州": {
        "all_count": 32,
        "sum": 10012,
        "ok_count": 32,
        "err_count": 0,
        "val": 313,
        "ok_rate": 100
      },
      "怒江傈僳族自治州": {
        "all_count": 31,
        "sum": 12262,
        "ok_count": 31,
        "err_count": 0,
        "val": 396,
        "ok_rate": 100
      },
      "临沧": {
        "all_count": 30,
        "sum": 22401,
        "ok_count": 30,
        "err_count": 0,
        "val": 747,
        "ok_rate": 100
      }
    },
    "sum": 1900875,
    "err_count": 5,
    "ok_count": 4089,
    "val": 465,
    "ok_rate": 99.878
  },
  "内蒙古": {
    "all_count": 4306,
    "city": {
      "呼和浩特": {
        "all_count": 1095,
        "sum": 288447,
        "err_count": 2,
        "ok_count": 1093,
        "val": 264,
        "ok_rate": 99.817
      },
      "乌兰察布": {
        "all_count": 659,
        "sum": 589885,
        "ok_count": 659,
        "err_count": 0,
        "val": 895,
        "ok_rate": 100
      },
      "阿拉善盟": {
        "all_count": 593,
        "sum": 100352,
        "ok_count": 593,
        "err_count": 0,
        "val": 169,
        "ok_rate": 100
      },
      "乌海": {
        "all_count": 421,
        "sum": 178259,
        "ok_count": 421,
        "err_count": 0,
        "val": 423,
        "ok_rate": 100
      },
      "包头": {
        "all_count": 304,
        "sum": 137267,
        "ok_count": 304,
        "err_count": 0,
        "val": 452,
        "ok_rate": 100
      },
      "鄂尔多斯": {
        "all_count": 283,
        "sum": 121642,
        "ok_count": 283,
        "err_count": 0,
        "val": 430,
        "ok_rate": 100
      },
      "呼伦贝尔": {
        "all_count": 273,
        "sum": 117588,
        "ok_count": 273,
        "err_count": 0,
        "val": 431,
        "ok_rate": 100
      },
      "兴安盟": {
        "all_count": 204,
        "sum": 82526,
        "ok_count": 204,
        "err_count": 0,
        "val": 405,
        "ok_rate": 100
      },
      "赤峰": {
        "all_count": 190,
        "sum": 76431,
        "ok_count": 190,
        "err_count": 0,
        "val": 402,
        "ok_rate": 100
      },
      "通辽": {
        "all_count": 119,
        "sum": 50373,
        "ok_count": 119,
        "err_count": 0,
        "val": 423,
        "ok_rate": 100
      },
      "巴彦淖尔": {
        "all_count": 103,
        "sum": 51694,
        "ok_count": 103,
        "err_count": 0,
        "val": 502,
        "ok_rate": 100
      },
      "锡林郭勒盟": {
        "all_count": 62,
        "sum": 21091,
        "ok_count": 62,
        "err_count": 0,
        "val": 340,
        "ok_rate": 100
      }
    },
    "sum": 1815555,
    "err_count": 2,
    "ok_count": 4304,
    "val": 422,
    "ok_rate": 99.954
  },
  "北京": {
    "all_count": 4540,
    "city": {
      "北京": {
        "all_count": 4540,
        "sum": 1198599,
        "err_count": 3,
        "ok_count": 4537,
        "val": 264,
        "ok_rate": 99.934
      }
    },
    "sum": 1198599,
    "err_count": 3,
    "ok_count": 4537,
    "val": 264,
    "ok_rate": 99.934
  },
  "吉林": {
    "all_count": 4585,
    "city": {
      "长春": {
        "all_count": 2693,
        "sum": 689717,
        "ok_count": 2693,
        "err_count": 0,
        "val": 256,
        "ok_rate": 100
      },
      "通化": {
        "all_count": 751,
        "sum": 230000,
        "ok_count": 751,
        "err_count": 0,
        "val": 306,
        "ok_rate": 100
      },
      "jilin": {
        "all_count": 400,
        "sum": 178009,
        "ok_count": 400,
        "err_count": 0,
        "val": 445,
        "ok_rate": 100
      },
      "延边朝鲜族自治州": {
        "all_count": 185,
        "sum": 81505,
        "ok_count": 185,
        "err_count": 0,
        "val": 441,
        "ok_rate": 100
      },
      "松原": {
        "all_count": 164,
        "sum": 87294,
        "ok_count": 164,
        "err_count": 0,
        "val": 532,
        "ok_rate": 100
      },
      "四平": {
        "all_count": 144,
        "sum": 57518,
        "ok_count": 144,
        "err_count": 0,
        "val": 399,
        "ok_rate": 100
      },
      "白山": {
        "all_count": 105,
        "sum": 46170,
        "ok_count": 105,
        "err_count": 0,
        "val": 440,
        "ok_rate": 100
      },
      "辽源": {
        "all_count": 74,
        "sum": 19971,
        "ok_count": 74,
        "err_count": 0,
        "val": 270,
        "ok_rate": 100
      },
      "白城": {
        "all_count": 69,
        "sum": 22290,
        "ok_count": 69,
        "err_count": 0,
        "val": 323,
        "ok_rate": 100
      }
    },
    "sum": 1412474,
    "ok_count": 4585,
    "err_count": 0,
    "val": 308,
    "ok_rate": 100
  },
  "四川": {
    "all_count": 4347,
    "city": {
      "成都": {
        "all_count": 1697,
        "sum": 646430,
        "ok_count": 1697,
        "err_count": 0,
        "val": 381,
        "ok_rate": 100
      },
      "自贡": {
        "all_count": 437,
        "sum": 248130,
        "err_count": 5,
        "ok_count": 432,
        "val": 574,
        "ok_rate": 98.856
      },
      "泸州": {
        "all_count": 398,
        "sum": 161124,
        "ok_count": 398,
        "err_count": 0,
        "val": 405,
        "ok_rate": 100
      },
      "德阳": {
        "all_count": 318,
        "sum": 80434,
        "ok_count": 318,
        "err_count": 0,
        "val": 253,
        "ok_rate": 100
      },
      "南充": {
        "all_count": 309,
        "sum": 89337,
        "ok_count": 309,
        "err_count": 0,
        "val": 289,
        "ok_rate": 100
      },
      "绵阳": {
        "all_count": 237,
        "sum": 102078,
        "ok_count": 237,
        "err_count": 0,
        "val": 431,
        "ok_rate": 100
      },
      "内江": {
        "all_count": 173,
        "sum": 89326,
        "ok_count": 173,
        "err_count": 0,
        "val": 516,
        "ok_rate": 100
      },
      "雅安": {
        "all_count": 127,
        "sum": 36779,
        "ok_count": 127,
        "err_count": 0,
        "val": 290,
        "ok_rate": 100
      },
      "达州": {
        "all_count": 110,
        "sum": 29866,
        "ok_count": 110,
        "err_count": 0,
        "val": 272,
        "ok_rate": 100
      },
      "资阳": {
        "all_count": 78,
        "sum": 23912,
        "ok_count": 78,
        "err_count": 0,
        "val": 307,
        "ok_rate": 100
      },
      "宜宾": {
        "all_count": 74,
        "sum": 33837,
        "ok_count": 74,
        "err_count": 0,
        "val": 457,
        "ok_rate": 100
      },
      "甘孜藏族自治州": {
        "all_count": 71,
        "sum": 35952,
        "ok_count": 71,
        "err_count": 0,
        "val": 506,
        "ok_rate": 100
      },
      "乐山": {
        "all_count": 66,
        "sum": 36376,
        "ok_count": 66,
        "err_count": 0,
        "val": 551,
        "ok_rate": 100
      },
      "广安": {
        "all_count": 55,
        "sum": 19500,
        "ok_count": 55,
        "err_count": 0,
        "val": 355,
        "ok_rate": 100
      },
      "巴中": {
        "all_count": 53,
        "sum": 77604,
        "ok_count": 53,
        "err_count": 0,
        "val": 1464,
        "ok_rate": 100
      },
      "广元": {
        "all_count": 35,
        "sum": 23811,
        "ok_count": 35,
        "err_count": 0,
        "val": 680,
        "ok_rate": 100
      },
      "眉山": {
        "all_count": 32,
        "sum": 14340,
        "ok_count": 32,
        "err_count": 0,
        "val": 448,
        "ok_rate": 100
      },
      "遂宁": {
        "all_count": 31,
        "sum": 12652,
        "ok_count": 31,
        "err_count": 0,
        "val": 408,
        "ok_rate": 100
      },
      "凉山彝族自治州": {
        "all_count": 29,
        "sum": 27183,
        "ok_count": 29,
        "err_count": 0,
        "val": 937,
        "ok_rate": 100
      },
      "攀枝花": {
        "all_count": 17,
        "sum": 5687,
        "ok_count": 17,
        "err_count": 0,
        "val": 335,
        "ok_rate": 100
      }
    },
    "sum": 1794358,
    "err_count": 5,
    "ok_count": 4342,
    "val": 413,
    "ok_rate": 99.885
  },
  "天津": {
    "all_count": 4473,
    "city": {
      "天津": {
        "all_count": 4473,
        "sum": 1314441,
        "err_count": 8,
        "ok_count": 4465,
        "val": 294,
        "ok_rate": 99.821
      }
    },
    "sum": 1314441,
    "err_count": 8,
    "ok_count": 4465,
    "val": 294,
    "ok_rate": 99.821
  },
  "宁夏": {
    "all_count": 4640,
    "city": {
      "银川": {
        "all_count": 2753,
        "sum": 1156164,
        "ok_count": 2753,
        "err_count": 0,
        "val": 420,
        "ok_rate": 100
      },
      "吴忠": {
        "all_count": 672,
        "sum": 271456,
        "ok_count": 672,
        "err_count": 0,
        "val": 404,
        "ok_rate": 100
      },
      "固原": {
        "all_count": 425,
        "sum": 670551,
        "err_count": 1,
        "ok_count": 424,
        "val": 1581,
        "ok_rate": 99.765
      },
      "石嘴山": {
        "all_count": 420,
        "sum": 247161,
        "ok_count": 420,
        "err_count": 0,
        "val": 588,
        "ok_rate": 100
      },
      "中卫": {
        "all_count": 370,
        "sum": 249016,
        "ok_count": 370,
        "err_count": 0,
        "val": 673,
        "ok_rate": 100
      }
    },
    "sum": 2594348,
    "err_count": 1,
    "ok_count": 4639,
    "val": 559,
    "ok_rate": 99.978
  },
  "安徽": {
    "all_count": 3865,
    "city": {
      "合肥": {
        "all_count": 1188,
        "sum": 344277,
        "ok_count": 1188,
        "err_count": 0,
        "val": 290,
        "ok_rate": 100
      },
      "滁州": {
        "all_count": 769,
        "sum": 131603,
        "ok_count": 769,
        "err_count": 0,
        "val": 171,
        "ok_rate": 100
      },
      "马鞍山": {
        "all_count": 408,
        "sum": 235930,
        "ok_count": 408,
        "err_count": 0,
        "val": 578,
        "ok_rate": 100
      },
      "黄山": {
        "all_count": 327,
        "sum": 100757,
        "ok_count": 327,
        "err_count": 0,
        "val": 308,
        "ok_rate": 100
      },
      "亳州": {
        "all_count": 268,
        "sum": 69754,
        "ok_count": 268,
        "err_count": 0,
        "val": 260,
        "ok_rate": 100
      },
      "宣城": {
        "all_count": 184,
        "sum": 70364,
        "ok_count": 184,
        "err_count": 0,
        "val": 382,
        "ok_rate": 100
      },
      "铜陵": {
        "all_count": 162,
        "sum": 71222,
        "ok_count": 162,
        "err_count": 0,
        "val": 440,
        "ok_rate": 100
      },
      "芜湖": {
        "all_count": 147,
        "sum": 70226,
        "ok_count": 147,
        "err_count": 0,
        "val": 478,
        "ok_rate": 100
      },
      "阜阳": {
        "all_count": 76,
        "sum": 32805,
        "ok_count": 76,
        "err_count": 0,
        "val": 432,
        "ok_rate": 100
      },
      "宿州": {
        "all_count": 75,
        "sum": 28658,
        "ok_count": 75,
        "err_count": 0,
        "val": 382,
        "ok_rate": 100
      },
      "六安": {
        "all_count": 68,
        "sum": 32498,
        "ok_count": 68,
        "err_count": 0,
        "val": 478,
        "ok_rate": 100
      },
      "安庆": {
        "all_count": 59,
        "sum": 26224,
        "ok_count": 59,
        "err_count": 0,
        "val": 444,
        "ok_rate": 100
      },
      "蚌埠": {
        "all_count": 53,
        "sum": 23499,
        "ok_count": 53,
        "err_count": 0,
        "val": 443,
        "ok_rate": 100
      },
      "淮南": {
        "all_count": 41,
        "sum": 13184,
        "ok_count": 41,
        "err_count": 0,
        "val": 322,
        "ok_rate": 100
      },
      "池州": {
        "all_count": 20,
        "sum": 7470,
        "ok_count": 20,
        "err_count": 0,
        "val": 374,
        "ok_rate": 100
      },
      "淮北": {
        "all_count": 20,
        "sum": 10130,
        "ok_count": 20,
        "err_count": 0,
        "val": 507,
        "ok_rate": 100
      }
    },
    "sum": 1268601,
    "ok_count": 3865,
    "err_count": 0,
    "val": 328,
    "ok_rate": 100
  },
  "山东": {
    "all_count": 4205,
    "city": {
      "青岛": {
        "all_count": 760,
        "sum": 269057,
        "ok_count": 760,
        "err_count": 0,
        "val": 354,
        "ok_rate": 100
      },
      "淄博": {
        "all_count": 695,
        "sum": 298286,
        "err_count": 1,
        "ok_count": 694,
        "val": 430,
        "ok_rate": 99.856
      },
      "潍坊": {
        "all_count": 427,
        "sum": 326736,
        "ok_count": 427,
        "err_count": 0,
        "val": 765,
        "ok_rate": 100
      },
      "临沂": {
        "all_count": 376,
        "sum": 99318,
        "ok_count": 376,
        "err_count": 0,
        "val": 264,
        "ok_rate": 100
      },
      "济宁": {
        "all_count": 331,
        "sum": 151651,
        "ok_count": 331,
        "err_count": 0,
        "val": 458,
        "ok_rate": 100
      },
      "济南": {
        "all_count": 307,
        "sum": 104253,
        "ok_count": 307,
        "err_count": 0,
        "val": 340,
        "ok_rate": 100
      },
      "菏泽": {
        "all_count": 213,
        "sum": 105585,
        "ok_count": 213,
        "err_count": 0,
        "val": 496,
        "ok_rate": 100
      },
      "烟台": {
        "all_count": 199,
        "sum": 64057,
        "ok_count": 199,
        "err_count": 0,
        "val": 322,
        "ok_rate": 100
      },
      "枣庄": {
        "all_count": 155,
        "sum": 71846,
        "ok_count": 155,
        "err_count": 0,
        "val": 464,
        "ok_rate": 100
      },
      "威海": {
        "all_count": 151,
        "sum": 60345,
        "ok_count": 151,
        "err_count": 0,
        "val": 400,
        "ok_rate": 100
      },
      "日照": {
        "all_count": 145,
        "sum": 58044,
        "ok_count": 145,
        "err_count": 0,
        "val": 400,
        "ok_rate": 100
      },
      "东营": {
        "all_count": 125,
        "sum": 40817,
        "ok_count": 125,
        "err_count": 0,
        "val": 327,
        "ok_rate": 100
      },
      "滨州": {
        "all_count": 96,
        "sum": 29230,
        "ok_count": 96,
        "err_count": 0,
        "val": 304,
        "ok_rate": 100
      },
      "德州": {
        "all_count": 83,
        "sum": 22880,
        "ok_count": 83,
        "err_count": 0,
        "val": 276,
        "ok_rate": 100
      },
      "泰安": {
        "all_count": 78,
        "sum": 23381,
        "ok_count": 78,
        "err_count": 0,
        "val": 300,
        "ok_rate": 100
      },
      "聊城": {
        "all_count": 64,
        "sum": 15838,
        "ok_count": 64,
        "err_count": 0,
        "val": 247,
        "ok_rate": 100
      }
    },
    "sum": 1741324,
    "err_count": 1,
    "ok_count": 4204,
    "val": 414,
    "ok_rate": 99.976
  },
  "山西": {
    "all_count": 4367,
    "city": {
      "太原": {
        "all_count": 1861,
        "sum": 381175,
        "err_count": 10,
        "ok_count": 1851,
        "val": 206,
        "ok_rate": 99.463
      },
      "运城": {
        "all_count": 1765,
        "sum": 543413,
        "err_count": 6,
        "ok_count": 1759,
        "val": 309,
        "ok_rate": 99.66
      },
      "长治": {
        "all_count": 260,
        "sum": 169989,
        "err_count": 5,
        "ok_count": 255,
        "val": 667,
        "ok_rate": 98.077
      },
      "晋城": {
        "all_count": 145,
        "sum": 58516,
        "ok_count": 145,
        "err_count": 0,
        "val": 404,
        "ok_rate": 100
      },
      "大同": {
        "all_count": 58,
        "sum": 25375,
        "ok_count": 58,
        "err_count": 0,
        "val": 438,
        "ok_rate": 100
      },
      "晋中": {
        "all_count": 53,
        "sum": 14732,
        "ok_count": 53,
        "err_count": 0,
        "val": 278,
        "ok_rate": 100
      },
      "阳泉": {
        "all_count": 52,
        "sum": 20725,
        "ok_count": 52,
        "err_count": 0,
        "val": 399,
        "ok_rate": 100
      },
      "忻州": {
        "all_count": 50,
        "sum": 28384,
        "ok_count": 50,
        "err_count": 0,
        "val": 568,
        "ok_rate": 100
      },
      "吕梁": {
        "all_count": 47,
        "sum": 21121,
        "ok_count": 47,
        "err_count": 0,
        "val": 449,
        "ok_rate": 100
      },
      "临汾": {
        "all_count": 46,
        "sum": 22731,
        "ok_count": 46,
        "err_count": 0,
        "val": 494,
        "ok_rate": 100
      },
      "朔州": {
        "all_count": 30,
        "sum": 11088,
        "ok_count": 30,
        "err_count": 0,
        "val": 370,
        "ok_rate": 100
      }
    },
    "sum": 1297249,
    "err_count": 21,
    "ok_count": 4346,
    "val": 298,
    "ok_rate": 99.519
  },
  "广东": {
    "all_count": 4156,
    "city": {
      "广州": {
        "all_count": 877,
        "sum": 427244,
        "ok_count": 877,
        "err_count": 0,
        "val": 487,
        "ok_rate": 100
      },
      "江门": {
        "all_count": 697,
        "sum": 303768,
        "ok_count": 697,
        "err_count": 0,
        "val": 436,
        "ok_rate": 100
      },
      "佛山": {
        "all_count": 610,
        "sum": 475842,
        "ok_count": 610,
        "err_count": 0,
        "val": 780,
        "ok_rate": 100
      },
      "深圳": {
        "all_count": 484,
        "sum": 157479,
        "ok_count": 484,
        "err_count": 0,
        "val": 325,
        "ok_rate": 100
      },
      "阳江": {
        "all_count": 269,
        "sum": 111853,
        "ok_count": 269,
        "err_count": 0,
        "val": 416,
        "ok_rate": 100
      },
      "肇庆": {
        "all_count": 218,
        "sum": 81017,
        "ok_count": 218,
        "err_count": 0,
        "val": 372,
        "ok_rate": 100
      },
      "湛江": {
        "all_count": 174,
        "sum": 61206,
        "ok_count": 174,
        "err_count": 0,
        "val": 352,
        "ok_rate": 100
      },
      "东莞": {
        "all_count": 169,
        "sum": 56919,
        "ok_count": 169,
        "err_count": 0,
        "val": 337,
        "ok_rate": 100
      },
      "梅州": {
        "all_count": 136,
        "sum": 44230,
        "ok_count": 136,
        "err_count": 0,
        "val": 325,
        "ok_rate": 100
      },
      "河源": {
        "all_count": 88,
        "sum": 29360,
        "ok_count": 88,
        "err_count": 0,
        "val": 334,
        "ok_rate": 100
      },
      "珠海": {
        "all_count": 81,
        "sum": 81306,
        "ok_count": 81,
        "err_count": 0,
        "val": 1004,
        "ok_rate": 100
      },
      "中山": {
        "all_count": 62,
        "sum": 23633,
        "ok_count": 62,
        "err_count": 0,
        "val": 381,
        "ok_rate": 100
      },
      "揭阳": {
        "all_count": 50,
        "sum": 29638,
        "ok_count": 50,
        "err_count": 0,
        "val": 593,
        "ok_rate": 100
      },
      "汕头": {
        "all_count": 44,
        "sum": 27433,
        "ok_count": 44,
        "err_count": 0,
        "val": 623,
        "ok_rate": 100
      },
      "茂名": {
        "all_count": 43,
        "sum": 18459,
        "ok_count": 43,
        "err_count": 0,
        "val": 429,
        "ok_rate": 100
      },
      "惠州": {
        "all_count": 42,
        "sum": 13080,
        "ok_count": 42,
        "err_count": 0,
        "val": 311,
        "ok_rate": 100
      },
      "潮州": {
        "all_count": 38,
        "sum": 14006,
        "ok_count": 38,
        "err_count": 0,
        "val": 369,
        "ok_rate": 100
      },
      "韶关": {
        "all_count": 28,
        "sum": 12518,
        "ok_count": 28,
        "err_count": 0,
        "val": 447,
        "ok_rate": 100
      },
      "汕尾": {
        "all_count": 16,
        "sum": 5266,
        "ok_count": 16,
        "err_count": 0,
        "val": 329,
        "ok_rate": 100
      },
      "云浮": {
        "all_count": 15,
        "sum": 5810,
        "ok_count": 15,
        "err_count": 0,
        "val": 387,
        "ok_rate": 100
      },
      "清远": {
        "all_count": 15,
        "sum": 6085,
        "ok_count": 15,
        "err_count": 0,
        "val": 406,
        "ok_rate": 100
      }
    },
    "sum": 1986152,
    "ok_count": 4156,
    "err_count": 0,
    "val": 478,
    "ok_rate": 100
  },
  "广西": {
    "all_count": 4507,
    "city": {
      "桂林": {
        "all_count": 1918,
        "sum": 789001,
        "ok_count": 1918,
        "err_count": 0,
        "val": 411,
        "ok_rate": 100
      },
      "南宁": {
        "all_count": 1081,
        "sum": 351938,
        "ok_count": 1081,
        "err_count": 0,
        "val": 326,
        "ok_rate": 100
      },
      "河池": {
        "all_count": 239,
        "sum": 108310,
        "ok_count": 239,
        "err_count": 0,
        "val": 453,
        "ok_rate": 100
      },
      "柳州": {
        "all_count": 237,
        "sum": 105906,
        "ok_count": 237,
        "err_count": 0,
        "val": 447,
        "ok_rate": 100
      },
      "玉林": {
        "all_count": 176,
        "sum": 79792,
        "ok_count": 176,
        "err_count": 0,
        "val": 453,
        "ok_rate": 100
      },
      "梧州": {
        "all_count": 156,
        "sum": 81253,
        "ok_count": 156,
        "err_count": 0,
        "val": 521,
        "ok_rate": 100
      },
      "钦州": {
        "all_count": 140,
        "sum": 68261,
        "ok_count": 140,
        "err_count": 0,
        "val": 488,
        "ok_rate": 100
      },
      "北海": {
        "all_count": 116,
        "sum": 60625,
        "ok_count": 116,
        "err_count": 0,
        "val": 523,
        "ok_rate": 100
      },
      "来宾": {
        "all_count": 103,
        "sum": 59644,
        "ok_count": 103,
        "err_count": 0,
        "val": 579,
        "ok_rate": 100
      },
      "贵港": {
        "all_count": 93,
        "sum": 33789,
        "ok_count": 93,
        "err_count": 0,
        "val": 363,
        "ok_rate": 100
      },
      "百色": {
        "all_count": 85,
        "sum": 39961,
        "ok_count": 85,
        "err_count": 0,
        "val": 470,
        "ok_rate": 100
      },
      "崇左": {
        "all_count": 69,
        "sum": 37910,
        "ok_count": 69,
        "err_count": 0,
        "val": 549,
        "ok_rate": 100
      },
      "贺州": {
        "all_count": 64,
        "sum": 37748,
        "ok_count": 64,
        "err_count": 0,
        "val": 590,
        "ok_rate": 100
      },
      "防城港": {
        "all_count": 30,
        "sum": 9779,
        "ok_count": 30,
        "err_count": 0,
        "val": 326,
        "ok_rate": 100
      }
    },
    "sum": 1863917,
    "ok_count": 4507,
    "err_count": 0,
    "val": 414,
    "ok_rate": 100
  },
  "新疆": {
    "all_count": 3476,
    "city": {
      "乌鲁木齐": {
        "all_count": 1938,
        "sum": 633285,
        "err_count": 1,
        "ok_count": 1937,
        "val": 327,
        "ok_rate": 99.948
      },
      "阿克苏地区": {
        "all_count": 598,
        "sum": 188255,
        "ok_count": 598,
        "err_count": 0,
        "val": 315,
        "ok_rate": 100
      },
      "昌吉回族自治州": {
        "all_count": 167,
        "sum": 89272,
        "ok_count": 167,
        "err_count": 0,
        "val": 535,
        "ok_rate": 100
      },
      "伊犁哈萨克自治州": {
        "all_count": 162,
        "sum": 71472,
        "ok_count": 162,
        "err_count": 0,
        "val": 441,
        "ok_rate": 100
      },
      "哈密": {
        "all_count": 117,
        "sum": 44437,
        "ok_count": 117,
        "err_count": 0,
        "val": 380,
        "ok_rate": 100
      },
      "巴音郭楞蒙古自治州": {
        "all_count": 109,
        "sum": 50298,
        "ok_count": 109,
        "err_count": 0,
        "val": 461,
        "ok_rate": 100
      },
      "塔城地区": {
        "all_count": 78,
        "sum": 35253,
        "ok_count": 78,
        "err_count": 0,
        "val": 452,
        "ok_rate": 100
      },
      "喀什地区": {
        "all_count": 78,
        "sum": 41346,
        "ok_count": 78,
        "err_count": 0,
        "val": 530,
        "ok_rate": 100
      },
      "和田地区": {
        "all_count": 55,
        "sum": 52591,
        "ok_count": 55,
        "err_count": 0,
        "val": 956,
        "ok_rate": 100
      },
      "石河子": {
        "all_count": 46,
        "sum": 17368,
        "ok_count": 46,
        "err_count": 0,
        "val": 378,
        "ok_rate": 100
      },
      "阿勒泰地区": {
        "all_count": 41,
        "sum": 25188,
        "ok_count": 41,
        "err_count": 0,
        "val": 614,
        "ok_rate": 100
      },
      "吐鲁番": {
        "all_count": 23,
        "sum": 12312,
        "ok_count": 23,
        "err_count": 0,
        "val": 535,
        "ok_rate": 100
      },
      "克拉玛依": {
        "all_count": 21,
        "sum": 12158,
        "ok_count": 21,
        "err_count": 0,
        "val": 579,
        "ok_rate": 100
      },
      "阿拉尔": {
        "all_count": 16,
        "sum": 8956,
        "ok_count": 16,
        "err_count": 0,
        "val": 560,
        "ok_rate": 100
      },
      "博尔塔拉蒙古自治州": {
        "all_count": 12,
        "sum": 5643,
        "ok_count": 12,
        "err_count": 0,
        "val": 470,
        "ok_rate": 100
      },
      "五家渠": {
        "all_count": 6,
        "sum": 1826,
        "ok_count": 6,
        "err_count": 0,
        "val": 304,
        "ok_rate": 100
      },
      "北屯": {
        "all_count": 5,
        "sum": 2094,
        "ok_count": 5,
        "err_count": 0,
        "val": 419,
        "ok_rate": 100
      },
      "图木舒克": {
        "all_count": 3,
        "sum": 1315,
        "ok_count": 3,
        "err_count": 0,
        "val": 438,
        "ok_rate": 100
      },
      "克孜勒苏柯尔克孜自治州": {
        "all_count": 1,
        "sum": 320,
        "ok_count": 1,
        "err_count": 0,
        "val": 320,
        "ok_rate": 100
      }
    },
    "sum": 1293389,
    "err_count": 1,
    "ok_count": 3475,
    "val": 372,
    "ok_rate": 99.971
  },
  "江苏": {
    "all_count": 3133,
    "city": {
      "泰州": {
        "all_count": 1042,
        "sum": 303624,
        "ok_count": 1042,
        "err_count": 0,
        "val": 291,
        "ok_rate": 100
      },
      "常州": {
        "all_count": 563,
        "sum": 400654,
        "ok_count": 563,
        "err_count": 0,
        "val": 712,
        "ok_rate": 100
      },
      "扬州": {
        "all_count": 291,
        "sum": 85646,
        "ok_count": 291,
        "err_count": 0,
        "val": 294,
        "ok_rate": 100
      },
      "苏州": {
        "all_count": 244,
        "sum": 57360,
        "ok_count": 244,
        "err_count": 0,
        "val": 235,
        "ok_rate": 100
      },
      "南京": {
        "all_count": 235,
        "sum": 62579,
        "ok_count": 235,
        "err_count": 0,
        "val": 266,
        "ok_rate": 100
      },
      "无锡": {
        "all_count": 180,
        "sum": 61714,
        "ok_count": 180,
        "err_count": 0,
        "val": 343,
        "ok_rate": 100
      },
      "南通": {
        "all_count": 145,
        "sum": 49753,
        "ok_count": 145,
        "err_count": 0,
        "val": 343,
        "ok_rate": 100
      },
      "宿迁": {
        "all_count": 101,
        "sum": 47446,
        "err_count": 1,
        "ok_count": 100,
        "val": 474,
        "ok_rate": 99.01
      },
      "镇江": {
        "all_count": 98,
        "sum": 22427,
        "ok_count": 98,
        "err_count": 0,
        "val": 229,
        "ok_rate": 100
      },
      "淮安": {
        "all_count": 87,
        "sum": 18968,
        "ok_count": 87,
        "err_count": 0,
        "val": 218,
        "ok_rate": 100
      },
      "徐州": {
        "all_count": 75,
        "sum": 28659,
        "ok_count": 75,
        "err_count": 0,
        "val": 382,
        "ok_rate": 100
      },
      "连云港": {
        "all_count": 47,
        "sum": 15205,
        "ok_count": 47,
        "err_count": 0,
        "val": 324,
        "ok_rate": 100
      },
      "盐城": {
        "all_count": 25,
        "sum": 14825,
        "ok_count": 25,
        "err_count": 0,
        "val": 593,
        "ok_rate": 100
      }
    },
    "sum": 1168860,
    "err_count": 1,
    "ok_count": 3132,
    "val": 373,
    "ok_rate": 99.968
  },
  "江西": {
    "all_count": 4497,
    "city": {
      "宜春": {
        "all_count": 1469,
        "sum": 616453,
        "err_count": 2,
        "ok_count": 1467,
        "val": 420,
        "ok_rate": 99.864
      },
      "南昌": {
        "all_count": 1366,
        "sum": 478696,
        "ok_count": 1366,
        "err_count": 0,
        "val": 350,
        "ok_rate": 100
      },
      "九江": {
        "all_count": 487,
        "sum": 115192,
        "ok_count": 487,
        "err_count": 0,
        "val": 237,
        "ok_rate": 100
      },
      "赣州": {
        "all_count": 443,
        "sum": 139847,
        "ok_count": 443,
        "err_count": 0,
        "val": 316,
        "ok_rate": 100
      },
      "上饶": {
        "all_count": 159,
        "sum": 62069,
        "ok_count": 159,
        "err_count": 0,
        "val": 390,
        "ok_rate": 100
      },
      "吉安": {
        "all_count": 157,
        "sum": 40880,
        "ok_count": 157,
        "err_count": 0,
        "val": 260,
        "ok_rate": 100
      },
      "鹰潭": {
        "all_count": 145,
        "sum": 35495,
        "ok_count": 145,
        "err_count": 0,
        "val": 245,
        "ok_rate": 100
      },
      "抚州": {
        "all_count": 97,
        "sum": 26845,
        "ok_count": 97,
        "err_count": 0,
        "val": 277,
        "ok_rate": 100
      },
      "萍乡": {
        "all_count": 63,
        "sum": 17891,
        "ok_count": 63,
        "err_count": 0,
        "val": 284,
        "ok_rate": 100
      },
      "景德镇": {
        "all_count": 58,
        "sum": 21157,
        "ok_count": 58,
        "err_count": 0,
        "val": 365,
        "ok_rate": 100
      },
      "新余": {
        "all_count": 53,
        "sum": 15059,
        "ok_count": 53,
        "err_count": 0,
        "val": 284,
        "ok_rate": 100
      }
    },
    "sum": 1569584,
    "err_count": 2,
    "ok_count": 4495,
    "val": 349,
    "ok_rate": 99.956
  },
  "河北": {
    "all_count": 4374,
    "city": {
      "石家庄": {
        "all_count": 1198,
        "sum": 261629,
        "ok_count": 1198,
        "err_count": 0,
        "val": 218,
        "ok_rate": 100
      },
      "保定": {
        "all_count": 1161,
        "sum": 413371,
        "err_count": 2,
        "ok_count": 1159,
        "val": 357,
        "ok_rate": 99.828
      },
      "邯郸": {
        "all_count": 632,
        "sum": 119845,
        "ok_count": 632,
        "err_count": 0,
        "val": 190,
        "ok_rate": 100
      },
      "沧州": {
        "all_count": 368,
        "sum": 89865,
        "ok_count": 368,
        "err_count": 0,
        "val": 244,
        "ok_rate": 100
      },
      "邢台": {
        "all_count": 257,
        "sum": 161403,
        "ok_count": 257,
        "err_count": 0,
        "val": 628,
        "ok_rate": 100
      },
      "承德": {
        "all_count": 224,
        "sum": 38997,
        "ok_count": 224,
        "err_count": 0,
        "val": 174,
        "ok_rate": 100
      },
      "唐山": {
        "all_count": 152,
        "sum": 65106,
        "ok_count": 152,
        "err_count": 0,
        "val": 428,
        "ok_rate": 100
      },
      "廊坊": {
        "all_count": 127,
        "sum": 54029,
        "ok_count": 127,
        "err_count": 0,
        "val": 425,
        "ok_rate": 100
      },
      "衡水": {
        "all_count": 99,
        "sum": 36884,
        "ok_count": 99,
        "err_count": 0,
        "val": 373,
        "ok_rate": 100
      },
      "张家口": {
        "all_count": 91,
        "sum": 27480,
        "ok_count": 91,
        "err_count": 0,
        "val": 302,
        "ok_rate": 100
      },
      "秦皇岛": {
        "all_count": 65,
        "sum": 30894,
        "ok_count": 65,
        "err_count": 0,
        "val": 475,
        "ok_rate": 100
      }
    },
    "sum": 1299503,
    "err_count": 2,
    "ok_count": 4372,
    "val": 297,
    "ok_rate": 99.954
  },
  "河南": {
    "all_count": 4081,
    "city": {
      "郑州": {
        "all_count": 1065,
        "sum": 382474,
        "err_count": 2,
        "ok_count": 1063,
        "val": 360,
        "ok_rate": 99.812
      },
      "开封": {
        "all_count": 677,
        "sum": 215869,
        "ok_count": 677,
        "err_count": 0,
        "val": 319,
        "ok_rate": 100
      },
      "新乡": {
        "all_count": 370,
        "sum": 79455,
        "ok_count": 370,
        "err_count": 0,
        "val": 215,
        "ok_rate": 100
      },
      "信阳": {
        "all_count": 300,
        "sum": 81947,
        "ok_count": 300,
        "err_count": 0,
        "val": 273,
        "ok_rate": 100
      },
      "洛阳": {
        "all_count": 236,
        "sum": 121248,
        "ok_count": 236,
        "err_count": 0,
        "val": 514,
        "ok_rate": 100
      },
      "南阳": {
        "all_count": 196,
        "sum": 67443,
        "err_count": 1,
        "ok_count": 195,
        "val": 346,
        "ok_rate": 99.49
      },
      "驻马店": {
        "all_count": 195,
        "sum": 84704,
        "ok_count": 195,
        "err_count": 0,
        "val": 434,
        "ok_rate": 100
      },
      "焦作": {
        "all_count": 171,
        "sum": 58813,
        "ok_count": 171,
        "err_count": 0,
        "val": 344,
        "ok_rate": 100
      },
      "许昌": {
        "all_count": 157,
        "sum": 39209,
        "ok_count": 157,
        "err_count": 0,
        "val": 250,
        "ok_rate": 100
      },
      "周口": {
        "all_count": 156,
        "sum": 49676,
        "ok_count": 156,
        "err_count": 0,
        "val": 318,
        "ok_rate": 100
      },
      "平顶山": {
        "all_count": 110,
        "sum": 46695,
        "ok_count": 110,
        "err_count": 0,
        "val": 425,
        "ok_rate": 100
      },
      "濮阳": {
        "all_count": 96,
        "sum": 32373,
        "ok_count": 96,
        "err_count": 0,
        "val": 337,
        "ok_rate": 100
      },
      "商丘": {
        "all_count": 90,
        "sum": 24330,
        "ok_count": 90,
        "err_count": 0,
        "val": 270,
        "ok_rate": 100
      },
      "鹤壁": {
        "all_count": 71,
        "sum": 54316,
        "ok_count": 71,
        "err_count": 0,
        "val": 765,
        "ok_rate": 100
      },
      "漯河": {
        "all_count": 66,
        "sum": 23219,
        "ok_count": 66,
        "err_count": 0,
        "val": 352,
        "ok_rate": 100
      },
      "安阳": {
        "all_count": 63,
        "sum": 32782,
        "ok_count": 63,
        "err_count": 0,
        "val": 520,
        "ok_rate": 100
      },
      "三门峡": {
        "all_count": 54,
        "sum": 12567,
        "ok_count": 54,
        "err_count": 0,
        "val": 233,
        "ok_rate": 100
      },
      "济源": {
        "all_count": 8,
        "sum": 3395,
        "ok_count": 8,
        "err_count": 0,
        "val": 424,
        "ok_rate": 100
      }
    },
    "sum": 1410515,
    "err_count": 3,
    "ok_count": 4078,
    "val": 346,
    "ok_rate": 99.926
  },
  "浙江": {
    "all_count": 4397,
    "city": {
      "杭州": {
        "all_count": 1945,
        "sum": 979745,
        "err_count": 2,
        "ok_count": 1943,
        "val": 504,
        "ok_rate": 99.897
      },
      "温州": {
        "all_count": 536,
        "sum": 232648,
        "ok_count": 536,
        "err_count": 0,
        "val": 434,
        "ok_rate": 100
      },
      "台州": {
        "all_count": 504,
        "sum": 157794,
        "ok_count": 504,
        "err_count": 0,
        "val": 313,
        "ok_rate": 100
      },
      "湖州": {
        "all_count": 501,
        "sum": 105513,
        "ok_count": 501,
        "err_count": 0,
        "val": 211,
        "ok_rate": 100
      },
      "绍兴": {
        "all_count": 278,
        "sum": 113893,
        "err_count": 1,
        "ok_count": 277,
        "val": 411,
        "ok_rate": 99.64
      },
      "嘉兴": {
        "all_count": 248,
        "sum": 105904,
        "ok_count": 248,
        "err_count": 0,
        "val": 427,
        "ok_rate": 100
      },
      "宁波": {
        "all_count": 152,
        "sum": 63009,
        "ok_count": 152,
        "err_count": 0,
        "val": 415,
        "ok_rate": 100
      },
      "金华": {
        "all_count": 143,
        "sum": 86845,
        "ok_count": 143,
        "err_count": 0,
        "val": 607,
        "ok_rate": 100
      },
      "舟山": {
        "all_count": 43,
        "sum": 18312,
        "ok_count": 43,
        "err_count": 0,
        "val": 426,
        "ok_rate": 100
      },
      "丽水": {
        "all_count": 27,
        "sum": 14078,
        "ok_count": 27,
        "err_count": 0,
        "val": 521,
        "ok_rate": 100
      },
      "衢州": {
        "all_count": 20,
        "sum": 6342,
        "ok_count": 20,
        "err_count": 0,
        "val": 317,
        "ok_rate": 100
      }
    },
    "sum": 1884083,
    "err_count": 3,
    "ok_count": 4394,
    "val": 429,
    "ok_rate": 99.932
  },
  "海南": {
    "all_count": 3751,
    "city": {
      "海口": {
        "all_count": 2711,
        "sum": 1154802,
        "err_count": 8,
        "ok_count": 2703,
        "val": 427,
        "ok_rate": 99.705
      },
      "文昌": {
        "all_count": 217,
        "sum": 98410,
        "ok_count": 217,
        "err_count": 0,
        "val": 454,
        "ok_rate": 100
      },
      "三亚": {
        "all_count": 174,
        "sum": 88614,
        "ok_count": 174,
        "err_count": 0,
        "val": 509,
        "ok_rate": 100
      },
      "儋州": {
        "all_count": 140,
        "sum": 74016,
        "ok_count": 140,
        "err_count": 0,
        "val": 529,
        "ok_rate": 100
      },
      "琼海": {
        "all_count": 139,
        "sum": 64048,
        "ok_count": 139,
        "err_count": 0,
        "val": 461,
        "ok_rate": 100
      },
      "澄迈县": {
        "all_count": 102,
        "sum": 38903,
        "ok_count": 102,
        "err_count": 0,
        "val": 381,
        "ok_rate": 100
      },
      "万宁": {
        "all_count": 52,
        "sum": 36493,
        "ok_count": 52,
        "err_count": 0,
        "val": 702,
        "ok_rate": 100
      },
      "乐东黎族自治县": {
        "all_count": 51,
        "sum": 30110,
        "ok_count": 51,
        "err_count": 0,
        "val": 590,
        "ok_rate": 100
      },
      "临高县": {
        "all_count": 44,
        "sum": 21019,
        "ok_count": 44,
        "err_count": 0,
        "val": 478,
        "ok_rate": 100
      },
      "陵水黎族自治县": {
        "all_count": 42,
        "sum": 19192,
        "ok_count": 42,
        "err_count": 0,
        "val": 457,
        "ok_rate": 100
      },
      "屯昌县": {
        "all_count": 28,
        "sum": 12853,
        "ok_count": 28,
        "err_count": 0,
        "val": 459,
        "ok_rate": 100
      },
      "东方": {
        "all_count": 27,
        "sum": 17402,
        "ok_count": 27,
        "err_count": 0,
        "val": 645,
        "ok_rate": 100
      },
      "定安县": {
        "all_count": 7,
        "sum": 1771,
        "ok_count": 7,
        "err_count": 0,
        "val": 253,
        "ok_rate": 100
      },
      "五指山": {
        "all_count": 7,
        "sum": 7017,
        "ok_count": 7,
        "err_count": 0,
        "val": 1002,
        "ok_rate": 100
      },
      "琼中黎族苗族自治县": {
        "all_count": 4,
        "sum": 1365,
        "ok_count": 4,
        "err_count": 0,
        "val": 341,
        "ok_rate": 100
      },
      "白沙黎族自治县": {
        "all_count": 3,
        "sum": 1060,
        "ok_count": 3,
        "err_count": 0,
        "val": 353,
        "ok_rate": 100
      },
      "昌江黎族自治县": {
        "all_count": 3,
        "sum": 1150,
        "ok_count": 3,
        "err_count": 0,
        "val": 383,
        "ok_rate": 100
      }
    },
    "sum": 1668225,
    "err_count": 8,
    "ok_count": 3743,
    "val": 446,
    "ok_rate": 99.787
  },
  "湖北": {
    "all_count": 4695,
    "city": {
      "武汉": {
        "all_count": 2690,
        "sum": 625914,
        "ok_count": 2690,
        "err_count": 0,
        "val": 233,
        "ok_rate": 100
      },
      "恩施土家族苗族自治州": {
        "all_count": 570,
        "sum": 94506,
        "ok_count": 570,
        "err_count": 0,
        "val": 166,
        "ok_rate": 100
      },
      "宜昌": {
        "all_count": 288,
        "sum": 97033,
        "ok_count": 288,
        "err_count": 0,
        "val": 337,
        "ok_rate": 100
      },
      "黄冈": {
        "all_count": 230,
        "sum": 179084,
        "ok_count": 230,
        "err_count": 0,
        "val": 779,
        "ok_rate": 100
      },
      "荆门": {
        "all_count": 151,
        "sum": 85830,
        "ok_count": 151,
        "err_count": 0,
        "val": 568,
        "ok_rate": 100
      },
      "荆州": {
        "all_count": 127,
        "sum": 40273,
        "ok_count": 127,
        "err_count": 0,
        "val": 317,
        "ok_rate": 100
      },
      "襄阳": {
        "all_count": 125,
        "sum": 54507,
        "ok_count": 125,
        "err_count": 0,
        "val": 436,
        "ok_rate": 100
      },
      "仙桃": {
        "all_count": 93,
        "sum": 26678,
        "ok_count": 93,
        "err_count": 0,
        "val": 287,
        "ok_rate": 100
      },
      "十堰": {
        "all_count": 77,
        "sum": 26894,
        "ok_count": 77,
        "err_count": 0,
        "val": 349,
        "ok_rate": 100
      },
      "咸宁": {
        "all_count": 74,
        "sum": 35012,
        "ok_count": 74,
        "err_count": 0,
        "val": 473,
        "ok_rate": 100
      },
      "孝感": {
        "all_count": 70,
        "sum": 30619,
        "ok_count": 70,
        "err_count": 0,
        "val": 437,
        "ok_rate": 100
      },
      "鄂州": {
        "all_count": 63,
        "sum": 26260,
        "ok_count": 63,
        "err_count": 0,
        "val": 417,
        "ok_rate": 100
      },
      "黄石": {
        "all_count": 49,
        "sum": 11751,
        "ok_count": 49,
        "err_count": 0,
        "val": 240,
        "ok_rate": 100
      },
      "潜江": {
        "all_count": 48,
        "sum": 16087,
        "ok_count": 48,
        "err_count": 0,
        "val": 335,
        "ok_rate": 100
      },
      "随州": {
        "all_count": 27,
        "sum": 10522,
        "ok_count": 27,
        "err_count": 0,
        "val": 390,
        "ok_rate": 100
      },
      "天门": {
        "all_count": 11,
        "sum": 2234,
        "ok_count": 11,
        "err_count": 0,
        "val": 203,
        "ok_rate": 100
      },
      "神农架林区": {
        "all_count": 2,
        "sum": 6001,
        "ok_count": 2,
        "err_count": 0,
        "val": 3001,
        "ok_rate": 100
      }
    },
    "sum": 1369205,
    "ok_count": 4695,
    "err_count": 0,
    "val": 292,
    "ok_rate": 100
  },
  "湖南": {
    "all_count": 3788,
    "city": {
      "长沙": {
        "all_count": 2160,
        "sum": 651595,
        "err_count": 6,
        "ok_count": 2154,
        "val": 303,
        "ok_rate": 99.722
      },
      "衡阳": {
        "all_count": 666,
        "sum": 265804,
        "err_count": 1,
        "ok_count": 665,
        "val": 400,
        "ok_rate": 99.85
      },
      "郴州": {
        "all_count": 179,
        "sum": 95374,
        "ok_count": 179,
        "err_count": 0,
        "val": 533,
        "ok_rate": 100
      },
      "株洲": {
        "all_count": 176,
        "sum": 52857,
        "ok_count": 176,
        "err_count": 0,
        "val": 300,
        "ok_rate": 100
      },
      "邵阳": {
        "all_count": 128,
        "sum": 85306,
        "ok_count": 128,
        "err_count": 0,
        "val": 666,
        "ok_rate": 100
      },
      "永州": {
        "all_count": 97,
        "sum": 27696,
        "ok_count": 97,
        "err_count": 0,
        "val": 286,
        "ok_rate": 100
      },
      "怀化": {
        "all_count": 75,
        "sum": 37318,
        "ok_count": 75,
        "err_count": 0,
        "val": 498,
        "ok_rate": 100
      },
      "娄底": {
        "all_count": 72,
        "sum": 19460,
        "ok_count": 72,
        "err_count": 0,
        "val": 270,
        "ok_rate": 100
      },
      "岳阳": {
        "all_count": 67,
        "sum": 19407,
        "ok_count": 67,
        "err_count": 0,
        "val": 290,
        "ok_rate": 100
      },
      "湘潭": {
        "all_count": 55,
        "sum": 17385,
        "ok_count": 55,
        "err_count": 0,
        "val": 316,
        "ok_rate": 100
      },
      "益阳": {
        "all_count": 51,
        "sum": 17350,
        "ok_count": 51,
        "err_count": 0,
        "val": 340,
        "ok_rate": 100
      },
      "常德": {
        "all_count": 45,
        "sum": 22553,
        "ok_count": 45,
        "err_count": 0,
        "val": 501,
        "ok_rate": 100
      },
      "张家界": {
        "all_count": 11,
        "sum": 5662,
        "ok_count": 11,
        "err_count": 0,
        "val": 515,
        "ok_rate": 100
      },
      "湘西土家族苗族自治州": {
        "all_count": 6,
        "sum": 1497,
        "ok_count": 6,
        "err_count": 0,
        "val": 250,
        "ok_rate": 100
      }
    },
    "sum": 1319264,
    "err_count": 7,
    "ok_count": 3781,
    "val": 349,
    "ok_rate": 99.815
  },
  "甘肃": {
    "all_count": 4122,
    "city": {
      "兰州": {
        "all_count": 1345,
        "sum": 572384,
        "ok_count": 1345,
        "err_count": 0,
        "val": 426,
        "ok_rate": 100
      },
      "天水": {
        "all_count": 1163,
        "sum": 226708,
        "ok_count": 1163,
        "err_count": 0,
        "val": 195,
        "ok_rate": 100
      },
      "酒泉": {
        "all_count": 432,
        "sum": 169478,
        "ok_count": 432,
        "err_count": 0,
        "val": 392,
        "ok_rate": 100
      },
      "白银": {
        "all_count": 166,
        "sum": 63161,
        "ok_count": 166,
        "err_count": 0,
        "val": 380,
        "ok_rate": 100
      },
      "定西": {
        "all_count": 156,
        "sum": 49940,
        "ok_count": 156,
        "err_count": 0,
        "val": 320,
        "ok_rate": 100
      },
      "张掖": {
        "all_count": 153,
        "sum": 50496,
        "ok_count": 153,
        "err_count": 0,
        "val": 330,
        "ok_rate": 100
      },
      "平凉": {
        "all_count": 142,
        "sum": 50890,
        "ok_count": 142,
        "err_count": 0,
        "val": 358,
        "ok_rate": 100
      },
      "金昌": {
        "all_count": 130,
        "sum": 45760,
        "ok_count": 130,
        "err_count": 0,
        "val": 352,
        "ok_rate": 100
      },
      "嘉峪关": {
        "all_count": 127,
        "sum": 79451,
        "ok_count": 127,
        "err_count": 0,
        "val": 626,
        "ok_rate": 100
      },
      "武威": {
        "all_count": 107,
        "sum": 49749,
        "ok_count": 107,
        "err_count": 0,
        "val": 465,
        "ok_rate": 100
      },
      "庆阳": {
        "all_count": 82,
        "sum": 23539,
        "ok_count": 82,
        "err_count": 0,
        "val": 287,
        "ok_rate": 100
      },
      "临夏回族自治州": {
        "all_count": 73,
        "sum": 19386,
        "ok_count": 73,
        "err_count": 0,
        "val": 266,
        "ok_rate": 100
      },
      "陇南": {
        "all_count": 46,
        "sum": 13262,
        "ok_count": 46,
        "err_count": 0,
        "val": 288,
        "ok_rate": 100
      }
    },
    "sum": 1414204,
    "ok_count": 4122,
    "err_count": 0,
    "val": 343,
    "ok_rate": 100
  },
  "福建": {
    "all_count": 4466,
    "city": {
      "泉州": {
        "all_count": 2236,
        "sum": 996747,
        "err_count": 8,
        "ok_count": 2228,
        "val": 447,
        "ok_rate": 99.642
      },
      "南平": {
        "all_count": 686,
        "sum": 321866,
        "ok_count": 686,
        "err_count": 0,
        "val": 469,
        "ok_rate": 100
      },
      "宁德": {
        "all_count": 360,
        "sum": 129726,
        "ok_count": 360,
        "err_count": 0,
        "val": 360,
        "ok_rate": 100
      },
      "福州": {
        "all_count": 350,
        "sum": 119798,
        "ok_count": 350,
        "err_count": 0,
        "val": 342,
        "ok_rate": 100
      },
      "莆田": {
        "all_count": 313,
        "sum": 98685,
        "ok_count": 313,
        "err_count": 0,
        "val": 315,
        "ok_rate": 100
      },
      "厦门": {
        "all_count": 288,
        "sum": 167861,
        "ok_count": 288,
        "err_count": 0,
        "val": 583,
        "ok_rate": 100
      },
      "漳州": {
        "all_count": 151,
        "sum": 81544,
        "ok_count": 151,
        "err_count": 0,
        "val": 540,
        "ok_rate": 100
      },
      "三明": {
        "all_count": 51,
        "sum": 21914,
        "ok_count": 51,
        "err_count": 0,
        "val": 430,
        "ok_rate": 100
      },
      "龙岩": {
        "all_count": 31,
        "sum": 13617,
        "ok_count": 31,
        "err_count": 0,
        "val": 439,
        "ok_rate": 100
      }
    },
    "sum": 1951758,
    "err_count": 8,
    "ok_count": 4458,
    "val": 438,
    "ok_rate": 99.821
  },
  "西藏": {
    "all_count": 4372,
    "city": {
      "拉萨": {
        "all_count": 2963,
        "sum": 1139485,
        "err_count": 3,
        "ok_count": 2960,
        "val": 385,
        "ok_rate": 99.899
      },
      "那曲": {
        "all_count": 496,
        "sum": 130601,
        "ok_count": 496,
        "err_count": 0,
        "val": 263,
        "ok_rate": 100
      },
      "林芝": {
        "all_count": 478,
        "sum": 149627,
        "ok_count": 478,
        "err_count": 0,
        "val": 313,
        "ok_rate": 100
      },
      "昌都": {
        "all_count": 340,
        "sum": 238111,
        "ok_count": 340,
        "err_count": 0,
        "val": 700,
        "ok_rate": 100
      },
      "阿里地区": {
        "all_count": 64,
        "sum": 140811,
        "ok_count": 64,
        "err_count": 0,
        "val": 2200,
        "ok_rate": 100
      },
      "日喀则": {
        "all_count": 22,
        "sum": 9121,
        "ok_count": 22,
        "err_count": 0,
        "val": 415,
        "ok_rate": 100
      },
      "山南": {
        "all_count": 9,
        "sum": 3252,
        "ok_count": 9,
        "err_count": 0,
        "val": 361,
        "ok_rate": 100
      }
    },
    "sum": 1811008,
    "err_count": 3,
    "ok_count": 4369,
    "val": 415,
    "ok_rate": 99.931
  },
  "贵州": {
    "all_count": 4675,
    "city": {
      "贵阳": {
        "all_count": 1711,
        "sum": 668909,
        "ok_count": 1711,
        "err_count": 0,
        "val": 391,
        "ok_rate": 100
      },
      "黔南布依族苗族自治州": {
        "all_count": 778,
        "sum": 244167,
        "ok_count": 778,
        "err_count": 0,
        "val": 314,
        "ok_rate": 100
      },
      "黔东南苗族侗族自治州": {
        "all_count": 664,
        "sum": 304903,
        "ok_count": 664,
        "err_count": 0,
        "val": 459,
        "ok_rate": 100
      },
      "遵义": {
        "all_count": 407,
        "sum": 255558,
        "ok_count": 407,
        "err_count": 0,
        "val": 628,
        "ok_rate": 100
      },
      "毕节": {
        "all_count": 345,
        "sum": 135293,
        "ok_count": 345,
        "err_count": 0,
        "val": 392,
        "ok_rate": 100
      },
      "黔西南布依族苗族自治州": {
        "all_count": 285,
        "sum": 119031,
        "ok_count": 285,
        "err_count": 0,
        "val": 418,
        "ok_rate": 100
      },
      "安顺": {
        "all_count": 202,
        "sum": 84032,
        "ok_count": 202,
        "err_count": 0,
        "val": 416,
        "ok_rate": 100
      },
      "六盘水": {
        "all_count": 176,
        "sum": 76699,
        "ok_count": 176,
        "err_count": 0,
        "val": 436,
        "ok_rate": 100
      },
      "铜仁": {
        "all_count": 107,
        "sum": 40698,
        "ok_count": 107,
        "err_count": 0,
        "val": 380,
        "ok_rate": 100
      }
    },
    "sum": 1929290,
    "ok_count": 4675,
    "err_count": 0,
    "val": 413,
    "ok_rate": 100
  },
  "辽宁": {
    "all_count": 4346,
    "city": {
      "沈阳": {
        "all_count": 2332,
        "sum": 458363,
        "ok_count": 2332,
        "err_count": 0,
        "val": 197,
        "ok_rate": 100
      },
      "本溪": {
        "all_count": 548,
        "sum": 113976,
        "ok_count": 548,
        "err_count": 0,
        "val": 208,
        "ok_rate": 100
      },
      "大连": {
        "all_count": 349,
        "sum": 127978,
        "ok_count": 349,
        "err_count": 0,
        "val": 367,
        "ok_rate": 100
      },
      "盘锦": {
        "all_count": 276,
        "sum": 74750,
        "ok_count": 276,
        "err_count": 0,
        "val": 271,
        "ok_rate": 100
      },
      "抚顺": {
        "all_count": 226,
        "sum": 75353,
        "ok_count": 226,
        "err_count": 0,
        "val": 333,
        "ok_rate": 100
      },
      "阜新": {
        "all_count": 166,
        "sum": 70469,
        "ok_count": 166,
        "err_count": 0,
        "val": 425,
        "ok_rate": 100
      },
      "鞍山": {
        "all_count": 127,
        "sum": 59065,
        "ok_count": 127,
        "err_count": 0,
        "val": 465,
        "ok_rate": 100
      },
      "丹东": {
        "all_count": 84,
        "sum": 40202,
        "ok_count": 84,
        "err_count": 0,
        "val": 479,
        "ok_rate": 100
      },
      "营口": {
        "all_count": 60,
        "sum": 17253,
        "ok_count": 60,
        "err_count": 0,
        "val": 288,
        "ok_rate": 100
      },
      "铁岭": {
        "all_count": 55,
        "sum": 14458,
        "ok_count": 55,
        "err_count": 0,
        "val": 263,
        "ok_rate": 100
      },
      "辽阳": {
        "all_count": 49,
        "sum": 19896,
        "ok_count": 49,
        "err_count": 0,
        "val": 406,
        "ok_rate": 100
      },
      "葫芦岛": {
        "all_count": 40,
        "sum": 25676,
        "ok_count": 40,
        "err_count": 0,
        "val": 642,
        "ok_rate": 100
      },
      "锦州": {
        "all_count": 23,
        "sum": 8376,
        "ok_count": 23,
        "err_count": 0,
        "val": 364,
        "ok_rate": 100
      },
      "朝阳": {
        "all_count": 11,
        "sum": 4963,
        "ok_count": 11,
        "err_count": 0,
        "val": 451,
        "ok_rate": 100
      }
    },
    "sum": 1110778,
    "ok_count": 4346,
    "err_count": 0,
    "val": 256,
    "ok_rate": 100
  },
  "重庆": {
    "all_count": 4397,
    "city": {
      "重庆": {
        "all_count": 4397,
        "sum": 1453226,
        "err_count": 1,
        "ok_count": 4396,
        "val": 331,
        "ok_rate": 99.977
      }
    },
    "sum": 1453226,
    "err_count": 1,
    "ok_count": 4396,
    "val": 331,
    "ok_rate": 99.977
  },
  "陕西": {
    "all_count": 4173,
    "city": {
      "西安": {
        "all_count": 2854,
        "sum": 1000550,
        "err_count": 1,
        "ok_count": 2853,
        "val": 351,
        "ok_rate": 99.965
      },
      "宝鸡": {
        "all_count": 472,
        "sum": 303641,
        "err_count": 1,
        "ok_count": 471,
        "val": 645,
        "ok_rate": 99.788
      },
      "汉中": {
        "all_count": 167,
        "sum": 111357,
        "ok_count": 167,
        "err_count": 0,
        "val": 667,
        "ok_rate": 100
      },
      "渭南": {
        "all_count": 154,
        "sum": 124798,
        "ok_count": 154,
        "err_count": 0,
        "val": 810,
        "ok_rate": 100
      },
      "咸阳": {
        "all_count": 145,
        "sum": 46291,
        "ok_count": 145,
        "err_count": 0,
        "val": 319,
        "ok_rate": 100
      },
      "榆林": {
        "all_count": 99,
        "sum": 67429,
        "ok_count": 99,
        "err_count": 0,
        "val": 681,
        "ok_rate": 100
      },
      "安康": {
        "all_count": 97,
        "sum": 61412,
        "ok_count": 97,
        "err_count": 0,
        "val": 633,
        "ok_rate": 100
      },
      "商洛": {
        "all_count": 85,
        "sum": 89492,
        "ok_count": 85,
        "err_count": 0,
        "val": 1053,
        "ok_rate": 100
      },
      "延安": {
        "all_count": 73,
        "sum": 34106,
        "ok_count": 73,
        "err_count": 0,
        "val": 467,
        "ok_rate": 100
      },
      "铜川": {
        "all_count": 27,
        "sum": 22413,
        "ok_count": 27,
        "err_count": 0,
        "val": 830,
        "ok_rate": 100
      }
    },
    "sum": 1861489,
    "err_count": 2,
    "ok_count": 4171,
    "val": 446,
    "ok_rate": 99.952
  },
  "青海": {
    "all_count": 4682,
    "city": {
      "西宁": {
        "all_count": 3540,
        "sum": 907235,
        "err_count": 2,
        "ok_count": 3538,
        "val": 256,
        "ok_rate": 99.944
      },
      "海西蒙古族藏族自治州": {
        "all_count": 820,
        "sum": 168776,
        "ok_count": 820,
        "err_count": 0,
        "val": 206,
        "ok_rate": 100
      },
      "海东": {
        "all_count": 212,
        "sum": 58665,
        "ok_count": 212,
        "err_count": 0,
        "val": 277,
        "ok_rate": 100
      },
      "海北藏族自治州": {
        "all_count": 110,
        "sum": 25172,
        "ok_count": 110,
        "err_count": 0,
        "val": 229,
        "ok_rate": 100
      }
    },
    "sum": 1159848,
    "err_count": 2,
    "ok_count": 4680,
    "val": 248,
    "ok_rate": 99.957
  },
  "黑龙江": {
    "all_count": 3534,
    "city": {
      "哈尔滨": {
        "all_count": 1573,
        "sum": 607028,
        "err_count": 7,
        "ok_count": 1566,
        "val": 388,
        "ok_rate": 99.555
      },
      "鸡西": {
        "all_count": 563,
        "sum": 109658,
        "ok_count": 563,
        "err_count": 0,
        "val": 195,
        "ok_rate": 100
      },
      "佳木斯": {
        "all_count": 548,
        "sum": 109123,
        "ok_count": 548,
        "err_count": 0,
        "val": 199,
        "ok_rate": 100
      },
      "大庆": {
        "all_count": 316,
        "sum": 122831,
        "err_count": 1,
        "ok_count": 315,
        "val": 390,
        "ok_rate": 99.684
      },
      "牡丹江": {
        "all_count": 121,
        "sum": 75617,
        "err_count": 1,
        "ok_count": 120,
        "val": 630,
        "ok_rate": 99.174
      },
      "齐齐哈尔": {
        "all_count": 118,
        "sum": 29404,
        "ok_count": 118,
        "err_count": 0,
        "val": 249,
        "ok_rate": 100
      },
      "双鸭山": {
        "all_count": 84,
        "sum": 32683,
        "ok_count": 84,
        "err_count": 0,
        "val": 389,
        "ok_rate": 100
      },
      "绥化": {
        "all_count": 78,
        "sum": 33027,
        "ok_count": 78,
        "err_count": 0,
        "val": 423,
        "ok_rate": 100
      },
      "黑河": {
        "all_count": 40,
        "sum": 25173,
        "ok_count": 40,
        "err_count": 0,
        "val": 629,
        "ok_rate": 100
      },
      "伊春": {
        "all_count": 28,
        "sum": 19146,
        "ok_count": 28,
        "err_count": 0,
        "val": 684,
        "ok_rate": 100
      },
      "七台河": {
        "all_count": 23,
        "sum": 5345,
        "ok_count": 23,
        "err_count": 0,
        "val": 232,
        "ok_rate": 100
      },
      "鹤岗": {
        "all_count": 22,
        "sum": 6020,
        "ok_count": 22,
        "err_count": 0,
        "val": 274,
        "ok_rate": 100
      },
      "大兴安岭地区": {
        "all_count": 20,
        "sum": 5810,
        "ok_count": 20,
        "err_count": 0,
        "val": 291,
        "ok_rate": 100
      }
    },
    "sum": 1180865,
    "err_count": 9,
    "ok_count": 3525,
    "val": 335,
    "ok_rate": 99.745
  }
}

const mc = {
  "name": "测试指标（ms）",
  "visualMap": {
    "min": 0,
    "max": 1000,
    "inRange": {
      "color": [ "springgreen", "yellowgreen", "yellow", "orange", "orangered" ]
    }
  }
}



const option: Record<string, any> = {
  title: {
    text: mc.name,
    subtextStyle: {
      fontSize: 15,
      lineHeight: 20,
      fontWeight: 'bold',
    },
  },
  tooltip: {
    trigger: 'item',
    confine: true,
    formatter: mapTooltipFormatter,
  },
  visualMap: mc.visualMap,
  series: [
    {
      name: mc.name,
      type: 'map',
      map: 'china',
      zoom: 1.25,
      itemStyle: {
        normal: {
          label: {
            show: true
          }
        },
      },
      data: calcData(),
    }
  ]
}

function calcData(): any[] {
  const ans: any[] = []
  for (let prov in data) {
    let d = data[prov]
    ans.push({ name: prov, value: d.val, originData: d })
  }
  return ans
}

function mapTooltipFormatter(p: any): string
{
    let valStr = p.value;
    if (mc.name == '可用性') {
        valStr += '%';
    }
    let ans = '<b class="pull-left">' + p.name + '：' + valStr + '</b><br/>';
    ans += '<b class="pull-left">错误点：' + p.data.originData.err_count + '</b><br/>';
    ans += '<b class="pull-left">总点数：' + p.data.originData.all_count + '</b><br/>';
    ans += '<hr/><table class="my-table">';
    ans += '<tr><th style="text-align: left">城市</th><th>指标</th><th>错误点</th><th>总点数</th><th>可用性</th></tr>';
    let cMap = p.data.originData.city;
    for (let c in cMap) {
        ans += '<tr><td style="text-align: left">' + c + '</td>';
        if (mc.name == '可用性') {
            ans += '<td>' + cMap[c]['val'] + '%</td>';
        } else {
            ans += '<td>' + cMap[c]['val'] + '</td>';
        }
        ans += '<td>' + cMap[c]['err_count'] + '</td>';
        ans += '<td>' + cMap[c]['all_count'] + '</td>';
        ans += '<td>' + cMap[c]['ok_rate'] + '%</td></tr>';
    }
    ans += '</table>';
    return ans;
}
</script>
