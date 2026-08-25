<template>
    <header class="bg-white shadow-sm sticky top-0 z-50 transition-all duration-300">
        <div class="container mx-auto px-4 sm:px-6 lg:px-8">
            <div class="flex justify-between items-center py-3 md:py-4">
                <div class="flex items-center">
                    <a href="/" class="flex items-center">
                        <img src="@/assets/images/logo.png" alt="don Logo" class="h-8 w-auto mr-2 w-[48px] h-[48px]">
                        <span class="text-primary font-bold text-lg">InsightHub</span>
                        <span class="text-xs text-gray-500 ml-2">洞见边缘视觉分析解决方案</span>
                    </a>
                </div>
                
                <nav class="md:flex items-center space-x-6">

                    <router-link to="/login" v-if="!userStore.haslogin">
                        <el-button link>登陆</el-button>
                    </router-link>
                    <router-link to="/login" v-else>
                        <el-button link>管理后台</el-button>
                    </router-link>
                </nav>
            </div>
        </div>
    </header>
       <router-view v-slot="{ Component }">
        <keep-alive>
            <component :is="Component" />
        </keep-alive>
        </router-view>
         <!-- 底部区域 -->
    

         <footer class="bg-[#070a13] text-slate-200 py-12">
    <div class="container mx-auto px-4 sm:px-6 lg:px-8">
        <div >
            
            <div class="border-t border-white/10 mt-8 pt-8 text-center text-sm text-slate-500">
            <p class="mb-2">
                <a href="#" class="hover:text-cyan-300 transition-colors">用户使用协议</a>
                &nbsp;|&nbsp;
                <a href="#" class="hover:text-cyan-300 transition-colors">隐私政策</a>
                &nbsp;|&nbsp;
                <a href="#" class="hover:text-cyan-300 transition-colors">第三方账号绑定协议</a>
            </p>
            <p>© 2026 湖南云立方智能科技有限公司 版权所有 | InsightHub AI视频分析平台</p>
        </div>

        </div>
        
    </div>
</footer>

      </template>
<script setup lang="ts">
import { useUserStore } from '@/store';
const userStore = useUserStore()
</script>