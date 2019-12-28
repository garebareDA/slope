<template>
  <div>
    <h1 class="settingTitle">slope</h1>

    <div class="prfile">
      <div>
        <div>
          <img class="icon" v-bind:src="icon" />
        </div>

        <div>{{name}}</div>

        <div v-if="!guest">
          <div>
            <button class="button settingButton" v-on:click="credential">ユーザー名の変更</button>
          </div>

          <div v-if="loginEmail">
            <div>
              <button class="button settingButton" v-on:click="credential">メールアドレスの変更</button>
            </div>

            <div>
              <button class="button settingButton" v-on:click="credential">>パスワードの変更</button>
            </div>
          </div>
        </div>

        <div>
          <button class="button logout" v-on:click="logout">ログアウト</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import * as firebase from "firebase/app";
import "firebase/auth";

export default Vue.extend({
  created() {
    const _this: any = this;
    firebase.auth().onAuthStateChanged(user => {
      if (user) {
        if (user.email == null) {
          _this.$data.guest = true;
        }

        if (user.displayName != null) {
          _this.$data.name = user.displayName;
        }

        if (user.photoURL != null) {
          _this.$data.icon = user.photoURL;
        }

        _this.$data.provider = user.providerData[0]!.providerId;
        if(_this.$data.provider == "firebase") {
          _this.$data.loginEmail = true;
        }

      } else {
        _this.$router.push("/");
      }
    });
  },

  methods: {
    logout(): void {
      firebase.auth().signOut();
    },

    credential(): void {
      const _this = this;
      const userProvider:String = _this.$data.provider;

      if(userProvider === "google.com"){
        const provider:firebase.auth.AuthProvider = new firebase.auth.GoogleAuthProvider();
        recertification(provider);
      }

      if(userProvider === "github.com"){
        const provider:firebase.auth.AuthProvider = new firebase.auth.GithubAuthProvider();
        recertification(provider);
      }
    }
  },

  data() {
    return {
      icon: "",
      name: "Guest",
      guest: false,
      provider: null,
      loginEmail:false,
    };
  }
});

//TODO引数に遷移先
function recertification(provider:firebase.auth.AuthProvider) {
  const user: firebase.User | null = firebase.auth().currentUser;

  firebase
    .auth()
    .signInWithPopup(provider)
    .then((result:firebase.auth.UserCredential) => {
      const credential: any = result.credential;

      user!
        .reauthenticateWithCredential(credential)
        .then(() => {

        })
        .catch((err: firebase.auth.AuthError) => {
          alert(err.message);
        });
    })
    .catch((err: firebase.auth.AuthError) => {
      alert(err.message);
      return;
    });
}
</script>