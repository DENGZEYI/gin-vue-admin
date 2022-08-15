<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <!-- 供应商 -->
        <el-form-item label="供应商" prop="ProviderID">
          <el-select v-model="searchInfo.ProviderID" clearable placeholder="请选择"
            @clear="() => { searchInfo.ProviderID = undefined }">
            <el-option v-for="(item, key) in providerOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <!-- 组别 -->
        <el-form-item label="组别" prop="GroupID">
          <el-select v-model="searchInfo.GroupID" clearable placeholder="请选择"
            @clear="() => { searchInfo.GroupID = undefined }">
            <el-option v-for="(item, key) in groupOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <!-- 名称 -->
        <el-form-item label="名称">
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
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="耗材ID" prop="ID" width="120" />
        <el-table-column align="left" label="耗材名称" prop="name" width="120" />
        <el-table-column align="left" label="价格" prop="price" width="120" />
        <el-table-column align="left" label="生产厂商" prop="factory" width="120" />
        <el-table-column align="left" label="规格" prop="specification" width="120" />
        <el-table-column align="left" label="组别" prop="groupId" width="120">
          <template #default="scope">{{ filterDict(scope.row.GroupID, groupOptions) }}</template>
        </el-table-column>
        <!-- ProviderID 注意名称大小写 -->
        <el-table-column align="left" label="供应商" prop="ProviderID" width="120">
          <template #default="scope">{{ filterDict(scope.row.ProviderID, providerOptions) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button"
              @click="updatebusGoodsDictFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            <el-button type="primary" link icon="DocumentAdd" size="small" @click="addRow(scope.row)">添加</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>

    <!-- 申请Goods表格 -->
    <div class="gva-table-box">
      <!-- 表格按钮 -->
      <div class="gva-btn-list">
        <el-popover v-model:visible="selectedDataSubVisible" placement="top" width="160">
          <p>确定要提交申请吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="selectedDataSubVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onSubmitSelectedData">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="checked" size="small" style="margin-left: 10px;" :disabled="!selectedData.length"
              @click="selectedDataSubVisible = true">提交申请</el-button>
          </template>
        </el-popover>
      </div>
      <!-- 表格主体 -->
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="selectedData" row-key="ID">
        <el-table-column align="left" label="耗材名称" prop="name" width="120" />
        <el-table-column align="left" label="数量" prop="number" width="120">
          <template #default="scope">
            <input type="number" v-model="scope.row.number" v-show="scope.row.iseditor" />
            <span v-show="!scope.row.iseditor">{{ scope.row.number }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button type="warning" @click="edit(scope.row, scope)">编辑</el-button>
            <el-button type="danger" @click="save(scope.row)">保存</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- Goods表格弹窗 -->
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="120px">
        <el-form-item label="耗材名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="价格:" prop="price">
          <el-input v-model.number="formData.price" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="生产厂商:" prop="factory">
          <el-input v-model="formData.factory" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="规格:" prop="specification">
          <el-input v-model="formData.specification" :clearable="true" placeholder="请输入" />
        </el-form-item>

        <el-form-item label="组别:" prop="GroupID">
          <el-select v-model="formData.GroupID" placeholder="请选择" :clearable="true">
            <el-option v-for="(item, key) in groupOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="供应商:" prop="ProviderID">
          <el-select v-model="formData.ProviderID" placeholder="请选择" :clearable="true">
            <el-option v-for="(item, key) in providerOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  applyBusGoodsByIds
} from '@/api/busGoodsDict'

// 全量引入格式化工具 请按需保留
import { formatDate, getBusDictFunc, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'


// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: '',
  price: 0,
  factory: '',
  specification: '',
  GroupID: undefined,
  ProviderID: undefined,
})

// 验证规则
const rule = reactive({
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const selectedData = ref([])
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
const setOptions = async () => {
  providerOptions.value = await getBusDictFunc('provider')
  groupOptions.value = await getBusDictFunc('group')
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

// selectedData提交标记
const selectedDataSubVisible = ref(false)

// 将所选行的内容添加到selectedData中
const addRow = (row) => {
  var selectedFlag = 0 //标识当前元素是否已经存在于字典中
  selectedData.value.forEach(element => {
    if (row.ID === element.ID) {
      selectedFlag = 1
    }
  });
  if (selectedFlag == 1) {
    return
  }
  selectedData.value.push({ ID: row.ID, name: row.name, number: 0, iseditor: false })
}

// 修改selectedData中的数量
const edit = (row) => {
  row.iseditor = true;
}
const save = (row) => {
  row.iseditor = false;
}

const onSubmitSelectedData = async () => {
  const ids = []
  selectedData.value && selectedData.value.map(item => {
    ids.push({ ID: item.ID, number: item.number })
  })
  const res = await applyBusGoodsByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '成功提交申请'
    })
    selectedDataSubVisible.value = false
  }
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
