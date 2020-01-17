<template>
  <div>
    <h1 class="title">slope</h1>

    <div class="buttons" v-if="login">
        <button class="postButton" v-on:click="showPost" >+</button>
        <button class="settingButton" v-on:click="transition('/setting')" >設定</button>
    </div>

    <div class="posts" v-for="(item, index) in list" :key="index">
      <div class="user" v-on:click="transition('/post/' + item.ID)">
        <p>
          <img class="userIcon" v-bind:src="item.photoURL" alt="アイコン">
        </p>
        <p class="userName">{{item.userName}}</p>
      </div>
      <p class="text">{{item.text}}</p>
    </div>

    <modal name="post" width="90%" height="auto">
      <div class="postModal">
        <div>
          <button class="closeButton" v-on:click="hideModal('post')">×</button>
          <button class="post" v-on:click="post" v-bind:disabled="isPush">投稿</button>
        </div>
          <textarea class="textarea" v-model="postText" v-bind:disabled="isPush" maxlength="500"></textarea>
          <div class="count">{{postText.length}}/500</div>
      </div>
    </modal>

    <modal name="loginModal" width="90%" height="auto">
      <div class="modalLogin">
        <h1>ログインしましょう</h1>
        <button class="loginButton" v-on:click="transition('/login')">ログイン</button>
        <button class="backButton" v-on:click="hideModal('loginModal')">後で</button>
      </div>
    </modal>

    <infinite-loading @infinite="infiniteGet">
      <div slot="no-more"></div>
      <div slot="no-results"></div>
    </infinite-loading>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios, { AxiosError, AxiosResponse } from "axios";
import * as firebase from "firebase/app";
import "firebase/auth";
import InfiniteLoading from 'vue-infinite-loading';
import ajax from "../src/ajax.ts";

export default Vue.extend({
  mounted() {
    firebase.auth().onAuthStateChanged((user:firebase.User | null) => {
      if (user) {
        this.$data.login = true;
      } else {
        this.$data.login = false;
        this.$modal.show("loginModal");
      }
    })
  },

  methods: {
    showPost(): void {
      this.$modal.show("post");
    },

    hideModal(hide:string): void {
      this.$modal.hide(hide);
    },

    transition(url: string): void {
      this.$router.push(url);
    },

    post(): void {
      this.$data.isPush = true;
      const text: string = this.$data.postText;
      let _this: any = this;
      firebase
        .auth()
        .currentUser!.getIdToken(true)
        .then((idToken: string) => {
          const url:string = "/postText";
          const params:object = {
            token: idToken,
            text: text
          };

          ajax.post(url, params, _this);
        })
        .catch((err: firebase.auth.Error) => {
          alert(err.message);
          _this.$data.isPush = false;
        });
    },

    infiniteGet($state:any): void {
      const _this: any = this;
      _this.getNumber += 10;
      const url: string = "/posts";
      const params:object = {
        params: {
          number:_this.getNumber
        }
      };

      ajax.get($state, url, _this, params);
    }
  },

  components: {
    InfiniteLoading,
  },

  data() {
    return {
      isModalActive: false,
      login: false,
      postText: "",
      isPush: false,
      getNumber:0,
      list:[]
    };
  }
});
</script>