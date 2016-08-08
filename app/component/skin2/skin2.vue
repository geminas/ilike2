<script type="text/javascript">

	Date.prototype.format = function(format) {
	       var date = {
	              "M+": this.getMonth() + 1,
	              "d+": this.getDate(),
	              "h+": this.getHours(),
	              "m+": this.getMinutes(),
	              "s+": this.getSeconds(),
	              "q+": Math.floor((this.getMonth() + 3) / 3),
	              "S+": this.getMilliseconds()
	       };
	       if (/(y+)/i.test(format)) {
	              format = format.replace(RegExp.$1, (this.getFullYear() + '').substr(4 - RegExp.$1.length));
	       }
	       for (var k in date) {
	              if (new RegExp("(" + k + ")").test(format)) {
	                     format = format.replace(RegExp.$1, RegExp.$1.length == 1
	                            ? date[k] : ("00" + date[k]).substr(("" + date[k]).length));
	              }
	       }
	       return format;
	}

	module.exports={
		data:function(){
			return {
				target:"app",
				scheme:{},
				id:"",
				phone:"",
				logo:"../public/img/Future_Forum.jpg",
				
			}
		},
		props:["scheme","id"],
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
			//var arr=location.href.split('/')
			//this.id=
			console.log(this.id)
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
				var phoneerr=false
				var self=this
				e.preventDefault()
				e.target.disabled=true;

				console.log("on submit")
			    var errnum=0;
				var errid=""
				console.log(this.scheme)
				var res={}
				//var msg=""
				try{

					for(var i in scheme.fields){
						console.log(scheme.fields[i].name,scheme.fields[i].data)
						if(scheme.fields[i].field_type=="checkboxes"){
							for(var j in scheme.fields[i].datas){
								scheme.fields[i].data+=scheme.fields[i].datas[j]
								scheme.fields[i].data+=";"
							}
						}
						if(this.scheme.fields[i].required==true&&this.scheme.fields[i].data==""){
							this.scheme.fields[i].error="此项不可以为空"
							this.scheme.fields[i].status = "has-error"
							if(errid==""){
								errid="#"+this.scheme.fields[i].name;
							}
							errnum++;

							continue
						}else{
							this.scheme.fields[i].error=""
							this.scheme.fields[i].status = ""
						}
						if(this.scheme.fields[i].field_type==="email"){
							var re = /^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$/i;
							//console.log("email validator")
		    				if(re.test(this.scheme.fields[i].data)==false){
			    				this.scheme.fields[i].error="邮箱输入有误,请按照example@website.xx的格式输入"
			    				this.scheme.fields[i].status = "has-error"
			    				if(errid==""){
								errid="#"+this.scheme.fields[i].name;
								}
			    				errnum++
			    				continue
			    			}else{
			    				this.scheme.fields[i].error=""
			    				this.scheme.fields[i].status = ""
			    			}
						}
						if(this.scheme.fields[i].field_type==="phone"){
							var re = /^(13[0-9]|15[012356789]|17[012678]|18[0-9]|14[57])[0-9]{8}$/;
		    				if(re.test(this.scheme.fields[i].data)==false){
			    				this.scheme.fields[i].error="电话输入有误,请按照真实的电话号码输入"
			    				this.scheme.fields[i].status = "has-error"
			    				if(errid==""){
								errid="#"+this.scheme.fields[i].name;
								}
								errnum++
								continue
							}else{
								//var iferr=false
								
									this.phone=this.scheme.fields[i].data
									this.scheme.fields[i].error=""
									this.scheme.fields[i].status = ""
								
							}
						}

						if(this.scheme.fields[i].field_type==="idcard"){
							var re = /^(^[1-9]\d{7}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}$)|(^[1-9]\d{5}[1-9]\d{3}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])((\d{4})|\d{3}[Xx])$)$/;
		    				if(re.test(this.scheme.fields[i].data)==false){
			    				this.scheme.fields[i].error="身份证号码输入有误"
			    				this.scheme.fields[i].status = "has-error"
			    				if(errid==""){
								errid="#"+this.scheme.fields[i].name;
								}
								errnum++
								continue
							}else{
								this.scheme.fields[i].error=""
								this.scheme.fields[i].status = ""
							}
						}
						

						res[scheme.fields[i].cid]=scheme.fields[i].data
					}
				}catch(e){
					console.log(e)
				}
				
				console.log(res)
				res["timestamp"]=new Date().format('yyyy-MM-dd hh:mm:ss')
				if(errnum==0){
					e.target.disabled="disable"
				console.log("ok to submit")
				
			$.ajax({
				type:'GET',
				url:"/phoneexist/"+this.id+"/"+this.phone,
				cache:false,
				processData: false,
	            success:function(data){
	                console.log(data)
	                if(data.status==0){
	                	console.log("phone not exist")
	                	$.ajax({
						type:'POST',
						url:window.location.href+"?phone="+self.phone,
						cache:false,
						data:JSON.stringify(res),
						contentType:"application/json",
						processData: false,
			            success:function(data){
			                console.log(data)
			                if(data.status==0){
			                	window.location.href="/thankyou";
			                }else{
			                	alert(data.msg);
			                	e.target.disabled=false;
			                }
			              	//alert(data)
			            },
			            error:function(data){
			                console.log("error");
			                console.log(data);
			               alert("error: "+data)
			               e.target.disabled=false;
			            }
					});

	                }else{
	                	alert("电话号码已经存在");
	                	e.target.disabled=false;
	                }
	              	//alert(data)
	            },
	            error:function(data){
	                console.log("error");
	                console.log(data);
	               alert("error: "+data)
	               e.target.disabled=false;
	            }
			});

									
				}else{
					e.target.disabled=false;
					e.preventDefault()
					console.log(errid)
					alert("请完成没有填完的必填选项,或按照提示修改错误")
					$('html, body').animate({
	                    scrollTop:($(errid).offset().top-50)
	                }, 200);
					return false
				}
					
			},
			jumpToWebsite:function(e){
				e.preventDefault()
				window.open("http://www.futureforum.org.cn")
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
					<div class="row">
						<div class="col-lg-12 text-center submit-button">
							<button type="submit" class="btn btn-success btn-lg" @click="onsubmit">点击报名</button>
						</div>
					</div>
					<div class="row mid-form">
					<div class="schemeclass" v-html="scheme.code">
						<!-- {{scheme.code}}</pre> -->
					</div>
					</div>
				</form>
			</div>
			
		</div>
		
		<div class="row">
			<div class="col-lg-8 col-lg-offset-2 text-center">
			<a href="http://www.futureforum.org.cn" @click="jumpToWebsite|prevent">
				<img class="logo" src="{{logo}}">
			</a>
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
  margin-top:0px;
  margin-bottom: 10px;
  padding-top:0px;
  color: white;
  display: none;
}

#contact .label-text-error{
	display: block;
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
#contact .mid-form{
	margin-top: 10px;

}

#contact .schemeclass{
	background-color: white; 
	padding-left: 10px;
	padding-right: 10px;
}

#contact .label-radio ~div {
  margin-right: 10px;
  display: inline;
}

#contact input[type="radio"]{
	margin-bottom:10px; 
}

#contact input[type="checkbox"]{
	margin-bottom:10px; 
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
  /*margin-top: 10px;*/
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