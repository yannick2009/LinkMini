<script setup la components: { Container },ng="ts">
import Header from './components/Header.vue';
import Container from './layout/Container.vue';
import Form from './components/Form.vue';
import { Toaster } from '@steveyuowo/vue-hot-toast';
import "@steveyuowo/vue-hot-toast/vue-hot-toast.css";
</script>

<template>
  <Toaster />
  <container>
    <Header />
    <Form />
  </container>
</template>

<style scoped></style>
