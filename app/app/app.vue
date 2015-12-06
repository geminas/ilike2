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
			url:"/public/json/tmptask.json",
			cache:false,
			//data:JSON.stringify(data),
			contentType:"application/json",
			processData: false,
            success:function(data){
              	self.$data.scheme=data
            },
            error:function(data){\
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
	'smartformvue':require('../component/smartform/smartform.vue')
	}
}
</script>

<template>

<section id="contact">
	<div class="container">
		
		<div class="row contact-title">
			<div class="col-lg-12 text-center">
				<h2>参会注册表</h2>
			</div>
		</div>
		<div class="row">
			<div class="col-lg-8 col-lg-offset-2">
				<form method="POST" action="/apply" id="contactForm" novalidate>
					<div>
						<component is="smartformvue" :scheme="scheme"/>
					</div>
					
					<div class="col-lg-12 text-center">
						<button type="submit" class="btn btn-success btn-lg">点我报名</button>
					</div>
				</form>
			</div>
		</div>
		
	</div>
</section>
</template>