<template>
    <el-tag v-for="item,index in value">{{ item.label }}</el-tag>
</template>
<script setup>
const appprops = defineProps({
    value: {
      type: Array,
      requied: true
    }
  })
</script>
  