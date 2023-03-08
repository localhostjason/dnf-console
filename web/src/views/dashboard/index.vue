<template>
  <div class="dashboard-container">
    <!--<panel-title title="概览" is-line></panel-title>-->
  </div>
</template>

<script>
export default {
  name: 'Dashboard',
  directives: {},
  components: {},
  computed: {},
  data() {
    return {}
  },
  created() {},
  methods: {}
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.el-row {
  .el-col {
    margin: 10px 0;
  }
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: '';
}

.clearfix:after {
  clear: both;
}

.box-card {
}
</style>
