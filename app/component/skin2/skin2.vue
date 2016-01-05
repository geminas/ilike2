<script type="text/javascript">
	module.exports={
		data:function(){
			return {
				target:"app",
				scheme:{},
				logo:"../public/img/Future_Forum.jpg",
				
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
		ready:function(){
			console.log("skin1 has been ready");
			// document.getElementsByTagName('body').style.backgroundImage = "url('"+this.background+"');";
			// for(var i in this.scheme.fields) {
			// 	this.scheme.fields[i].placeholder = this.scheme.fields[i].label;
			// }

			// console.log(this.scheme);
		},
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
<section id="contact" style="background:url('{{scheme.bgimg}}');">
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
<style>
	#contact .contact-title h2{
  text-align: center;
}
#contact .contact-title{
  margin-bottom: 30px;
  color:white;
}

#contact {
    padding-top: 30px;
}

#contact .has-error .form-control{
  border-color: #A72585;
  border-width: 3px;
}

#contact{
 /* background: url('{{background}}');*/
  background-size: cover !important;
  background-position-y: 40% !important;
}

#contact .contact-title h3{
  font-size: 3em;
  margin-bottom: 20px;
}

#contact .contact-title h2{
  font-size: 5em;
  margin-bottom: 20px;
}

#contact .contact-title h6{
  margin-top: 0px;
}

#contact .conference-form,#contact .contents{
  background-color: rgba(255, 255, 255, 0.35);
  color:black;
}
/*#contact .container{
  max-width: 600px;
}*/
#contact row{}
#contact .btn{
  margin-bottom: 0px;
}
#contact label{
  width: 100%;
  margin: 13px auto;
  padding-top:10px;
  color: white;
  display: none;
}

#contact .label-source {
  margin-right: 10px;
  display: inline;
}

#contact .label-checkboxes{
  display: block;
}

#contact .label-checkboxes ~ div{
  margin-right: 10px;
  display: inline;
}

#contact .label-radio {
  display: block;
}

#contact .label-radio ~div {
  margin-right: 10px;
  display: inline;
}

#contact input[type="checkbox"]{
	/*margin-left:10px; */
}
#contact .required{
	color:red;
}

#contact .submit-button{
	margin:15px 0px 18px 0px;
}

#contact .submit-button .btn {
    background: #F52B2B;
    width: 250px;
}

#contact .logo {
  width: 200px;
  margin-bottom: 10px;
}

#contact select {
  margin-top: 10px;
  margin-bottom: 10px;
}

#contact textarea {
  resize:none;
}

@media (max-width: 700px) {
  #contact .logo {
    width: 100px;
  }

  #contact .contact-title{
    margin-bottom: 30px;
    color:white;
  }

  #contact .submit-button .btn {
    background: #F52B2B;
    width: 180px;
  }

  #contact .contact-title h3 {
    font-size: 1.8em;
    margin-top: 0px;
    margin-bottom: 0px;
  }

  #contact .contact-title h2 {
    font-size: 2em;
  }
}

</style>