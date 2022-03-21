<template>
  <div class="header">   
    <h1><NuxtLink to="/">SNPT</NuxtLink></h1>
       <div class="nav">
        <NuxtLink to="/home">Home</NuxtLink>
        <NuxtLink to="/upload">Upload</NuxtLink>
        <NuxtLink to="/search">Search</NuxtLink>
      </div>
  </div>
</template>

<script>
export default {
  name: 'HeaderComponent'
}
</script> 

<style src="~/assets/styles/header.css"></style>