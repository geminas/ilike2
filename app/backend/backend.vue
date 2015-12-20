<script>

module.exports={
	el:"#show",
	data:{
		schemes:[],
		taskMap:{},
		bindform:{},
		bindid:""
	},
	created:function(){
		console.log("backend has been created");
		this.sync()
	},
	methods:{
		sync:function(){
			var self=this
			$.ajax({
				type:'GET',
				url:"/gettasks",
				cache:false,
				//data:JSON.stringify(data),
				contentType:"application/json",
				processData: false,
	            success:function(data){
	            	console.log(data)
	            	self.schemes=[]
	            	for(key in data){
	            		self.schemes.push({
	            			id:key,
	            			scheme:JSON.parse(data[key])
	            		})
	            	}
	            	self.taskMap=data
	            	self.bindform=self.schemes[0].scheme
	            	self.bindid=self.schemes[0].id
	              	//console.log(self.schemes)
	            },
	            error:function(data){
	                console.log("error");
	                //console.log(data);
	               alert("error: "+data)
	            }
			});	
		},
		choosetask:function(index){
			this.bindform=this.schemes[index].scheme
			this.bindid=this.schemes[index].id
		},
		save:function(){
			//var item=JSON.parse(this.taskMap[this.bindid])
			var j=this.$refs.formbuilder.getCollection()
			//console.log()
			//console.log(item)
			this.bindform.fields=j
			var id=this.bindid
			var self=this
			$.ajax({
			type:'POST',
			url:"/posttask?id="+id,
			cache:false,
			data:JSON.stringify(self.bindform),
			contentType:"application/json",
			processData: false,
            success:function(data){
                console.log(data)
              	//alert(data)
              	//var d=JSON.parse(data)
              	//console.log(d)
              	//self.$data.scheme=d
              	self.sync()
            },
            error:function(data){
                console.log("error");
                //console.log(data);
               alert("error: "+data)
            }
		});
		}
	},
	ready:function(){
	console.log("The smartform test main is loaded")
	},
	components:{
	'smartformvue':require('../component/smartform/smartform.vue'),
	'formbuildervue':require('../component/formbuilder/formbuilder.vue')
	}
}
</script>

<template>
<header></header>

<div class="container">
	<div class="row">
		<table class="table table-hover">
			<thead>
				<th>活动ID</th>
				<th>活动名</th>
				<th>活动描述</th>
			</thead>
			<tbody>
				<tr v-for="t in schemes" @click="choosetask($index)">
					<th>{{t.id}}</th>
					<th>{{t.scheme.name}}</th>
					<th>{{t.scheme.describe}}</th>
				</tr>
			</tbody>
		</table>
	</div>
	<div class="row">
		<button @click="save()">保存设置</button>
	</div>

	<div class="row">
		<!-- <div>
			<component is="smartformvue" :scheme="bindform.items" v-ref="smarttable"/>
		</div>
 -->		
 		<div class="row text-center">
 			<label for="">活动名</label><input type="text" v-model="bindform.name">
 			<label for="">活动描述</label><input type="text" v-model="bindform.describe">

 		</div>
 		
 		<div class="row">
 			<component is="formbuildervue" v-ref:formbuilder :scheme="bindform.fields"/>
 		</div>
 		
	</div>
</div>


<!-- <div>
	<ul>
		<li v-for="(k,v) in schemes">
			{{v.id}}
		</li>
	</ul>
</div> -->

</template>

<style>
	
</style>