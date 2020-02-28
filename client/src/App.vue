<template>
  <div id="app">
    <el-container>
    <el-header>This header</el-header>
    <el-main>
        <el-table
            :data="posts"
            style="width: 100%">
              <el-table-column
                prop="id"
                label="id">
              </el-table-column>
              <el-table-column
                prop="Name"
                label="name">
              </el-table-column>
              <el-table-column
                fixed="left"
                label="action"
                width="200">
                <template slot-scope="scope">
                  <el-button @click="handleClick" type="text" size="small"><i class="el-icon-edit"></i></el-button>
                  <el-button type="text" size="small"><i class="el-icon-delete"></i></el-button>
                </template>
              </el-table-column>
          </el-table>
    </el-main>
    <el-footer>Copyright TuPT</el-footer>
  </el-container>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'app',
  data () {
    return {
      posts: [],
      errors: []
    }
  },

  // lấy dữ liệu khi component được tạo thành công
  mounted () {
    axios.get(`http://localhost:1323/users`).then(response => {
      console.log(response.data)
      this.posts = response.data
    })
      .catch(e => {
        this.errors.push(e)
      })
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.el-header, .el-footer {
    background-color: #B3C0D1;
    color: #333;
    text-align: center;
    line-height: 60px;
  }

.el-aside {
  background-color: #D3DCE6;
  color: #333;
  text-align: center;
  line-height: 200px;
}

.el-main {
  background-color: #E9EEF3;
  color: #333;
  text-align: center;
  line-height: 160px;
}

body > .el-container {
  margin-bottom: 40px;
}

.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  line-height: 260px;
}

.el-container:nth-child(7) .el-aside {
  line-height: 320px;
}

</style>
