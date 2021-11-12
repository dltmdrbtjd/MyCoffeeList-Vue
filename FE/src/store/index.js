import axios from 'axios'
import { createStore } from 'vuex'

export default createStore({
  state: {
    coffeeList: []
  },
  getters: {
    getCoffee: (state) => (id) => {
      return state.coffeeList.find((coffee) => coffee.id === id)
    }
  },
  mutations: {
    loadCoffeeList(state, list) {
      const coffee = list || []
      state.coffeeList.push(coffee)
    }
  },
  actions: {
    async loadCoffeeList({ commit }) {
      const res = await axios.get('http://localhost:5000')
      commit('loadCoffeeList', res.data)
    },

    pushCoffee({ commit }, coffee) {
      commit('loadCoffeeList', coffee)
    }
  }
})
