<template>
  <div class="home-wrapper">
    <h1>My CoffeeList</h1>
    <HomeContent />
    <router-link to="/create">
    </router-link>
  </div>
</template>

<script>
import HomeContent from '../components/HomeContent.vue'
import { APIS } from '../libs/axios'
export default {
  name: 'Home',
  components: {
    HomeContent
  },
  created () {
    APIS.TestLoad().then((res) => console.log(res.data))
  }
}
</script>
<style scoped>

h1 {
  text-align: center;
  padding-bottom: 10px;
  border-bottom: 1px solid white;
}

.home-wrapper {
  max-width: 480px;
  width: 100%;
  height: 70vh;
  background-color: #575757;
  margin: 0 auto;
  padding: 30px;
  border-radius: 15px;
  box-shadow: 4px 4px 4px rgba(0,0,0,0.3);
  position: relative;
}
</style>
