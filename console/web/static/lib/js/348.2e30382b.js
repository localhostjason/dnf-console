"use strict";(self["webpackChunkvue3_template"]=self["webpackChunkvue3_template"]||[]).push([[348],{491:function(e,l,a){a.d(l,{K5:function(){return d},T8:function(){return u},ZC:function(){return m},Zr:function(){return n},p:function(){return o},ss:function(){return i},tm:function(){return r}});var t=a(6569);const u=e=>t.d.request({url:"/gm/account/list",method:"get",params:e}),o=(e,l)=>t.d.request({url:`/gm/account/${e}/recharge`,method:"post",data:l}),d=e=>t.d.request({url:`/gm/account/${e}/reset_create_charac`,method:"post"}),r=e=>t.d.request({url:`/gm/account/${e}`,method:"delete"}),n=(e,l)=>t.d.request({url:`/gm/account/${e}/change_password`,method:"post",data:l}),m=(e,l)=>t.d.request({url:`/gm/account/${e}`,method:"put",data:l}),i=e=>t.d.request({url:"/gm/account",method:"post",data:e})},9456:function(e,l,a){a.d(l,{F:function(){return u}});var t=a(6569);const u=e=>t.d.request({url:`/gm/account/${e}/roles`,method:"get"})},2870:function(e,l,a){a.d(l,{Z:function(){return w}});var t=a(7477),u=a(3396),o=a(7139);const d={class:"bot",style:{height:"20px","margin-bottom":"13px",background:"#fff"}},r={class:"bline-dotted"},n={class:"category-index"},m={class:"category-name"},i={style:{padding:"10px 30px","margin-bottom":"30px"}};function s(e,l,a,t,s,p){const c=(0,u.up)("el-col"),f=(0,u.up)("el-row");return(0,u.wg)(),(0,u.iD)("div",null,[(0,u.Wm)(f,{class:(0,o.C_)(a.isLast?"last-step":"")},{default:(0,u.w5)((()=>[(0,u.Wm)(c,{span:24},{default:(0,u.w5)((()=>[(0,u._)("div",d,[(0,u._)("div",r,[(0,u._)("div",n,(0,o.zw)(a.num),1),(0,u._)("div",m,(0,o.zw)(a.title),1)])])])),_:1})])),_:1},8,["class"]),(0,u._)("div",null,[(0,u._)("div",i,[(0,u.WI)(e.$slots,"default",{},void 0,!0)])])])}var p={name:"PanelStep",props:{num:{type:[String,Number],required:!0,default:1},title:{type:String},isLast:{type:Boolean,default:!1}}},c=a(89);const f=(0,c.Z)(p,[["render",s],["__scopeId","data-v-af9f6fb6"]]);var g=f;const _=(0,t.nz)(g);var w=_},7032:function(e,l,a){a.r(l),a.d(l,{default:function(){return q}});var t=a(3396),u=a(4870),o=a(2507),d=a(2870),r=a(612),n=a(9242),m=a(7192),i=a(6569);const s=(e,l)=>i.d.request({url:`/gm/roles/${e}/email`,method:"post",data:l}),p=e=>i.d.request({url:"/gm/gold/list",method:"get",params:e});var c=a(6881);const f=(0,t._)("span",{class:"text-warning sm"},"ss装备不要勾选！！！",-1),g=(0,t._)("div",{class:"card-header"},[(0,t._)("span",null,"物品代码")],-1);var _=(0,t.aZ)({__name:"EmailForm",setup(e,{expose:l}){const a=(0,u.iH)(),o=(0,u.qj)({code:null,number:0,seperate_upgrade:0,upgrade:0,is_amplify:!1,amplify_option:3,amplify_value:0,gold:0,seal_flag:!1}),d=(0,u.qj)({code:[{required:!0,message:"请输入物品代码",trigger:"blur"}],number:[{required:!0,message:"请输入数量",trigger:"blur"}],upgrade:[{type:"integer",min:0,max:31,message:"强化/增幅等级在0至31",trigger:"blur"}]}),r=(0,u.iH)(null),i=[{label:"体力",value:1},{label:"精神",value:2},{label:"力量",value:3},{label:"智力",value:4}],_=(0,u.qj)({data:[],total:0,loading:!1}),w=(0,u.qj)({page:1,per_page:10});let b=(0,u.qj)({name:""});const v=(0,u.iH)(),W=async()=>{const e=await(0,m.Gu)(a);if(e)if(r.value)try{await s(r.value,o),(0,c.PE)("发送成功，请小退一下！")}catch(l){}else(0,c.N3)("未获取到角色ID")},y=e=>{r.value=e},h=async()=>{_.loading=!0;const e={...b,...w},{items:l,total:a}=await p(e);_.data=l,_.total=a,_.loading=!1},V=e=>{w.per_page=e,h()},U=e=>{w.page=e,h()},q=()=>{w.page=1,h()},k=()=>{v.value.resetFields(),w.page=1,h()};return h(),l({setCharacNo:y}),(e,l)=>{const r=(0,t.up)("el-input"),m=(0,t.up)("el-form-item"),s=(0,t.up)("el-checkbox"),p=(0,t.up)("el-option"),c=(0,t.up)("el-select"),y=(0,t.up)("el-button"),h=(0,t.up)("el-form"),x=(0,t.up)("el-col"),C=(0,t.up)("el-row"),S=(0,t.up)("el-table-column"),Z=(0,t.up)("el-table"),z=(0,t.up)("el-pagination"),j=(0,t.up)("el-card"),D=(0,t.Q2)("loading");return(0,t.wg)(),(0,t.iD)("div",null,[(0,t.Wm)(C,{gutter:20},{default:(0,t.w5)((()=>[(0,t.Wm)(x,{span:8},{default:(0,t.w5)((()=>[(0,t.Wm)(h,{model:o,rules:d,ref_key:"formRef",ref:a,"label-width":"100px"},{default:(0,t.w5)((()=>[(0,t.Wm)(m,{label:"物品代码",prop:"code"},{default:(0,t.w5)((()=>[(0,t.Wm)(r,{modelValue:o.code,"onUpdate:modelValue":l[0]||(l[0]=e=>o.code=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),(0,t.Wm)(m,{label:"数量",prop:"number"},{default:(0,t.w5)((()=>[(0,t.Wm)(r,{modelValue:o.number,"onUpdate:modelValue":l[1]||(l[1]=e=>o.number=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),(0,t.Wm)(m,{label:"武器锻造等级",prop:"seperate_upgrade"},{default:(0,t.w5)((()=>[(0,t.Wm)(r,{modelValue:o.seperate_upgrade,"onUpdate:modelValue":l[2]||(l[2]=e=>o.seperate_upgrade=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),(0,t.Wm)(m,{label:"强化/增幅等级",prop:"upgrade"},{default:(0,t.w5)((()=>[(0,t.Wm)(r,{modelValue:o.upgrade,"onUpdate:modelValue":l[3]||(l[3]=e=>o.upgrade=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),(0,t.Wm)(m,{label:"具备异界属性"},{default:(0,t.w5)((()=>[(0,t.Wm)(s,{modelValue:o.is_amplify,"onUpdate:modelValue":l[4]||(l[4]=e=>o.is_amplify=e)},null,8,["modelValue"])])),_:1}),(0,t.wy)((0,t.Wm)(m,{label:"红字"},{default:(0,t.w5)((()=>[(0,t.Wm)(c,{modelValue:o.amplify_option,"onUpdate:modelValue":l[5]||(l[5]=e=>o.amplify_option=e)},{default:(0,t.w5)((()=>[((0,t.wg)(),(0,t.iD)(t.HY,null,(0,t.Ko)(i,(e=>(0,t.Wm)(p,{key:e.value,label:e.label,value:e.value},null,8,["label","value"]))),64))])),_:1},8,["modelValue"])])),_:1},512),[[n.F8,o.is_amplify]]),(0,t.wy)((0,t.Wm)(m,{label:"追加属性",prop:"amplify_value"},{default:(0,t.w5)((()=>[(0,t.Wm)(r,{modelValue:o.amplify_value,"onUpdate:modelValue":l[6]||(l[6]=e=>o.amplify_value=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1},512),[[n.F8,o.is_amplify]]),(0,t.Wm)(m,{label:"金币",prop:"gold"},{default:(0,t.w5)((()=>[(0,t.Wm)(r,{modelValue:o.gold,"onUpdate:modelValue":l[7]||(l[7]=e=>o.gold=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),(0,t.Wm)(m,{label:"是否封装"},{default:(0,t.w5)((()=>[(0,t.Wm)(s,{modelValue:o.seal_flag,"onUpdate:modelValue":l[8]||(l[8]=e=>o.seal_flag=e)},null,8,["modelValue"]),f])),_:1}),(0,t.Wm)(m,null,{default:(0,t.w5)((()=>[(0,t.Wm)(y,{type:"primary",size:"small",onClick:W},{default:(0,t.w5)((()=>[(0,t.Uk)("发送")])),_:1})])),_:1})])),_:1},8,["model","rules"])])),_:1}),(0,t.Wm)(x,{span:13,offset:3},{default:(0,t.w5)((()=>[(0,t.Wm)(j,{class:"box-card"},{header:(0,t.w5)((()=>[g])),default:(0,t.w5)((()=>[(0,t._)("div",null,[(0,t.Wm)(C,null,{default:(0,t.w5)((()=>[(0,t.Wm)(h,{inline:!0,model:(0,u.SU)(b),ref_key:"filterFormRef",ref:v,"label-width":"100px"},{default:(0,t.w5)((()=>[(0,t._)("div",null,[(0,t.Wm)(m,{label:"物品名称:",prop:"name"},{default:(0,t.w5)((()=>[(0,t.Wm)(r,{modelValue:(0,u.SU)(b).name,"onUpdate:modelValue":l[9]||(l[9]=e=>(0,u.SU)(b).name=e),clearable:""},null,8,["modelValue"])])),_:1}),(0,t.Wm)(m,null,{default:(0,t.w5)((()=>[(0,t.Wm)(y,{type:"primary",onClick:q},{default:(0,t.w5)((()=>[(0,t.Uk)("查询")])),_:1}),(0,t.Wm)(y,{onClick:(0,n.iM)(k,["prevent"])},{default:(0,t.w5)((()=>[(0,t.Uk)("重置")])),_:1},8,["onClick"])])),_:1})])])),_:1},8,["model"])])),_:1}),(0,t.Wm)(C,null,{default:(0,t.w5)((()=>[(0,t.wy)(((0,t.wg)(),(0,t.j4)(Z,{data:_.data,border:"",fit:""},{default:(0,t.w5)((()=>[(0,t.Wm)(S,{prop:"code",label:"物品编号",width:"150"}),(0,t.Wm)(S,{prop:"name",label:"物品名称"})])),_:1},8,["data"])),[[D,_.loading]])])),_:1}),(0,t.Wm)(C,null,{default:(0,t.w5)((()=>[(0,t.Wm)(z,{"current-page":w.page,"page-sizes":[10,20,30,40,50],"page-size":w.per_page,total:_.total,background:"",layout:"total, sizes, prev, pager, next, jumper",onSizeChange:V,onCurrentChange:U},null,8,["current-page","page-size","total"])])),_:1})])])),_:1})])),_:1})])),_:1})])}}});const w=_;var b=w,v=a(9456),W=a(2899);const y={class:"tc-step-border"},h={class:"l"};var V=(0,t.aZ)({__name:"index",setup(e){const l=(0,u.qj)({data:[],loading:!1}),a=async e=>{try{l.loading=!0,l.data=await(0,v.F)(e),l.loading=!1}catch(a){l.loading=!1}},n=(0,u.iH)(null),m=e=>{n.value.setCharacNo(e)};return(e,i)=>{const s=(0,t.up)("el-row");return(0,t.wg)(),(0,t.iD)("div",null,[(0,t.Wm)(s,null,{default:(0,t.w5)((()=>[(0,t.Wm)((0,u.SU)(o.Z),{title:"邮件工具"})])),_:1}),(0,t._)("div",y,[(0,t._)("div",h,[(0,t.Wm)((0,u.SU)(d.Z),{num:"1",title:"账号选择"},{default:(0,t.w5)((()=>[(0,t.Wm)((0,u.SU)(r.Z),{onSetUid:a})])),_:1}),(0,t.Wm)((0,u.SU)(d.Z),{num:"2",title:"角色选择"},{default:(0,t.w5)((()=>[(0,t.Wm)((0,u.SU)(W.Z),{ref:"roleTableRef",loading:l.loading,data:l.data,onSetCharacNo:m},null,8,["loading","data"])])),_:1}),(0,t.Wm)((0,u.SU)(d.Z),{num:"3",title:"发送邮件","is-last":""},{default:(0,t.w5)((()=>[(0,t.Wm)((0,u.SU)(b),{ref_key:"emailFormRef",ref:n},null,512)])),_:1})])])])}}});const U=V;var q=U},2899:function(e,l,a){a.d(l,{Z:function(){return m}});var t=a(3396),u=a(7139),o=a(4870),d=a(5134),r=(0,t.aZ)({__name:"RoleTable",props:{data:{type:Array,required:!0,default:()=>[]},loading:{type:Boolean,required:!0,default:!1}},emits:["setCharacNo"],setup(e,{expose:l,emit:a}){const r=(0,o.iH)(null),n=()=>{a("setCharacNo",r.value)},m=()=>{r.value=null};return l({resetCharacNo:m}),(l,a)=>{const m=(0,t.up)("el-radio"),i=(0,t.up)("el-table-column"),s=(0,t.up)("el-table"),p=(0,t.Q2)("loading");return(0,t.wg)(),(0,t.iD)("div",null,[(0,t.wy)(((0,t.wg)(),(0,t.j4)(s,{data:e.data,ref:"tableRef",border:""},{default:(0,t.w5)((()=>[(0,t.Wm)(i,{width:"50",label:"#",align:"center"},{default:(0,t.w5)((e=>[(0,t.Wm)(m,{modelValue:r.value,"onUpdate:modelValue":a[0]||(a[0]=e=>r.value=e),label:e.row.charac_no,class:"radio",onChange:n},null,8,["modelValue","label"])])),_:1}),(0,t.Wm)(i,{prop:"charac_no",label:"角色编号",width:"180"}),(0,t.Wm)(i,{prop:"charac_name",label:"角色名称",width:"180"}),(0,t.Wm)(i,{prop:"create_time",label:"创建时间",width:"180"},{default:(0,t.w5)((e=>[(0,t._)("span",null,(0,u.zw)((0,o.SU)(d.vc)(e.row.create_time)),1)])),_:1}),(0,t.Wm)(i,{prop:"lev",label:"等级"}),(0,t.Wm)(i,{prop:"m_id",label:"uid",width:"150"})])),_:1},8,["data"])),[[p,e.loading]])])}}});const n=r;var m=n},612:function(e,l,a){a.d(l,{Z:function(){return p}});var t=a(3396),u=a(7139),o=a(4870),d=a(491),r=a(7192);const n={style:{float:"left"}},m={style:{float:"right",color:"var(--el-text-color-secondary)","font-size":"13px"}};var i=(0,t.aZ)({__name:"SelectAccount",emits:["setUid"],setup(e,{emit:l}){const a=(0,o.iH)(),i=(0,o.qj)({uid:null}),s=(0,o.qj)({uid:[{required:!0,message:"请选择账号",trigger:"blur"}]}),p=(0,o.qj)({data:[]}),c=(0,o.iH)(!1),f=async()=>{c.value=!0,p.data=await(0,d.T8)({}),c.value=!1},g=async()=>{const e=await(0,r.Gu)(a);e&&l("setUid",i.uid)};return f(),(e,l)=>{const o=(0,t.up)("el-option"),d=(0,t.up)("el-select"),r=(0,t.up)("el-form-item"),f=(0,t.up)("el-button"),_=(0,t.up)("el-form");return(0,t.wg)(),(0,t.iD)("div",null,[(0,t.Wm)(_,{model:i,rules:s,ref_key:"formRef",ref:a,"label-width":"100px"},{default:(0,t.w5)((()=>[(0,t.Wm)(r,{label:"账号ID",prop:"uid"},{default:(0,t.w5)((()=>[(0,t.Wm)(d,{modelValue:i.uid,"onUpdate:modelValue":l[0]||(l[0]=e=>i.uid=e),filterable:"",clearable:"",placeholder:"选择账号id",loading:c.value,"loading-text":"正在加载账号"},{default:(0,t.w5)((()=>[((0,t.wg)(!0),(0,t.iD)(t.HY,null,(0,t.Ko)(p.data,(e=>((0,t.wg)(),(0,t.j4)(o,{key:e.uid,label:e.uid,value:e.uid},{default:(0,t.w5)((()=>[(0,t._)("span",n,(0,u.zw)(e.uid),1),(0,t._)("span",m," 角色数:"+(0,u.zw)(e.roles),1)])),_:2},1032,["label","value"])))),128))])),_:1},8,["modelValue","loading"])])),_:1}),(0,t.Wm)(r,null,{default:(0,t.w5)((()=>[(0,t.Wm)(f,{type:"primary",onClick:g,size:"small"},{default:(0,t.w5)((()=>[(0,t.Uk)("查询")])),_:1})])),_:1})])),_:1},8,["model","rules"])])}}});const s=i;var p=s}}]);