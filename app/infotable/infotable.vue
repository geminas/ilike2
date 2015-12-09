<script>

module.exports={
	el:"#show",
	data:{
	target:"infotable",
	d:window.data
	},
	ready:function(){
	console.log("The infotable test main is loaded")
	},
	components:{
	'infotablevue':require('../component/infotable/infotable.vue')
	}
}
</script>

<template>
<div>
<component is="infotablevue" :data="d"/>
</div>
</template>