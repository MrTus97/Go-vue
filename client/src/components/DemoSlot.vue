<template>
	<div class="demo-slot">
		<comp-tabs title="Tiêu đề của tabs">
			<ul class="tabs-header">
				<li 
				v-for="(item,index) in tabsHeader" 
				v-bind:key="index"
				v-on:click="setActiveTab(index)"
				v-bind:class="{active: tabActive === index}"
				>{{item.text}}</li>
			</ul>
			<div class="tab-content tab1" v-if="tabActive === 0">
				<p>Tab1 orem ipsum dolor sit, amet consectetur adipisicing elit. Ab mollitia aliquid dignissimos excepturi labore in tempora. Sapiente corporis laborum, odit doloremque molestias non, laudantium earum eaque iure provident aspernatur voluptatum.</p>
			</div>
			<div class="tab-content tab2" v-if="tabActive === 1">
				<p> tab2 Lorem ipsum dolor sit amet, consectetur adipisicing elit. Quos minima natus illum accusamus itaque odit assumenda ab hic voluptates adipisci, molestiae ipsam quaerat tempore porro modi, fugit delectus quae possimus.</p>
			</div>
			<div class="tab-content tab3" v-if="tabActive === 2">
				<p> tab3 Lorem ipsum dolor sit amet consectetur, adipisicing elit. Labore necessitatibus aut quas vel, veritatis quaerat, ex cumque, itaque iste magni delectus et eius voluptatum repellendus magnam! Magnam et veritatis labore!</p>
			</div>
		</comp-tabs>
	</div>
</template>

<script>
import CompTabs from './CompTabs.vue'
export default {
	name: 'demo-slot',
	components: {
		CompTabs,
	},
	data(){
		return {
			tabsHeader:[
				{text: 'Tab1'},
				{text: 'Tab2'},
				{text: 'Tab3'},
			],
			tabActive: 0, 
		}
	},
	methods: {
		setActiveTab(index){
			this.tabActive = index;
		}
	}
}
</script>

<style>
	.demo-slot{
		margin-top: 100px;
		padding: 100px 30px;
		border: 1px solid #000;
	}
</style>