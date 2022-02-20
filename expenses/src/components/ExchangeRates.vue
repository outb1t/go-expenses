<template>
  <div>{{ msg }}</div>
  <div>{{ this.rates }}</div>
  <div>{{ store.count }}</div>
  <div>{{ store.doubleCount }}</div>
  <button v-on:click="store.increment()">Double count</button>
  <button v-on:click="store.waitAndIncrement()">Wait and increment</button>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import http from "@/http";
import { counterStore } from "@/stores/counter";

export default defineComponent({
  props: {
    msg: String,
  },
  data() {
    return {
      rates: {},
      store: counterStore(),
    };
  },
  mounted() {
    this.loadRates();
  },
  methods: {
    loadRates(): void {
      http.get("/rates").then((response) => {
        this.rates = response.data;
      });
    },
  },
});
</script>

<style scoped></style>
