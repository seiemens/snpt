<template>
  <div>
    <Header />
    <div class="flex">
      <h3 v-if="$fetchState.pending">Fetching Snpts...</h3>
      <h3 v-else-if="$fetchState.error">
        An Error occured while fetching ideas.
      </h3>
    </div>
    <div class="flex" v-if="snpts">
      <div v-for="snpt in snpts" :key="snpt.id">
        <form
          action=""
          method="post"
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
          <input type="submit" value="Edit" />
        </form>
        <div class="flex">
          <button @click="deleteSnpt()">Delete</button>
        </div>
      </div>
    </div>
    <div class="flex" v-else>
      <h3>There is no snippet with this id.</h3>
    </div>
    <div class="flex">
      <h4 v-if="success">
        <i>{{ success }}</i>
      </h4>
    </div>
  </div>
</template>

<style scoped src="~/assets/styles/snpt.css">
</style>

<script>
import axios from "axios";

export default {
  data() {
    return {
      snpts: [],
      success: false,
    };
  },
  props: {
    id: {
      type: String,
    },
    title: {
      type: String,
    },
    content: {
      type: String,
    },
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
        id: this.snpts[0].id,
        title: this.snpts[0].title,
        content: this.snpts[0].content,
        cookie: document.cookie.split("=").pop(),
      };
      console.log(data);
      axios
        .post("http://localhost:3333/edit", data, {
          headers: {
            Accept: "application/json",
          },
        })
        .then((res) => {
          console.log(res);
          this.success = res.data.answer;
        });
    },
    deleteSnpt() {
      let data = {
        id: this.snpts[0].id,
        cookie: document.cookie.split("=").pop(),
      };
      console.log(data);
      axios.post("http://localhost:3333/delete", data).then((res) => {
        console.log(res);
        this.success = res.data.answer;
      });
    },
  },
};
</script>