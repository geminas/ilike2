<script>

module.exports={
	el:"#show",
	data:{
		schemes:[],
		taskMap:{},
		bindform:{},
		binddatapath:"",
		bindid:"",
		bg_img:""
	},
	created:function(){
		console.log("backend has been created");
		this.sync()
	},
	ready:function(){
		$('#uploader').on('change',function(e){
				console.log("changed")
				e.preventDefault();
				//self.$parent.$.mainframe.progress="正在处理中..."
				var fileupload=document.getElementById('uploader')
				//console.log(fileupload.files.length)
				self.uploadfile();
			})

	},
	methods:{
		uploadfile:function(e){
				e.preventDefault();
				var self=this;
				var fd=new FormData();
				var fileupload=document.getElementById('uploader')
				fd.append("uploadfile", fileupload.files[0]);

				$.ajax({
					type:'POST',
					url:"/upload",
					data:fd,
					cache:false,
					contentType:false,
					processData: false,
		            success:function(data){
		                //console.log(data)
		              	if(data.status==0){
		              		self.bindform.bgimg=data.msg;
		              	}
		            },
		            error: function(data){
		                console.log("error");
		                console.log(data);
		                
		            }
				});
				
			},
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
	            	 console.log(data)
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
	            	self.schemes.sort(function(a,b){return b.id-a.id})
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
		deletetask:function(index){
			this.schemes[index].scheme.deleted=true
			this.updatetask(index)
		},
		choosetask:function(index){
			this.bindform=this.schemes[index].scheme
			this.bindid=this.schemes[index].id
			$('#myModal').modal()
		},
		checktask:function(index){
			// this.bindform=this.schemes[index].scheme
			// this.bindid=this.schemes[index].id
			// if(this.schemes[index].id!=""){
			// this.binddatapath="/viewdata/"+this.schemes[index].id
			// }
			// $('#myModal2').modal()
			// 
			window.open("/viewdata/"+this.schemes[index].id)
		},
		updatetask:function(index){
			var id=this.schemes[index].id
			if(id ===""||id===undefined){
				console.log("Id is invalid")
				return 
			}
			var self=this;
			$.ajax({
			type:'POST',
			url:"/posttask?id="+id,
			cache:false,
			data:JSON.stringify(self.schemes[index].scheme),
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
		},
		seecode:function(){
			var markupStr = $('#summernote').summernote('code');
			console.log(markupStr)
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
<div class="page-header text-center">
  <h1>活动报名系统后台</h1>
</div>
<div id="myTabContent" class="tab-content">
   <div class="tab-pane fade in active" id="home">
      <div class="container">
	<div class="row">
		<table class="table table-hover activity-table">
			<thead>
				<th>活动ID</th>
				<th>活动名</th>
				<th>操作</th>	
				<!-- <th>活动描述</th> -->
			</thead>
			<tbody>
				<template v-for="t in schemes">
				<tr v-if="t.scheme.deleted===false" >
					<td @click="choosetask($index)">{{t.id}}</td>
					<td @click="choosetask($index)">{{t.scheme.name}}</td>
					<!-- <td @click="choosetask($index)">{{t.scheme.describe}}</td> -->
					<td><span @click="choosetask($index)">设置</span><span @click="checktask($index)">查看</span><span ><a target="_blank" href="/mainframe/{{t.id}}">浏览</a></span><span @click="deletetask($index)">删除</span></td>	
				</tr>
				</template>
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
	      		<div class="row">
	      			<label for="">活动名</label><input type="text" v-model="bindform.name">
	      		</div>
	      		
				<div class="row">
					<!-- <div id="summernote">Hello Summernote</div>
					<button type="button" @click="seecode()">ClickMe</button> -->
				</div>
				

				  

	      	</div>
	      	<!-- <div class="row">
	      		<label for="">活动描述</label><input type="text" v-model="bindform.describe">
	      	</div> -->
			<div class="row">
				<form action="/upload" method="post" enctype="multipart/form-data" id="form">
					
					<input id="uploader" type="file" name="uploadfile" style="display:inline-block;width:70%"/>
					<input type="submit" value="upload" @click="uploadfile($event)" style="display:inline-block;width:23%"/>

				</form>
				<img src="{{bindform.bgimg}}" alt="">
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
	
	<div class="modal fade" id="myModal2" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
	  <div class="modal-dialog" role="document">
	    <div class="modal-content">
	      <div class="modal-header">
	        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
	        <h4 class="modal-title" id="myModalLabel">查看活动数据</h4>
	      </div>
	      <div class="modal-body">
	      	
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
	.modal img{
		width: 100%;
	}
	iframe{
		width: 100%;
	}
	.activity-table span{
		margin-right: 5px;
		color: cadetblue;
	}
	.activity-table a{
		color: cadetblue;
	}
	.activity-table span:hover{
		cursor: pointer;
	}

</style>