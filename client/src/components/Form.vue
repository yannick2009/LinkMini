<script setup lang="ts">
import { ref, watch } from 'vue';
import { urlRegex } from '../constant';
import { toast } from '@steveyuowo/vue-hot-toast';


const isValid = ref<boolean>(true);
const longUrl = ref<string>('');
function onSubmit(event: Event) {
    event.preventDefault();
    toast.success('Form submitted');
}

watch(longUrl, (newVal) => {
    if (longUrl.value.length > 0 && !urlRegex.test(newVal)) {
        isValid.value = false;
    } else {
        isValid.value = true;
    }
});

</script>

<template>
    <form>
        <label>QR Code</label>
        <div class="qrcode"></div>
        <label for="short">short URL</label>
        <input disabled id="short" type="text">
    </form>
    <form @submit="onSubmit">
        <label for="long">Your long URL</label>
        <input :class="!isValid ? 'invalid' : ''" v-model="longUrl" id="long" type="text">
        <button :disabled="!isValid || longUrl.trim().length == 0">Generate</button>
    </form>
</template>

<style scoped>
form {
    width: 50%;
    width: 310px;
    display: flex;
    margin: .5rem auto;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 2rem;
    background: white;
    border-radius: 6px;
    box-shadow: 0 1px 2px 0 rgb(0 0 0 / 0.05);
}

form .qrcode {
    width: 100%;
    height: 170px;
    background: #F3F4F6;
    border-radius: 5px;
    margin-bottom: 1rem;
    border: 1px dashed #2563EB;
}

form>* {
    width: 100%;
    outline: none;
}

form>label {
    font-size: 1.2rem;
    font-weight: 700;
    color: #334155;
    margin: .5rem 0;
}

form>input {
    padding: 1.2rem;
    margin: .5rem 0;
    border: 1px solid #2563EB;
    border-radius: 5px;
    background: #DBEAFE;
}

form input:disabled {
    background: #EFF6FF;
    border: 1px dashed #2563EB;
}

form>button {
    padding: 1.2rem;
    background: #2563EB;
    color: white;
    border: none;
    margin: .5rem 0;
    cursor: pointer;
    border-radius: 5px;
}

form>button:hover {
    background: #1E40AF;
}

.invalid {
    border: 1px solid #E11D48;
    background: #FECDD3;
}
</style>
