const Login = {
  template: '#login-page',
  data: () => ({
    username: '',
    password: ''
  }),
  methods: {
    doLogin: async function() {
      let response = await fetch('/api/admin/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          'username': this.username,
          'password': this.password
        })
      })
      
      if (response.status == 200) {
        let basicAuth = btoa(`${this.username}:${this.password}`)
        localStorage.setItem('basic-auth', basicAuth)
        this.$router.push('/')
      } else if (response.status == 401) {
        UIkit.notification({
          message: `Invalid login`,
          status: 'danger',
          pos: 'bottom-center',
          timeout: 3000
        })
      } else {
        UIkit.notification({
          message: `Sorry! Something error`,
          status: 'danger',
          pos: 'bottom-center',
          timeout: 3000
        })
      }
    }
  }
}
