<template>
  <div>
    <h1 class="title">slope</h1>
    <div class="repry">
      <div class="user">
        <p>
          <img class="userIcon" v-bind:src="photoURL" alt="アイコン" />
        </p>
        <p class="userName">{{userName}}</p>
      </div>
      <p class="text">{{text}}</p>
    </div>
    <div class="loginCenter">
      <button class="button emailLogin" v-if="isLogin" v-on:click="showPost">反応</button>
    </div>
    <hr />

    <div class="repry" v-for="(item, index) in list" :key="index">
      <div class="user">
        <p>
          <img class="userIcon" v-bind:src="item.photoURL" alt="アイコン" />
        </p>
        <p class="userName">{{item.userName}}</p>
      </div>
      <p class="text">{{item.text}}</p>
    </div>
    <infinite-loading @infinite="infiniteGet">
      <div slot="no-more"></div>
      <div slot="no-results"></div>
    </infinite-loading>

    <modal name="post" width="90%" height="auto">
      <div class="postModal">
        <div>
          <button class="closeButton" v-on:click="hidePost">×</button>
          <button class="post" v-on:click="post" v-bind:disabled="isPush">投稿</button>
        </div>
        <textarea class="textarea" v-model="postText" v-bind:disabled="isPush"></textarea>
      </div>
    </modal>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios, { AxiosError, AxiosResponse } from "axios";
import * as firebase from "firebase/app";
import "firebase/auth";
import InfiniteLoading from "vue-infinite-loading";

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

  mounted(): void {
    const _this = this;
    firebase.auth().onAuthStateChanged((user: firebase.User | null) => {
      if (user) {
        _this.$data.isLogin = true;
      } else {
        _this.$data.isLogin = false;
      }
    });
  },

  methods: {
    post(): void {
      this.$data.isPush = true;
      const text: string = this.$data.postText;
      const _this: any = this;
      firebase
        .auth()
        .currentUser!.getIdToken(true)
        .then((idToken: string) => {
          axios
            .post("/postText/repry", {
              token: idToken,
              text: text,
              repryID: _this.$route.params.id
            })
            .then((res: AxiosResponse) => {
              _this.$data.postText = "";
              _this.$modal.hide("post");
              _this.$data.isPush = false;
              _this.$router.go({
                path: _this.$router.currentRoute.path,
                force: true
              });
            })
            .catch((err: AxiosError) => {
              alert(err.message);
              _this.$data.isPush = false;
            });
        })
        .catch((err: firebase.auth.Error) => {
          alert(err.message);
          _this.$data.isPush = false;
        });
    },

    infiniteGet($state: any): void {
      const _this: any = this;
      _this.getNumber += 10;
      axios
        .get("/posts/reprys", {
          params: {
            number: _this.getNumber,
            repry: _this.$route.params.id
          }
        })
        .then((result: AxiosResponse) => {
          const get: any = result.data;
          get.reverse();
          _this.$data.list.push(...get);
          console.log(_this.$data.list);
          if (result.data.length != 10) {
            $state.complete();
          } else {
            $state.loaded();
          }
        })
        .catch((err: AxiosError) => {
          $state.complete();
          alert(err.message);
        });
    },

    showPost(): void {
      this.$modal.show("post");
    },

    hidePost(): void {
      this.$modal.hide("post");
    }
  },

  components: {
    InfiniteLoading
  },

  data() {
    return {
      photoURL: "",
      userName: "",
      text: "",

      isModalActive: false,
      login: false,
      postText: "",
      isPush: false,
      getNumber: 0,
      list: [],

      isLogin: false
    };
  }
});
</script>