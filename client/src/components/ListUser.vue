<template>
    <div class="list-user">
        <user
        v-for="user in listUser"
        v-bind:key="user.id"
        v-bind:user="user"
        v-on:deleteUser="deleteUser"
        />
    </div>
</template>

<script>
export default {
  name: 'list-user',
  props: {
    listUser: {
      type: Array,
      default: []
    }
  },
  data () {
    return {

    }
  },
  components: {
    User
  },
  methods: {
    deleteUser (data) {
      this.$emit('deleteUserEvent', data)
    }
  }
}
</script>

<style>

</style>
