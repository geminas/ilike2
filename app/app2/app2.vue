<script>

module.exports={
	el:"#show",
	data:{
	target:"app",
	scheme:{},
	logo:"../public/img/Future_Forum.jpg"
	},
	created:function(){
		console.log("app has been created");
		var self=this;
		$.ajax({
			type:'GET',
			url:"/public/json/scheme2.json",
			cache:false,
			//data:JSON.stringify(data),
			contentType:"application/json",
			processData: false,
            success:function(data){
            	for(var i in data){
            		data[i].data="",
            		data[i].error="",
            		data[i].status=""
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
					this.scheme[i].status = "has-error"
					if(errid==""){
						errid="#"+this.scheme[i].name;
					}
					errnum++;
					continue
				}else{
					this.scheme[i].error=""
					this.scheme[i].status = ""
				}
				//console.log(this.scheme[i].validator)
				if(this.scheme[i].validator==="email"){
					var re = /^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$/i;
					//console.log("email validator")
    				if(re.test(this.scheme[i].data)==false){
	    				this.scheme[i].error="邮箱输入有误,请按照example@website.xx的格式输入"
	    				this.scheme[i].status = "has-error"
	    				if(errid==""){
						errid="#"+this.scheme[i].name;
						}
	    				errnum++
	    				continue
	    			}else{
	    				this.scheme[i].error=""
	    				this.scheme[i].status = ""
	    			}
				}
				if(this.scheme[i].validator==="phone"){
					var re = /^(13[0-9]|15[012356789]|17[012678]|18[0-9]|14[57])[0-9]{8}$/;
    				if(re.test(this.scheme[i].data)==false){
	    				this.scheme[i].error="电话输入有误,请按照真实的电话号码输入"
	    				this.scheme[i].status = "has-error"
	    				if(errid==""){
						errid="#"+this.scheme[i].name;
						}
						errnum++
						continue
					}else{
						this.scheme[i].error=""
						this.scheme[i].status = ""
					}
				}

				if(this.scheme[i].validator==="idcard"){
					var re = /^(^[1-9]\d{7}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}$)|(^[1-9]\d{5}[1-9]\d{3}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])((\d{4})|\d{3}[Xx])$)$/;
    				if(re.test(this.scheme[i].data)==false){
	    				this.scheme[i].error="身份证号码输入有误"
	    				this.scheme[i].status = "has-error"
	    				if(errid==""){
						errid="#"+this.scheme[i].name;
						}
						errnum++
						continue
					}else{
						this.scheme[i].error=""
						this.scheme[i].status = ""
					}
				}

				// if(this.scheme[i].validator==="phone"){
				// 	var re = /^(13[0-9]|15[012356789]|17[012678]|18[0-9]|14[57])[0-9]{8}$/;
    // 				if(re.test(this.scheme[i].data)==false){
	   //  				this.scheme[i].error="电话输入有误,请按照真实的电话号码输入"
	   //  				this.scheme[i].status = "has-error"
	   //  				if(errid==""){
				// 		errid="#"+this.scheme[i].name;
				// 		}
				// 		errnum++
				// 		continue
				// 	}else{
				// 		this.scheme[i].error=""
				// 		this.scheme[i].status = ""
				// 	}
				// }
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
				<h3>诚挚邀请</h3>
			</div>
		</div>
		<div class="row conference-form">
			<div class="col-lg-8 col-lg-offset-2">
				<form method="POST" action="/2/apply" id="contactForm" novalidate>
					<div class="mid-div">
					<component is="smartformvue" :scheme="scheme" v-ref="smarttable"/>
					</div>
					<div class="col-lg-12 text-center submit-button">
						<button type="submit" class="btn btn-success btn-lg" @click="onsubmit">提交</button>
					</div>

				</form>
			</div>
			
		</div>
	
		<div class="row contents">
			<div class="col-lg-8 col-lg-offset-2 text-center">
				<img class="logo" src="{{logo}}">
			</div>
		</div>
	</div>
</section>
</template>