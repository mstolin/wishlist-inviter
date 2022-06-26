<script>
import UserHelper from "../helper/UserHelper";

const url = "";

export default {
  name: "UserInfo",
  props: {
    userId: String,
  },
  emits: ["receivedUserId"],
  methods: {
    getUserId() {
      UserHelper.fetchUserId()
        .then((json) => {
          if (json.hasOwnProperty("id")) {
            this.$emit("receivedUserId", json.id);
          } else {
            throw new Error(`ID is not given in response ${json}`);
          }
        })
        .catch((error) => alert(error));
    },
    async _fetchUserId() {
      return fetch(url, { method: "POST" }).then((data) => data.json());
    },
  },
};
</script>

<template>
  <div v-if="userId">
    Welcome <strong>{{ userId }}</strong>
  </div>
  <div v-else>
    Log in or <button @click="getUserId">create new user</button>
  </div>
</template>
