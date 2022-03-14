<template v-if="clicc = 0">
  <div class="popup" id="popup">
    <h5><b>Hello!</b></h5>
    <p>
      This website is using biscuits to store data. By using it you agree to the
      <a
        href="https://www.youtube.com/watch?v=dQw4w9WgXcQ"
        target="_blank"
        id="tac"
        >terms and conditions</a
      >.
      <button @click="disablePopup()">I agrê.</button>
      <a href="https://google.com">I refoos.</a>
    </p>
  </div>
</template>

<script>
export default {
  name: "PopupComponent",

  methods: {
    disablePopup() {
      document.getElementById("popup").style.opacity = 0;
    },
  },
};
</script> 

<style src="~/assets/styles/popup.css"></style>