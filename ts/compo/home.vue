<template>
  <div>
    <h1 class="title">slope</h1>

    <div class="buttons" v-if="login">
        <button class="postButton" v-on:click="showPost" >+</button>
        <button class="settingButton" v-on:click="settingButton" >設定</button>
    </div>

    <div class="posts" v-for="(item, index) in list" :key="index">
      <div>
        <div class="user">
          <img class="userIcon" v-bind:src="item.photoURL" alt="アイコン">
          <p class="userName">{{item.userName}}</p>
        </div>
        <p class="text">{{item.text}}</p>
      </div>
    </div>

    <modal name="post" width="90%" height="auto">
      <div class="postModal">
        <div>
          <button class="closeButton" v-on:click="hidePost">×</button>
          <button class="post" v-on:click="post" v-bind:disabled="isPush">投稿</button>
        </div>
        <textarea class="textarea" v-model="postText" v-bind:disabled="isPush"></textarea>
      </div>
    </modal>

    <modal name="loginModal" width="90%" height="auto">
      <div class="modalLogin">
        <h1>ログインしましょう</h1>
        <button class="loginButton" v-on:click="loginButton">ログイン</button>
        <button class="backButton" v-on:click="hideLogin">後で</button>
      </div>
    </modal>

    <infinite-loading @infinite="infiniteGet"></infinite-loading>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios, { AxiosError, AxiosResponse } from "axios";
import * as firebase from "firebase/app";
import "firebase/auth";
import InfiniteLoading from 'vue-infinite-loading';

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

    hidePost(): void {
      this.$modal.hide("post");
    },

    hideLogin(): void {
      this.$modal.hide("loginModal");
    },

    post(): void {
      this.$data.isPush = true;
      const text: string = this.$data.postText;
      let _this: any = this;
      firebase
        .auth()
        .currentUser!.getIdToken(true)
        .then((idToken: string) => {
          axios
            .post("/postText", {
              token: idToken,
              text: text
            })
            .then((res: AxiosResponse) => {
              console.log(res);
              _this.$data.postText = "";
              _this.$modal.hide("post");
              _this.$data.isPush = false;
              _this.$router.go({path: this.$router.currentRoute.path, force: true});
            })
            .catch((err: AxiosError) => {
              alert(err);
              _this.$data.isPush = false;
            });
        })
        .catch((err: firebase.auth.Error) => {
          alert(err);
          _this.$data.isPush = false;
        });
    },

    loginButton(): void {
      this.$router.push("/login");
    },
    settingButton():void {
      this.$router.push("/setting");
    },

    infiniteGet($state:any): void {
      const _this: any = this;
      _this.getNumber += 10;
      axios.get("/posts", {
        params: {
          number:_this.getNumber
        }
      }).then((result:AxiosResponse) => {
        const get:any = result.data
        get.reverse();
        _this.$data.list.push(...get);
        console.log(_this.$data.list);
        if(result.data.length != 10){
          $state.complete();
        }else{
          $state.loaded();
        }
      }).catch((err: AxiosError) => {
        $state.complete();
        alert(err.message);
      });
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