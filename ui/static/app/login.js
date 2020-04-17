const Login = {
  template: '#login-page',
  data: () => ({
    code: ''
  }),
  methods: {
    doLogin: async function() {
      let response = await fetch('/api/voter/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          'code': this.code
        })
      })
      
      if (response.status == 200) {
        let voterCode = this.code
        localStorage.setItem('voter-code', voterCode)
        this.$router.push('/')
      } else if (response.status == 401) {
        UIkit.notification({
          message: `<span class="uk-text-light">Invalid login</span>`,
          status: 'danger',
          pos: 'bottom-center',
          timeout: 3000
        })
      } else {
        UIkit.notification({
          message: `<span class="uk-text-light">Sorry! Something error</span>`,
          status: 'danger',
          pos: 'bottom-center',
          timeout: 3000
        })
      }
    }
  }
}
