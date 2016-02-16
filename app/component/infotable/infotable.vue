<script type="text/javascript">
	function dataCom(a, b) {
		return a["timestamp"].localeCompare(b["timestamp"]) * -1;
	}

	module.exports={
		data:function(){
			return {
				//hello:"world"
				scheme:{},
				data:[],
				xlsx:[],
				pagesize:100,
				page:[],
				index:0
			}
		},
		props:["data","scheme"],
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
		ready:function(){
			console.log("infotable has been ready");
			var fs=[]
			this.data.reverse()
			console.log("data")
			console.log(this.data)
			for(var i in this.scheme.fields){
				fs.push(this.scheme.fields[i].label)
			}
			fs.push("提交时间")
			this.xlsx.push(fs)
			this.data.sort(dataCom)

			for(var d in this.data){
				var f=[]
				for(var i in this.scheme.fields){
					f.push( this.data[d][this.scheme.fields[i].cid])
				}
				f.push(this.data[d]["timestamp"])
				this.xlsx.push(f)
			}
			var tmp;
			for( var dd in this.data){
				if(dd%this.pagesize==0){
					tmp=[];
					tmp.push(this.data[dd])
					continue;
				}
				if(dd%this.pagesize==this.pagesize-1){
					tmp.push(this.data[dd])
					this.page.push(tmp)
					tmp=[];
					continue;
				}
				tmp.push(this.data[dd])
			}
			if(tmp.length!==0){
				this.page.push(tmp)
			}
			console.log("page")
			console.log(this.page)
			console.log("xlsx")
			console.log(this.xlsx)
		},
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
			pageto:function(i){
				this.index=i;
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
		        //console.log(window)
		        // var t=Date.now()
		        // window.tableExport('data-table', t, 'xls');
		        // 
		        // 
		        // 
		        // 
		        // ////////////////////////////////////
		        function datenum(v, date1904) {
					if(date1904) v+=1462;
					var epoch = Date.parse(v);
					return (epoch - new Date(Date.UTC(1899, 11, 30))) / (24 * 60 * 60 * 1000);
				}
				 
				function sheet_from_array_of_arrays(data, opts) {
					var ws = {};
					var range = {s: {c:10000000, r:10000000}, e: {c:0, r:0 }};
					for(var R = 0; R != data.length; ++R) {
						for(var C = 0; C != data[R].length; ++C) {
							if(range.s.r > R) range.s.r = R;
							if(range.s.c > C) range.s.c = C;
							if(range.e.r < R) range.e.r = R;
							if(range.e.c < C) range.e.c = C;
							var cell = {v: data[R][C] };
							if(cell.v == null) continue;
							var cell_ref = XLSX.utils.encode_cell({c:C,r:R});
							
							if(typeof cell.v === 'number') cell.t = 'n';
							else if(typeof cell.v === 'boolean') cell.t = 'b';
							else if(cell.v instanceof Date) {
								cell.t = 'n'; cell.z = XLSX.SSF._table[14];
								cell.v = datenum(cell.v);
							}
							else cell.t = 's';
							
							ws[cell_ref] = cell;
						}
					}
					if(range.s.c < 10000000) ws['!ref'] = XLSX.utils.encode_range(range);
					return ws;
				}
				 
				/* original data */
				var data = this.xlsx
				var ws_name = "SheetJS";
				 
				function Workbook() {
					if(!(this instanceof Workbook)) return new Workbook();
					this.SheetNames = [];
					this.Sheets = {};
				}
				 
				var wb = new Workbook(), ws = sheet_from_array_of_arrays(data);
				 
				/* add worksheet to workbook */
				wb.SheetNames.push(ws_name);
				wb.Sheets[ws_name] = ws;
				var wbout = XLSX.write(wb, {bookType:'xlsx', bookSST:true, type: 'binary'});

				function s2ab(s) {
					var buf = new ArrayBuffer(s.length);
					var view = new Uint8Array(buf);
					for (var i=0; i!=s.length; ++i) view[i] = s.charCodeAt(i) & 0xFF;
					return buf;
				}
				var t=Date.now()

				saveAs(new Blob([s2ab(wbout)],{type:"application/octet-stream"}), ""+t+".xlsx")
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

 	<!-- <div v-for="p in page" id="tagle-{{$index}}" style="display:none;"> -->
	 <table class="table" id="data-table">
	 	<thead>
	 		<th>number</th>
	 		<th v-for="f in scheme.fields">{{f.label}}</th>
	 		<th>time</th>
	 	</thead>
	 	<tbody>
	 		<tr v-for="d in page[index]">
	 			<td>{{index*pagesize+$index+1}}</td>
	 			<td v-for="f in scheme.fields">{{d[f.cid]}}</td>
	 			<td>{{d["timestamp"]}}</td>
	 		</tr>
	 	</tbody>
	 </table>
	 </div>
	 <div class="row text-center">
	 	<ul class="pagination">
		  <li v-for="i in page"><a href="#" @click="pageto($index)">{{$index+1}}</a></li>
		</ul>
	 </div>
	 
<!-- </div> -->
<button type="button" @click="onsave">保存成xls</button>
 

 </template>
<style type="text/css"></style>