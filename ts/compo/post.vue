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
    <hr>

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
        <textarea class="textarea" v-model="postText" v-bind:disabled="isPush" maxlength="500"></textarea>
        <div class="count">{{postText.length}}/500</div>
      </div>
    </modal>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios, { AxiosError, AxiosResponse } from "axios";
import * as firebase from "firebase/app";
import "firebase/auth";
import InfiniteLoading, {StateChanger} from "vue-infinite-loading";
import ajax from "../src/ajax.ts";

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
          const url:string = "/postText/repry";
          const params:object = {
            token: idToken,
            text: text,
            repryID: _this.$route.params.id
          };

          ajax.post(url, params, _this);
        })
        .catch((err: firebase.auth.Error) => {
          alert(err.message);
          _this.$data.isPush = false;
        });
    },

    infiniteGet($state: StateChanger): void {
      const _this: any = this;
      const url: string = "/posts/reprys";
      _this.getNumber += 10;
      const params:object = {
        params: {
          number: _this.getNumber,
          repry: _this.$route.params.id
        }
      };

      ajax.get($state, url, _this, params);
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