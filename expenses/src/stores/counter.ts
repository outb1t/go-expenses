import { defineStore } from "pinia";

export const counterStore = defineStore({
  id: "counterStore",
  state: () => ({
    count: 0,
  }),
  getters: {
    doubleCount: (state) => state.count * 2,
  },
  actions: {
    increment() {
      this.count++;
    },
    waitAndIncrement() {
      setTimeout(() => {
        this.increment();
      }, 1000);
    },
  },
});
