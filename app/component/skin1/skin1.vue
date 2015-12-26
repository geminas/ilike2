<script type="text/javascript">
	module.exports={
		data:function(){
			return {
				target:"app",
				scheme:{},
				logo:"../public/img/Future_Forum.jpg"
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
				<h3>{{scheme.describe}} </h3>
				<h2>{{scheme.name}}</h2>
				<h6>{{scheme.time}}</h6>	
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
			
			<div class="col-lg-8 col-lg-offset-2 bottom-form">
				<p  style="color:gray;font-size:13px;padding:0 15px;"> 
				<br>
					<strong>一、购票来宾权益说明：</strong>	<br><br>
					1、	贵宾坐席A级（2880元/人）<br>
						享受1月17日全天会议专属坐席区域（购票贵宾区1-2排）<br>
						享受1月17日午间商务自助餐（国贸酒店餐厅）<br>
						享受1月17日全天会间茶歇服务（国贸酒店精品服务）<br>
						获得年会定制限量版资料及纪念品（全套，如：刊物、书籍、礼券、精美礼品等）<br>
					2、	贵宾坐席B级（1880元/人）<br>
						享受1月17日全天会议专属坐席区域（购票贵宾区3-4排）<br>
						享受1月17日午间商务自助餐午餐（国贸酒店餐厅）<br>
						享受1月17日全天会间茶歇服务（国贸酒店精品服务）<br>
						获得年会定制限量版资料及纪念品（两份，如：刊物、书籍、礼券、精美礼品等）<br><br>

					<strong>二、免费参会申请：</strong><br><br>
					本次年会同时开放免费入场机会，提交报名后，组委会将于五个工作日内进行资格审核，通过后将以邮件回复发出“通关门票”。<br>
					免费座位处于主会场后区，限制400席，火热抢票进行中！<br>
					备注：申请免费参会人员，请务必满足以下几个条件中任意一项<br>
					1、在校大学生，本科以上类别，所在学校属于全国前十位名优秀大学且成绩优异。<br>
					2、企业管理者，涉及科学及科技领域的各行业企业高级领导，CTO以上级别。<br>
					3、职场精英，对于科学、人文、生命、宇宙等未来话题深感兴趣的跨龄职场人士。<br><br>

					<strong>三、购票报名截止日期：</strong><br><br>
					2015年12月31日 (周四)，请在此日期之前完成购票报名。<br><br>

					<strong>四、联系方式：<strong/><br><br>
					Email：FF2016@futureforum.org.cn<br>
					Tel：010-58751635<br>
					刘女士： 18511296094<br>
					徐女士： 18301029183<br>
					<br>
				</p>
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
<style type="text/css">
#contact .contact-title h2{
  text-align: center;
}
#contact .contact-title{
  margin-bottom: 0px;
  color:white;
  background-color: cadetblue;
}

#contact .has-error .form-control{
  border-color: #A72585;
  border-width: 3px;
}

body{
  font-family: 'Avenir Next',Avenir,'Helvetica Neue',Helvetica,'Lantinghei SC','Hiragino Sans GB','Microsoft YaHei',"微软雅黑",STHeiti,'WenQuanYi Micro Hei',SimSun,sans-serif;
}

section {
  padding: 0px 0;
}

#contact .contact-title h3{
  font-size: 1.8em;
  margin-bottom: 20px;
}

#contact .contact-title h2{
  font-size: 4em;
  margin-bottom: 20px;
}

#contact .contact-title h6{
  margin-top: 0px;
}

#contact .conference-form,#contact .contents{
  background-color: beige;
  color:black;
}
/*#contact .container{
  max-width: 600px;
}*/
#contact row{}
#contact .btn{
  margin-bottom: 20px;
}
#contact label{
  width: 100%;
  border-top: thin solid #eee;
  margin: 13px auto;
  padding-top:10px;
  font-weight: 500;
}
#contact input[type="checkbox"]{
	/*margin-left:10px; */
}
#contact .required{
	color:red;
}

#contact .submit-button{
	margin:15px 0px 45px 0px;
}

#contact .bottom-form {
  background-color: wheat;
}

#contact .contents p {
  background-color: honeydew;
}

#contact .logo {
  width: 200px;
}

#contact textarea {
  resize:none;
}

@media (max-width: 700px) {
  #contact .contact-title h3{
    font-size: 1.6em;
    margin-bottom: 20px;
  }

  #contact .contact-title h2 {
    font-size: 2em;
  }

  #contact .logo {
    width: 100px;
  }
}
</style>