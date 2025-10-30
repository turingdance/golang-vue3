/** @type {import('tailwindcss').Config} */
module.exports = {
    // 配置需要使用tailwindcss的文件
    content: ['./index.html','./src/**/**/*.{vue,ts,js,ts,jsx,tsx}'],
    theme: {
      extend: {
        colors: {},
        fontFamily: {
            sans: ['Inter', 'system-ui', 'sans-serif'],
        },
    },
    },
    plugins: [],
  }
  
  