<template>
    <h1>{{ Keymsg }}
        <button v-on:click="changeTitleheader">Thay đổi tiêu đề </button>
    </h1>
</template>

<script>
export default {
  name: 'comp-header',
  props: {
    Keymsg: {
      type: String,
      default: 'giá trị mặc định'
    }
  },
  data () {
    return {
      textHeader: 'Hello Header'
    }
  },
  methods: {
    changeTitleheader (e) {
      this.$emit('changeTitleEvent')
    }
  }
}
</script>

<style>

</style>
