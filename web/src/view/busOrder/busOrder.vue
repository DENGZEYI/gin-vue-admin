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
            </el-table>
          </template>
        </el-table-column>
        <el-table-column align="left" label="申请表ID" prop="ID" width="120" />
        <el-table-column align="left" label="申请提交日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="申请人" prop="Applicant.nickName" width="120" />
        <el-table-column align="left" label="申请审批日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.UpdatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="审批人" prop="Approver.nickName" width="120" />
        <el-table-column align="left" label="申请状态" prop="state" width="120">
          <template #default="scope">{{ filterDict(scope.row.state, applyStateOptions) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" @click="approveBusOrderFunc(scope.row)">审批
            </el-button>
            <el-button type="primary" link icon="edit" size="small" @click="ingressBusOrderFunc(scope.row)">入库
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <!-- 审批状态 -->
    <el-dialog v-model="approveDialogFormVisible" :before-close="approveCloseDialog" title="申请表详情">
      <el-form :model="approveFormData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="申请状态:" prop="state">
          <el-select v-model="approveFormData.state" placeholder="请选择" :clearable="true">
            <el-option v-for="(item, key) in applyStateOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="approveCloseDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="approveEnterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'BusOrder'
}
</script>

<script setup>
import {
  createBusOrder,
  deleteBusOrder,
  deleteBusOrderByIds,
  updateBusOrder,
  findBusOrder,
  getBusOrderList,
  getOrderDetails
} from '@/api/busOrder'

// 全量引入格式化工具 请按需保留
import { getDictFunc, getBusDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const approveFormData = ref({
  ID: 0,
  state: 0,
})

// 验证规则
const rule = reactive({
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

// 更新行
// 更新审批状态
const approveBusOrderFunc = async (row) => {
  const res = await findBusOrder({ ID: row.ID })
  if (res.code === 0) {
    approveFormData.value = res.data.rebusOrder
    approveDialogFormVisible.value = true
  }
}


// 弹窗控制标记
const approveDialogFormVisible = ref(false)

// 关闭弹窗
const approveCloseDialog = () => {
  approveDialogFormVisible.value = false
  approveFormData.value = {
    ID: 0,
    state: 0,
  }
}

// 弹窗确定
const approveEnterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    res = await updateBusOrder(approveFormData.value)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      approveCloseDialog()
      getTableData()
    }
  })
}
</script>

<style>
</style>
