<template>
  <div>
    <div>
      <Header />
      <h3 class="flex">{{ this.snpt[0].title }}</h3>
    </div>
  </div>
</template>

<script>
export default {
  data: {
    snpt: {
      id: {
        type: String,
      },
      title: {
        type: String,
      },
      content: {
        type: String,
      },
      cookie: {
        type: String,
      },
    },
  },
  async fetch() {
    this.snpt = await fetch(`http://localhost:3333/s/${this.$route.query.id}`, {
      mode: "cors",
    }).then((res) => res.json());
    console.log(this.snpt);
  },
  head() {
    return {
      title: `SNPT | Snippet`,
    };
  },
};
</script>