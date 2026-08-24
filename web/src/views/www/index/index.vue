<template>
   <main class="bg-techBg text-slate-100 bg-tech-grid font-sans pt-[70px]">
        <!-- Hero首屏 -->
        <section id="home" class="relative min-h-[calc(100vh-70px)] flex flex-col items-center justify-center overflow-hidden px-6">
            <div class="absolute w-[900px] h-[900px] border border-techPrimary/20 rounded-full"></div>
            <div class="absolute w-[700px] h-[700px] border border-techPrimary/10 rounded-full"></div>
            <div class="relative z-10 text-center max-w-5xl">
                <h1 class="text-[clamp(2rem,5vw,3.5rem)] font-bold leading-tight">
                    视频智能分析 <span class="text-gradient-blue">软硬件一体化</span>解决方案
                </h1>
                <p class="text-techGray text-lg md:text-xl mt-6 leading-relaxed">
                    我们是视频采集、解码处理、AI智能分析、视频结构化全链路服务商。<br>自研边缘推理硬件网关 + 可视化管理平台软件，支持多路视频流接入、多深度学习模型并行调度，面向工业视觉、智慧安防、园区监测、安全生产等行业交付整套可落地的视频AI解决方案。
                </p>
                <div class="mt-10 flex flex-wrap justify-center gap-5">
                    <router-link class="bg-techPrimary px-8 py-3 rounded-md hover:bg-techLight transition" to="/register">注册使用</router-link>
                    <!-- <a href="#scene" class="border border-techPrimary px-8 py-3 rounded-md hover:bg-techPrimary/10 transition">查看场景案例</a> -->
                </div>
            </div>
        </section>

        <!-- 关于我们 / 企业简介 -->
        <section id="about" class="max-w-7xl mx-auto px-6 py-24">
            <div class="text-center mb-14">
                <h2 class="text-3xl font-bold mb-3">关于我们</h2>
                <p class="text-techGray">专注视频AI技术，软硬一体，边云协同</p>
            </div>
            <div class="bg-techCard border border-techBorder rounded-2xl p-8 md:p-10">
                <p class="text-techGray text-base leading-8">
                    公司深耕机器视觉与视频智能分析领域，是集硬件研发、算法开发、平台软件定制于一体的高新技术服务商。我们自主研发边缘AI推理网关设备，搭载高性能视频解码芯片，可同时接入多路RTSP/RTMP/ONVIF摄像头与视频文件；配套可视化管理软件平台，支持多AI模型仓库管理、任务调度、实时预览、分析结果存储、告警推送与数据报表导出。
                    <br><br>
                    区别于纯算法软件厂商，我们实现硬件驱动、视频解码、模型推理、上层管理平台全栈自研，软硬件深度优化适配，大幅降低项目集成调试成本。产品广泛应用于智慧园区、工厂安全生产、工地监管、河道监测、智慧农业、仓储视觉巡检等场景，可根据客户业务需求完成算法模型定制、功能二次开发与私有化部署交付。
                </p>
            </div>
        </section>

        <!-- 硬件产品板块 -->
        <section id="hardware" class="max-w-7xl mx-auto px-6 pb-24">
            <div class="text-center mb-14">
                <h2 class="text-3xl font-bold mb-3">边缘硬件产品</h2>
                <p class="text-techGray">高性能视频AI推理网关</p>
            </div>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <svg class="w-11 h-11 text-techPrimary mb-5" viewBox="0 0 40 40" fill="none" stroke="currentColor" stroke-width="2">
                        <rect x="6" y="12" width="28" height="16" rx="3"/>
                        <line x1="20" y1="28" x2="20" y2="34"/>
                    </svg>
                    <h3 class="text-lg font-semibold mb-3">AI视频推理网关</h3>
                    <p class="text-techGray text-sm">搭载国产加速芯片，多路视频硬解码，本地运行YOLO、分割、追踪等多类模型，实现边缘侧实时分析。</p>
                </div>
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <svg class="w-11 h-11 text-techPrimary mb-5" viewBox="0 0 40 40" fill="none" stroke="currentColor" stroke-width="2">
                        <circle cx="20" cy="20" r="12"/>
                        <path d="M20 8v-4 M20 36v-4 M8 20h-4 M36 20h-4"/>
                    </svg>
                    <h3 class="text-lg font-semibold mb-3">多路视频接入设备</h3>
                    <p class="text-techGray text-sm">支持网络摄像头、NVR、视频流服务器接入，最多可并发处理数十路视频通道。</p>
                </div>
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <svg class="w-11 h-11 text-techPrimary mb-5" viewBox="0 0 40 40" fill="none" stroke="currentColor" stroke-width="2">
                        <rect x="8" y="8" width="24" height="24" rx="4"/>
                        <path d="M14 20 h12 M20 14 v12"/>
                    </svg>
                    <h3 class="text-lg font-semibold mb-3">工业级扩展终端</h3>
                    <p class="text-techGray text-sm">丰富IO接口，可外接声光告警、继电器输出，适配工厂、工地等复杂工业现场环境。</p>
                </div>
            </div>
        </section>

        <!-- 软件平台板块 -->
        <section id="software" class="max-w-7xl mx-auto px-6 pb-24">
            <div class="text-center mb-14">
                <h2 class="text-3xl font-bold mb-3">软件管理平台</h2>
                <p class="text-techGray">设备、模型、任务、数据一体化管控</p>
            </div>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <h3 class="text-lg font-semibold mb-3">设备集中管理</h3>
                    <p class="text-techGray text-sm">网关设备远程注册、状态监控、在线升级、通道配置、视频预览，实现海量边缘设备统一运维。</p>
                </div>
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <h3 class="text-lg font-semibold mb-3">AI模型仓库管理</h3>
                    <p class="text-techGray text-sm">模型上传、版本管理、参数配置、任务下发；支持同一视频通道加载多个算法模型并行分析。</p>
                </div>
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <h3 class="text-lg font-semibold mb-3">实时告警与事件推送</h3>
                    <p class="text-techGray text-sm">识别到违规事件后自动抓拍图片、保存录像片段，支持WEB弹窗、消息推送等多种告警方式。</p>
                </div>
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <h3 class="text-lg font-semibold mb-3">数据统计与报表导出</h3>
                    <p class="text-techGray text-sm">智能分析结果持久化存储，提供可视化统计看板，支持按时间段导出Excel报表，便于事后追溯。</p>
                </div>
            </div>
        </section>

        <!-- 核心优势板块 -->
        <section id="advantage" class="max-w-7xl mx-auto px-6 pb-24">
            <div class="text-center mb-14">
                <h2 class="text-3xl font-bold mb-3">核心优势</h2>
                <p class="text-techGray">软硬一体 · 高性能 · 易部署 · 可定制</p>
            </div>
            <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4 gap-6">
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <svg class="w-11 h-11 text-techPrimary mb-5" viewBox="0 0 40 40" fill="none" stroke="currentColor" stroke-width="2">
                        <rect x="5" y="8" width="30" height="24" rx="3"/>
                        <path d="M10,20 L30,20"/>
                    </svg>
                    <h3 class="text-lg font-semibold mb-3">软硬深度协同适配</h3>
                    <p class="text-techGray text-sm">自研边缘硬件与软件系统深度耦合，降低兼容性风险，方案整体运行更加稳定可靠。</p>
                </div>
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <svg class="w-11 h-11 text-techPrimary mb-5" viewBox="0 0 40 40" fill="none" stroke="currentColor" stroke-width="2">
                        <circle cx="20" cy="20" r="14"/>
                        <circle cx="20" cy="20" r="6"/>
                    </svg>
                    <h3 class="text-lg font-semibold mb-3">多AI模型灵活调度</h3>
                    <p class="text-techGray text-sm">支持同时部署、切换、管理多种视觉算法模型，按需开启检测、追踪、分割任务。</p>
                </div>
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <svg class="w-11 h-11 text-techPrimary mb-5" viewBox="0 0 40 40" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M8 30 L20 10 L32 30"/>
                        <circle cx="20" cy="24" r="4"/>
                    </svg>
                    <h3 class="text-lg font-semibold mb-3">多路视频实时处理</h3>
                    <p class="text-techGray text-sm">并发接入多路视频流，边缘本地完成解码推理，减少视频上传带宽压力，响应速度快。</p>
                </div>
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <svg class="w-11 h-11 text-techPrimary mb-5" viewBox="0 0 40 40" fill="none" stroke="currentColor" stroke-width="2">
                        <rect x="7" y="7" width="11" height="11" rx="2"/>
                        <rect x="22" y="7" width="11" height="11" rx="2"/>
                        <rect x="7" y="22" width="11" height="11" rx="2"/>
                        <rect x="22" y="22" width="11" height="11" rx="2"/>
                    </svg>
                    <h3 class="text-lg font-semibold mb-3">行业场景快速定制</h3>
                    <p class="text-techGray text-sm">模块化软硬件架构，可快速针对工业检测、安全生产、智慧园区等项目定制开发。</p>
                </div>
            </div>
        </section>

        <!-- 应用场景板块 -->
        <section id="scene" class="max-w-7xl mx-auto px-6 pb-24">
            <div class="text-center mb-14">
                <h2 class="text-3xl font-bold mb-3">行业应用场景</h2>
                <p class="text-techGray">多行业落地，满足不同视频智能监管需求</p>
            </div>
            <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <h3 class="text-lg font-semibold mb-3">智慧园区安防</h3>
                    <p class="text-techGray text-sm">区域入侵检测、人员越界、徘徊检测、烟火识别，实现园区全天候智能安防。</p>
                </div>
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <h3 class="text-lg font-semibold mb-3">工厂安全生产监管</h3>
                    <p class="text-techGray text-sm">安全帽、防护服佩戴识别、危险区域闯入、离岗睡岗、违规烟火行为智能识别告警。</p>
                </div>
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <h3 class="text-lg font-semibold mb-3">建筑工地智能监管</h3>
                    <p class="text-techGray text-sm">人员实名制识别、渣土车监管、裸土覆盖检测、高空危险作业行为监测。</p>
                </div>
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <h3 class="text-lg font-semibold mb-3">智慧农业监测</h3>
                    <p class="text-techGray text-sm">大棚作物状态监控、野生动物识别、水库河道水位、漂浮物、垃圾巡检识别。</p>
                </div>
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <h3 class="text-lg font-semibold mb-3">仓储物流巡检</h3>
                    <p class="text-techGray text-sm">货物堆积异常、通道占用、人员违规进入、叉车运行安全智能检测。</p>
                </div>
                <div class="bg-techCard border border-techBorder rounded-xl p-7 card-lift">
                    <h3 class="text-lg font-semibold mb-3">私有化项目部署</h3>
                    <p class="text-techGray text-sm">支持内网离线部署，数据本地留存，满足政府、国企等项目的数据安全合规要求。</p>
                </div>
            </div>
        </section>
    </main>

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