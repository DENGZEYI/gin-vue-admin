<template>
    <div>
        <div class="gva-search-box">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <!-- 供应商 -->
                <el-form-item label="供应商" prop="provider_id">
                    <el-select v-model="searchInfo.provider_id" clearable placeholder="请选择"
                        @clear="() => { searchInfo.provider_id = undefined }">
                        <el-option v-for="(item, key) in providerOptions" :key="key" :label="item.label"
                            :value="item.value" />
                    </el-select>
                </el-form-item>
                <!-- 组别 -->
                <el-form-item label="专业组" prop="group_id">
                    <el-select v-model="searchInfo.group_id" clearable placeholder="请选择"
                        @clear="() => { searchInfo.group_id = undefined }">
                        <el-option v-for="(item, key) in groupOptions" :key="key" :label="item.label"
                            :value="item.value" />
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
            <el-table style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID">
                <el-table-column align="left" label="耗材ID" prop="ID" width="120" />
                <el-table-column align="left" label="耗材名称" prop="name" width="120" />
                <el-table-column align="left" label="专业组" prop="group_id" width="120">
                    <template #default="scope">{{ filterDict(scope.row.group_id, groupOptions) }}</template>
                </el-table-column>
                <el-table-column align="left" label="单位" prop="unit" width="120" />
                <el-table-column align="left" label="规格" prop="specification" width="120" />
                <el-table-column align="left" label="生产商" prop="factory_id" width="120">
                    <template #default="scope">{{ filterDict(scope.row.factory_id, factoryOptions) }}</template>
                </el-table-column>
                <el-table-column align="left" label="供应商" prop="provider_id" width="120">
                    <template #default="scope">{{ filterDict(scope.row.provider_id, providerOptions) }}</template>
                </el-table-column>
                <el-table-column align="left" label="按钮组">
                    <template #default="scope">
                        <el-button type="primary" link icon="DocumentAdd" size="small" @click="addRow(scope.row)">申购耗材
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="gva-pagination">
                <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page"
                    :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
                    @current-change="handleCurrentChange" @size-change="handleSizeChange" />
            </div>
        </div>

        <el-divider>
            <span>耗材申购单</span>
        </el-divider>

        <!-- 申请Goods表格 -->
        <el-form :inline="true" :model="applyFormData" ref="applyFormRef" label-width="80px" :rules="rules">
            <div class="gva-table-box">
                <!-- 表格按钮 -->
                <div class="gva-btn-list">
                    <el-popover v-model:visible="applyBtnVisible" placement="top" width="160">
                        <p>确定要提交申购单吗？</p>
                        <div style="text-align: right; margin-top: 8px;">
                            <el-button size="small" type="primary" link @click="applyBtnVisible = false">取消</el-button>
                            <el-button size="small" type="primary" @click="applyFunc">确定</el-button>
                        </div>
                        <template #reference>
                            <el-button icon="checked" size="small" style="margin-left: 10px;"
                                :disabled="!applyFormData.apply_details.length" @click="applyBtnVisible = true">提交申购单
                            </el-button>
                        </template>
                    </el-popover>
                </div>
                <!-- 表格主体 -->
                <el-table style="width: 100%" tooltip-effect="dark" :data="applyFormData.apply_details" row-key="ID">
                    <el-table-column align="left" label="耗材ID" prop="goods_dict_id" width="200" />
                    <el-table-column align="left" label="耗材名称" prop="goods_dict_name" width="200" />
                    <el-table-column align="left" label="组别" prop="group" width="200" />
                    <el-table-column align="left" label="生产商" prop="factory" width="200" />
                    <el-table-column align="left" label="单位" prop="unit" width="120" />
                    <el-table-column align="left" label="耗材规格" prop="specification" width="200" />
                    <el-table-column align="left" label="申请数量" prop="apply_number" width="200">
                        <template #default="scope">
                            <el-form-item label=" " :prop="'apply_details.' + scope.$index + '.apply_number'"
                                :rules='rules.apply_number'>
                                <el-input onkeyup="value=value.replace(/^(0+)|[^\d]+/g,'')"
                                    v-model.number="scope.row.apply_number" placeholder="请输入入库数量" />
                            </el-form-item>
                        </template>
                    </el-table-column>
                </el-table>
            </div>
        </el-form>
    </div>
</template>
  
<script>
export default {
    name: 'busApply'
}
</script>
  
<script setup>
import {
    getBusGoodsDictList,
    applyBusGoods
} from '@/api/busGoodsDict'

// 全量引入格式化工具 请按需保留
import { formatDate, getBusDictFunc, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'


// 验证规则
const rules = reactive({
    name: [{
        required: true,
        message: '必填项',
        trigger: ['blur'],
    }],
    apply_number: [{
        required: true,
        message: '必填项',
        trigger: ['blur'],
    }],
})


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


const applyBtnVisible = ref(false)
const applyFormRef = ref()
const applyFormData = ref({
    apply_details: []
})

// 判断是否重复添加需要申请的耗材
const isRepeat = (row) => {
    for (var i = 0; i < applyFormData.value.apply_details.length; i++) {
        if (applyFormData.value.apply_details[i].goods_dict_id == row.ID) {
            return true
        }
    }
    return false
}

// 添加耗材以完成申请操作
const addRow = async (row) => {
    // 避免重复添加
    if (isRepeat(row)) {
        ElMessage({
            type: 'warning',
            message: '不允许重复添加'
        })
        return
    }
    const applly_detail = {
        goods_dict_name: row.name,
        goods_dict_id: row.ID,
        unit: row.unit,
        specification: row.specification,
        group: row.group.name,
        factory: row.factory.name,
        apply_number: row.apply_number
    }
    applyFormData.value.apply_details.push(applly_detail)
    ElMessage({
        type: 'success',
        message: '添加成功'
    })
}

// 提交申请
const applyFunc = async () => {
    applyFormRef.value.validate(async (valid) => {
        if (!valid) return
        const rst = await applyBusGoods(applyFormData.value)
        if (rst.code === 0) {
            ElMessage({
                type: 'success',
                message: '成功提交申购单'
            })
        }
        applyBtnVisible.value = false
    })
}




</script>
  
<style>

</style>
  