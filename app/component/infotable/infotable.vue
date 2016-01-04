<script type="text/javascript">
	module.exports={
		data:function(){
			return {
				//hello:"world"
				data:[]
			}
		},
		props:["data"],
		/////////////Life Span/////////////////////
		// created:function(){
		// 	console.log("infotable has been created");
		// },
		// beforeCompile:function(){
		// 	console.log("infotable beforeCompiled");
		// },
		// compiled:function(){
		// 	console.log("infotable has been compiled");
		// },
		// ready:function(){
		// 	console.log("infotable has been ready");
		// },
		// attached:function(){
		// 	console.log("infotable has been attached");
		// },
		// detached:function(){
		// 	console.log("infotable has been detached");
		// },
		// beforeDestory:function(){
		// 	console.log("before infotable destoryed");
		// },
		// destoryed:function(){
		// 	console.log("infotable has been destoryed");
		// },
		methods:{
			onsave:function(){
				this.savexls(this.data)
			},
			savexls:function(objArray){
		        // var array = typeof objArray != 'object' ? JSON.parse(objArray) : objArray;

		        // var str = '';

		        // for (var i = 0; i < array.length; i++) {
		        //     var line = '';

		        //     for (var index in array[i]) {
		        //         line += array[i][index] + ',';
		        //     }

		        //     // Here is an example where you would wrap the values in double quotes
		        //     // for (var index in array[i]) {
		        //     //    line += '"' + array[i][index] + '",';
		        //     // }

		        //     line.slice(0,line.Length-1); 

		        //     str += line + '\r\n';
		        // }
		        // window.open( "data:text/csv;charset=utf-8," + escape(str))
		        console.log(window)
		        var t=Date.now()
		        window.tableExport('data-table', t, 'xls');
		    }
			// sayHello:function(){
			// 	console.log("Hello,This is the infotable component");
			// }
			// sendurl:function(){
			// 	var data={hello:"world"}
			// 	$.ajax({
			// 		type:'POST',
			// 		url:"testinfotable",
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
			// 	console.log('infotable has created again');
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
		}

	}
</script>
<template>

<!-- <div>This is component infotable</div>
 -->
 <div class="container-fluid">
	 <table class="table" id="data-table">
	 	<thead>
	 		<th>number</th>
	 		<th v-for="(key,val) in data[0]">{{key}}</th>
	 	</thead>
	 	<tbody>
	 		<tr v-for="d in data">
	 			<td>{{$index+1}}</td>
	 			<td v-for="(k,v) in d">{{v}}</td>
	 		</tr>
	 	</tbody>
	 </table>
</div>
<button type="button" @click="onsave">保存成xls</button>
 </template>
<style type="text/css"></style>