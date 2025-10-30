<template>
    <header class="bg-white shadow-sm sticky top-0 z-50 transition-all duration-300">
        <div class="container mx-auto px-4 sm:px-6 lg:px-8">
            <div class="flex justify-between items-center py-3 md:py-4">
                <div class="flex items-center">
                    <a href="/" class="flex items-center">
                        <img src="https://picsum.photos/40/40" alt="实验空间 Logo" class="h-8 w-auto mr-2">
                        <span class="text-primary font-bold text-lg">实验空间</span>
                        <span class="text-xs text-gray-500 ml-2">虚拟仿真教学课程平台</span>
                    </a>
                </div>
                
                <nav class="md:flex items-center space-x-6">
                    <BtnLogin class="text-gray-600 hover:text-primary transition-colors"></BtnLogin>
                </nav>
                
                <!-- <div class="md:hidden">
                    <button type="button" class="text-gray-600 hover:text-primary">
                        <i class="fa fa-bars text-xl"></i>
                    </button>
                </div> -->
            </div>
        </div>
    </header>
       <router-view v-slot="{ Component }">
        <keep-alive>
            <component :is="Component" />
        </keep-alive>
        </router-view>
         <!-- 底部区域 -->
    <footer class="bg-darkblue text-white py-12">
        <div class="container mx-auto px-4 sm:px-6 lg:px-8">
            <div class="grid grid-cols-4 md:grid-cols-4 gap-8">
                <div>
                    <h3 class="text-lg font-bold mb-4">实验空间</h3>
                    <p class="text-gray-300 text-sm">
                        虚拟仿真教学课程平台，致力于为高校师生提供高质量的虚拟仿真实验教学资源。
                    </p>
                </div>
                <div>
                    <h4 class="font-semibold mb-4">快速链接</h4>
                    <ul class="space-y-2 text-sm text-gray-300">
                        <li><a href="#" class="hover:text-white transition-colors">首页</a></li>
                        <li><a href="#" class="hover:text-white transition-colors">实验中心</a></li>
                        <li><a href="#" class="hover:text-white transition-colors">智能实验室</a></li>
                        <li><a href="#" class="hover:text-white transition-colors">数字教材</a></li>
                    </ul>
                </div>
                <div>
                    <h4 class="font-semibold mb-4">支持与服务</h4>
                    <ul class="space-y-2 text-sm text-gray-300">
                        <li><a href="#" class="hover:text-white transition-colors">帮助中心</a></li>
                        <li><a href="#" class="hover:text-white transition-colors">联系我们</a></li>
                        <li><a href="#" class="hover:text-white transition-colors">关于我们</a></li>
                        <li><a href="#" class="hover:text-white transition-colors">用户反馈</a></li>
                    </ul>
                </div>
                <div>
                    <h4 class="font-semibold mb-4">关注我们</h4>
                    <div class="flex space-x-4 mb-4">
                        <a href="#" class="w-10 h-10 rounded-full bg-white/10 flex items-center justify-center hover:bg-white/20 transition-colors">
                            <i class="fa fa-weixin text-xl"></i>
                        </a>
                        <a href="#" class="w-10 h-10 rounded-full bg-white/10 flex items-center justify-center hover:bg-white/20 transition-colors">
                            <i class="fa fa-weibo text-xl"></i>
                        </a>
                        <a href="#" class="w-10 h-10 rounded-full bg-white/10 flex items-center justify-center hover:bg-white/20 transition-colors">
                            <i class="fa fa-youtube-play text-xl"></i>
                        </a>
                    </div>
                    <div class="text-sm text-gray-300">
                        <div class="flex items-center mb-2">
                            <i class="fa fa-phone mr-2"></i> 
                        </div>
                        <div class="flex items-center">
                            <i class="fa fa-envelope mr-2"></i> 
                        </div>
                    </div>
                </div>
            </div>
            <!-- <div class="border-t border-white/10 mt-8 pt-8 text-center text-sm text-gray-400">
                <p>© 2025 实验空间-X.com 版权所有 | 国家级虚拟仿真教学课程平台</p>
            </div> -->
        </div>
    </footer>
      </template>