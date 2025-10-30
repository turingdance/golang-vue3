<template>
  <div class="footer" >
    <div class="foot-container">
      <div class="row">
        <div class="col">
          <div class="title">{{ configStore.sitemap['site.orgname'] }}</div>
          <ul class="items">
            <li>营业执照：{{ configStore.sitemap['site.orgno'] }}</li>
            <li>备案号码：<a href="http://beian.miit.gov.cn/" target="_blank">{{ configStore.sitemap['site.beianno'] }}</a></li>
            <li>地址：{{ configStore.sitemap['site.orgaddr'] }}</li>
          </ul>
        </div>
        <div class="col">
          <div class="title">联系我们</div>
          <ul class="items">
            <li>电话：{{ configStore.sitemap['site.linkphone'] }}</li>
            <li>邮箱：{{ configStore.sitemap['site.linkemail'] }}</li>
            <li>Q&nbsp;&nbsp;Q：{{ configStore.sitemap['site.linkqq'] }}</li>
          </ul>
        </div>
        <div class="col">
          <p class="code">
            <AuthImage :src="configStore.sitemap['site.qrcodewx']" v-if="configStore.sitemap['site.qrcodewx']" class="qrcode kefu" alt="二维码" />
            <img src="../../../../assets/common/kefu.jpg" v-else class="qrcode kefu" alt="二维码" />
          </p>
          <p class="small">扫一扫 加客服</p>
        </div>
      </div>
    </div>
    <div class="copyright">Copyright © 2018-2030 {{ configStore.sitemap['site.orgname'] }} All Rights Reserved.</div>
  </div>
</template>
<script lang="ts" setup>
import { useConfigStore } from '@/store';
const configStore = useConfigStore()
</script>
<style lang="scss" scoped>
.footer {
  display: flex;
  flex-direction: column;
  background-color: #ffffff;
  display: flex;
  justify-content: flex-end;
  align-items: center;
  background-color: #ffffff;
  .foot-container {
    width: 80vw;
    .row {
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      align-items: center;
      .col {
        color: #6c757d;
        .title {
          margin-bottom: 0.5rem;
          font-weight: 500;
          line-height: 1.2;
          font-size: 1.25rem;
        }
        .items {
          margin: 0px;
          padding: 0px;
          font-size: 1rem;
          line-height: 2rem;
        }
      }
    }
  }
  .copyright {
    display: flex;
    justify-content: center;
    align-content: center;
    line-height: 3;
    font-size: 1rem;
    border-top: 1px solid #dee2e6 !important;
    flex:1;
    width:100%;
  }
}
.kefu {
  width: 80px;
  height: 80px;
}
</style>
