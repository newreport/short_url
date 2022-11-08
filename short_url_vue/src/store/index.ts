import { createStore } from 'vuex'

const store = createStore({
  state () {
    return {
      count: 5
    }
  },
  mutations: {
    increment (state) {
      state.count++
    }
  }
})
export default store