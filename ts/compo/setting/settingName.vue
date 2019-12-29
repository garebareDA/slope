<template>
  <div class="loginCenter">
    <h1 class="title">Slope</h1>
    <h3 class="login">名前を変更</h3>

    <div>
      <input placeholder="新しい名前を入力" v-model="name" class="loginText" />
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

  methods: {
    change(): void {
      const _this = this;
      let name:string;
      const user:firebase.User | null = firebase.auth().currentUser;

      if (user === null) {
        _this.$router.push("/");
      }

      if (_this.$data.name != ""){
        name = _this.$data.name;
      }else{
        alert("名前を入力してください");
        return;
      }

      user!.updateProfile({
        displayName: name,
      }).then(() =>{
        _this.$router.push("/");
      }).catch((err:firebase.FirebaseError) => {
        alert(err.message);
      });
    }
  },

  data() {
    return {
      name: "",
    };
  }
});
</script>