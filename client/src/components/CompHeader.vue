<template>
		<h1>{{ Keymsg }}
		<button v-on:click="changeTitleheader">Thay doi title từ com-header </button>
		</h1>
</template>

<script>
export default {
	name: 'comp-header',
	props: {
		Keymsg : {
			type: String,
			default: "giá trị mặc định"
		}
	},
	data() {
		return {
			textHeader: 'Hello Header'
		}
	},
	methods: {
		changeTitleheader(e){
			console.log('before')
			this.$emit('changeTitleEvent')
			console.log('after')
		}
	}
}
</script>

<style>

</style>