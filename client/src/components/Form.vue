<script setup>
function onSubmit(event) {
    event.preventDefault();
}
</script>

<template>
    <form @submit="onSubmit">
        <label for="url">Lien URL</label>
        <input id="url" type="text">
        <button>Generer</button>
    </form>
</template>

<style scoped>
form {
    width: 50%;
    width: 310px;
    display: flex;
    margin: 2rem auto;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 2rem;
    background: white;
    border-radius: 6px;
    box-shadow: 0 1px 2px 0 rgb(0 0 0 / 0.05);
}

form>* {
    width: 100%;
    outline: none;
}

form>label {
    font-size: 1.2rem;
    font-weight: 700;
    color: #334155;
    margin-bottom: .5rem;
}

form>input {
    padding: 1rem;
    margin: .5rem 0;
    border: 1px solid #2563EB;
    border-radius: 5px;
    background: #DBEAFE;
}

form>button {
    padding: 1rem;
    background: #2563EB;
    color: white;
    border: none;
    margin: .5rem 0;
    /* padding: .5rem; */
    cursor: pointer;
    border-radius: 5px;
}

form>button:hover {
    background: #1E40AF;
}
</style>
