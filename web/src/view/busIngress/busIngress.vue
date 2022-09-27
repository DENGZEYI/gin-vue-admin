<template>
    <div>
        <div class="gva-table-box">
            <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
                @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55" />
                <el-table-column align="left" label="采购时间" prop="purchase_time" width="180">
                    <template #default="scope">{{ formatDate(scope.row.purchase_date) }}</template>
                </el-table-column>
                <el-table-column align="left" label="申请单ID" prop="ID" width="120" />
                <el-table-column align="left" label="耗材名称" prop="goods_name" />
                <el-table-column align="left" label="规格" prop="specification" />
                <el-table-column align="left" label="申请数量" prop="apply_number" />
                <el-table-column align="left" label="已到数量" prop="arrival_number" />
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
                <el-table-column align="left" fixed="right" label="选择入库耗材" width="120">
                    <template #default="scope">
                        <el-button type="primary" @click="addFunc(scope.row)">添加</el-button>
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
            <span>入库单</span>
        </el-divider>
        <!-- 入库详情 -->
        <div>
            <el-container>
                <el-aside width="400px">
                    <el-form :model="ingressFormData" label-width="80px">
                        <el-form-item label="收货仓库">
                            <el-input v-model="ingressFormData.group_name" disabled />
                        </el-form-item>
                        <el-form-item label="送货单号">
                            <el-input v-model="ingressFormData.delivery_number"></el-input>
                        </el-form-item>
                        <el-form-item label="送货日期">
                            <el-date-picker v-model="ingressFormData.deliveryDate" placeholder="请输入日期" type="date"
                                style="width: 100%" />
                        </el-form-item>
                        <el-form-item label="入库日期">
                            <el-date-picker v-model="ingressFormData.ingress_date" type="date" style="width: 100%" />
                        </el-form-item>
                        <el-form-item label="入库按钮">
                            <el-button type="primary" @click="ingressFunc">入库</el-button>
                        </el-form-item>
                    </el-form>
                </el-aside>
                <el-main>
                    <div class="gva-table-box">
                        <el-table :data="ingressFormData.ingress_details">
                            <el-table-column align="left" label="耗材名称" prop="goods_name" width="120" />
                            <el-table-column align="left" label="规格" prop="specification" />
                            <el-table-column align="left" label="批号" prop="batch">
                                <template #default="scope">
                                    <el-form-item>
                                        <el-input ref="input" size='mini' v-model="scope.row.batch" placeholder="请输入批号">
                                        </el-input>
                                    </el-form-item>
                                </template>
                            </el-table-column>
                            <el-table-column align="left" label="有效期" prop="expiration_date">
                                <template #default="scope">
                                    <el-form-item>
                                        <el-date-picker :clearable="false" v-model="scope.row.expiration_date"
                                            type="date" placeholder="请输入有效期">
                                        </el-date-picker>
                                    </el-form-item>
                                </template>
                            </el-table-column>
                            <el-table-column align="left" label="入库数量" prop="ingress_number" >
                                <template #default="scope">
                                    <el-form-item>
                                        <el-input type="number" size='mini' v-model.number="scope.row.ingress_number" placeholder="请输入入库数量"/>
                                    </el-form-item>
                                </template>
                            </el-table-column>
                        </el-table>
                    </div>
                </el-main>
            </el-container>
        </div>
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

const ingressFormData = ref({
    order_id: 0,
    group_name: "",
    delivery_number: "",
    ingress_date: new Date(),
    ingress_details: []
})

// 添加耗材以完成入库操作
const addFunc = async (row) => {
    if (ingressFormData.value.order_id === 0) {
        ingressFormData.value.order_id = row.ID
    }
    if (ingressFormData.value.order_id !== row.ID) {
        ElMessage({
            type: 'fail',
            message: '不允许添加属于不同申请单的耗材'
        })
        return
    }
    ingressFormData.value.group_name = row.group_name
    const goods = {
        goods_name: row.goods_name,
        specification: row.specification,
        batch: "",
        expiration_date: "",
        ingress_number: 0,
    }
    ingressFormData.value.ingress_details.push(goods)
    ElMessage({
        type: 'success',
        message: '添加成功'
    })
}
// 入库
const ingressFunc = async () => {
    const rst = await ingress(ingressFormData.value)
    if (table.code === 0) {
        ElMessage({
            type: 'success',
            message: '入库成功'
        })
    }else{
        ElMessage({
            type: 'fail',
            message: '入库失败'
        })
    }
}

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}
</script>
  
<style>

</style>