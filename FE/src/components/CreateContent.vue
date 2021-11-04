<template>
  <div>
    <div class="text-box">
      <h3>종류</h3>
      <input type="text" v-model="contents.kinds" placeholder="종류를 작성해주세요.">
    </div>
    <div class="text-box">
      <h3>설명</h3>
      <textarea type="text" v-model="contents.explain" placeholder="설명을 작성해주세요."></textarea>
    </div>
    <div class="text-box">
      <h3>추천가게</h3>
      <input type="text" v-model="contents.recommend" placeholder="추천가게를 작성해주세요.">
    </div>
    <button>작성완료</button>
  </div>
</template>
<script>
export default {
  name: '',
  components: {},
  data () {
    return {
      contents: {
      }
    }
  }
}
</script>
<style scoped>
.text-box {
  padding: 20px;
  margin-top: 15px;
  background-color: white;
  border-radius: 15px;
}

h3 {
  color: black;
}

input,
textarea {
  border: 1px solid #c4c4c4;
  width: 90%;
  padding: 4px 8px;
  margin-top: 6px;
  color: black;
  outline: none;
}

button {
  display: block;
  margin: 0 auto;
  margin-top: 20px;
  padding: 10px 20px;
  background-color: burlywood;
  color: white;
  border-radius: 10px;
  cursor: pointer;
}
</style>
