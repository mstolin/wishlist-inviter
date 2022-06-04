<script>
export default {
    props: {
        userId: String,
    },
};
</script>

<template>
    <div v-if="userId">Welcome {{ userId }}</div>
    <div v-else>Log in or create new user</div>
</template>
