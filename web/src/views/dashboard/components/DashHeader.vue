<template>
  <div class="dash-header">
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card class="box-card" shadow="hover">
          <div>
            <span class="fs-14">活跃用户</span>
            <p><span>6,300</span></p>
            <div class="right-icon">
              <svg-icon icon-class="person"></svg-icon>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="box-card" shadow="hover">
          <div>
            <span class="fs-14">活跃角色</span>
            <p><span>16,300</span></p>
            <div class="right-icon">
              <svg-icon icon-class="nested"></svg-icon>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="box-card" shadow="hover">
          <div>
            <span class="fs-14">充值D币</span>
            <p><span>9,900</span></p>
            <div class="right-icon">
              <svg-icon icon-class="dollar" class-name="rela"></svg-icon>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="box-card" shadow="hover">
          <div>
            <span class="fs-14">充值D点</span>
            <p><span>26,300</span></p>
            <div class="right-icon">
              <svg-icon icon-class="yen" class-name="rela"></svg-icon>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts"></script>

<style lang="scss" scoped>
.dash-header {
  .el-card {
    position: relative;
  }

  p {
    font-size: 28px;
  }

  .right-icon {
    position: absolute;
    top: 53%;
    right: 25px;
    width: 60px;
    height: 60px;
    line-height: 60px;
    color: var(--el-color-primary);
    text-align: center;
    background: var(--el-color-primary-light-9);
    border-radius: 50%;
    transform: translateY(-50%);

    .svg-icon {
      font-size: 28px;
      &.rela {
        position: relative;
        left: 3px;
        top: 4px;
      }
    }
  }
}
</style>
