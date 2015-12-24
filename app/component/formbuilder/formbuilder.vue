<script type="text/javascript">
	
	module.exports={
		data:function(){
			return {
				fb:{},
				scheme:[],
				payload:{}
				//hello:"world"
			}
		},
		props:["scheme"],
		/////////////Life Span/////////////////////
		created:function(){
			console.log("formbuilder has been created");
			
      	},
		// beforeCompile:function(){
		// 	console.log("formbuilder beforeCompiled");
		// },
		// compiled:function(){
		// 	console.log("formbuilder has been compiled");
		// },
		ready:function(){
			console.log("formbuilder has been ready");
			var self=this
			this.fb = new Formbuilder({
	        selector: '#formbuilder',
		        bootstrapData:JSON.stringify(self.scheme)
			});
			//console.log()
		    // this.fb.on('save', function(payload){
		    //     console.log(payload);
		    // })
		},
		// attached:function(){
		// 	console.log("formbuilder has been attached");
		// },
		// detached:function(){
		// 	console.log("formbuilder has been detached");
		// },
		// beforeDestory:function(){
		// 	console.log("before formbuilder destoryed");
		// },
		// destoryed:function(){
		// 	console.log("formbuilder has been destoryed");
		// },
		methods:{

			// toJSON:function(){
			// 	//var j=this.fb.mainView.collection.toJSON()
			// 	//console.log(j)
			// 	return  JSON.stringify(j)
			// },
			getCollection:function(){
				//console.log("payload is below")
				//console.log(this.payload)
				this.fb.mainView.saveForm()
				return  this.fb.mainView.collection.toJSON()
			}
		},
		computed:{
			// foo:function(){
			// 	return "bar"
			// }
		},
		events:{
			//////////Events 
			// 'hook:created':function(){
			// 	console.log('formbuilder has created again');
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
			'scheme':function(){
				var self=this
				//ex.push(this.scheme[1])
				//console.log(ex)
				this.fb = new Formbuilder({
		        selector: '#formbuilder',
			        bootstrapData:self.scheme});
				//console.log(this.fb)
				//console.log("lalalalala")
				//console.log(this.fb.mainView.collection.toJSON())
			    this.fb.on('save', function(payload){
			        //console.log("saved");
			        //console.log(payload)
			        self.payload=payload
			    })

			}
		}

	}
</script>
<template>
<div id='formbuilder'></div>

 </template>
<style type="text/css"></style>