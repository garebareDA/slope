<template>
  <div class="loginCenter">
    <h1 class="title">Slope</h1>
    <h3 class="login">再認証</h3>

    <div>
      <input type="text" placeholder="Eメール" v-model="email" class="loginText" />
    </div>

    <div>
      <input type="password" placeholder="パスワード" v-model="password" class="loginText" />
    </div>

    <div>
      <button v-on:click="login" class="button emailLogin">再認証</button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import * as firebase from "firebase/app";
import "firebase/auth";
import { ErrorHandler } from "vue-router/types/router";

export default Vue.extend({
  created() {
    const _this = this;
    firebase.auth().onAuthStateChanged(function(user) {
      if (!user) {
        _this.$router.push("/");
      }
    });
  },

  methods: {
    recertification(): void {
      const _this = this;
      const data = this.$data;
      const user: firebase.User | null = firebase.auth().currentUser;
      const credential: firebase.auth.AuthCredential = firebase.auth.EmailAuthProvider.credential(
        data.email,
        data.password
      );

      user!.reauthenticateWithCredential(credential).then(() => {
        const params:string = _this.$route.params['change'];
        _this.$router.push('/setting/' + params);
      }).catch((err:firebase.FirebaseError) => {
        alert(err.message);
      });
    }
  },

  data() {
    return {
      email: "",
      password: ""
    };
  }
});
</script>