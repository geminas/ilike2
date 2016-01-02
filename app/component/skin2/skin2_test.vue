<script>

module.exports={
	el:"#show",
	data:{
	target:"skin2"
	},
	ready:function(){
	console.log("The skin2 test main is loaded")
	},
	components:{
	'skin2vue':require('./skin2.vue')
	}
}
</script>

<template>
<div><h1>Test--{{target}} is below</h1></div>
<component is="skin2vue" />
</template>