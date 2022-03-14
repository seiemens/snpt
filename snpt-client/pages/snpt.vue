<template>
  <div>
    <Header />
    <div class="flex">
      <h3 v-if="$fetchState.pending">Fetching Snpts...</h3>
      <h3 v-else-if="$fetchState.error">
        An Error occured while fetching ideas.
      </h3>
    </div>
    <div v-for="snpt in snpts" :key="snpt.id" v-on:submit.prevent="onSubmit()">
      <form
        action=""
        v-on:submit.prevent="onSubmit()"
        id="editForm"
        class="flex snpt"
      >
        <h3>{{ snpt.id }}</h3>
        <input type="text" v-model="snpt.title" id="title" />
        <textarea
          type="text"
          v-model="snpt.content"
          form="editForm"
          id="desc"
        />
        <input type="submit" value="edit" />
      </form>
    </div>
  </div>
</template>

<style scoped src="~/assets/styles/snpt.css">
</style>

<script>
export default {
  data() {
    return {
      snpts: [],
    };
  },
  async fetch() {
    this.snpts = await fetch(
      `http://localhost:3333/s/${this.$route.query.id}`,
      {
        mode: "cors",
      }
    ).then((res) => res.json());
    console.log(this.snpts);
  },
  head() {
    return {
      title: `SNPT | Snippet`,
    };
  },
  methods: {
    onSubmit() {
      let data = {
        title: this.title,
        content: this.content,
      };
      axios
        .post("http://localhost:3333/create", data, {
          headers: {
            Accept: "application/json",
          },
        })
        .then((res) => {
          console.log(res);
          this.success = res.data.success ? true : false;
        });
    },
  },
};
</script>