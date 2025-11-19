<template>
  <a-result status="500" title="抱歉，出错了！">
    <template #extra>
      <a href="#" class="btn btn-warning">
        <i class="fa fa-commenting"></i> 报障
      </a>
      <router-link to="/" class="link">
        <button class="btn btn-primary">
          <i class="fa fa-home"></i> 返回首页
        </button>
      </router-link>
    </template>
  </a-result>
</template>
