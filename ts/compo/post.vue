<template>
  <div>
    <h1 class="title">slope</h1>
    <div class="posts" >
      <div class="user">
        <img class="userIcon" v-bind:src="photoURL" alt="アイコン" />
        <p class="userName">{{userName}}</p>
      </div>
      <p class="text">{{text}}</p>
    </div>
    <button>反応</button>
    <hr>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios, { AxiosError, AxiosResponse } from "axios";
export default Vue.extend({
  created(): void {
    const _this = this;
    axios
      .get("/posts/post", {
        params: {
          number: _this.$route.params.id
        }
      })
      .then((result: AxiosResponse) => {
        console.log(result.data);
        const data = result.data;
        _this.$data.photoURL = data.photoURL;
        _this.$data.userName = data.userName;
        _this.$data.text = data.text;
      })
      .catch((err: AxiosError) => {
        alert(err.message);
      });
  },

  data() {
    return {
      photoURL: "",
      userName: "",
      text: ""
    };
  }
});
</script>