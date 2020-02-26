<template>
  <div id="app">
    <!-- Header -->
    <comp-header
      v-bind:Keymsg="msg"
      v-on:changeTitleEvent="handleChangeTitle"
    />

    <!-- Body -->
  </div>
</template>

<script>
import CompHeader from './components/CompHeader'
export default {
  name: 'app',
  components: {
    CompHeader
  },
  method: {
    changeTitleEvent () {
      this.msg = 'hello 123'
    }
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

</style>
