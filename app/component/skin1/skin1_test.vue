<script>

module.exports={
	el:"#show",
	data:{
	target:"skin1"
	},
	ready:function(){
	console.log("The skin1 test main is loaded")
	},
	components:{
	'skin1vue':require('./skin1.vue')
	}
}
</script>

<template>
<div><h1>Test--{{target}} is below</h1></div>
<component is="skin1vue" />
</template>