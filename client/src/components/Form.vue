<script setup lang="ts">
import { ref, watch } from 'vue';
import { urlRegex } from '../constant';
import { toast } from '@steveyuowo/vue-hot-toast';

// const qrcode = ref<string>('iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAABo0lEQVR42uyYsbG0MAyE5SEgdAmU4tLs0lzKlUDogPH+s7L473h3l7551qAM8yVYq5WQ3HHHHX82IhhNUgkAHpJDO49cAQ8RWVqsufMhldDsyBmQ0JcWUXCs+1YFzY6cAuFYYffgFniqGk6BU8IB2Dfg9cgRMAwp1hxU1d9MbHLg5Gq2s/6l8cwNnFbMPI/iHdo+pgJE+LSiCp12fOZG2/UFRE4EwJ4QeJag93DNpgcAlZ8+ahOgdneR7A1gunkPkmUBxwNaUH1xWh8Am6QI9QxiSV0WJcAbwHtoLOHxSrSt9B+ynx0Qe+LgA6jEVehPPbgAOAtQz5I7k56K3sDrRfkBZN2pbTOsCJTrmOQBqEPVnGlpxX3U8+IMeCRwWEjoHJNEtHgvndcHMHqN9tKoJoxrumcAntoN0O1HaPzmT4utmYH/Gy2E0VHsr9kbYL/IqVg5Ms+phMMbYDsQvQGdaRnXMckPoH9gzH3n4Jffd0E+gKTFiyom7g/pnhuw9VXuNGQ2FL7/sLibGzBfwnDfVPq6byhvJjY5cMcdd/x6/AsAAP//HgYF/pjpdhcAAAAASUVORK5CYII=');
const isValid = ref<boolean>(true);
const longUrl = ref<string>('');
const qrcode = ref<string>('');
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
        <div class="qrcode">
            <div v-if="!qrcode" class="no-qrcode">
                <i class="fa-solid fa-qrcode"></i>
                No QR Code
            </div>
            <img v-if="qrcode" :src="'data:image/png;base64,' + qrcode" alt="" srcset="">
        </div>
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
