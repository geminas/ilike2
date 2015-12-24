<script>

module.exports={
	el:"#show",
	data:{
	target:"formbuilder"
	},
	ready:function(){
	console.log("The formbuilder test main is loaded")
	},
	components:{
	'formbuildervue':require('./formbuilder.vue')
	}
}
</script>

<template>
<div><h1>Test--{{target}} is below</h1></div>
<component is="formbuildervue" />
</template>