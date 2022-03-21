<template>
  <div>
    <Header />
    <div class="flex home">
      <h3>Welcome to SNPT!</h3>
      <div class="home">
        <h4>
          Enjoy our Service, which is definitely better than pastebin.
          <br />Legal note: By using this page you automatically agree to the
          terms of service and use of biscuits, even if you dont accept the
          fancy little popup.
        </h4>
      </div>
    </div>
    <Biscuitpopup />
  </div>
</template>

<script>
export default {
  head() {
    return {
      Biscuitpopuptitle: "SNPT | Home",
    };
  },
};
</script>