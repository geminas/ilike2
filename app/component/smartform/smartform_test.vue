<script>

module.exports={
	el:"#show",
	data:{
	target:"smartform",
	scheme:{}
	},
	created:function(){
		console.log("smartform test app has been created");
		var self=this;
		$.ajax({
			type:'GET',
			url:"/getjson",
			cache:false,
			//data:JSON.stringify(data),
			contentType:"application/json",
			processData: false,
            success:function(data){
                //console.log(data)
              	//alert(data)
              	var d=JSON.parse(data)
              	//console.log(d)
              	self.$data.scheme=d
            },
            error:function(data){
                console.log("error");
                //console.log(data);
               alert("error: "+data)
            }
		});	
	},
	ready:function(){
	console.log("The smartform test main is loaded")
	},
	components:{
	'smartformvue':require('./smartform.vue')
	}
}
</script>

<template>
<div><h1>Test--{{target}} is below</h1></div>
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
			<div class="col-lg-8 col-lg-offset-2">
				<form method="POST" action="/apply" id="contactForm" novalidate>
					<component is="smartformvue" :scheme="scheme"/>
				</form>
			</div>
		</div>
		<div class="row">
			<div class="col-lg-8 col-lg-offset-2">
				<p style="color:gray;font-size:13px;padding:0 15px;"> 
					1.报名截止日期：2015年12月31日 <br>
					2.付费嘉宾请将《付费参会注册表》于2015年12月25日前发至组委会邮箱candy.liu@futureforum.org.cn；<br>
					3.申请免费参会的嘉宾请将《免费参会申请表》于2015年12月31日前发至组委会邮箱candy.liu@futureforum.org.cn <br>
					4.参会咨询：刘女士  18511296094  010-58751635 
				</p>
			</div>
		</div>
			
	</div>
</section>
</template>