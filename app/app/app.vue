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
			for(var i in this.scheme){
				if(this.scheme[i].required==true&&this.scheme[i].data==""){
					this.scheme[i].error="  此项不可以为空"
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
	    				errnum++
	    				continue
	    			}else{
	    				this.scheme[i].error=""
	    			}
				}
				if(this.scheme[i].validator==="phone"){
					var re = /^(\+\d{1,3}[- ]?)?\d{10}$/;
    				if(re.test(this.scheme[i].data)==false){
	    				this.scheme[i].error="电话输入有误,请按照+8612345678901"
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
			<div class="col-lg-8 col-lg-offset-2">
				<form method="POST" action="/apply" id="contactForm" novalidate>
					<div>
					<component is="smartformvue" :scheme="scheme" v-ref="smarttable"/>
					</div>
					<div class="col-lg-12 text-center">
						<button type="submit" class="btn btn-success btn-lg" @click.prevent="onsubmit">点我报名</button>
					</div>

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