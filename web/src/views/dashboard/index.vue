<template>
  <div class="dashboard-container">
    <div class="dash-panel">
      <dash-header></dash-header>
      <dash-body></dash-body>
      <dash-foot></dash-foot>
    </div>
  </div>
</template>

<script setup lang="ts">
import DashHeader from './components/DashHeader'
import DashBody from './components/DashBody'
import DashFoot from './components/DashFoot'
</script>

<style lang="scss" scoped>
@import './dash.scss';

.dashboard-container {
  padding: 0;
  background: #f6f8f9;
  height: 100%;
  margin: 0 -15px;

  .dash-panel {
    padding: 20px;
  }
}
</style>
