<template>
  <div class="create-wrapper">
    <h1>Post Create</h1>
    <CreateContent />
  </div>
</template>
<script>
import CreateContent from '../components/CreateContent.vue';
export default {
  name: '',
  components: {CreateContent},
  data() {
    return {
      example: ''
    }
  },
  setup() {},
  created() {},
  mounted() {},
  unmounted() {},
  methods: {},
}
</script>
<style scoped>

h1 {
  text-align: center;
  padding-bottom: 10px;
  border-bottom: 1px solid white;
}

.create-wrapper {
  max-width: 480px;
  width: 100%;
  background-color: #575757;
  margin: 0 auto;
  padding: 30px;
  border-radius: 15px;
  box-shadow: 4px 4px 4px rgba(0,0,0,0.3);
  position: relative;
}
</style>