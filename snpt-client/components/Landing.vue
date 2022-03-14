<template>
  <div>
    <div class="landing flex">
      <div class="logo">SNPT</div>
      <div class="sub">F*#! Pastebin</div>
      <NuxtLink to="/upload" class="arrow">&#8609;</NuxtLink>
    </div>
  </div>
</template>

<script>
export default {
  name: "LandingComponent",
  async mounted() {
    await this.$axios
      .get("http://localhost:3333/cookie", {
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Methods": "GET,PUT,POST,DELETE,PATCH,OPTIONS",
      })
      .then((res) => {
        if (!document.cookie.match("biscuit")) {
          document.cookie = `biscuit=${res.data.cookie}`;
        }
      });
  },
};
</script> 

<style src="~/assets/styles/landingpage.css"></style>