<template>
  <div>
    <div>
      <Header />
      <h3 class="flex">Search!</h3>
      <div class="flex">
        <form action="" method="get" v-on:submit.prevent="onSubmit()" id="frm">
          <input
            type="text"
            id="name"
            placeholder="Search String"
            v-model="title"
            required
          />
          <input type="submit" value="🔎" />
        </form>
      </div>
    </div>
  </div>
</template>

<style src="~/assets/styles/search.css"></style>

<script>
export default {
  data() {
    return {
      title: "",
    };
  },
  methods: {
    async onSubmit() {
      this.$router.push(`/snpt?id=${this.title}`);
    },
  },
  head() {
    return {
      title: "SNPT | Search",
    };
  },
};
</script>