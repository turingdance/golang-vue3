<template>
    <div class="flex flex-row align-center gap-1">
        <label>
            <el-checkbox v-model="modelvalue" label="注册即视同意" checked disabled value="true" size="small" />
        </label>
        <el-button type="primary" @click="renderuser()" link>洞见用户服务条款</el-button>
        <!-- <el-button type="primary" link>用户隐私政策</el-button> -->
    </div>
    <el-dialog title="请查看协议" v-model="_show" append-to-body :width="960">
        <div v-html="html" style="height: 480px;overflow-y: scroll;"></div>
        <template #footer>
            <div class="flex flex-row-reverse">
                <el-button  @click="_show=false">关闭</el-button>
                <el-button type="primary" @click="confirm">确认</el-button>
            </div>
        </template>
    </el-dialog>
</template>
<script lang="ts" setup>
import {ref} from "vue"
import userMdText from './user.md?raw'
import MarkdownIt from "markdown-it"
const _show = ref(false)
const md = new (MarkdownIt as any)();
const html = ref('')
const renderuser = (str)=>{
    html.value = md.render(userMdText)
    _show.value = true
}

const modelvalue = defineModel({
    default:false,
    type:Boolean,
})
const confirm = ()=>{
    _show.value = false;
    modelvalue.value = true
}
</script>