<template>
  <div>
    <Header />
    <h3 class="flex">New SNPT</h3>
    <div class="flex">
      <form action="" method="post" v-on:submit.prevent="onSubmit()" id="frm">
        <input
          type="text"
          name="snpttitle"
          id="name"
          placeholder="Title"
          v-model="title"
          required
        />
        <textarea
          name="snptdesc"
          id="desc"
          placeholder="Paste your snippet here"
          v-model="content"
          required
          cols="60"
          rows="10"
          form="frm"
        ></textarea>
        <input type="submit" value="Submit" />
      </form>
    </div>
    <div class="flex">
      <h4 v-if="success">
        <i>SNPT uploaded! (id: {{ success }})</i>
      </h4>
    </div>
  </div>
</template>

<style scoped src="~/assets/styles/upload.css"></style>

<script>
import axios from "axios";
export default {
  data() {
    return {
      success: false,
    };
  },
  props: {
    title: {
      type: String,
      required: true,
    },
    content: {
      type: String,
      required: true,
    },
  },
  methods: {
    onSubmit() {
      let data = {
        title: this.title,
        content: this.content,
        cookie: document.cookie.split("=").pop(),
      };
      console.log(data);
      axios
        .post("http://localhost:3333/create", data, {
          headers: {
            Accept: "application/json",
          },
        })
        .then((res) => {
          console.log(res);
          this.success = res.data.answer.id;
        });
    },
  },
  head() {
    return {
      title: "SNPT | Upload",
    };
  },
};
</script>

