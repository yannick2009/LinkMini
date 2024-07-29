<template>
 <section class="container">
    <slot></slot>
 </section>
</template>

<style scoped lang="css">
    .container {
        width: 100%;
        height: 100dvh;
        background: #EFF6FF;
        display: flex;
        flex-direction: column;
        flex-wrap: wrap;
        justify-content: center;
        align-items: center;
    }
</style>
