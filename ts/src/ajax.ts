import axios, { AxiosError, AxiosResponse } from "axios";
import { StateChanger } from "vue-infinite-loading";

function get($state: StateChanger, url: string, _this: any, params: object): void {
  axios.get(url, params)
  .then((result: AxiosResponse) => {
    const get: any = result.data
    get.reverse();
    _this.$data.list.push(...get);
    console.log(_this.$data.list);
    if (result.data.length < 9) {
      $state.complete();
    } else {
      $state.loaded();
    }
  }).catch((err: AxiosError) => {
    $state.complete();
    alert(err.message);
  });
}

function post(url: string, params: object, _this: any): void {
  axios
    .post(url, params)
    .then(() => {
      _this.$data.postText = "";
      _this.$modal.hide("post");
      _this.$data.isPush = false;
      _this.$data.list = [];
      location.reload(true);
    })
    .catch((err: AxiosError) => {
      alert(err.message);
      _this.$data.isPush = false;
    });
}

export default{
  post,
  get
}