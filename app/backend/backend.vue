<script>

module.exports={
	el:"#show",
	data:{
		schemes:[],
		taskMap:{},
		bindform:{},
		binddatapath:"",
		bindid:""
	},
	created:function(){
		console.log("backend has been created");
		this.sync()
	},
	methods:{
		addtask:function(){
			var self=this
			$.ajax({
				type:'GET',
				url:"/newtask",
				cache:false,
				//data:JSON.stringify(data),
				//contentType:"application/json",
				processData: false,
	            success:function(data){
	            	// console.log(data)
	            	// self.schemes=[]
	            	// for(key in data){
	            	// 	self.schemes.push({
	            	// 		id:key,
	            	// 		scheme:JSON.parse(data[key])
	            	// 	})
	            	// }
	            	// self.taskMap=data
	            	// self.bindform=self.schemes[0].scheme
	            	// self.bindid=self.schemes[0].id
	              	//console.log(self.schemes)
	              	self.sync()
	            },
	            error:function(data){
	                console.log("error");
	                //console.log(data);
	               alert("error: "+data)
	            }
			});	
		},
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
			$('#myModal').modal()
		},
		checktask:function(index){
			// this.bindform=this.schemes[index].scheme
			// this.bindid=this.schemes[index].id
			if(this.schemes[index].id!=""){
			this.binddatapath="/viewdata/"+this.schemes[index].id
			}
			$('#myModal2').modal()
		},
		save:function(e){
			//var item=JSON.parse(this.taskMap[this.bindid])
			e.preventDefault()
			e.stopPropagation()
			//console.log("okoik")
			$('#myModal').modal('hide')
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




<!-- <div>
	<ul>
		<li v-for="(k,v) in schemes">
			{{v.id}}
		</li>
	</ul>
</div> -->
<ul id="myTab" class="nav nav-tabs">
   <li class="active">
      <a href="#home" data-toggle="tab">
         活动设置
      </a>
   </li>
   <li><a href="#ios" data-toggle="tab">查看活动数据</a></li>

</ul>
<div id="myTabContent" class="tab-content">
   <div class="tab-pane fade in active" id="home">
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
	<button type="button" class="btn btn-primary btn-lg" @click="addtask">
	  新增活动
	</button>
	<!-- Button trigger modal -->
	<!-- <button type="button" class="btn btn-primary btn-lg" data-toggle="modal" data-target="#myModal">
	  Launch demo modal
	</button> -->

	<!-- Modal -->
	<div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
	  <div class="modal-dialog" role="document">
	    <div class="modal-content">
	      <div class="modal-header">
	        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
	        <h4 class="modal-title" id="myModalLabel">修改活动设置</h4>
	      </div>
	      <div class="modal-body">
	      	<div class="row">
	      		<label for="">活动名</label><input type="text" v-model="bindform.name">
	      	</div>
	      	<div class="row">
	      		<label for="">活动描述</label><input type="text" v-model="bindform.describe">
	      	</div>
	        <div class="row">
 				<component is="formbuildervue" v-ref:formbuilder :scheme="bindform.fields"/>
 			</div>
	      </div>
	      <div class="modal-footer">
	        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
	        <button type="button" class="btn btn-primary" @click="save($event)">保存设置</button>
	      </div>
	    </div>
	  </div>
	</div>
</div>
   </div>
   <div class="tab-pane fade" id="ios">
      <div class="container">
	<div class="row">
		<table class="table table-hover">
			<thead>
				<th>活动ID</th>
				<th>活动名</th>
				<th>活动描述</th>
			</thead>
			<tbody>
				<tr v-for="t in schemes" @click="checktask($index)">
					<th>{{t.id}}</th>
					<th>{{t.scheme.name}}</th>
					<th>{{t.scheme.describe}}</th>
				</tr>
			</tbody>
		</table>
	</div>
	<!-- <button type="button" class="btn btn-primary btn-lg" @click="addtask">
	  新增活动
	</button>
 -->	<!-- Button trigger modal -->
	<!-- <button type="button" class="btn btn-primary btn-lg" data-toggle="modal" data-target="#myModal">
	  Launch demo modal
	</button> -->

	<!-- Modal -->
	<div class="modal fade" id="myModal2" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
	  <div class="modal-dialog" role="document">
	    <div class="modal-content">
	      <div class="modal-header">
	        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
	        <h4 class="modal-title" id="myModalLabel">查看活动数据</h4>
	      </div>
	      <div class="modal-body">
	      	<!-- <div class="row">
	      		<label for="">活动名</label><input type="text" v-model="bindform.name">
	      	</div>
	      	<div class="row">
	      		<label for="">活动描述</label><input type="text" v-model="bindform.describe">
	      	</div> -->
	        <div class="row">
 				<iframe src="{{binddatapath}}" frameborder="0"></iframe>
 			</div>
	      </div>
	      <!-- <div class="modal-footer">
	        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
	        <button type="button" class="btn btn-primary" @click="save($event)">保存设置</button>
	      </div> -->
	    </div>
	  </div>
	</div>
</div>
   </div>
  
</div>
</template>

<style>
	iframe{
		width: 100%;
	}
</style>