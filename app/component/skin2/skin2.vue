<script type="text/javascript">
	module.exports={
		data:function(){
			return {
				target:"app",
				scheme:{},
				logo:"../public/img/Future_Forum.jpg",
				background:"../public/img/ljwl_background.jpg"
			}
		},
		props:["scheme"],
		/////////////Life Span/////////////////////
		// created:function(){
		// 	console.log("skin1 has been created");
		// },
		// beforeCompile:function(){
		// 	console.log("skin1 beforeCompiled");
		// },
		// compiled:function(){
		// 	console.log("skin1 has been compiled");
		// },
		// ready:function(){
		// 	console.log("skin1 has been ready");
		// },
		// attached:function(){
		// 	console.log("skin1 has been attached");
		// },
		// detached:function(){
		// 	console.log("skin1 has been detached");
		// },
		// beforeDestory:function(){
		// 	console.log("before skin1 destoryed");
		// },
		// destoryed:function(){
		// 	console.log("skin1 has been destoryed");
		// },
		methods:{
			onsubmit:function(e){
				e.preventDefault()
				console.log("on submit")
				//console.log(this.scheme)
				var res={}
				try{

					for(var i in scheme.fields){
						console.log(scheme.fields[i].name,scheme.fields[i].data)
						res[scheme.fields[i].name]=scheme.fields[i].data
					}
				}catch(e){
					console.log(e)
				}
				console.log(res)
					$.ajax({
					type:'POST',
					url:window.location.href,
					cache:false,
					data:JSON.stringify(res),
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
			// sayHello:function(){
			// 	console.log("Hello,This is the skin1 component");
			// }
			// sendurl:function(){
			// 	var data={hello:"world"}
			// 	$.ajax({
			// 		type:'POST',
			// 		url:"testskin1",
			// 		cache:false,
			// 		data:JSON.stringify(data),
			// 		contentType:"application/json",
			// 		processData: false,
		 //            success:function(data){
		 //                console.log(data)
		 //              	alert(data)
		 //            },
		 //            error:function(data){
		 //                console.log("error");
		 //                console.log(data);
		 //               alert("error: "+data)
		 //            }
			// 	});			
			// }
		},
		computed:{
			// foo:function(){
			// 	return "bar"
			// }
		},
		events:{
			//////////Events 
			// 'hook:created':function(){
			// 	console.log('skin1 has created again');
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
		},
		components:{
		'smartformvue':require('../smartform/smartform.vue')
		}

	}
</script>
<template>
<section id="contact">
	<div class="container">
		<div class="row contact-title">
			<div class="col-lg-12 text-center">
				<h3>{{scheme.name}} </h3>
			</div>
		</div>
		<div class="row conference-form">
			<div class="col-lg-8 col-lg-offset-2 ">
				<form method="POST" action="/apply" id="contactForm" novalidate>
					<div class="mid-form">
					<component is="smartformvue" :scheme="scheme.fields" v-ref="smarttable"/>
					</div>

					<!-- <div class="col-lg-12">
						<textarea rows="3" cols="20">
							xxxxx
						</textarea>
					</div> -->

					<div class="col-lg-12 text-center submit-button">
						<button type="submit" class="btn btn-success btn-lg" @click="onsubmit">点击报名</button>
					</div>

				</form>
			</div>
			
		</div>
		<div class="row">
			<div class="col-lg-8 col-lg-offset-2 text-center">
				<img class="logo" src="{{logo}}">
			</div>
		</div>
		
	</div>
</section>

 </template>
