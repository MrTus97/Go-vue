<template>
	<div class="compo-tabs">
		<div class="tabs-title">
			<span>{{title}}</span>
			<ion-icon name="heart"></ion-icon>
		</div>
		<div class="tabs-content">
			<slot></slot>
		</div>
	</div>
</template>

<script>
export default {
	name: 'comp-tabs',
	data(){
		return{

		}
	},
	props:{
		title: String,
	},
	methods: {
		

	}
}
</script>

<style>
	.tabs-title{
		padding: 15px 30px;
		display: flex;
		justify-content: space-between;
		border: 1px solid #000;
	}
	.tabs-content{
		margin-top: 2px;
		padding: 2px;
		border: 2px solid blue;
	}
	.tabs-header{
		list-style: none;
		margin: 0px;
		padding: 0px;
		display: flex;
	}

	.tabs-header li {
		padding: 8px 15px;
		background-color: #ccc;
	}
	.tabs-header li:hover{
		color: #fff;
		background-color: #333;
	}
	.tab-content {
		padding: 15px;
		border: 2px solid;
		text-align: left;
	}
</style>