<template>
    <div>
        <div class="gva-table-box">
            <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
                @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55" />
                <el-table-column align="left" label="申请单ID" prop="ID" width="120" />
                <el-table-column align="left" label="申请单提交日期" width="180">
                    <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
                </el-table-column>
                <el-table-column align="left" label="耗材名称"/>
                <el-table-column align="left" label="规格"/>
                <el-table-column align="left" label="申请数量"/>
                <el-table-column align="left" label="已到数量"/>
                <el-table-column align="left" label="库存点" prop="name" width="140" />
                <el-table-column align="left" label="申请人" prop="applicant.nickName" width="120" />
                <el-table-column align="left" label="申请审批日期" width="180">
                    <template #default="scope">{{ formatDate(scope.row.UpdatedAt) }}</template>
                </el-table-column>
                <el-table-column align="left" label="审批人" prop="approver.nickName" width="120" />
                <el-table-column align="left" label="申请状态" prop="state" width="120">
                    <template #default="scope">{{ filterDict(scope.row.state, applyStateOptions) }}</template>
                </el-table-column>
            </el-table>
            <div class="gva-pagination">
                <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page"
                    :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
                    @current-change="handleCurrentChange" @size-change="handleSizeChange" />
            </div>
        </div>
        <!-- 入库详情 -->
        <div class="layout-container-demo">
            <el-container>
                <el-aside width="200px">
                    <el-input v-model="input" placeholder="Please input" disabled />
                </el-aside>
                <el-main>Main</el-main>
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
    getBusGroupList
} from '@/api/busGroup'


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
    const table = await getBusGroupList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
    }
}

getTableData()

// ============== 表格控制部分结束 ===============


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}
</script>
  
<style>

</style>