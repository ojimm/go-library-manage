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
          <el-form-item
            label="图书名称"
            prop="name"
          ><el-input
            v-model="queryParams.name"
            placeholder="请输入图书名称"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
          />
          </el-form-item>
          <el-form-item
            label="书籍类别"
            prop="cid"
          ><el-select
            v-model="queryParams.cid"
            placeholder="请选择"
            clearable
            size="small"
          >
            <el-option
              v-for="dict in cidOptions"
              :key="dict.key"
              :label="dict.value"
              :value="dict.key"
            />
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
              v-permisaction="['admin:book:add']"
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
          :data="bookList"
          @selection-change="handleSelectionChange"
        >
          <el-table-column
            type="selection"
            width="55"
            align="center"
          /><el-table-column
            label="图书名称"
            align="center"
            prop="name"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="书籍类别"
            align="center"
            prop="cid"
            :formatter="cidFormat"
            width="100"
          >
            <template slot-scope="scope">
              {{ cidFormat(scope.row) }}
            </template>
          </el-table-column>
          <el-table-column
            label="作者"
            align="center"
            prop="author"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="出版社"
            align="center"
            prop="publish"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="ISBN"
            align="center"
            prop="isbn"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="简介"
            align="center"
            prop="intro"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="语言"
            align="center"
            prop="language"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="价格"
            align="center"
            prop="price"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="出版日期"
            align="center"
            prop="pubdate"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="库存"
            align="center"
            prop="count"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="操作"
            align="center"
            class-name="small-padding fixed-width"
          >
            <template slot-scope="scope">
              <el-popconfirm
                class="delete-popconfirm"
                title="确认要修改吗?"
                confirm-button-text="修改"
                @confirm="handleUpdate(scope.row)"
              >
                <el-button
                  slot="reference"
                  v-permisaction="['admin:book:edit']"
                  size="mini"
                  type="text"
                  icon="el-icon-edit"
                >修改
                </el-button>
              </el-popconfirm>
              <el-popconfirm
                class="delete-popconfirm"
                title="确认要删除吗?"
                confirm-button-text="删除"
                @confirm="handleDelete(scope.row)"
              >
                <el-button
                  slot="reference"
                  v-permisaction="['admin:book:remove']"
                  size="mini"
                  type="text"
                  icon="el-icon-delete"
                >删除
                </el-button>
              </el-popconfirm>
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
          <el-form ref="form" :model="form" :rules="rules" label-width="80px">
            <el-form-item label="图书名称" prop="name">
              <el-input v-model="form.name" placeholder="图书名称" />
            </el-form-item>
            <el-form-item label="书籍类别" prop="cid">
              <el-select v-model="form.cid" placeholder="请选择">
                <el-option
                  v-for="dict in cidOptions"
                  :key="dict.key"
                  :label="dict.value"
                  :value="dict.key"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="作者" prop="author">
              <el-input v-model="form.author" placeholder="作者" />
            </el-form-item>
            <el-form-item label="出版社" prop="publish">
              <el-input v-model="form.publish" placeholder="出版社" />
            </el-form-item>
            <el-form-item label="ISBN" prop="isbn">
              <el-input v-model="form.isbn" placeholder="ISBN" />
            </el-form-item>
            <el-form-item label="语言" prop="language">
              <el-select v-model="form.language" placeholder="语言">
                <el-option label="中文" value="中文" />
                <el-option label="英文" value="英文" />
              </el-select>
            </el-form-item>
            <el-form-item label="价格" prop="price">
              <el-input v-model="form.price" placeholder="价格" />
            </el-form-item>
            <el-form-item label="出版日期" prop="pubdate">
              <el-input v-model="form.pubdate" placeholder="出版日期" />
            </el-form-item>
            <el-form-item label="库存" prop="count">
              <el-input-number v-model="form.count" :min="1" label="库存" />
            </el-form-item>
            <el-form-item label="简介" prop="intro">
              <el-input
                v-model="form.intro"
                type="textarea"
                placeholder="简介"
              />
            </el-form-item>
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
  addBook,
  delBook,
  getBook,
  listBook,
  updateBook
} from '@/api/admin/book'

import { listBookClass } from '@/api/admin/book-class'
export default {
  name: 'Book',
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
      bookList: [],

      // 关系表类型
      cidOptions: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        name: undefined,
        cid: undefined
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        name: [
          { required: true, message: '图书名称不能为空', trigger: 'blur' }
        ],
        cid: [
          { required: true, message: '书籍类别id不能为空', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getList()
    this.getBookClassItems()
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listBook(this.addDateRange(this.queryParams, this.dateRange)).then(
        (response) => {
          this.bookList = response.data.list
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
        name: undefined,
        author: undefined,
        publish: undefined,
        isbn: undefined,
        intro: undefined,
        language: undefined,
        price: undefined,
        pubdate: undefined,
        count: 1,
        cid: undefined
      }
      this.resetForm('form')
    },
    getImgList: function() {
      this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
    },
    fileClose: function() {
      this.fileOpen = false
    },
    cidFormat(row) {
      return this.selectItemsLabel(this.cidOptions, row.cid)
    },
    // 关系
    getBookClassItems() {
      this.getItems(listBookClass, undefined).then((res) => {
        this.cidOptions = this.setItems(res, 'id', 'name')
        console.log(this.cidOptions)
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
      this.title = '添加图书'
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
      getBook(id).then((response) => {
        this.form = response.data
        this.open = true
        this.title = '修改图书'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate((valid) => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateBook(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addBook(this.form).then((response) => {
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
          return delBook({ ids: Ids })
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
