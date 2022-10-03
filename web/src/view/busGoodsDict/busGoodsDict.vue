<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <!-- 供应商 -->
        <el-form-item label="供应商" prop="provider_id">
          <el-select v-model="searchInfo.provider_id" clearable placeholder="请选择"
            @clear="() => { searchInfo.provider_id = undefined }">
            <el-option v-for="(item, key) in providerOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <!-- 组别 -->
        <el-form-item label="组别" prop="group_id">
          <el-select v-model="searchInfo.group_id" clearable placeholder="请选择"
            @clear="() => { searchInfo.group_id = undefined }">
            <el-option v-for="(item, key) in groupOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <!-- 名称 -->
        <el-form-item label="耗材名称">
          <el-input v-model="searchInfo.name" placeholder="请输入名称" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <!-- Goods表格 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length"
              @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="耗材ID" prop="ID" width="120" />
        <el-table-column align="left" label="耗材名称" prop="name" width="120" />
        <el-table-column align="left" label="组别" prop="group_id" width="120">
          <template #default="scope">{{ filterDict(scope.row.group_id, groupOptions) }}</template>
        </el-table-column>
        <el-table-column align="left" label="价格" prop="price" width="120" />
        <el-table-column align="left" label="单位" prop="unit" width="120" />
        <el-table-column align="left" label="规格" prop="specification" width="120" />
        <el-table-column align="left" label="合同代码" prop="contract_code" width="120" />
        <el-table-column align="left" label="生产商" prop="factory_id" width="120">
          <template #default="scope">{{ filterDict(scope.row.factory_id, factoryOptions) }}</template>
        </el-table-column>
        <el-table-column align="left" label="供应商" prop="provider_id" width="120">
          <template #default="scope">{{ filterDict(scope.row.provider_id, providerOptions) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button"
              @click="updatebusGoodsDictFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>

    <!-- Goods表格弹窗 -->
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rules" label-width="120px">
        <el-form-item label="耗材名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="组别:" prop="group_id">
          <el-select v-model="formData.group_id" placeholder="请选择" :clearable="true">
            <el-option v-for="(item, key) in groupOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="生产商:" prop="factory_id">
          <el-select v-model="formData.factory_id" placeholder="请选择" :clearable="true">
            <el-option v-for="(item, key) in factoryOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
          <el-form-item label="供应商:" prop="provider_id">
            <el-select v-model="formData.provider_id" placeholder="请选择" :clearable="true">
              <el-option v-for="(item, key) in providerOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
          </el-form-item>
        </el-form-item>
        <el-form-item label="规格:" prop="specification">
          <el-input v-model="formData.specification" :clearable="true" placeholder="请输入" />
        </el-form-item>

        <el-form-item label="合同代码:" prop="contract_code">
          <el-input v-model="formData.contract_code" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="价格:" prop="price">
          <el-input v-model.number="formData.price" :clearable="true" placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'busGoodsDict'
}
</script>

<script setup>
import {
  createBusGoodsDict,
  deleteBusGoodsDict,
  deleteBusGoodsDictByIds,
  updateBusGoodsDict,
  findBusGoodsDict,
  getBusGoodsDictList,
} from '@/api/busGoodsDict'

// 全量引入格式化工具 请按需保留
import { formatDate, getBusDictFunc, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'


// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: '',
  price: 0,
  specification: '',
  factory_id: undefined,
  group_id: undefined,
  provider_id: undefined,
})

// 验证规则
const rules = reactive({
  name: [{
    required: true,
    message: '必填项',
    trigger: ['blur'],
  }],
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
  const table = await getBusGoodsDictList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const providerOptions = ref([])
const groupOptions = ref([])
const factoryOptions = ref([])
const setOptions = async () => {
  providerOptions.value = await getBusDictFunc('provider')
  groupOptions.value = await getBusDictFunc('group')
  factoryOptions.value = await getBusDictFunc("factory")
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteBusGoodsDictFunc(row)
  })
}


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async () => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.ID)
    })
  const res = await deleteBusGoodsDictByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updatebusGoodsDictFunc = async (row) => {
  const res = await findBusGoodsDict({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reBusGoodsDict
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteBusGoodsDictFunc = async (row) => {
  const res = await deleteBusGoodsDict({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: '',
    price: 0,
    factory: '',
    specification: '',
    GroupID: undefined,
    ProviderID: undefined,
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createBusGoodsDict(formData.value)
        break
      case 'update':
        res = await updateBusGoodsDict(formData.value)
        break
      default:
        res = await createBusGoodsDict(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}



</script>

<style>

</style>
