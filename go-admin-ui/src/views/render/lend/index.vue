<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form
          ref="queryForm"
          :model="queryParams"
          :inline="true"
          label-width="68px"
        >
          <!-- <el-form-item
            label="读者"
            prop="renderId"
          ><el-select
            v-model="queryParams.renderId"
            filterable
            placeholder="请选择"
            clearable
            size="small"
          >
            <el-option
              v-for="dict in renderIdOptions"
              :key="dict.key"
              :label="dict.value"
              :value="dict.key"
            />
          </el-select>
          </el-form-item> -->
          <el-form-item
            label="图书"
            prop="bookId"
          ><el-select
            v-model="queryParams.bookId"
            filterable
            placeholder="请选择"
            clearable
            size="small"
          >
            <el-option
              v-for="dict in bookIdOptions"
              :key="dict.key"
              :label="dict.value"
              :value="dict.key"
            />
          </el-select>
          </el-form-item>
          <el-form-item label="状态" prop="state">
            <el-select
              v-model="queryParams.state"
              placeholder="借阅状态"
              clearable
              size="small"
            >
              <el-option label="未归还" value="0" />
              <el-option label="已归还" value="1" />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              icon="el-icon-search"
              size="mini"
              @click="handleQuery"
            >搜索</el-button>
            <el-button
              icon="el-icon-refresh"
              size="mini"
              @click="resetQuery"
            >重置</el-button>
            <el-button
              v-permisaction="['admin:lend:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
            >新增
            </el-button>
          </el-form-item>
        </el-form>

        <el-table
          v-loading="loading"
          :data="lendList"
          @selection-change="handleSelectionChange"
        >
          <el-table-column
            type="selection"
            width="55"
            align="center"
          /><el-table-column
            label="读者"
            align="center"
            prop="renderId"
            :formatter="renderIdFormat"
            width="100"
          >
            <template slot-scope="scope">
              {{ renderIdFormat(scope.row) }}
            </template> </el-table-column><el-table-column
            label="图书"
            align="center"
            prop="bookId"
            :formatter="bookIdFormat"
          >
            <template slot-scope="scope">
              {{ bookIdFormat(scope.row) }}
            </template> </el-table-column><el-table-column
            label="借阅日期"
            align="center"
            prop="lendDate"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="最后归还日期"
            align="center"
            prop="backDate"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="状态"
            align="center"
            prop="state"
            :show-overflow-tooltip="true"
          >
            <template slot-scope="scope">
              {{ scope.row.state == '0' ? '未归还' : '已归还' }}
            </template>
          </el-table-column>
          <el-table-column
            label="是否逾期"
            align="center"
            prop="state"
            :show-overflow-tooltip="true"
          >
            <template slot-scope="scope">
              {{
                scope.row.state == '0' &&
                  new Date().getTime > new Date(scope.row.backDate).getTime
                  ? '已逾期'
                  : '未逾期'
              }}
            </template>
          </el-table-column>
          <el-table-column
            label="操作"
            align="center"
            class-name="small-padding fixed-width"
          >
            <template slot-scope="scope">
              <el-popconfirm
                :disabled="scope.row.state == '1'"
                class="delete-popconfirm"
                title="确认要归还吗?"
                confirm-button-text="归还"
                @confirm="changeState(scope.row.id)"
              >
                <el-button
                  slot="reference"
                  v-permisaction="['admin:lend:edit']"
                  :disabled="scope.row.state == '1'"
                  type="success"
                >{{ scope.row.state == '1' ? '已归还' : '确认归还' }}
                </el-button>
              </el-popconfirm>

              <!-- <span
                v-if="scope.row.state == '1'"
                style="color: green"
              >已归还</span>
              <span v-else style="color: red">未归还</span> -->

              <!-- <el-popconfirm
                class="delete-popconfirm"
                title="确认要删除吗?"
                confirm-button-text="删除"
                @onConfirm="handleDelete(scope.row)"
              >
                <el-button
                  slot="reference"
                  v-permisaction="['admin:lend:remove']"
                  size="mini"

                  icon="el-icon-delete"
                >删除
                </el-button>
              </el-popconfirm> -->
            </template>
          </el-table-column>
        </el-table>

        <pagination
          v-show="total > 0"
          :total="total"
          :page.sync="queryParams.pageIndex"
          :limit.sync="queryParams.pageSize"
          @pagination="getList"
        />

        <!-- 添加或修改对话框 -->
        <el-dialog :title="title" :visible.sync="open" width="500px">
          <el-form ref="form" :model="form" :rules="rules" label-width="120px">
            <el-form-item label="读者" prop="renderId">
              <el-select
                v-model="form.renderId"
                filterable
                placeholder="请选择"
              >
                <el-option
                  v-for="dict in renderIdOptions"
                  :key="dict.key"
                  :label="dict.value"
                  :value="dict.key"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="图书" prop="bookId">
              <el-select v-model="form.bookId" filterable placeholder="请选择">
                <el-option
                  v-for="dict in bookIdOptions"
                  :key="dict.key"
                  :label="dict.value"
                  :value="dict.key"
                />
              </el-select>
            </el-form-item>
            <!-- <el-form-item label="借阅日期" prop="lendDate">
              <el-date-picker
                v-model="form.lendDate"
                type="date"
                value-format="yyyy-MM-dd HH:mm:ss"
                placeholder="借阅日期"
              />
            </el-form-item>
            <el-form-item label="最后归还日期" prop="backDate">
              <el-date-picker
                v-model="form.backDate"
                type="date"
                value-format="yyyy-MM-dd HH:mm:ss"
                placeholder="最后归还日期"
              />
            </el-form-item> -->
            <!-- <el-form-item label="状态" prop="state">
              <el-select v-model="form.state" placeholder="请选择">
                <el-option
                  v-for="dict in stateOptions"
                  :key="dict.value"
                  :label="dict.label"
                  :value="dict.value"
                />
              </el-select>
            </el-form-item> -->
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import {
  addLend,
  delLend,
  getLend,
  listLend,
  updateLend
} from '@/api/admin/lend'
import moment from 'moment'

import { listRender } from '@/api/admin/render'
import { listBook } from '@/api/admin/book'
export default {
  name: 'Lend',
  components: {},
  data() {
    return {
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      isEdit: false,
      // 类型数据字典
      typeOptions: [],
      lendList: [],

      // 关系表类型
      renderIdOptions: [],
      bookIdOptions: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        renderId: undefined,
        bookId: undefined,
        state: undefined
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        renderId: [
          { required: true, message: '读者不能为空', trigger: 'blur' }
        ],
        bookId: [{ required: true, message: '图书不能为空', trigger: 'blur' }]
        // lendDate: [
        //   { required: true, message: '借阅日期不能为空', trigger: 'blur' }
        // ],
        // backDate: [
        //   { required: true, message: '归还日期不能为空', trigger: 'blur' }
        // ],
        // state: [{ required: true, message: '状态不能为空', trigger: 'blur' }]
      }
    }
  },
  created() {
    this.getList()
    this.getRenderItems()
    this.getBookItems()
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listLend(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.lendList = response.data.list
          this.total = response.data.count
          this.loading = false
        }
      )
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.form = {
        id: undefined,
        renderId: undefined,
        bookId: undefined,
        lendDate: undefined,
        backDate: undefined
      }
      this.resetForm('form')
    },
    getImgList: function() {
      this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
    },
    fileClose: function() {
      this.fileOpen = false
    },
    renderIdFormat(row) {
      return this.selectItemsLabel(this.renderIdOptions, row.renderId)
    },
    bookIdFormat(row) {
      return this.selectItemsLabel(this.bookIdOptions, row.bookId)
    },
    // 关系
    getRenderItems() {
      this.getItems(listRender, undefined).then((res) => {
        this.renderIdOptions = this.setItems(res, 'id', 'name')
      })
    },
    getBookItems() {
      this.getItems(listBook, undefined).then((res) => {
        this.bookIdOptions = this.setItems(res, 'id', 'name')
      })
    },
    // 文件
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = []
      this.resetForm('queryForm')
      this.handleQuery()
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset()
      this.open = true
      this.title = '添加借阅记录'
      this.isEdit = false
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map((item) => item.id)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      const id = row.id || this.ids
      getLend(id).then((response) => {
        this.form = response.data
        this.open = true
        this.title = '修改借阅记录'
        this.isEdit = true
      })
    },
    changeState(id) {
      updateLend({ id, state: 1 }).then((response) => {
        if (response.code === 200) {
          this.msgSuccess('归还成功')
          this.open = false
          this.getList()
        } else {
          this.msgError(response.msg)
        }
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate((valid) => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateLend(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addLend({
              ...this.form,
              state: 0,
              lendDate: moment().format('yyyy-MM-DD HH:mm:ss'),
              backDate: moment().add(20, 'day').format('yyyy-MM-DD HH:mm:ss')
            }).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          }
        }
      })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      var Ids = (row.id && [row.id]) || this.ids

      this.$confirm('是否确认删除编号为"' + Ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(function() {
          return delLend({ ids: Ids })
        })
        .then((response) => {
          if (response.code === 200) {
            this.msgSuccess(response.msg)
            this.open = false
            this.getList()
          } else {
            this.msgError(response.msg)
          }
        })
        .catch(function() {})
    }
  }
}
</script>
