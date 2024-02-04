<template>
  <div>
    <panel-title title="拍卖行"></panel-title>

    <h5>TODO: 测试</h5>

    <el-row>
      <el-col :span="24">
        <el-alert title="立即开启后，请重新登录游戏" type="warning" />
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="24">
        <div>
          <el-table :data="state" border fit>
            <el-table-column label="名称" prop="desc" width="120"></el-table-column>
            <el-table-column label="状态" prop="state" width="100">
              <template #default="scope">
                <el-tag type="warning" v-if="!scope.row.visible">未开启</el-tag>
                <el-tag v-else type="success">开启</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="定时检查-开启" prop="-" width="130">
              <template #default="scope">
                <el-button type="primary" link @click="openIntervalAu(scope.row.name, scope.row.visible)"
                  >配置
                </el-button>
              </template>
            </el-table-column>
            <el-table-column label="操作" prop="-">
              <template #default="scope">
                <el-button type="primary" link @click="openAu(scope.row.name, scope.row.visible)">立即开启</el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import PanelTitle from '@/components/PanelTitle'
import { ref } from 'vue'
import { AuctionState } from '@/views/gm/auction/model'
import { getAuctionState, openAuction } from '@/api/gm/auction'
import { errorMessage } from '@/utils/element/message'
import { confirmWarning } from '@/utils/element/messageBox'

const state = ref<AuctionState[]>([])

const getAuctionStateInfo = async () => {
  try {
    state.value = await getAuctionState()
  } catch (e) {}
}

const openAu = async (name: string, visible: boolean) => {
  try {
    await confirmWarning('此操作将开启，是否继续')
    if (visible) {
      errorMessage('已开启，无需再点开启')
      return
    }

    await openAuction(name)
    await getAuctionStateInfo()
  } catch (e) {}
}

const openIntervalAu = async (name: string, visible: boolean) => {
  alert('todo：每个月 几号 定时检查，并开启')
}
getAuctionStateInfo()
</script>

<style scoped></style>
