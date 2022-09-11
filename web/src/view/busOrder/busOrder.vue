<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @expand-change="expandChange">
        <el-table-column label="详情" type="expand" width="120">
          <template #default="props">
            <h3>采购申请单详情</h3>
            <el-table :data="props.row.details">
              <el-table-column label="耗材名称" prop="goods_dict.name" width="120" />
              <el-table-column label="申请数量" prop="number" width="120" />
              <el-table-column label="已入库数量" prop="number_already_in" width="160" />
            </el-table>
          </template>
        </el-table-column>
        <el-table-column align="left" label="申请单ID" prop="ID" width="120" />
        <el-table-column align="left" label="申请提交日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="申请人" prop="applicant.nickName" width="120" />
        <el-table-column align="left" label="申请审批日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.UpdatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="审批人" prop="approver.nickName" width="120" />
        <el-table-column align="left" label="申请状态" prop="state" width="120">
          <template #default="scope">{{ filterDict(scope.row.state, applyStateOptions) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="success" @click="approveFunc(scope.row,'success')">审批通过</el-button>
            <el-button type="danger" @click="approveFunc(scope.row,'fail')">审批退回</el-button>
            <el-button type="primary" @click="purchaseFunc(scope.row)">采购</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BusOrder'
}
</script>

<script setup>
import {
  getBusOrderList,
  getOrderDetails,
  approve,
  purchase
} from '@/api/busOrder'

// 全量引入格式化工具 请按需保留
import { getDictFunc, getBusDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage } from 'element-plus'
import { async } from 'q';
import { ref, reactive } from 'vue'
// 审批采购单所需的数据
const approveData = ref({
  ID: 0,
  state: 0,
})
// 采购采购单所需的数据
const purchaseData = ref({
  ID: 0,
})


const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getBusOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============
const applyStateOptions = ref([])
// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
  applyStateOptions.value = await getBusDictFunc('applyState')
}

// 获取需要的字典 可能为空 按需保留
setOptions()

const expandChange = async (row) => {
  const res = await getOrderDetails({ ID: row.ID })
  if (res.code === 0) {
    row.details = res.data.rebusOrder.bus_order_details
  }
}

// 审批订购单
const approveFunc = async (row, type) => {
  if (type === 'success') {
    approveData.state = 2
  } else if (type === 'fail') {
    approveData.state = 3
  }else{
    approveData.state = 1
  }
  approveData.ID = row.ID
  const res = await approve(approveData)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '订单状态更新成功'
    })
  } else {
    ElMessage({
      type: 'fail',
      message: '订单状态更新失败'
    })
  }
  getTableData()
}

// 采购订购单
const purchaseFunc = async (row) => {
  purchaseData.ID = row.ID
  const res = await purchase(purchaseData)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '状态更新成功'
    })
  } else {
    ElMessage({
      type: 'fail',
      message: '状态更新失败'
    })
  }
  getTableData()
}


</script>

<style>

</style>
