<script setup lang="ts">
import { ref, watch } from 'vue';
import Header from '../components/Header.vue';
import Container from '../layout/Container.vue';
import { toast } from '@steveyuowo/vue-hot-toast';
import { shortUrlRegex } from '../constant';
import { getShortUrlStats } from '../api';

const isValid = ref<boolean>(true);
const shortURL = ref<string>('');
const clickNum = ref<number>(0);

function onSubmit(event: Event) {
    event.preventDefault();
    const urlHash = shortURL.value.split('/').pop();
    getShortUrlStats(urlHash ?? '')
        .then((res) => {
            clickNum.value = res?.data?.clickNum || 0;
        })
        .catch((err) => {
            console.error(err);
            toast.error('Failed to get short URL stats');
        });
}

watch(shortURL, (newVal) => {
    if (!shortUrlRegex.test(newVal)) {
        isValid.value = false;
    } else {
        isValid.value = true;
    }
});
</script>

<template>
    <Container>
        <Header />
        <form @submit="onSubmit">
            <label for="short">Short URL</label>
            <input :class="!isValid && shortURL.trim().length > 0 ? 'invalid' : ''" v-model="shortURL" type="text"
                name="" id="">
            <button :disabled="!isValid || shortURL.trim().length == 0">View stats</button>
        </form>
        <div class="stats">
            <h6>{{ clickNum }} <i class="fa-solid fa-arrow-up-right-from-square"></i></h6>
            <p>Clicks</p>
        </div>
    </Container>
</template>

<style scoped lang="css">
form,
.stats {
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

.stats h6 {
    font-size: 3rem;
    color: #2563EB;
    margin: 1rem 0;
}

.stats i {
    margin-left: .5rem;
    font-size: 1.5rem;
}

.stats p {
    font-size: 2rem;
    color: #334155;
}

.invalid {
    border: 1px solid #E11D48;
    background: #FECDD3;
}
</style>
