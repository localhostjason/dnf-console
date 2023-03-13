<template>
  <div>
    <el-row>
      <panel-title title="邮件工具"></panel-title>
    </el-row>

    <div class="tc-step-border">
      <div class="l">
        <panel-step num="1" title="账号选择"> </panel-step>
        <panel-step num="2" title="角色选择"> </panel-step>
        <panel-step num="3" title="AAA" class="last-step"></panel-step>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import PanelTitle from '@/components/PanelTitle'
import PanelStep from '@/components/PanelStep'
</script>

<style scoped></style>
