<script lang="ts" setup>
import { ref, watch } from "vue"
import Register from "./components/register/index.vue"
import Login from "./components/login/index.vue"
import ResetPwd from "./components/resetpwd/index.vue"
import { useRoute } from "vue-router"
const componentName = ref<string>("login")
const route = useRoute()
watch(
  () => route.name,
  (v) => {
    componentName.value = v as string
    console.log('route.name',v)
  },
  {
    immediate: true
  }
)
</script>
<template>
<div class="page-container">
        <div class="main">
            <div class="main-container">
                <div class="left">
                    <div style="width: 40vw;">
                        <div class="grid-bg"></div>
                        <div class="glow-ball glow-ball-1"></div>
                        <div class="glow-ball glow-ball-2"></div>

                        <div class="slogan-wrap">
                            <h1>洞见</h1>
                            <p>AI视频智能分析平台</p>
                        </div>

                        <div class="feature-list">
                            <div class="feature-item">
                                <div class="feature-icon">
                                    <i class="fa fa-eye"></i>
                                </div>
                                <div class="feature-text">
                                    <h3>智能视频识别</h3>
                                    <p>AI视觉算法，自动识别画面目标与异常事件</p>
                                </div>
                            </div>
                            <div class="feature-item">
                                <div class="feature-icon">
                                    <i class="fa fa-bell"></i>
                                </div>
                                <div class="feature-text">
                                    <h3>实时告警推送</h3>
                                    <p>异常瞬间预警，快速响应处置</p>
                                </div>
                            </div>
                            <div class="feature-item">
                                <div class="feature-icon">
                                    <i class="fa fa-database"></i>
                                </div>
                                <div class="feature-text">
                                    <h3>数据分析回溯</h3>
                                    <p>视频数据结构化存储，事件可查可追溯</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="right">
                  <Register v-if="componentName == 'register'" />
                  <ResetPwd v-else-if="componentName == 'resetpwd'" />
                  <Login v-else />
                </div>
            </div>
        </div>
</div>
</template>

<style lang="scss" scoped>
        .page-container {
            width: 100%;
            overflow: hidden;
            background-color: #0b0f1a;
            display: flex;
            flex:1;
            background-image: radial-gradient(rgba(14, 165, 233, 0.08) 1px, transparent 1px);
        }
        .main {
            width: 100%;
            height: 100%;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        .main-container {
            width: 100%;
            height: 100%;
            display: flex;
        }

        /* 左侧 40vw */
        .left {
            width: 40vw;
            height: 100%;
            position: relative;
            overflow: hidden;
            background-color: #0b0f1a;
        }
        .left > div{
            width: 40vw;
            height: 100%;
            position: relative;
            padding: 0 60px;
            display: flex;
            flex-direction: column;
            justify-content: center;
            z-index: 2;
        }
        /* 网格背景 */
        .grid-bg{
            position:absolute;
            inset:0;
            background-image: radial-gradient(rgba(14, 165, 233, 0.08) 1px, transparent 1px);
            background-size:30px 30px;
            z-index:1;
        }
        /* 光晕光球 */
        .glow-ball{
            position:absolute;
            width:320px;
            height:320px;
            border-radius:50%;
            filter: blur(80px);
            opacity: 0.22;
            z-index: 1;
        }
        .glow-ball-1{
            background:#0ea5e9;
            top:10%;
            left:5%;
        }
        .glow-ball-2{
            background:#22d3ee;
            bottom:5%;
            right:5%;
        }
        .slogan-wrap h1{
            color:#fff;
            font-size:42px;
            margin-bottom:14px;
        }
        .slogan-wrap p{
            color:#94a3b8;
            font-size:18px;
            margin-bottom: 50px;
            line-height: 1.7;
        }
        /* 特性列表 */
        .feature-item{
            display:flex;
            gap:18px;
            margin-bottom:28px;
        }
        .feature-icon{
            width:46px;
            height:46px;
            border-radius:10px;
            flex-shrink: 0;
            background: rgba(14, 165, 233, 0.15);
            border: 1px solid rgba(14, 165, 233, 0.2);
            color:#22d3ee;
            display:flex;
            align-items:center;
            justify-content:center;
            font-size:20px;
        }
        .feature-text h3{
            color:#ffffff;
            font-size:18px;
            margin-bottom:6px;
        }
        .feature-text p{
            color:#8492a6;
            font-size:14px;
            line-height: 1.7;
        }

        /* 右侧表单区（更深底色） */
        .right {
            flex: 1;
            height: 100%;
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: flex-start;
            background-color: #070a13;
        }
        .login-card{
            width:420px;
            padding:24px 36px;
            margin-top: 20px;
            min-height: calc( 100% - 40px );
            border-radius:16px;
            background:rgba(17,22,34,0.7);
            backdrop-filter: blur(12px);
            border: 1px solid rgba(14,165,233,0.15);
            display: flex;
            flex-direction: column;
        }
        ::v-deep(.login-card){
          h2,h3,h4,h1{
            text-align:center;
            color:#fff;
          }
          .box-card{
            display: flex;
            gap: 5px;
            width: 100%;
          }
        }
        /* 小屏：992px以下隐藏左侧，右侧撑满全屏 */
        @media (max-width:992px){
            .left{
                display: none;
            }
            .right{
                width: 100%;
            }
            .login-card{
                width: calc(100% - 40px);
                max-width: 400px;
            }
        }
</style>
