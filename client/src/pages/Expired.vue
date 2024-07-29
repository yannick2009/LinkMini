<script setup lang="ts">
import { RouterLink } from 'vue-router';
import Container from '../layout/Container.vue';
</script>

<template>
    <Container>
        <i class="fa-solid fa-hourglass-end"></i>
        <h1>Link has expire</h1>
        <p>Lorem ipsum dolor sit, amet consectetur adipisicing elit. Perspiciatis nam reiciendis at, quidem sitsimil</p>
        <RouterLink to="/"><button>Generate another</button></RouterLink>
    </Container>
</template>

<style scoped lang="css">
i {
    margin: 2rem 0;
    font-size: 10rem;
    color: #2563EB;
}

h1 {
    font-size: 3rem;
    margin: 1rem 0;
}

p {
    width: 80%;
    margin: 1rem auto;
    max-width: 600px;
    font-size: 1rem;
    color: #334155;
    text-align: center;
}

button {
    margin-top: 1.5rem;
    padding: 1.5rem;
    background-color: #2563EB;
    color: white;
    border: none;
    border-radius: 6px;
    cursor: pointer;
}
</style>
