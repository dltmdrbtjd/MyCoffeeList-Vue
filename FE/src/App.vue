<template>
  <div class="main-wrapper">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',

  data: () => ({
    //
  })
}
</script>
<style>

* {
  color: #DEDEDE;
  text-decoration: none;
  margin: 0;
  padding: 0;
}

.main-wrapper {
  width: 100%;
  height: 100vh;
  background-color: #262626;
  padding-top: 50px;
}

</style>
