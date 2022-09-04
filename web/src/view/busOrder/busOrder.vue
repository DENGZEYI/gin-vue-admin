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
              <el-table-column align="left" label="按钮组">
                <template #default="scope">
                  <el-button type="primary" link icon="edit" size="small" @click="addBusIngressFunc(props.row,scope.row)">入库
                  </el-button>
                </template>
              </el-table-column>
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
            <el-button type="primary" link icon="edit" size="small" @click="approveBusOrderFunc(scope.row)">审批
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
    <!-- 入库 -->
    <el-dialog v-model="ingressDialogFormVisible" :before-close="ingressCloseDialog" title="入库详情">
      <el-form :model="ingressFormData" label-position="right" ref="elFormRef" :rules="rule" label-width="120px">
        <el-form-item label="耗材名称:" prop="name">
          <el-input v-model="ingressFormData.goods_dict.name" :disabled="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="入库数量:" prop="ingressNumber">
          <el-input v-model.number="ingressFormData.ingress_number" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="批号:" prop="batch">
          <el-input v-model="ingressFormData.batch" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="有效日期:" prop="expiration_date">
          <el-date-picker v-model="ingressFormData.expiration_date" type="date" :default-time="defaultTime" placeholder="请输入有效日期" />
        </el-form-item>
        
        <el-form-item label="送货单号:" prop="invoice_number">
          <el-input v-model="ingressFormData.delivery_number" :clearable="true" placeholder="请输入" />
        </el-form-item>

        <el-form-item label="发票号:" prop="invoice_number">
          <el-input v-model="ingressFormData.invoice_number" :clearable="true" placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="ingressCloseDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="ingressEnterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
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
  approveBusOrder,
  ingressBusOrder,
  findBusOrder,
  getBusOrderList,
  getOrderDetails,
  getOrderDetail
} from '@/api/busOrder'

// 全量引入格式化工具 请按需保留
import { getDictFunc, getBusDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

// 审批提交的表单
const approveFormData = ref({
  ID: 0,
  state: 0,
})
// 入库提交的表单
const ingressFormData = ref({
})

// 有效日期
const value1 = ref('')

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

// 入库
const addBusIngressFunc = async (pprow,prow) => {
  const res = await getOrderDetail({
    bus_order_id: pprow.ID,
    goods_dict_id: prow.goods_dict.ID,
  })
  if (res.code === 0) {
    ingressFormData.value = res.data.reOrderDetail

    ingressDialogFormVisible.value = true
  }
}


// 弹窗控制标记
const approveDialogFormVisible = ref(false)
const ingressDialogFormVisible = ref(false)

// 审批：关闭弹窗
const approveCloseDialog = () => {
  approveDialogFormVisible.value = false
  approveFormData.value = {
    ID: 0,
    state: 0,
  }
}

// 入库：关闭弹窗
const ingressCloseDialog = () => {
  ingressDialogFormVisible.value = false
}

// 审批:弹窗确定
const approveEnterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    res = await approveBusOrder(approveFormData.value)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '审批完成'
      })
      approveCloseDialog()
      getTableData()
    }
  })
}

// 入库：弹窗确认
const ingressEnterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    res = await ingressBusOrder(ingressFormData.value)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '入库成功'
      })
      ingressCloseDialog()
    }
  })
}
</script>

<style>
</style>
