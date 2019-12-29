<template>
  <div class="loginCenter">
    <h1 class="title">Slope</h1>
    <h3 class="login">{{changeText}}を変更</h3>

    <div>
      <input v-bind:placeholder="'新しい' + changeText + 'を入力'" v-model="text" class="loginText" />
    </div>

    <div>
      <button v-on:click="change" class="button emailLogin">変更</button>
    </div>

  </div>
</template>

<script lang="ts">
import Vue from "vue";
import * as firebase from "firebase/app";
import "firebase/auth";
import { ErrorHandler } from "vue-router/types/router";

export default Vue.extend({
  created(){
    const _this = this;
    firebase.auth().onAuthStateChanged(function(user) {
      if (!user) {
        _this.$router.push("/");
      }
    });
  },

  beforeMount(){
    const _this = this;
    const params:string = _this.$route.params['change'];

    if (params === "email") {
      _this.$data.changeText = "Eメール"
    }else if (params === "password"){
      _this.$data.changeText = "パスワード"
    }else{
      _this.$router.push("/");
    }
  },

  methods: {
    change(): void {
      const _this = this;
      let name:string;
      const user:firebase.User | null = firebase.auth().currentUser;
      const params:string = _this.$route.params['change'];
      const text:string = _this.$data.text;

      if (user === null) {
        _this.$router.push("/");
      }

      if (text === ""){
        alert("入力してください");
        return;
      }

      if (params === "email") {

        user!.updateEmail(text).then(() => {
          _this.$router.push("/");
        }).catch((err:firebase.FirebaseError) => {
          alert(err.message);
        });

      }else if (params === "password"){

        user!.updatePassword(text).then(() => {
          _this.$router.push("/");
        }).catch((err:firebase.FirebaseError) => {
          alert(err.message);
        });

      }else{
        _this.$router.push("/");
      }
    }
  },

  data() {
    return {
      text: "",
      changeText:"",
    };
  }
});
</script>