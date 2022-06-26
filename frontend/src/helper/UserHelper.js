export default class UserHelper {
  static async fetchUserId() {
    const host = "http://localhost:8080";
    let url = `${host}/users`;
    return fetch(url, { method: "POST" }).then((response) => {
      if (!response.ok) {
        throw new Error(`HTTP status: ${response.status}`);
      }
      return response.json();
    });
  }
}
