<template>
  <div class="detail-wrapper">
    <h2>종류</h2>
    <p>에스프레소</p>
    <h2>설명</h2>
    <p>에스프레소 설명 쏼라쏼라</p>
    <h2>추천가게</h2>
    <p>커피빈</p>
    <div class="btn-wrapper">
      <button>수정하기</button>
      <button class="del-btn">삭제하기</button>
    </div>
  </div>
</template>
<script>
export default {
  name: '',
  components: {},
  data () {
    return {
      example: ''
    }
  }
}
</script>
<style scoped>
p {
  margin-bottom: 15px;
}

.btn-wrapper {
  display: flex;
  justify-content: center;
  align-content: center;
  margin-top: 50px;
}

button {
  padding: 10px 20px;
  color: #262626;
  background-color: whitesmoke;
  border-radius: 12px;
}

.del-btn {
  margin-left: 60px;
}
</style>
