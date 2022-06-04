<script>
import ImportWishlist from "./components/ImportWishlist.vue";
import UserInfo from "./components/UserInfo.vue";
import Wishlist from "./components/Wishlist.vue";

export default {
    name: "App",
    components: {
        ImportWishlist,
        UserInfo,
        Wishlist,
    },
    data() {
        return {
            userId: null,
        };
    },
};
</script>

<template>
    <header>
        <h1>Present-Roulette</h1>
        <UserInfo v-bind:userId="userId" />
    </header>

    <main>
        <div v-if="userId">
            <ImportWishlist />
            <Wishlist />
        </div>
        <div>
            <p>You are not logged in</p>
        </div>
    </main>
</template>

<style></style>
