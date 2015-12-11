<script>

module.exports={
	el:"#show",
	data:{
	target:"app",
	scheme:{}
	},
	created:function(){
		console.log("app has been created");
		var self=this;
		$.ajax({
			type:'GET',
			url:"/public/json/scheme.json",
			cache:false,
			//data:JSON.stringify(data),
			contentType:"application/json",
			processData: false,
            success:function(data){
            	for(var i in data){
            		data[i].data="",
            		data[i].error=""
            	}
              	self.$data.scheme=data
              	//console.log(JSON.stringify(data))
            },
            error:function(data){
                console.log("error");
                //console.log(data);
               alert("error: "+data)
            }
		});	
	},
	methods:{
		onsubmit:function(e){
			console.log(e)
			//console.log(a,b,c)
			//console.log(this.$)
			//console.log(this.scheme)

			var errnum=0;
			var errid=""
			for(var i in this.scheme){
				if(this.scheme[i].required==true&&this.scheme[i].data==""){
					this.scheme[i].error="  此项不可以为空"
					if(errid==""){
						errid="#"+this.scheme[i].name;
					}
					errnum++;
					continue
				}else{
					this.scheme[i].error=""
				}
				//console.log(this.scheme[i].validator)
				if(this.scheme[i].validator==="email"){
					var re = /^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$/i;
					//console.log("email validator")
    				if(re.test(this.scheme[i].data)==false){
	    				this.scheme[i].error="邮箱输入有误,请按照example@website.xx的格式输入"
	    				if(errid==""){
						errid="#"+this.scheme[i].name;
						}
	    				errnum++
	    				continue
	    			}else{
	    				this.scheme[i].error=""
	    			}
				}
				if(this.scheme[i].validator==="phone"){
					var re = /^(0|86|17951)?(13[0-9]|15[012356789]|17[678]|18[0-9]|14[57])[0-9]{8}$/;
    				if(re.test(this.scheme[i].data)==false){
	    				this.scheme[i].error="电话输入有误,请按照真实的电话号码输入"
	    				if(errid==""){
						errid="#"+this.scheme[i].name;
						}
						errnum++
						continue
					}else{
						this.scheme[i].error=""
					}
				}

				if(this.scheme[i].validator==="idcard"){
					var re = /^(^[1-9]\d{7}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}$)|(^[1-9]\d{5}[1-9]\d{3}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])((\d{4})|\d{3}[Xx])$)$/;
    				if(re.test(this.scheme[i].data)==false){
	    				this.scheme[i].error="身份证号码输入有误"
	    				if(errid==""){
						errid="#"+this.scheme[i].name;
						}
						errnum++
						continue
					}else{
						this.scheme[i].error=""
					}
				}
			}
			if(errnum==0){
				console.log("ok to submit")
				//return true
				$(e.target).trigger(e);
			}else{
				e.preventDefault()
				console.log(errid)
				$('html, body').animate({
                    scrollTop:($(errid).offset().top-50)
                }, 200);
				return false
			}
		}
	},
	ready:function(){
	console.log("The smartform test main is loaded")
	},
	components:{
	'smartformvue':require('../component/smartform/smartform.vue')
	}
}
</script>

<template>

<section id="contact">
	<div class="container">
		<div class="row contact-title">
			<div class="col-lg-12 text-center">
				<h3>未来论坛2016年会报名</h3>
				<h2>人类认知新百年</h2>
				<h6>2016年1月17日   中国·北京·国贸三期</h6>	
			</div>
		</div>
		<div class="row">
			<div class="col-lg-8 col-lg-offset-2 conference-form">
				<form method="POST" action="/apply" id="contactForm" novalidate>
					<div>
					<component is="smartformvue" :scheme="scheme" v-ref="smarttable"/>
					</div>
					<div class="col-lg-12 text-center submit-button">
						<button type="submit" class="btn btn-success btn-lg" @click="onsubmit">点击报名</button>
					</div>

				</form>
			</div>
			
		</div>
	
		<div class="row">
			<div class="col-lg-8 col-lg-offset-2 contents">
				<p style="color:gray;font-size:13px;padding:0 15px;"> 
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
					备注：申请免费人员，请务必满足以下几个条件中任意一项<br>
					1、在校大学生，本科以上类别，所在学校属于全国前十位名优秀大学且成绩优异。<br>
					2、企业管理者，涉及科学及科技领域的各行业企业高级领导，CTO以上级别。<br>
					3、职场精英，对于科学、人文、生命、宇宙等未来话题深感兴趣的跨龄职场人士。<br><br>

					<strong>三、购票报名截止日期：</strong><br>
					2015年12月31日 (周四)，请在此日期之前完成购票报名。<br><br>

					<strong>四、联系方式：<strong/><br><br>
					Tel：010-58751635<br>
					刘女士： 18511296094<br>
					徐女士： 18301029183<br>
				</p>
			</div>
		</div>
	</div>
</section>
</template>