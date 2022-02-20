<template>
  <div>{{ msg }}</div>
  <div>{{ this.rates }}</div>
</template>

<script lang="ts">

import {defineComponent} from "vue";
import http from "@/http";

export default defineComponent({
  props: {
    msg: String,
  },
  data() {
    return {
      rates: {}
    }
  },
  mounted() {
    this.loadRates()
  },
  methods: {
    loadRates(): void {
      http.get('/rates').then(response => {
        this.rates = response.data
      })
    }
  }
})

</script>

<style scoped>

</style>