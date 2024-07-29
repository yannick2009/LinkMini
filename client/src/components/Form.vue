<script setup lang="ts">
import { ref, watch } from 'vue';
import { urlRegex } from '../constant';
import { toast } from '@steveyuowo/vue-hot-toast';
import { createShortUrl } from '../api';

const serverUrl: string = import.meta.env.VITE_SERVER_URL;
const isValid = ref<boolean>(true);
const longUrl = ref<string>('');
const shortURL = ref<string>('');
const qrcode = ref<string>('');

function onSubmit(event: Event) {
    event.preventDefault();
    createShortUrl(longUrl.value)
        .then((res) => {
            shortURL.value = serverUrl + '/' + res?.data.urlHash || '';
            qrcode.value = res?.data.qrCode || '';
        })
        .catch((err) => {
            console.error(err);
            toast.error('Failed to generate short URL');
        });
}

const copyLink = (event: MouseEvent) => {
    const target = event.target as HTMLElement | null;
    if (target?.innerText.length === 0) return;
    navigator.clipboard.writeText(target?.innerText || '');
    toast.success('Link copied');
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
        <div class="qrcode">
            <div v-if="!qrcode" class="no-qrcode">
                <i class="fa-solid fa-qrcode"></i>
                No QR Code
            </div>
            <img v-if="qrcode" :src="'data:image/png;base64,' + qrcode" alt="" srcset="">
        </div>
        <label for="short">short URL</label>
        <p id="short" @click="copyLink">{{ shortURL }}</p>
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

form .no-qrcode {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
    width: 100%;
    font-size: 1.2rem;
    font-weight: 700;
    background-color: #F3F4F6;
    color: #2563EB;
}

form .no-qrcode i {
    font-size: 5rem;
    margin-bottom: 1rem;
}

form .qrcode {
    width: 100%;
    height: 200px;
    background: white;
    border-radius: 5px;
    margin-bottom: 1rem;
    border: 1px dashed #2563EB;
}

form .qrcode img {
    width: 100%;
    height: 100%;
    object-fit: contain;
}

form>* {
    width: 100%;
    outline: none;
}

form label {
    font-size: 1.2rem;
    font-weight: 700;
    color: #334155;
    margin: .5rem 0;
}

form input {
    padding: 1.2rem;
    margin: .5rem 0;
    border: 1px solid #2563EB;
    border-radius: 5px;
    background: #DBEAFE;
}

form p {
    padding: 1.2rem;
    margin: .5rem 0;
    border: 1px dashed #2563EB;
    border-radius: 5px;
    background: #DBEAFE;
    cursor: pointer;
}

form p:active {
    background: #93C5FD;
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
