<template>
    <div>
        <div class="gva-table-box">
            <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID">
                <el-table-column align="left" label="采购时间" prop="purchase_time" width="180">
                    <template #default="scope">{{ formatDate(scope.row.purchase_date) }}</template>
                </el-table-column>
                <el-table-column align="left" label="申请单ID" prop="ID" width="120" />
                <el-table-column align="left" label="耗材名称" prop="goods_name" />
                <el-table-column align="left" label="规格" prop="specification" />
                <el-table-column align="left" label="申请数量" prop="apply_number" />
                <el-table-column align="left" label="已入库数量" prop="arrival_number" width="140" />
                <el-table-column align="left" label="库存点" prop="group_name" width="140" />
                <el-table-column align="left" label="供应商名称" prop="provider_name" width="140" />
                <el-table-column align="left" label="申请时间" width="180">
                    <template #default="scope">{{ formatDate(scope.row.apply_date) }}</template>
                </el-table-column>
                <el-table-column align="left" label="申请人" prop="applicant_name" width="120" />
                <el-table-column align="left" label="申请审批时间" width="180">
                    <template #default="scope">{{ formatDate(scope.row.approve_date) }}</template>
                </el-table-column>
                <el-table-column align="left" label="审批人" prop="approver_name" width="120" />
                <el-table-column align="left" label="采购人" prop="purchaser_name" width="120" />
                <el-table-column align="left" fixed="right" label="按钮组" width="120">
                    <template #default="scope">
                        <el-button type="primary" @click="addRow(scope.row)">添加到入库单</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="gva-pagination">
                <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page"
                    :page-size="pageSize" :page-sizes="[5,10, 30, 50, 100]" :total="total"
                    @current-change="handleCurrentChange" @size-change="handleSizeChange" />
            </div>
        </div>
        <el-divider>
            <span>入库单</span>
        </el-divider>
        <!-- 入库详情 -->
        <el-form :inline="true" :model="ingressFormData" ref="ingressFormRef" label-width="80px" :rules="rules">
            <el-form-item label="收货仓库">
                <el-input v-model="ingressFormData.group_name" disabled />
            </el-form-item>
            <el-form-item label="送货单号" prop="delivery_number">
                <el-input v-model="ingressFormData.delivery_number" placeholder="请输入送货单号"></el-input>
            </el-form-item>
            <el-form-item label="送货日期" prop="deliveryDate">
                <el-date-picker v-model="ingressFormData.deliveryDate" placeholder="请输入日期" type="date"
                    style="width: 100%" />
            </el-form-item>
            <el-form-item label="入库日期">
                <el-date-picker v-model="ingressFormData.ingress_date" type="date" style="width: 100%" disabled />
            </el-form-item>
            <el-form-item>
                <el-popover v-model:visible="ingressBtnVisible" placement="top" width="160">
                    <p>确定要入库吗？</p>
                    <div style="text-align: right; margin-top: 8px;">
                        <el-button size="small" type="primary" @click="ingressBtnVisible = false">取消</el-button>
                        <el-button size="small" type="primary" @click="ingressFunc">确定入库</el-button>
                    </div>
                    <template #reference>
                        <el-button type="primary" :disabled="!ingressFormData.ingress_details.length"
                            @click="ingressBtnVisible = true">入库</el-button>
                    </template>
                </el-popover>
            </el-form-item>
            <div class="gva-table-box">
                <el-table :data="ingressFormData.ingress_details">
                    <el-table-column align="left" label="耗材名称" prop="goods_name" width="120" />
                    <el-table-column align="left" label="规格" prop="specification" />
                    <el-table-column align="left" label="批号" prop="batch">
                        <template #default="scope">
                            <el-form-item label=" " :prop=" 'ingress_details.' + scope.$index + '.batch' "
                                :rules='rules.batch'>
                                <el-input ref="input" v-model="scope.row.batch" placeholder="请输入批号" style="width: 100%">
                                </el-input>
                            </el-form-item>
                        </template>
                    </el-table-column>
                    <el-table-column align="left" label="有效期" prop="expiration_date">
                        <template #default="scope">
                            <el-form-item label=" " :prop=" 'ingress_details.' + scope.$index + '.expiration_date' "
                                :rules='rules.expiration_date'>
                                <el-date-picker :clearable="false" v-model="scope.row.expiration_date" type="date"
                                    placeholder="请输入有效期">
                                </el-date-picker>
                            </el-form-item>
                        </template>
                    </el-table-column>
                    <el-table-column align="left" label="入库数量" prop="ingress_number">
                        <template #default="scope">
                            <el-form-item label=" " :prop=" 'ingress_details.' + scope.$index + '.ingress_number' "
                                :rules='rules.ingress_number'>
                                <el-input onkeyup="value=value.replace(/^(0+)|[^\d]+/g,'')"
                                    v-model.number="scope.row.ingress_number" placeholder="请输入入库数量" />
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
    name: 'BusGroup'
}
</script>

<script setup>
import {
    getBusIngressList,
    ingress
} from '@/api/busIngress'


import { getDictFunc, getBusDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})


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
    const table = await getBusIngressList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
    }
}

getTableData()

// ============== 表格控制部分结束 ===============


const ingressBtnVisible = ref(false)
const ingressFormData = ref({
    order_id: undefined,
    group_name: undefined,
    delivery_number: undefined,
    ingress_date: new Date(),
    ingress_details: []
})
const ingressFormRef = ref()

const rules = reactive({
    delivery_number: [{
        required: true,
        message: '必填项',
        trigger: ['blur'],
    }],
    deliveryDate: [{
        required: true,
        message: '必填项',
        trigger: ['blur'],
    }],
    batch: [{
        required: true,
        message: '必填项',
        trigger: ['blur'],
    }],
    expiration_date: [{
        required: true,
        message: '必填项',
        trigger: ['blur'],
    }],
    ingress_number: [{
        required: true,
        message: '必填项',
        trigger: ['blur'],
    }],
})
// 判断是否重复添加需要入库的耗材
const isRepeat = (row) => {
    for (var i = 0; i < ingressFormData.value.ingress_details.length; i++) {
        if (ingressFormData.value.ingress_details[i].goods_dict_id == row.goods_dict_id) {
            return true
        }
    }
    return false
}
// 添加耗材以完成入库操作
const addRow = async (row) => {
    if (ingressFormData.value.order_id === undefined) {
        ingressFormData.value.order_id = row.ID
        ingressFormData.value.group_name = row.group_name
    }
    if (ingressFormData.value.order_id !== row.ID) {
        ElMessage({
            type: 'warning',
            message: '不允许添加属于不同申请单的耗材'
        })
        return
    }
    // 避免重复添加
    if (isRepeat(row)) return
    // 要入库的耗材详情
    const ingress_detail = {
        goods_dict_id: row.goods_dict_id,
        goods_name: row.goods_name,
        specification: row.specification,
        batch: undefined,
        expiration_date: undefined,
        ingress_number: undefined,
    }
    ingressFormData.value.ingress_details.push(ingress_detail)
    ElMessage({
        type: 'success',
        message: '添加成功'
    })
}
// 提交入库单
const ingressFunc = async () => {
    ingressFormRef.value.validate(async (valid) => {
        if (!valid) return
        const rst = await ingress(ingressFormData.value)
        if (rst.code === 0) {
            ElMessage({
                type: 'success',
                message: '入库成功'
            })
        } else {
            ElMessage({
                type: 'error',
                message: '入库失败'
            })
        }
        getTableData()
        ingressBtnVisible.value = false
        // 清空
        ingressFormData.value = {
            ingress_details: []
        }
    })
}

</script>
  
<style>

</style>