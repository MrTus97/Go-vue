<template>
	<div class="user">
		<p>{{user.name}}</p>
		<el-button type="danger" icon="el-icon-delete" circle v-on:click="handleDelete"></el-button>
		<!-- <button
			
			>Delete</button> -->
	</div>
</template>

<script>

export default {
	name: 'user',
	props: {
		user: {
			type: Object,
			default: null
		}
	},
	data() {
		return {
			
		}
	},
	methods: {
		handleDelete(){
			// Call to listUser
			var data ={
				id: this.user.id
			}
			this.$emit('deleteUser', data)
		}
	}
}
</script>

<style>
.user {
	border: 1px solid;
	margin-top: 5px;
}
</style>