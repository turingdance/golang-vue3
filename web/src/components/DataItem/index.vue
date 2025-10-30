<template>
  <div class="dataitem">
    <div class="data">{{ count }}</div>
    <label class="title">{{ title }}</label>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from "vue"
import request from "@/utils/request"
export default defineComponent({
  props: {
    title: {
      type: String
    },
    server: {
      type: String
    },
    cond: {
      type: Object
    }
  },
  setup(props) {
    const count = ref(0)
    const loaddata = () => {
      request({
        url: props.server,
        method: "post",
        data: props.cond
      }).then((res) => {
        count.value = res.data
      })
    }
    onMounted(() => {
      //console.log("props.cond", props.cond)
      loaddata()
    })
    //watch(props,()=>loaddata())
    return {
      count
    }
  }
})
</script>
<style lang="scss">
.dataitem {
  display: flex;
  flex-direction: column;
  width: 100px;
  margin: 10px;
  .title {
    font-size: 15px;
    line-height: 30px;
  }
  .data {
    font-size: 14px;
  }
}
</style>
