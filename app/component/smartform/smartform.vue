<script type="text/javascript">
	module.exports={
		data:function(){
			return {
				//hello:"world"
				scheme:{}
			}
		},
		props:["scheme"],
		/////////////Life Span/////////////////////
		
		// beforeCompile:function(){
		// 	console.log("smartform beforeCompiled");
		// },
		// compiled:function(){
		// 	console.log("smartform has been compiled");
		// },
		created:function(){
		console.log("app has been created");
		var self=this;
		// $.ajax({
		// 	type:'GET',
		// 	url:"/getjson",
		// 	cache:false,
		// 	//data:JSON.stringify(data),
		// 	contentType:"application/json",
		// 	processData: false,
  //           success:function(data){
  //               //console.log(data)
  //             	//alert(data)
  //             	//var d=JSON.parse(data)
  //             	//console.log(d)
  //             	//self.$data.scheme=d
  //           },
  //           error:function(data){
  //               console.log("error");
  //               //console.log(data);
  //              alert("error: "+data)
  //           }
		// });	
	},
		ready:function(){
			//console.log("smartform has been ready");
			//console.log(this.scheme)
		},
		// attached:function(){
		// 	console.log("smartform has been attached");
		// },
		// detached:function(){
		// 	console.log("smartform has been detached");
		// },
		// beforeDestory:function(){
		// 	console.log("before smartform destoryed");
		// },
		// destoryed:function(){
		// 	console.log("smartform has been destoryed");
		// },
		methods:{
			// sayHello:function(){
			// 	console.log("Hello,This is the smartform component");
			// }
			sendurl:function(){
				var data={hello:"world"}
				$.ajax({
					type:'POST',
					url:"testsmartform",
					cache:false,
					data:JSON.stringify(data),
					contentType:"application/json",
					processData: false,
		            success:function(data){
		                console.log(data)
		              	alert(data)
		            },
		            error:function(data){
		                console.log("error");
		                console.log(data);
		               alert("error: "+data)
		            }
				});			
			}
		},
		computed:{
			// foo:function(){
			// 	return "bar"
			// }
		},
		events:{
			//////////Events 
			// 'hook:created':function(){
			// 	console.log('smartform has created again');
			// },
			// greeting:function(msg){
			// 	console.log(msg);
			// },
			// hello:'sayHello'
			///////////Can be Triggered by
			//////vm.$emit('hello')
		},
		watch:{
			//////////Data Watcher
			// 'a':function(val,oldval){
			// 	console.log('new: %s, old: %s', val, oldVal);
			// }
		}

	}
</script>
<template>


<div v-for="item in scheme" class="row">
	<div v-if="item.field_type=='text'" class="col-xs-12 ">
	<label>{{item.label}}</label>
	<input type="text" class="form-control" name="{{item.name}}" placeholder="{{item.placeholder}}" id="name" required data-validation-required-message="Please enter your name.">
    <p class="help-block text-danger"></p>
	</div>

	<div v-if="item.field_type=='select'" class="col-xs-12 ">
		<label>{{item.label}}</label>
		<select name="{{item.name}}">
			<option v-for="op in item.field_options.options">
				{{op.label}}
			</option>
		</select>
	</div>
	<div v-if="item.field_type=='radio'" class="col-xs-12 ">
		<label>{{item.label}}</label>
		<div v-for="op in item.field_options.options">
			<input type="radio" name="{{item.name}}">
			<span>{{op.label}}</span>
		</div>
	</div>
	<div v-if="item.field_type=='checkbox'" class="col-xs-12 ">
		<label>{{item.label}}</label>
		<div v-for="op in item.field_options.options">
			<input type="checkbox" name="{{item.name}}">
			<span>{{op.label}}</span>
		</div>
	</div>
</div>
				

<!-- <div>This is component smartform</div>
 -->
 <!-- <button v-on:click="sendurl()">ClickMe</button> -->
 </template>
<style type="text/css"></style>