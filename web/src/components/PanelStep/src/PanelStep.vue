<template>
  <div>
    <el-row :class="isLast ? 'last-step' : ''">
      <el-col :span="24">
        <div class="bot" style="height: 20px; margin-bottom: 13px; background: #fff">
          <div class="bline-dotted">
            <div class="category-index">{{ num }}</div>
            <div class="category-name">{{ title }}</div>
          </div>
        </div>
      </el-col>
    </el-row>
    <div>
      <div style="padding: 10px 30px; margin-bottom: 30px">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PanelStep',
  props: {
    num: {
      type: [String, Number],
      required: true,
      default: 1
    },
    title: {
      type: String
    },
    isLast: {
      type: Boolean,
      default: false
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
//.bline-dotted {
//  border-bottom: 1px dotted #000;
//}

.category-index {
  position: absolute;
  padding-top: 2px;
  text-align: center;
  color: #fff;
  font-size: 14px;
  left: 0;
  top: 6px;
  width: 20px;
  height: 20px;
  background: #4385f5;
  border-radius: 50px;
}

.category-name {
  position: absolute;
  left: 26px;
  font-size: 12px;
  top: 1px;
  background: #fff;
  height: 28px;
  padding: 8px 8px;
  color: #303133;
  font-weight: 600;
}
</style>

<style lang="scss">
.tc-step-border {
  border-left: 1px dashed #d8dce5;
  margin-left: 10px;

  .l {
    position: relative;
    left: -13px;
  }

  .last-step {
    margin-bottom: 20px;

    ~ div {
      background: #fff;
      margin-top: -30px;
    }
  }

  .bline-dotted {
    border-bottom: 0 !important;
  }
}
</style>
