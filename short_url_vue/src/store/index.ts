import { createStore } from "vuex";
import base64 from "@utils/base64";
import jwt from "@utils/jwt";

const store = createStore({
  state() {
    return {
      token: {
        id: -1,
        refreshToken: "",
        acessToken: "",
        userInfo: {
          name: "noLogin",
          nickname: "noLogin",
          role: -1,
          urlLength: -1,
        },
      },
    };
  },
  mutations: {
    refreshToken(state, refreshToken: string) {
      //eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaXNzIjoibmV3cmVwb3J0Iiwic3ViIjoic29tZWJvZHkiLCJhdWQiOlsic29tZWJvZHlfZWxzZSJdLCJleHAiOjE2NzAwNTM4NTUsIm5iZiI6MTY2ODc1Nzg1NSwiaWF0IjoxNjY4NzU3ODU1fQ.TOaGid3KggjZVDu5mhbzNVtgOYT8hGBqHgyM0p8jKrk
      state.token.refreshToken = refreshToken;
      let info = new jwt().analysisClaims(refreshToken);
      localStorage.setItem("refreshToken", new base64().encode(refreshToken));
      state.token.id = info.id;
    },
    accessToken(state, accessToken: string) {
      //eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6ImFkbWluIiwibmlja25hbWUiOiJhZG1pbiIsInJvbGUiOjEsInVybExlbmd0aCI6OSwiaXNzIjoibmV3cmVwb3J0Iiwic3ViIjoic29tZWJvZHkiLCJhdWQiOlsic29tZWJvZHlfZWxzZSJdLCJleHAiOjE2Njg3NDU1NTIsIm5iZiI6MTY2ODc0NDk1MiwiaWF0IjoxNjY4NzQ0OTUyfQ.Dud765THBe4A3zudOooTUore_E8tILGm_3NEZFW54ZI
      state.token.acessToken = accessToken;
      let info = new jwt().analysisClaims(accessToken);
      state.token.userInfo.name = info.name;
      state.token.userInfo.nickname = info.nickname;
      state.token.userInfo.role = info.role;
      state.token.userInfo.urlLength = info.urlLength;
      localStorage.setItem("accessToken", new base64().encode(accessToken));
    },
    cleanToken(state) {
      localStorage.clear(); //清空本地存储
      state.token.id = -1;
      state.token.refreshToken = "";
      state.token.acessToken = "";
      state.token.userInfo.name = "noLogin";
      state.token.userInfo.nickname = "noLogin";
      state.token.userInfo.role = -1;
      state.token.userInfo.urlLength = -1;
    },
  },
});
export default store;
