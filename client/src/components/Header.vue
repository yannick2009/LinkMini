<template>
    <h1>Linkmini</h1>
    <p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Repudiandae dolores commodi nobis iure pariatur
        itaque.Lorem ipsum dolor sit amet consectetur, adipisicing elit. Repudiandae dolores commodi nobis iure pariatur
        itaque.
    </p>
    <nav class="menu">
        <RouterLink to="/"><button><i class="fa-solid fa-house-chimney"></i></button></RouterLink>
        <RouterLink to="/stats"><button><i class="fa-solid fa-chart-simple"></i></button></RouterLink>
    </nav>
</template>

<style scoped lang="css">
h1 {
    color: #2563EB;
}

p {
    width: 80%;
    margin: 1rem auto;
    max-width: 600px;
    font-size: .9rem;
    color: #334155;
    text-align: center;
}

.menu {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 1rem;
    margin: 1rem auto;
    width: fit-content;
}

.menu button {
    padding: 1rem;
    color: #2563EB;
    background: white;
    border: none;
    cursor: pointer;
    border-radius: 5px;
    font-size: 1.5rem;
}
</style>
