<template>
  <div class="home-content-wrapper" v-on:click="MoreData">
    <h2>종류</h2>
    <p>espresso</p>
    <h2>설명</h2>
    <p>에스프레소 또는 카페 에스프레소는 고압,고온 하의 물을 미세하게 분쇄한 커피 가루에 가해 추출해내는 고농축 커피의 일종이다.</p>
    <h2>추천가게</h2>
    <p class="recommend">커피빈</p>
    <div v-if="isMoreData">
      <h2>종류</h2>
      <p>espresso</p>
      <h2>설명</h2>
      <p>에스프레소 또는 카페 에스프레소는 고압,고온 하의 물을 미세하게 분쇄한 커피 가루에 가해 추출해내는 고농축 커피의 일종이다.</p>
      <h2>추천가게</h2>
      <p class="recommend">커피빈</p>
    </div>
  </div>
  <button class="create-btn" :class="bottomBtnClass">게시글 작성하기</button>
</template>
<script>
export default {
  name: 'HomeContent',
  data () {
    return {
      example: '',
      isMoreData: false,
      height: 0,
      bottomBtnClass: 'fixed'
    }
  },
  methods: {
    MoreData () {
      if (this.isMoreData) {
        this.isMoreData = false
      } else {
        this.isMoreData = true
      }
    },
    Height () {
      this.height = window.innerHeight
      console.log(this.height)
    }
  },
  mounted () {
    window.addEventListener('resize', this.Height)
  },
  watch: {
    height () {
      const el = document.querySelector('.home-content-wrapper')
      if (!el) {
        return
      }
      this.bottomBtnClass = this.height - 350 < el.offsetHeight ? '' : 'fixed'
    }
  }
}
</script>
<style scoped>
.home-content-wrapper {
  max-width: 400px;
  width: 100%;
  background-color: #262626;
  padding: 20px;
  margin: 0 auto;
  margin-top: 20px;
  border-radius: 15px;
  box-shadow: 4px 4px 4px rgba(0,0,0,0.5);
}

p {
  margin-bottom: 15px;
}

.recommend {
  color:bisque;
  font-weight: 600;
}

.create-btn {
  background-color: #262626;
  padding: 10px 20px;
  border-radius: 12px;
  cursor: pointer;
}

.fixed {
  position: absolute;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
}
</style>
