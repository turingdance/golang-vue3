<template>
   <!-- 主横幅区域 -->
   <section class="bg-gradient-blue text-white py-8 md:py-12 relative overflow-hidden">
        <div class="absolute inset-0 opacity-10">
            <div class="absolute top-0 left-0 w-full h-full bg-[url('https://picsum.photos/1200/400')] bg-cover bg-center"></div>
        </div>
        <div class="container mx-auto px-4 sm:px-6 lg:px-8 relative z-10">
            <div class="flex flex-row md:flex-row md:items-center">
                <div class="md:w-2/3 mb-6 md:mb-0">
                    <h1 class="text-[clamp(1.5rem,3vw,2.5rem)] font-bold mb-3 text-shadow">金属激光切割加工虚拟仿真实训</h1>
                    <div class="flex flex-wrap items-center gap-4 mb-4">
                        <span class="bg-secondary text-white text-sm px-3 py-1 rounded-full" v-if="false">国家级一流课程</span>
                        <div class="flex items-center text-sm">
                            <i class="fa fa-user mr-1"></i> 张灵 | 湖南工业大学
                        </div>
                        <div class="flex items-center text-sm">
                            <i class="fa fa-calendar mr-1"></i> 上线时间: 2025年
                        </div>
                    </div>

                    <div class="grid grid-cols-4 md:grid-cols-5 gap-3 md:gap-5 mb-6">
                        <div class="bg-white/10 p-3 rounded-lg">
                            <div class="text-xs opacity-80">所属专业类</div>
                            <div class="font-medium">金工实习</div>
                        </div>
                        <div class="bg-white/10 p-3 rounded-lg">
                            <div class="text-xs opacity-80">对应专业</div>
                            <div class="font-medium">机械设计制造及其自动化
                            </div>
                        </div>
                        <div class="bg-white/10 p-3 rounded-lg">
                            <div class="text-xs opacity-80">课程类型</div>
                            <div class="font-medium">专业基础课</div>
                        </div>
                        <div class="bg-white/10 p-3 rounded-lg">
                            <div class="text-xs opacity-80">实验类型</div>
                            <div class="font-medium">综合设计型</div>
                        </div>
                        <div class="bg-white/10 p-3 rounded-lg hidden md:block">
                            <div class="text-xs opacity-80">所属课程</div>
                            <div class="font-medium">金工实习</div>
                        </div>
                    </div>
                    
                    <p class="text-white/90 mb-6 max-w-3xl">
                        钣金类产品加工虚拟仿真实验教学项目以本科工程训练基础课中冲、锻、数控加工教学模块为对象，利用数字化技术建立虚拟车间、加工设备及加工过程数字化学境，学生可以在学习、操作机床、制作产品。实验以一种典型的电柜为载体，将工程识图、零件加工工艺、实训基础知识和操作有机结合，通过设置不同知识点和考核点，考察所学知识的运用能力，使学生了解产品的制造工艺，培养学生的工程素养，为学生的后续课程学习和创新奠定基础。
                    </p>
                    
                    <div class="flex flex-wrap gap-4">
                        <button @click="learnit" class="bg-secondary hover:bg-secondary/90 text-white px-6 py-3 rounded-lg font-medium transition-all shadow-lg hover:shadow-xl flex items-center">
                            <i class="fa fa-play-circle mr-2"></i> 我要做实验
                        </button>
                        <div class="flex items-center bg-white/100 px-4 py-2 rounded-lg">
                            <span class="flex items-center mr-4">
                                <i class="fa fa-eye mr-1"></i> 2623
                            </span>
                            <span class="flex items-center mr-4">
                                <i class="fa fa-heart mr-1"></i> 2339
                            </span>
                            <span class="flex items-center">
                                <i class="fa fa-share-alt mr-1"></i> 分享
                            </span>
                        </div>
                    </div>
                </div>
                
                <div class="md:w-1/3">
                    <div class="relative rounded-xl overflow-hidden shadow-2xl group ml-2">
                        <video src="/about.mp4" class="w-full h-auto object-cover rounded-xl transition-transform duration-500 group-hover:scale-105" controls autoplay="false"></video>
                    </div>
                </div>
            </div>
        </div>
    </section>

    <!-- 评分和统计数据 -->
    <section class="py-8 bg-white">
        <div class="container mx-auto px-4 sm:px-6 lg:px-8">
            <div class="flex flex-row lg:flex-row gap-8">
                <div class="lg:w-2/3">
                    <div class="flex items-center mb-6">
                        <div class="text-4xl font-bold text-primary mr-3">4.8</div>
                        <div class="flex text-yellow-400">
                            <i class="fa fa-star"></i>
                            <i class="fa fa-star"></i>
                            <i class="fa fa-star"></i>
                            <i class="fa fa-star"></i>
                            <i class="fa fa-star-half-o"></i>
                        </div>
                        <div class="text-gray-500 ml-3">3.25万人评价</div>
                        <a href="#" v-if="false" class="ml-4 text-primary hover:text-darkblue transition-colors text-sm font-medium">去评价 <i class="fa fa-angle-right ml-1"></i></a>
                    </div>
                    
                    <div class="bg-neutral rounded-xl p-6 mb-8">
                        <div class="flex flex-wrap justify-between gap-4 md:gap-8">
                            <div class="text-center">
                                <div class="text-sm text-gray-500 mb-1">浏览量</div>
                                <div class="text-2xl font-bold text-primary">133,199</div>
                            </div>
                            <div class="text-center">
                                <div class="text-sm text-gray-500 mb-1">实验人次</div>
                                <div class="text-2xl font-bold text-primary">331,94</div>
                            </div>
                            <div class="text-center">
                                <div class="text-sm text-gray-500 mb-1">实做人数</div>
                                <div class="text-2xl font-bold text-primary">10,574</div>
                            </div>
                            <div class="text-center">
                                <div class="text-sm text-gray-500 mb-1">实验平均用时</div>
                                <div class="text-2xl font-bold text-primary">94'</div>
                            </div>
                            <div class="text-center">
                                <div class="text-sm text-gray-500 mb-1">实验完成率</div>
                                <div class="text-2xl font-bold text-primary">99%</div>
                            </div>
                            <div class="text-center">
                                <div class="text-sm text-gray-500 mb-1">实验通过率</div>
                                <div class="text-2xl font-bold text-primary">97.5%</div>
                            </div>
                            <div class="text-center hidden md:block">
                                <div class="text-sm text-gray-500 mb-1">接入率</div>
                                <div class="text-2xl font-bold text-primary">98.5%</div>
                            </div>
                        </div>
                        <div class="mt-4 text-sm text-gray-500 text-center">
                            共享应用 <span class="mx-2">|</span> 每日0点更新
                        </div>
                    </div>
                    
                    <!-- 实验介绍标签页 -->
                    <div class="mb-8">
                        <div class="border-b border-gray-200">
                            <ul class="flex flex-wrap -mb-px">
                                <li class="mr-2" role="presentation" v-for="item in groupList">
                                    <button @click="chosegroup(item.index)" :class="groupIndex==item.index?btngroupchosed:btngroupnotchosed"  type="button" role="tab" aria-selected="true">
                                        {{ item.title }}
                                    </button>
                                </li>
                            </ul>
                        </div>
                        
                        <!-- 实验介绍内容 -->
                        <div class="tab-panel p-6 bg-white" role="tabpanel" v-if="groupIndex=='jieshao'">
                            <div class="border-b border-gray-200 pb-3 mb-4" >
                                <ul class="flex flex-wrap gap-2 md:gap-4" id="subTabList">
                                    <li v-for="item in tabList.filter(oo=>oo.group==1)">
                                        <button @click="choseindex(item.index)" :class="currentindex==item.index?btnchosed:btnnotchosed" >
                                            {{item.title}}
                                        </button>
                                    </li>
                                </ul>
                            </div>
                            
                            <!-- 实验教学目标 -->
                            <div  class="sub-tab-panel" v-if="currentindex=='mubiao'">
                             <p>   该虚拟仿真实验教学旨在通过数字化实验平台，将钣金加工过程中所涉及到的机床设备软件化，借助虚拟仿真系统进行实验操作，引导学生通过计算机最大程度地模拟激光切割加工的完整过程。本实验以“工程训练大赛-新能源车”典型零件作为加工对象，在加工过程中考察学生安全生产、制造工艺等理论知识的应用能力，使学生掌握特种加工技术之一——金属激光切割的加工方法与工艺特点、制造过程及相关设备的操作。</p>
                             <p>该虚拟仿真实验项目可作为全校工科类基础实践课程《金工实习》中的一个实训项目，也可支撑机械设计制造及其自动化专业课程《精密加工技术》开展实验。</p>

                            </div>
                            
                            <!-- 其他子标签内容（默认隐藏） -->
                            <div  class="sub-tab-panel" v-if="currentindex=='yuanli'">
                                <p>金属激光切割加工虚拟仿真实训，以“工程训练大赛-新能源车”典型零件作为加工对象，要求学生掌握激光切割原理、加工方法和加工工艺特点、制造过程及激光切割机设备的操作。
进入虚拟实验系统后，在默认的工厂环境下，根据任务路线导航：首先进入课程理论知识的学习，包括激光切割原理、设备组成介绍、影响因素、应用情况等；在进入实操界面后，学习金属激光切割机安全操作规程，了解安全生产知识；从产品库中选定加工对象，确认正确的加工图纸，核对给定的加工工艺路线，做好加工准备；根据工艺路线，按照设备操作规程和步骤进行参数设置，完成零件的虚拟加工训练；完成前期实验环节后，再进行产品检验，最终完成零件—组件—产品的虚拟仿真训练； 系统对学生整个实验过程的操作痕迹进行记录，给出实验评价。
根据实验流程，依次涉及到机械制图、机械加工工艺、激光切割等理论和实践环节。
知识点：共4个
（1）智能制造和激光切割相关的常规主流应用软件（CAD/ CypCut）
（2）常见表面加工方案，包括尺寸精度、材料及工艺性
（3）特种加工——激光加工的分类、方法及工艺特点
（4）激光切割的原理、设备操作使用	
金属激光切割加工虚拟仿真实训中的设备模型与光纤激光金属切割机实物一致、加工零件选用我校本科生参加“工程训练大赛-新能源车”优秀作品中的典型加工零件（包括主动轮、从动轮和底板），1:1还原设计图纸。学生通过虚拟仿真实验，即可身临其境式地面向实际加工进行机床仿真操作，解决了传统实践教学受设备台数、时间、地点及耗费材料等制约的难题，同时可充分保证学生学习体验，激发学生学习兴趣和热情。
</p>
                            </div>
                            <div id="methods" class="sub-tab-panel" v-if="currentindex=='fangfa'">
                                <p>① 数据库建设：根据教学目标要求，选择典型的产品零件案例，完成零件二维图纸设计和工艺流程分析，以及工厂环境、机床设备的三维模型建立，导入虚拟仿真实验教学系统中。
② 知识点预习：鼓励学生课前登录虚拟仿真实验教学系统，通过虚拟实验系统中的文字、视频和虚拟交互等方式自学；同时根据本教学团队编写的教材《工程实训教程》自主学习相关理论知识，坚持线上和线下结合开展课前预习。
③ 虚实结合操作：教师团队讲解并示范指导操作后，学生进行虚拟实验。项目采用图形用户界面和对话框驱动，沉浸感强、导航明确，通过标准化步骤，学生可以反复练习以强化记忆。
④ 个体化评估：本实验要求学生个体完成，不进行分组考核。系统可以提供实验过程中操作环节和理论知识点考核和评价，并对误操作提示。
⑤ 课后复习重温：通过虚拟操作训练，学生能很快地熟悉实际设备操作、了解产品加工过程、掌握相关的知识点。今后开展学科竞赛或者创新制作需要对实物进行加工，可在课后登录虚拟仿真系统，重温熟悉实验过程。
</p>
                            </div>
                            <div id="steps" class="sub-tab-panel" v-if="currentindex=='buzhou'">
                                <p><a href="/buzhou.pdf" target="_blank"> <el-image src="/images/buzhou.png"   /></a></p>
                            </div>
                            <div id="results" class="sub-tab-panel" v-if="currentindex=='yaoqiu'">
                                <p>专业与年级要求
实验项目面向参加金工实习A/B的机械设计制造及其自动化、机械工程、工业设计（理）、金属材料、冶金工程、测控技术等工科类学生，年级为大学二年级，根据学分和周数安排合理安排本实验环节的时间。
基本知识和能力要求
学生要掌握机械制图、机械原理、计算机辅助设计、机械制造技术基础等课程的相关知识和理论。学生要求具备计算机的基本操作能力，能够读懂零部件工程图纸，能够较熟练地运用Creo、Solidworks、AutoCAD等绘图软件进行模型构建等。
</p>
                            </div>
                        </div>
                        
                        <!-- 实验必读内容（默认隐藏） -->
                        <div class="tab-panel p-6 bg-white" role="tabpanel" v-if="groupIndex=='bidu'">
                            <p><a href="/buzhou.pdf" target="_blank"> <el-image src="/images/buzhou.png"   /></a></p>
                        </div>
                    </div>
                </div>
                
                <!-- 右侧边栏 -->
                <div class="lg:w-1/3">
                    <div class="space-y-6">
                        <!-- 成绩分布 -->
                        <div class="bg-white rounded-xl shadow-sm p-6 card-hover" v-if="false">
                            <h3 class="text-lg font-semibold mb-4 flex items-center">
                                <i class="fa fa-pie-chart text-primary mr-2"></i> 成绩分布
                            </h3>
                            <div class="flex justify-center mb-4">
                                <canvas id="scoreChart" width="200" height="200"></canvas>
                            </div>
                            <div class="space-y-2 text-sm">
                                <div class="flex items-center">
                                    <span class="w-3 h-3 bg-primary rounded-full mr-2"></span>
                                    <span>优秀: 45.31%</span>
                                </div>
                                <div class="flex items-center">
                                    <span class="w-3 h-3 bg-secondary rounded-full mr-2"></span>
                                    <span>良好: 52.16%</span>
                                </div>
                                <div class="flex items-center">
                                    <span class="w-3 h-3 bg-red-500 rounded-full mr-2"></span>
                                    <span>不及格: 2.54%</span>
                                </div>
                            </div>
                        </div>
                        
                        <!-- 实验系统浏览器要求 -->
                        <div class="bg-white rounded-xl shadow-sm p-6 card-hover">
                            <h3 class="text-lg font-semibold mb-4 flex items-center" style="color:#0a3d91">
                                <i class="fa fa-browser text-primary mr-2"></i> 实验系统浏览器要求
                            </h3>
                            <div class="flex items-center p-3 bg-lightblue rounded-lg mb-4">
                                <img src="https://picsum.photos/50/50" alt="浏览器图标" class="w-10 h-10 mr-3">
                                <div>
                                    <div class="font-medium color-#0a3d91" style="color:#0a3d91">推荐浏览器</div>
                                    <div class="text-sm text-gray-600">Chrome 90.0及以上版本</div>
                                </div>
                            </div>
                            <ul class="space-y-2 text-sm">
                                <li class="flex items-start">
                                    <i class="fa fa-check-circle text-green-500 mt-1 mr-2"></i>
                                    <span>支持WebGL 1.0及以上版本</span>
                                </li>
                                <li class="flex items-start">
                                    <i class="fa fa-check-circle text-green-500 mt-1 mr-2"></i>
                                    <span>浏览器需开启JavaScript功能</span>
                                </li>
                                <li class="flex items-start">
                                    <i class="fa fa-check-circle text-green-500 mt-1 mr-2"></i>
                                    <span>建议屏幕分辨率1920×1080及以上</span>
                                </li>
                                <li class="flex items-start">
                                    <i class="fa fa-check-circle text-green-500 mt-1 mr-2"></i>
                                    <span>网络环境要求：带宽≥10Mbps</span>
                                </li>
                            </ul>
                        </div>
                        
                        <!-- 教学支持热线 -->
                        <div class="bg-white rounded-xl shadow-sm p-6 card-hover">
                            <h3 class="text-lg font-semibold mb-4 flex items-center color-#0a3d91" style="color:#0a3d91">
                                <i class="fa fa-phone text-primary mr-2"></i> 教学支持热线
                            </h3>
                            <div class="bg-lightblue rounded-lg p-4 text-center">
                                <div class="text-2xl font-bold text-primary mb-2">0731-22183137</div>
                                <div class="text-sm text-gray-600">工作日 9:00-17:00</div>
                            </div>
                            <div class="mt-4 text-center">
                                <a href="#" class="text-primary hover:text-darkblue transition-colors text-sm">
                                    <i class="fa fa-envelope mr-1"></i> 发送邮件咨询
                                </a>
                            </div>
                        </div>
                        
                        <!-- 相关实验推荐 -->
                        <div class="bg-white rounded-xl shadow-sm p-6 card-hover" v-if="false">
                            <h3 class="text-lg font-semibold mb-4 flex items-center">
                                <i class="fa fa-th-large text-primary mr-2"></i> 相关实验推荐
                            </h3>
                            <div class="space-y-3">
                                <a href="#" class="block p-3 border border-gray-100 rounded-lg hover:border-primary hover:bg-lightblue/50 transition-colors">
                                    <div class="font-medium text-gray-800">数控车床虚拟仿真实验</div>
                                    <div class="text-xs text-gray-500 mt-1">清华大学 | 4.7分</div>
                                </a>
                                <a href="#" class="block p-3 border border-gray-100 rounded-lg hover:border-primary hover:bg-lightblue/50 transition-colors">
                                    <div class="font-medium text-gray-800">焊接工艺虚拟仿真训练</div>
                                    <div class="text-xs text-gray-500 mt-1">哈尔滨工业大学 | 4.6分</div>
                                </a>
                                <a href="#" class="block p-3 border border-gray-100 rounded-lg hover:border-primary hover:bg-lightblue/50 transition-colors">
                                    <div class="font-medium text-gray-800">3D打印技术虚拟实验</div>
                                    <div class="text-xs text-gray-500 mt-1">上海交通大学 | 4.9分</div>
                                </a>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>


</template>
<script setup>
    import { useUserStore } from "@/store/modules/user"
    import {showlogin} from "@/utils/index"
    import {ref} from "vue"

    import {   useRouter } from "vue-router";
    
    const btnchosed = ref('sub-tab active px-3 py-1 rounded-full bg-primary text-white text-sm')
    const btnnotchosed = ref('sub-tab px-3 py-1 rounded-full bg-gray-100 text-gray-700 text-sm hover:bg-gray-200')
    const btngroupchosed = ref('inline-block p-4 border-b-2 border-primary text-primary font-medium')
    const btngroupnotchosed = ref('inline-block p-4 border-b-2 border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300')
    const currentindex = ref('mubiao')
    const groupIndex = ref('jieshao')
    const  contextStore = useUserStore()
    const groupList = ref([
        {index:'jieshao',title:'实验介绍',group:1},
        {index:'bidu',title:'实验必读',group:1},
    ])
    const tabList  = ref([
        {index:'mubiao',title:'实验教学目标',group:1},
        {index:'yuanli',title:'实验原理',group:1},
        {index:'fangfa',title:'实验教学过程与实验方法',group:1},
        {index:'buzhou',title:'实验步骤',group:1},
        {index:'yaoqiu',title:'面向学生要求',group:1},
    ])

    function learnit(){
        console.log('contextStore.haslogin',contextStore.haslogin)
        if(!contextStore.haslogin){
            showlogin()
        }else{
            router.push("/learn")
        }
    }
    function choseindex(index){
        currentindex.value = index
    }

    function chosegroup(index){
        groupIndex.value = index
    }


   
const router = useRouter()
    </script>

    <style scoped scss>
span{
    color: #0a3d91;
}
.tab-panel{
    color:#1f2937;
}
</style>