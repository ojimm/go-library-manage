<template>
  <iframe class="iframe-fullscreen" :src="iframeURL"></iframe>
</template>
<script setup>
const iframeURL = `${import.meta.env.VITE_APP_BASE_URL}/swagger/admin/index.html`;
</script>