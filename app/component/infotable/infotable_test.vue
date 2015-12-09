<script>

module.exports={
	el:"#show",
	data:{
	target:"infotable",
	d:[{
		a:1,
		b:2,
		c:3
	},{
		a:4,
		b:5,
		c:6
	}]
	},
	ready:function(){
	console.log("The infotable test main is loaded")
	},
	components:{
	'infotablevue':require('./infotable.vue')
	}
}
</script>

<template>
<div><h1>Test--{{target}} is below</h1></div>
<component is="infotablevue" :data="d"/>
</template>