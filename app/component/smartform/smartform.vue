<script type="text/javascript">
	module.exports={
		data:function(){
			return {
				//hello:"world"
				scheme:{}
			}
		},
		props:["scheme"],
		created:function(){
		console.log("app has been created");
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
	<div v-if="item.field_type=='text'||item.field_type=='email'||item.field_type=='phone'||item.field_type=='idcard'" class="col-xs-12 {{item.status}}">
	<label class="label-{{item.name}}">{{item.label}}<span v-if="item.required==true" style="color:red;">*</span><span style="color:red;">{{item.error}}</span></label>
	<label class="label-text-error" for=""><span style="color:red;">{{item.error}}</span></label></label>
	<input type="text" class="form-control" name="{{item.name}}" placeholder="{{item.placeholder}}" id="{{item.name}}" required data-validation-required-message="Please enter your name." v-model="item.data">
    <p class="help-block text-danger"></p>
	</div>
	<div v-if="item.field_type=='textarea'" class="col-xs-12 {{item.status}}">
	<label class="label-{{item.name}}">{{item.label}}<span v-if="item.required==true" style="color:red;">*</span><span style="color:red;">{{item.error}}</span></label>
	<textarea class="form-control" name="{{item.name}}" placeholder="{{item.placeholder}}" id="{{item.name}}" required data-validation-required-message="Please enter your name." v-model="item.data"></textarea>
    <p class="help-block text-danger"></p>
	</div>
	<div v-if="item.field_type=='paragraph'" class="col-xs-12 {{item.status}}">
	<label class="label-{{item.name}}">{{item.label}}<span v-if="item.required==true" style="color:red;">*</span><span style="color:red;">{{item.error}}</span></label>
	<textarea class="form-control" name="{{item.name}}" placeholder="{{item.placeholder}}" id="{{item.name}}" required data-validation-required-message="Please enter your name." v-model="item.data"></textarea>
    <p class="help-block text-danger"></p>
	</div>
	<div v-if="item.field_type=='dropdown'" class="col-xs-12 ">
		<label class="label-source">{{item.label}}<span v-if="item.required==true" style="color:red;">*</span><span style="color:red;">{{item.error}}</span></label>
		<select name="{{item.name}}" id="{{item.name}}" v-model="item.data">
			<option v-for="op in item.field_options.options" selected>
				{{op.label}}
			</option>
		</select>
	</div>
	<div v-if="item.field_type=='select'" class="col-xs-12 ">
		<label class="label-{{item.name}}">{{item.label}}<span v-if="item.required==true" style="color:red;">*</span><span style="color:red;">{{item.error}}</span></label>
		<select name="{{item.name}}" id="{{item.name}}" v-model="item.data">
			<option v-for="op in item.field_options.options" selected>
				{{op.label}}
			</option>
		</select>
	</div>
	<div v-if="item.field_type=='radio-inline'" class="col-xs-12 ">
		<label class="label-{{item.name}}">{{item.label}}<span v-if="item.required==true" style="color:red;">*</span><span style="color:red;">{{item.error}}</span></label>
		
		<div v-for="op in item.field_options.options" class="radio-inline">
			<input type="radio" name="{{item.name}}" v-model="item.data" :value="op.label" >
			<span>{{op.label}}</span>
		</div>
	</div>

	<div v-if="item.field_type=='radio'" class="col-xs-12 " id="{{item.name}}">
		<label class="label-radio">{{item.label}}<span v-if="item.required==true" style="color:red;">*</span><span style="color:red;">{{item.error}}</span></label>
		<div v-for="op in item.field_options.options" class="">
			<input type="radio" name="{{item.name}}"  v-model="item.data" :value="op.label" checked>
			<span>{{op.label}}</span>
		</div>
	</div>
	<div v-if="item.field_type=='checkboxes'" class="col-xs-12 " id="{{item.name}}">
		<label class="label-checkboxes">{{item.label}}<span v-if="item.required==true" style="color:red;">*</span><span style="color:red;">{{item.error}}</span></label>
		<!-- <label for="">{{item.data}}</label> -->
		<div v-for="op in item.field_options.options">
			<input type="checkbox" name="{{item.name}}" v-model="item.datas" :true-value="op.label" value="{{op.label}}">
			<span>{{op.label}}</span>
		</div>
	</div>
	
	<div v-if="item.field_type=='checkbox'" class="col-xs-12 " id="{{item.name}}">
		<label class="label-{{item.name}}">{{item.label}}</label>
		<label for="">{{item.data}}</label>
		<div v-for="op in item.field_options.options">
			<input type="checkbox" name="{{item.name}}" v-model="item.data" :true-value="op.label">
			<span>{{op.label}}</span>
		</div>
	</div>
</div>
				

<!-- <div>This is component smartform</div>
 -->
 <!-- <button v-on:click="sendurl()">ClickMe</button> -->
 </template>
<style type="text/css"></style>