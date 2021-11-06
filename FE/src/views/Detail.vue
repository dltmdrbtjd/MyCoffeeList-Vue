<template>
  <div class="detail-wrapper">
    <router-link to="/">
      <h1>My Coffee</h1>
    </router-link>
    <DetailContent />
  </div>
</template>
<script>
import DetailContent from '../components/DetailContent'
export default {
  name: '',
  components: {
    DetailContent
  },
  data () {
    return {
      example: ''
    }
  }
}
</script>
<style scoped>

h1 {
  text-align: center;
  padding-bottom: 10px;
  border-bottom: 1px solid white;
}

.detail-wrapper {
  max-width: 600px;
  width: 100%;
  margin: 0 auto;
  background-color: steelblue;
  padding: 20px;
  border-radius: 15px;
}

</style>
