<template>
  <a-result status="403" title="抱歉，您无权限查看此页面！">
    <template #extra>
      <router-link to="/" class="link">
        <button class="btn btn-primary">
          <i class="fa fa-home"></i> 返回首页
        </button>
      </router-link>
    </template>
  </a-result>
</template>
