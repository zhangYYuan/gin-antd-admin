<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="名称">
          <el-input placeholder="名称" v-model="searchInfo.name"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog('addTreatType')" type="primary">新增请客分类</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="id" min-width="60" prop="ID"></el-table-column>
      <el-table-column label="名称" min-width="150" prop="name"></el-table-column>
      <el-table-column label="等级" min-width="150" prop="level"></el-table-column>
      <el-table-column label="简介" min-width="150" prop="describe"></el-table-column>
      <el-table-column label="区域" min-width="150" prop="region"></el-table-column>

      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editTreatType(scope.row)" size="small" type="text">编辑</el-button>
          <el-button @click="deleteTreatType(scope.row)" size="small" type="text">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :style="{float:'right',padding:'20px'}"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增TreatType">
      <el-form :inline="true" :model="form" label-width="80px">
        <el-form-item label="名称">
          <el-input autocomplete="off" v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="等级">
        <el-select placeholder="请选择" v-model="form.level">
          <el-option
                  :key="item.value"
                  :label="`${item.label}(${item.value})`"
                  :value="item.value"
                  v-for="item in levelOptions"
          ></el-option>
        </el-select>
      </el-form-item>

        <el-form-item label="区域">
          <el-select placeholder="请选择" v-model="form.region">
            <el-option
                    :key="item.value"
                    :label="`${item.label}(${item.value})`"
                    :value="item.value"
                    v-for="item in regionOptions"
            ></el-option>
          </el-select>
        </el-form-item>


        <el-form-item label="分类简介">
          <el-input autocomplete="off" v-model="form.describe"></el-input>
        </el-form-item>
      </el-form>
      <div class="warning">
          新增请客分类之后就可以在请客单中使用
      </div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

import {
  getById,
  getList,
  createTreatType,
  updataTreatType,
  deleteTreatType
} from '@/api/treattype'
import infoList from '@/components/mixins/infoList'

const levelOptions = [
  {
    value: 1,
    label: '一阶'
  },
  {
    value: 3,
    label: '三阶'
  }
]

const regionOptions = [
  {
    value: 1,
    label: '国内'
  },
  {
    value: 2,
    label: '海外'
  }
]

export default {
  name: 'Api',
  mixins: [infoList],
  data() {
    return {
      listApi: getList,
      listKey: 'list',
      dialogFormVisible: false,
      form: {
        name: '',
        level: undefined,
        region: undefined,
        describe: ''
      },
      levelOptions: levelOptions,
      regionOptions: regionOptions,
      type: ''
    }
  },
  methods: {
    //条件搜索前端看此方法
      onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    initForm() {
      this.form = {
        name: '',
        level: undefined,
        region: undefined,
        describe: ''
      }
    },
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    openDialog(type) {
      this.type = type
      this.dialogFormVisible = true
    },
    async editTreatType(row) {
      const res = await getById({ id: row.ID })
      this.form = res.data.treatType
      this.openDialog('edit')
    },
    async deleteTreatType(row) {
      this.$confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          const res = await deleteTreatType(row)
          if (res.success) {
            this.$message({
              type: 'success',
              message: '删除成功!'
            })
            this.getTableData()
          }
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    },
    async enterDialog() {
      switch (this.type) {
        case 'addTreatType':
          {
            const res = await createTreatType(this.form)
            if (res.success) {
              this.$message({
                type: 'success',
                message: '添加成功',
                showClose: true
              })
            }
            this.getTableData()
            this.closeDialog()
          }

          break
        case 'edit':
          {
            const res = await updataTreatType(this.form)
            if (res.success) {
              this.$message({
                type: 'success',
                message: '添加成功',
                showClose: true
              })
            }
            this.getTableData()
            this.closeDialog()
          }
          break
        default:
          {
            this.$message({
              type: 'error',
              message: '未知操作',
              showClose: true
            })
          }
          break
      }
    }
  }
}
</script>
<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.warning {
    color: #DC143C;
}
</style>
