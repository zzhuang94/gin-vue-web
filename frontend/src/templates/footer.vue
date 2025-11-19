<template>
  <div>
    <span class="text-muted">
      Copyright &copy; 2025 zzhuang94
      <a href="#" class="link">首页地址</a>
    </span>

    <a-space :size="20" class="pull-right">
      <a href="#" target="_blank">
        <span class="text-brand">
          <i class="fa fa-book"></i> 使用文档
        </span>
      </a>
      <a href="#"  target="_blank">
        <span class="text-brand">
          <i class="fa fa-share-alt"></i> API文档
        </span>
      </a>
      <a href="#" target="_blank" >
        <span class="text-brand">
          <i class="fa fa-commenting"></i> 提交反馈
        </span>
      </a>
    </a-space>
  </div>
</template>

<style scoped>
a>span {
  font-weight: bold;
}
</style>
