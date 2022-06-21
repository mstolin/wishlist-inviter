<script>
export default {
    props: {
        userId: String,
    },
    emits: ['receivedUserId'],
    methods: {
        getUserId() {
            this.$emit('receivedUserId', 'testetst');
        }
    }
};
</script>

<template>
    <div v-if="userId">Welcome {{ userId }}</div>
    <div v-else>Log in or <button @click="getUserId">create new user</button></div>
</template>
